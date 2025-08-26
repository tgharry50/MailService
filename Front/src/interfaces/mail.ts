// Mail model
export interface Mail {
    id: number;
    uuid: string;
    topic: string;
    content: string;
}
export type MailCreate = Omit<Mail, 'id' | 'uuid'>
export type MailUpdate = Omit<Mail, 'id'> 
// Mail Recipents model
export interface MailRecipents {
    id: number;
    uuid: string;
    taskUUID: string;
    recipent: string;
    active: boolean;
}

export type MailRecipentsCreate = Omit<MailRecipents, 'id' | 'uuid' | 'active'>
export type MailRecipentsUpdate = Omit<MailRecipents, 'id' | 'taskUUID'>