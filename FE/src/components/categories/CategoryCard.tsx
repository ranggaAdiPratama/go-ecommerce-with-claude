/* eslint-disable @typescript-eslint/no-explicit-any */
import type { Category } from "../../types/category.types";
import * as LucideIcons from 'lucide-react';

interface CategoryCardProps {
    category: Category;
    onClick?: (category: Category) => void;
}

export const CategoryCard = ({ category, onClick }: CategoryCardProps) => {
    const IconComponent = (LucideIcons as any)[category.icon] || LucideIcons.Tag;

    return (
        <div
            onClick={() => onClick?.(category)}
            className="cursor-pointer flex flex-col items-center text-center group"
            style={{ minWidth: '100px', width: '100px' }}
        >
            <div className="w-16 h-16 bg-white rounded-full flex items-center justify-center mb-3 group-hover:scale-110 transition-transform duration-200 shadow-md">
                <IconComponent className="w-7 h-7 text-indigo-600" />
            </div>
            <h3 className="font-medium text-gray-800 text-xs">{category.name}</h3>
        </div>
    )
}