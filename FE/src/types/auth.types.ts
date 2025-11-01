import type { MetaData } from "./api.types";

export interface RegisterRequest {
    name: string;
    email: string;
    username: string;
    password: string;
    role?: 'user' | 'admin';
}

export interface RegisterApiResponse {
    meta_data: MetaData
}