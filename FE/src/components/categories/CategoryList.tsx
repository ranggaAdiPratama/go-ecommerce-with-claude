import { Loader2 } from "lucide-react";
import { useCategories } from "../../hooks/useCategories";
import type { Category } from "../../types/category.types";
import { CategoryCard } from "./CategoryCard";

export const CategoryList = () => {
    const { categories, loading, error } = useCategories();

    const handleCategoryClick = (category: Category) => {
        console.log('Category clicked:', category.name, 'Slug:', category.slug);
    };

    if (loading) {
        return (
            <div className="py-12">
                <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                    <div className="flex items-center justify-center">
                        <Loader2 className="w-8 h-8 text-indigo-600 animate-spin" />
                    </div>
                </div>
            </div>
        );
    }

    if (error) {
        return (
            <div className="py-12">
                <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                    <p className="text-center text-red-600">{error}</p>
                </div>
            </div>
        );
    }

    if (categories.length === 0) {
        return null;
    }

    return (
        <section className="py-8 bg-red-50">
            <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <div className="mb-6 text-center">
                    <h2 className="text-2xl font-bold text-gray-900 mb-2">Browse Categories</h2>
                    <p className="text-gray-600">Explore products by category</p>
                </div>

                <div className="overflow-x-auto pb-4">
                    <div style={{ display: 'flex', gap: '24px' }}>
                        {categories.map((category) => (
                            <CategoryCard
                                key={category.id}
                                category={category}
                                onClick={handleCategoryClick}
                            />
                        ))}
                    </div>
                </div>
            </div>
        </section>
    );
}