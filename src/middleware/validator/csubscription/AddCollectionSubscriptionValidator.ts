import { NextFunction, Request, Response } from "express";
import { param } from "express-validator";
import { ExpressMiddlewareInterface } from "routing-controllers";
import Container, { Service } from "typedi";

import { JwtPayloadData } from "../../../helper/decorators/jwtPayloadData";
import { CollectionService } from "../../../service/CollectionService";
import { AccountStatus, JwtPayload } from "../../../service/dto/AuthServiceDto";
import { validate } from "../../../utils/validate";

@Service()
export class AddCollectionSubscriptionValidator implements ExpressMiddlewareInterface {
	public constructor(private collectionService: CollectionService, @JwtPayloadData() private jwtPayloadData: JwtPayload) {}


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
						const collection = await collectionService.getCollectionById(value, AccountStatus.PERMANENT); // jwtPayloadData.accountStatus);
						if(collection === undefined) {
							throw new Error("Collection doesn't exist");	
						}
					}
					return true;
				})
		];

		await validate(validators, req);
		next();
	}
} 