import { Store } from 'lucide-react';

export const EmptyState = () => {
    return (
        <div className="text-center py-12">
            <Store className="w-16 h-16 text-gray-400 mx-auto mb-4" />
            <p className="text-gray-600 text-lg">No shops available at the moment</p>
        </div>
    );
};