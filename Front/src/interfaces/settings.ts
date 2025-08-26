export interface Settings {
    id: number;
    uuid: string;
    name: string;
    address: string;
    port: string;
    login: string;
    password: string;
    active: boolean;
}

export type MailServiceCreate = Omit<Settings, 'id' | 'uuid' | 'active'>
export type MailServiceUpdate = Omit<Settings, 'id'>