"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const core_1 = require("@nestjs/core");
const common_1 = require("@nestjs/common");
const swagger_1 = require("@nestjs/swagger");
const appConfig_1 = require("./appConfig");
const appModule_1 = require("./appModule");
const { host, port } = appConfig_1.apiConfig;
async function bootstrap() {
    const app = await core_1.NestFactory.create(appModule_1.default, {
        cors: true,
    });
    app.useGlobalPipes(new common_1.ValidationPipe());
    const options = new swagger_1.DocumentBuilder()
        .setTitle('Note API')
        .setDescription('Note API description')
        .setVersion('1.0')
        .addTag('PoC')
        .build();
    const documentation = swagger_1.SwaggerModule.createDocument(app, options);
    swagger_1.SwaggerModule.setup('doc', app, documentation);
    await app.listen(port);
    console.log(`Server listen on http://${host}:${port}`);
}
bootstrap();
//# sourceMappingURL=main.js.map