import type { Shop } from '../types/shop.types';

interface ShopCardProps {
    shop: Shop;
    onVisit?: (shop: Shop) => void;
}

const rankColors = {
    bronze: 'bg-amber-700',
    silver: 'bg-gray-400',
    gold: 'bg-yellow-500',
    platinum: 'bg-cyan-400',
};

const rankTextColors = {
    bronze: 'text-amber-700',
    silver: 'text-gray-600',
    gold: 'text-yellow-600',
    platinum: 'text-cyan-600',
};

export const ShopCard = ({ shop, onVisit }: ShopCardProps) => {
    const handleImageError = (e: React.SyntheticEvent<HTMLImageElement>) => {
        e.currentTarget.src = 'https://placehold.co/400x300?text=No+Image';
    };

    return (
        <div className="bg-white rounded-xl shadow-md overflow-hidden hover:shadow-xl transition-shadow duration-300 cursor-pointer w-80 flex-shrink-0">
            <div className="relative bg-gray-100 flex items-center justify-center p-8" style={{ width: '400px', height: '300px' }}>
                <img
                    src={shop.logo}
                    alt={shop.name}
                    style={{ maxWidth: '350px', maxHeight: '250px', objectFit: 'contain' }}
                    className="group-hover:scale-105 transition-transform duration-300"
                    onError={handleImageError}
                />
                <div className="absolute top-3 right-3">
                    <span className={`${rankColors[shop.rank]} text-white text-xs font-bold px-3 py-1 rounded-full uppercase shadow-lg`}>
                        {shop.rank}
                    </span>
                </div>
            </div>
            <div className="p-2">
                <h3 className="text-lg font-semibold text-gray-900 mb-2">
                    {shop.name}
                </h3>
                <div className="flex items-center gap-2 mb-3">
                    <span className={`text-sm font-medium ${rankTextColors[shop.rank]} capitalize`}>
                        {shop.rank} Seller
                    </span>
                </div>
                <button
                    onClick={() => onVisit?.(shop)}
                    className="w-full mt-2 bg-indigo-600 hover:bg-indigo-700 text-white font-medium py-2 px-4 rounded-lg transition-colors"
                >
                    Visit Shop
                </button>
            </div>
        </div>
    );
};