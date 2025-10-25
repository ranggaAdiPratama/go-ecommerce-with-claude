import { Store } from 'lucide-react';
import { useShopDisplay } from '../hooks/useShopDisplay';
import { ShopCard } from './ShopCard';
import { LoadingSpinner } from './LoadingSpinner';
import { ErrorMessage } from './ErrorMessage';
import { EmptyState } from './EmptyState';
import type { Shop } from '../types/shop.types';

export const ShopList = () => {
    const { shops, loading, error, refetch } = useShopDisplay();

    const handleVisitShop = (shop: Shop) => {
        console.log('Visiting shop:', shop.name, 'ID:', shop.id);
    };

    if (loading) return <LoadingSpinner />;
    if (error) return <ErrorMessage message={error} onRetry={refetch} />;

    return (
        <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
            {/* Header (centered) */}
            <header className="bg-white shadow-sm text-center">
                <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6 text-center">
                    <div className="flex flex-col items-center gap-3 text-center">
                        <Store className="w-8 h-8 text-indigo-600" />
                        <div>
                            <h1 className="text-3xl font-bold text-gray-900">Featured Shops</h1>
                            <p className="text-sm text-gray-600 mt-1">Discover our latest partner stores</p>
                        </div>
                    </div>
                </div>
            </header>

            <main className="py-12">
                <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                    {shops.length === 0 ? (
                        <EmptyState />
                    ) : (

                        <div className="flex flex-nowrap overflow-x-auto gap-6 pb-4 scrollbar-hide">
                            {shops.map((shop) => (
                                <div key={shop.id} className="flex-shrink-0">
                                    <ShopCard shop={shop} onVisit={handleVisitShop} />
                                </div>
                            ))}
                        </div>
                    )}
                </div>
            </main>

            {/* Footer (centered) */}
            <footer className="bg-white mt-12 border-t border-gray-200 text-center">
                <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6 text-center">
                    <p className="text-center text-gray-600">Products coming soon! Stay tuned for our full catalog.</p>
                </div>
            </footer>
        </div>
    );
};