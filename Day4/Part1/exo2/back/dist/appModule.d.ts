import { MiddlewareConsumer, NestModule } from '@nestjs/common';
export default class AppModule implements NestModule {
    configure(consumer: MiddlewareConsumer): void;
}
