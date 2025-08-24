export interface User{
    id: number;
    uuid: string;
    firstname: string;
    lastname: string;
    username: string;
    password: string;
    email: string;
    phone: string;
    active: boolean;
}

export type UserAdd = Omit<User, 'id' | 'active' | 'uuid'>;

export type UserEdit = Omit<User, 'id'>

