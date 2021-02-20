"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
const common_1 = require("@nestjs/common");
const note_module_1 = require("./notes/note.module");
const appLogger_1 = require("./appLogger");
const appCors_1 = require("./appCors");
let AppModule = class AppModule {
    configure(consumer) {
        consumer.apply(appLogger_1.default).forRoutes('*');
        consumer.apply(appCors_1.CorsMiddleware).forRoutes('*');
    }
};
AppModule = __decorate([
    common_1.Module({
        imports: [note_module_1.default],
    })
], AppModule);
exports.default = AppModule;
//# sourceMappingURL=appModule.js.map