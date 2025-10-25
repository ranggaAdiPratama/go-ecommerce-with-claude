import { useFeaturedShops } from '../../hooks/useShops';
import { ShopCard } from './ShopCard';
import { LoadingSpinner } from '../common/LoadingSpinner';
import { ErrorMessage } from '../common/ErrorMessage';
import { EmptyState } from '../common/EmptyState';
import type { Shop } from '../../types/shop.types';

export const ShopList = () => {
    const { shops, loading, error, refetch } = useFeaturedShops(4);

    const handleVisitShop = (shop: Shop) => {
        console.log('Visiting shop:', shop.name, 'ID:', shop.id);
        // TODO: Add navigation logic here
        // Example: navigate(`/shops/${shop.id}`);
    };

    if (loading) {
        return <LoadingSpinner />;
    }

    if (error) {
        return <ErrorMessage message={error} onRetry={refetch} />;
    }

    return (
        <div className="py-12">
            <div className="px-4 sm:px-6 lg:px-8">
                {shops.length === 0 ? (
                    <EmptyState />
                ) : (
                    <div className="overflow-x-auto pb-4">
                        <div style={{ display: 'flex', gap: '24px' }}>
                            {shops.map((shop) => (
                                <ShopCard key={shop.id} shop={shop} onVisit={handleVisitShop} />
                            ))}
                        </div>
                    </div>
                )}
            </div>
        </div>
    );
};
