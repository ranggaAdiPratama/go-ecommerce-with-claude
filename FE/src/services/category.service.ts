import apiConfig from "../config/api.config";
import type { CategoriesResponse, Category } from "../types/category.types";
import { apiService } from "./api.service";

class CategoryService {
    async getCategories(): Promise<Category[]> {
        const response = await apiService
            .get<CategoriesResponse>(apiConfig.endpoints.categories);

        if (response.meta_data.code !== 200) {
            throw new Error(response.meta_data.message || 'Failed to fetch categories');
        }

        return response.data;
    }
}

export const categoryService = new CategoryService();
