import { Store } from 'lucide-react';

interface EmptyStateProps {
    message?: string;
}

export const EmptyState = ({ message = 'No shops available at the moment' }: EmptyStateProps) => {
    return (
        <div className="text-center py-12">
            <Store className="w-16 h-16 text-gray-400 mx-auto mb-4" />
            <p className="text-gray-600 text-lg">{message}</p>
        </div>
    );
};