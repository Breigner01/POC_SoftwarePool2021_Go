import { NestMiddleware } from '@nestjs/common';
import { Request, Response, NextFunction } from 'express';
export default class AppLogger implements NestMiddleware {
    private logger;
    use(request: Request, response: Response, next: NextFunction): void;
}
