import { NoteRepository } from './note.repository';
import { Note } from './note.entity';
import { NoteDto } from './note.dto';
export declare class NoteService {
    private readonly noteRepository;
    constructor(noteRepository: NoteRepository);
    getNotes(): Promise<Note[]>;
    getNote(title: string): Promise<Note>;
    createNote(noteDto: NoteDto): Promise<Note>;
    updateNote(noteDto: Partial<NoteDto>): Promise<Note>;
    deleteNote(title: string): Promise<string>;
}
