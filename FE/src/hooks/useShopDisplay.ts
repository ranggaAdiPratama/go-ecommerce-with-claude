import { useEffect, useState } from "react";
import { shopService } from "../services/shop.service";
import type { Shop } from "../types/shop.types";

export const useShopDisplay = () => {
    const [shops, setShops] = useState<Shop[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    const fetchShopDsiplay = async () => {
        try {
            setLoading(true)
            setError(null);

            const data = await shopService.getShopDisplay()

            setShops(data)
        } catch (err) {
            setError(err instanceof Error ? err.message : 'Failed to fetch shops');
            console.error('Error fetching shops:', err);
        } finally {
            setLoading(false);
        }
    }

    useEffect(() => {
        fetchShopDsiplay()
    }, [])

    return {
        shops,
        loading,
        error,
        refetch: fetchShopDsiplay
    }
}