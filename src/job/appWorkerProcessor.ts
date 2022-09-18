import { Job } from "bullmq";
import Mercury from "@postlight/mercury-parser";
import axios from 'axios';
import Container from "typedi";
import * as tf from '@tensorflow/tfjs';
import { DataSource } from "typeorm";

import db from '../config/db';
import { AppJobData, AppJobNames, IndexDocumentData } from "./types";
import { AquilaClientService } from "../lib/AquilaClientService";
import { AccountStatus } from "../service/dto/AuthServiceDto";
import { CollectionTemp } from "../entity/CollectionTemp";
import { Collection } from "../entity/Collection";
import { Bookmark, BookmarkStatus } from "../entity/Bookmark";
import { BookmarkPara } from "../entity/BookmarkPara";
import { BookmarkParaTemp } from "../entity/BookmarkParaTemp";
import { BookmarkTemp, BookmarkTempStatus } from "../entity/BookmarkTemp";


async function summarize(html: string) {
	const response = await axios.post('http://localhost:5008/process', { html });	
	return response.data.result;
}

let aquilaClient, dataSource: DataSource;

export default async function(job: Job<AppJobData, void, AppJobNames>) {
	if(job.name === AppJobNames.INDEX_DOCUMENT) {
			const { data  } = <{data: IndexDocumentData}>job;
			// extract metadata from html
			const parsedHtml = await Mercury.parse(data.bookmark.url, { html: data.bookmark.html});
			console.log("From worker: parsedHtml", job.data, parsedHtml);
			// generate array summary from text content
			const summary = await summarize(parsedHtml.content || "");

			// connect to db
			if(!dataSource) {
				dataSource = await db.initialize();
			}

			// connect to aquila
			const aquilaClient = Container.get(AquilaClientService)
			await aquilaClient.connect();

			// load collection
			let collection: Collection | CollectionTemp | null;
			if(data.accountStatus === AccountStatus.TEMPORARY) {
				collection = await CollectionTemp.findOne({ where: { id: data.bookmark.collectionId }});
			}else {
				collection = await Collection.findOne({ where: { id: data.bookmark.collectionId }});
			}
			if(collection === null) {
				throw new Error('Invalid collection');
			}
			// generate vector from array of paragraph 
			const vectorArray = await aquilaClient.getHubServer().compressDocument(collection.aquilaDbName, summary);
			console.log("From worker job: Hub output", summary, vectorArray);
			// bulk insert into aquiladb
			const documents = summary.map((para: string, index: number) => {
				return {
					metadata: {
						para
					},
					code: vectorArray[index]
				}
			})

			let bookmarkParas: BookmarkPara[] | BookmarkParaTemp[];
			await db.transaction(async transactionalEntityManager => {
				// insert all para to bookmark_para table
				if(data.accountStatus === AccountStatus.TEMPORARY) {
					bookmarkParas = summary.map((para: string) => {
						const bookmarkPara = new BookmarkParaTemp()
						bookmarkPara.bookmarkId = data.bookmark.id,
						bookmarkPara.content = para;
						return bookmarkPara;
					})
					await transactionalEntityManager.save(bookmarkParas);
				}else {
						bookmarkParas = summary.map((para: string) => {
						const bookmarkPara = new BookmarkPara()
						bookmarkPara.bookmarkId = data.bookmark.id,
						bookmarkPara.content = para;
						return bookmarkPara;
					})
					await transactionalEntityManager.save(bookmarkParas);
				}
				// create documents on aquiladb
				const documents = summary.map((para: string, index: number) => {
					return {
						metadata: {
							para,
							bookmark_para_id: bookmarkParas[index].id
						},
						code: vectorArray[index]
					}
				})
				console.log("COLLECTIONS*******: ", collection, documents);
				const output = await aquilaClient.getDbServer().createDocuments((collection as Collection | CollectionTemp).aquilaDbName, documents)
				console.log("From worker: output", output, documents);

				// update status as PROCESSED
				let bookmark: Bookmark | BookmarkTemp | null;
				if(data.accountStatus === AccountStatus.TEMPORARY) {
					bookmark = await BookmarkTemp.findOne({ where: { id: data.bookmark.id}})
					if(!bookmark) {
						throw new Error("Invalid bookmark id");
					}
					bookmark.status = BookmarkTempStatus.PROCESSED;
				}else {
					bookmark = await Bookmark.findOne({ where: { id: data.bookmark.id}})
					if(!bookmark) {
						throw new Error("Invalid bookmark id");
					}
					bookmark.status = BookmarkStatus.PROCESSED;
				}

				await transactionalEntityManager.save(bookmark);
			});

			const searchData = tf.randomUniform([1, 10]).arraySync() as number[][];
			const result = await aquilaClient.getDbServer().searchKDocuments(collection.aquilaDbName, searchData, 100);
			console.log("Result from server", JSON.stringify(result));
	}
}