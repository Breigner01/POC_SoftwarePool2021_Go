import { DeleteResult, Repository } from 'typeorm';
import { Note } from './note.entity';
import { NoteDto } from './note.dto';
export declare class NoteRepository extends Repository<Note> {
    createNote(note: NoteDto): Promise<Note>;
    getNotes(): Promise<Note[]>;
    getNote(title: string): Promise<Note | undefined>;
    updateNote(noteDto: Partial<NoteDto>): Promise<Note>;
    deleteNote(title: string): Promise<DeleteResult>;
}
