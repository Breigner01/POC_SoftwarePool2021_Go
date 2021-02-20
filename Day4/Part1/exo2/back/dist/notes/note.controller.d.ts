import { NoteDto } from './note.dto';
import { Note } from './note.entity';
import { NoteService } from './note.service';
export default class NoteController {
    private readonly noteService;
    constructor(noteService: NoteService);
    getNotes(): Promise<Note[]>;
    getNote(title: string): Promise<Note>;
    postNote(noteDto: NoteDto): Promise<Note>;
    putNote(noteDto: Partial<NoteDto>): Promise<Note>;
    deleteNote(title: string): Promise<string>;
}
