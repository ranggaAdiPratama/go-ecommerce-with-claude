import type { ApiResponse } from "./api.types";

export interface Category {
    id: string;
    name: string;
    icon: string;
    slug: string;
}

export type CategoriesResponse = ApiResponse<Category[]>;