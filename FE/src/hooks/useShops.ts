import { useState, useEffect } from 'react';
import { shopService } from '../services/shop.service';
import type { Shop } from '../types/shop.types';

interface UseShopsResult {
    shops: Shop[];
    loading: boolean;
    error: string | null;
    refetch: () => Promise<void>;
}

export const useFeaturedShops = (limit: number = 4): UseShopsResult => {
    const [shops, setShops] = useState<Shop[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    const fetchShops = async (limit: number = 4) => {
        try {
            setLoading(true);
            setError(null);
            const data = await shopService.getFeaturedShops(limit);
            setShops(data);
        } catch (err) {
            setError(err instanceof Error ? err.message : 'Failed to fetch shops');
            console.error('Error fetching shops:', err);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchShops(limit);
    }, [limit]);

    return {
        shops,
        loading,
        error,
        refetch: fetchShops
    };
};
