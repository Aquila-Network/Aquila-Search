import { NextFunction, Request, Response } from 'express';
import { ExpressMiddlewareInterface } from 'routing-controllers';
import { Service } from 'typedi';
import { AuthService } from '../../service/AuthService';

@Service()
export class AuthMiddleware implements ExpressMiddlewareInterface {
	public constructor(authService: AuthService) {

	}

	public use(req: Request, res: Response, next: NextFunction): void {
		const token = req.headers.authorization;

		next();
	}
}