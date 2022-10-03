import { NextFunction, Request, Response } from "express";
import { param } from "express-validator";
import { ExpressMiddlewareInterface } from "routing-controllers";
import Container, { Service } from "typedi";

import { JwtPayloadData } from "../../../helper/decorators/jwtPayloadData";
import { CollectionService } from "../../../service/CollectionService";
import { CollectionSubscriptionService } from "../../../service/CollectionSubscriptionService";
import { AccountStatus, JwtPayload } from "../../../service/dto/AuthServiceDto";
import { validate } from "../../../utils/validate";

@Service()
export class SubscribeCollectionValidator implements ExpressMiddlewareInterface {
	public constructor(private collectionService: CollectionService, private collectionSubService: CollectionSubscriptionService, @JwtPayloadData() private jwtPayloadData: JwtPayload) {}


	public async use(req: Request, res: Response, next: NextFunction) {
		const validators = [
			param('collectionId')
				.trim().not().isEmpty()
				.withMessage("Collection Id is required")
				.bail()
				.isUUID()
				.withMessage("Invalid Collection Id")
				.bail()
				.custom(async (value) => {
					const jwtPayloadData = req.jwtTokenPayload;
					if(jwtPayloadData) {
						const collectionService = Container.get(CollectionService);
						const collection = await collectionService.getCollectionById(value, AccountStatus.PERMANENT);
						if(collection === undefined) {
							throw new Error("Collection doesn't exist");	
						}
					}
					return true;
				})
				.bail()
				.custom(async (value) => {
					const jwtPayloadData = req.jwtTokenPayload;
					if(jwtPayloadData) {
						const collectionSubService = Container.get(CollectionSubscriptionService);
						const collection = await collectionSubService.isCollectionSubscribedByCustomer(value, jwtPayloadData.customerId, jwtPayloadData.accountStatus);
						if(collection !== null) {
							throw new Error("Collection is subscribed already");	
						}
					}
					return true;
				})
		];

		await validate(validators, req);
		next();
	}
} 