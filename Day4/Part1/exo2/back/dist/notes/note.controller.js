"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
var __param = (this && this.__param) || function (paramIndex, decorator) {
    return function (target, key) { decorator(target, key, paramIndex); }
};
Object.defineProperty(exports, "__esModule", { value: true });
const common_1 = require("@nestjs/common");
const note_dto_1 = require("./note.dto");
const note_service_1 = require("./note.service");
let NoteController = class NoteController {
    constructor(noteService) {
        this.noteService = noteService;
    }
    async getNotes() {
        return this.noteService.getNotes();
    }
    async getNote(title) {
        return this.noteService.getNote(title);
    }
    async postNote(noteDto) {
        return this.noteService.createNote(noteDto);
    }
    async putNote(noteDto) {
        return this.noteService.updateNote(noteDto);
    }
    async deleteNote(title) {
        return this.noteService.deleteNote(title);
    }
};
__decorate([
    common_1.Get(),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", []),
    __metadata("design:returntype", Promise)
], NoteController.prototype, "getNotes", null);
__decorate([
    common_1.Get(':title'),
    __param(0, common_1.Param('title')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String]),
    __metadata("design:returntype", Promise)
], NoteController.prototype, "getNote", null);
__decorate([
    common_1.Post(),
    __param(0, common_1.Body()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [note_dto_1.NoteDto]),
    __metadata("design:returntype", Promise)
], NoteController.prototype, "postNote", null);
__decorate([
    common_1.Put(),
    __param(0, common_1.Body()),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [Object]),
    __metadata("design:returntype", Promise)
], NoteController.prototype, "putNote", null);
__decorate([
    common_1.Delete(':title'),
    __param(0, common_1.Param('title')),
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [String]),
    __metadata("design:returntype", Promise)
], NoteController.prototype, "deleteNote", null);
NoteController = __decorate([
    common_1.Controller('notes'),
    __metadata("design:paramtypes", [note_service_1.NoteService])
], NoteController);
exports.default = NoteController;
//# sourceMappingURL=note.controller.js.map