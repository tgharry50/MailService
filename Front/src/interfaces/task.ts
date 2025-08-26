export interface Task {
    id: number;
    uuid: string;
    mailUUID: string;
    mailServiceUUID: string;
    weekDays: string;
    everyDay: number;
    everyWeek: number;
    everyYear: number;
    repeat: boolean;
    active: boolean;
}

export type TaskCreate = Omit<Task, 'id' | 'active' | 'uuid'>
export type TaskUpdate = Omit<Task, 'id' | 'mailUUID' | 'mailServiceUUID'>