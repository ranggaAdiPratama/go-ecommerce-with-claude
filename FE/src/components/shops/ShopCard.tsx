import { RANK_COLORS, RANK_TEXT_COLORS } from '../../constants/ranks';
import type { Shop } from '../../types/shop.types';

interface ShopCardProps {
    shop: Shop;
    onVisit?: (shop: Shop) => void;
}

export const ShopCard = ({ shop, onVisit }: ShopCardProps) => {
    const handleImageError = (e: React.SyntheticEvent<HTMLImageElement>) => {
        e.currentTarget.src = 'https://placehold.co/400x300?text=No+Image';
    };

    return (
        <div
            className="bg-white rounded-xl shadow-md overflow-hidden hover:shadow-xl transition-shadow duration-300 cursor-pointer"
            style={{ minWidth: '400px', width: '400px' }}
            onClick={() => onVisit?.(shop)}
        >
            {/* Image Container */}
            <div
                className="relative bg-gray-100 flex items-center justify-center"
                style={{ height: '300px', width: '400px' }}
            >
                <img
                    src={shop.logo}
                    alt={shop.name}
                    style={{ width: '100%', height: '100%', objectFit: 'cover' }}
                    className="transition-transform duration-300 hover:scale-105"
                    onError={handleImageError}
                />
                <div className="absolute top-3 right-3">
                    <span className={`${RANK_COLORS[shop.rank]} text-white text-xs font-bold px-3 py-1 rounded-full uppercase shadow-lg`}>
                        {shop.rank}
                    </span>
                </div>
            </div>

            {/* Content */}
            <div className="p-6">
                <h3 className="text-lg font-semibold text-gray-900 mb-2">
                    {shop.name}
                </h3>
                <div className="flex items-center gap-2 mb-3">
                    <span className={`text-sm font-medium ${RANK_TEXT_COLORS[shop.rank]} capitalize`}>
                        {shop.rank} Seller
                    </span>
                </div>
                <button className="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-medium py-2 px-4 rounded-lg transition-colors">
                    Visit Shop
                </button>
            </div>
        </div>
    );
};