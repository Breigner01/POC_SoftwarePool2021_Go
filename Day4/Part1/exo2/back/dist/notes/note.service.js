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
exports.NoteService = void 0;
const common_1 = require("@nestjs/common");
const typeorm_1 = require("@nestjs/typeorm");
const note_repository_1 = require("./note.repository");
let NoteService = class NoteService {
    constructor(noteRepository) {
        this.noteRepository = noteRepository;
    }
    async getNotes() {
        return this.noteRepository.getNotes();
    }
    async getNote(title) {
        const note = await this.noteRepository.getNote(title);
        if (!note) {
            throw new common_1.NotFoundException('Note not found');
        }
        return note;
    }
    async createNote(noteDto) {
        const note = await this.noteRepository.getNote(noteDto.title);
        if (note) {
            throw new common_1.NotFoundException('Note already exist');
        }
        return this.noteRepository.createNote(noteDto);
    }
    async updateNote(noteDto) {
        const note = await this.noteRepository.getNote(noteDto.title);
        if (!note) {
            throw new common_1.NotFoundException('Note already exist');
        }
        return this.noteRepository.updateNote(noteDto);
    }
    async deleteNote(title) {
        const note = await this.noteRepository.getNote(title);
        if (!note) {
            throw new common_1.NotFoundException(`Note '${title}' doesn't exist`);
        }
        await this.noteRepository.deleteNote(title);
        return `Note '${title}' has been delete`;
    }
};
NoteService = __decorate([
    common_1.Injectable(),
    __param(0, typeorm_1.InjectRepository(note_repository_1.NoteRepository)),
    __metadata("design:paramtypes", [note_repository_1.NoteRepository])
], NoteService);
exports.NoteService = NoteService;
//# sourceMappingURL=note.service.js.map