import type { Shop, ShopsResponse } from "../types/shop.types";
import { apiService } from './api.service';
import { apiConfig } from "../config/api.config";

class ShopService {
    async getShopDisplay(): Promise<Shop[]> {
        const response = await apiService.get<ShopsResponse>(apiConfig.endpoints.shops + '?limit=4')

        if (response.meta_data.code !== 200) {
            throw new Error(response.meta_data.message || 'Failed to fetch shops');
        }

        return response.data
    }
}

export const shopService = new ShopService()