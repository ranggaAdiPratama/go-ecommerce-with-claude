import type { MetaData } from "./api.types";

export interface LoginData {
    user: User;
    token: string;
    refresh_token: string;
}

export interface LoginRequest {
    username: string;
    password: string;
}

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

export interface User {
    id: string;
    name: string;
    username: string;
    email: string;
    role: string;
    created_at: string;
    updated_at: string;
}
