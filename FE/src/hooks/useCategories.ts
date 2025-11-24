import { useEffect, useState } from "react";
import type { Category } from "../types/category.types";
import { categoryService } from "../services/category.service";

interface UseCategoriesResult {
    categories: Category[];
    loading: boolean;
    error: string | null;
    refetch: () => Promise<void>;
}

export const useCategories = (): UseCategoriesResult => {
    const [categories, setCategories] = useState<Category[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    const fetchCategories = async () => {
        try {
            setLoading(true);
            setError(null);
            const data = await categoryService.getCategories();
            setCategories(data);
        } catch (err) {
            setError(err instanceof Error ? err.message : 'Failed to fetch categories');
            console.error('Error fetching categories:', err);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchCategories();
    }, []);

    return {
        categories,
        loading,
        error,
        refetch: fetchCategories
    };
}