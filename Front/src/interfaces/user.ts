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

export type UserCreate = Omit<User, 'id' | 'active' | 'uuid'>;

export type UserUpdate = Omit<User, 'id'>

