"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.NoteRepository = void 0;
const typeorm_1 = require("typeorm");
const note_entity_1 = require("./note.entity");
let NoteRepository = class NoteRepository extends typeorm_1.Repository {
    async createNote(note) {
        return this.save(note);
    }
    async getNotes() {
        return this.find();
    }
    async getNote(title) {
        return this.findOne({ title });
    }
    async updateNote(noteDto) {
        await this.update({ title: noteDto.title }, noteDto);
        return this.getNote(noteDto.title);
    }
    async deleteNote(title) {
        return this.delete({ title });
    }
};
NoteRepository = __decorate([
    typeorm_1.EntityRepository(note_entity_1.Note)
], NoteRepository);
exports.NoteRepository = NoteRepository;
//# sourceMappingURL=note.repository.js.map