import apiConfig from "../config/api.config";
import type { RegisterApiResponse, RegisterRequest } from "../types/auth.types";
import { apiService } from "./api.service";

class AuthService {
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