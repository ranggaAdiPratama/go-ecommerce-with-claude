import apiConfig from '../config/api.config';

class ApiService {
    private baseUrl: string;

    constructor() {
        this.baseUrl = apiConfig.baseUrl;
    }

    async fetch<T>(endpoint: string, options?: RequestInit): Promise<T> {
        const url = `${this.baseUrl}${endpoint}`;

        const response = await fetch(url, {
            ...options,
            headers: {
                'Content-Type': 'application/json',
                ...options?.headers,
            },
        });

        return response.json();
    }

    async get<T>(endpoint: string): Promise<T> {
        return this.fetch<T>(endpoint, { method: 'GET' });
    }

    async post<T>(endpoint: string, data: unknown, options?: RequestInit): Promise<T> {
        return this.fetch<T>(endpoint, {
            method: 'POST',
            body: JSON.stringify(data),
            ...options
        });
    }

    async put<T>(endpoint: string, data: unknown): Promise<T> {
        return this.fetch<T>(endpoint, {
            method: 'PUT',
            body: JSON.stringify(data),
        });
    }

    async delete<T>(endpoint: string): Promise<T> {
        return this.fetch<T>(endpoint, { method: 'DELETE' });
    }
}

export const apiService = new ApiService();