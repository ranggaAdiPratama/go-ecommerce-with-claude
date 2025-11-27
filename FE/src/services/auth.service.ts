/* eslint-disable @typescript-eslint/no-explicit-any */
import apiConfig from "../config/api.config";
import type { ApiResponse } from "../types/api.types";
import type { LoginData, LoginRequest, RegisterApiResponse, RegisterRequest } from "../types/auth.types";
import { apiService } from "./api.service";

class AuthService {
    async login(data: LoginRequest): Promise<string | LoginData> {
        const response = await apiService.post<ApiResponse<LoginData>>(
            apiConfig.endpoints.login,
            data
        );

        if (response.meta_data.code !== 200) {
            throw new Error(response.meta_data.message || 'Login failed');
        }

        localStorage.setItem('token', response.data.token);
        localStorage.setItem('refresh_token', response.data.refresh_token);
        localStorage.setItem('user', JSON.stringify(response.data.user));

        return response.data;
    }

    async logout(): Promise<void> {
        const token = localStorage.getItem('token')

        const response = await apiService.post<ApiResponse<null>>(
            apiConfig.endpoints.logout,
            null,
            {
                headers: {
                    Authorization: `Bearer ${token}`,
                    "Content-Type": "application/json",
                }
            }
        );

        if (response.meta_data.code !== 200) {
            console.log(response.meta_data.message || 'Login failed');
        }

        localStorage.removeItem('token');
        localStorage.removeItem('refresh_token');
        localStorage.removeItem('user');
    }

    getToken(): string | null {
        return localStorage.getItem('token');
    }

    getUser(): any | null {
        const user = localStorage.getItem('user');

        return user ? JSON.parse(user) : null;
    }

    isAuthenticated(): boolean {
        return !!this.getToken();
    }

    async register(data: RegisterRequest): Promise<string> {
        const response = await apiService.post<RegisterApiResponse>(
            apiConfig.endpoints.register,
            data
        )

        if (response.meta_data.code !== 201) {
            throw new Error(response.meta_data.message);
        }

        return response.meta_data.message;
    }
}

export const authService = new AuthService()