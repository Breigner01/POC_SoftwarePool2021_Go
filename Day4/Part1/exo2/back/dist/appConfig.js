"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.apiConfig = exports.dbConfig = void 0;
const env_var_1 = require("env-var");
const env = (name, required = true) => env_var_1.get(name).required(required);
exports.dbConfig = {
    username: env('DB_USER').asString(),
    password: env('DB_PASS').asString(),
    host: env('DB_HOST').asString(),
    port: env('DB_PORT').asPortNumber(),
    database: env('DB_NAME').asString(),
};
exports.apiConfig = {
    port: env('API_PORT').asPortNumber(),
    host: env('API_HOST').asString(),
};
//# sourceMappingURL=appConfig.js.map