import { apiService } from './api.service';
import apiConfig from '../config/api.config';
import type { ShopsResponse, Shop } from '../types/shop.types';

class ShopService {
    async getFeaturedShops(limit: number = 4): Promise<Shop[]> {
        const endpoint = `${apiConfig.endpoints.shops}?limit=${limit}`;
        const response = await apiService.get<ShopsResponse>(endpoint);

        if (response.meta_data.code !== 200) {
            throw new Error(response.meta_data.message || 'Failed to fetch shops');
        }

        return response.data;
    }

    // Future methods
    // async getAllShops(page: number, limit: number): Promise<ShopsResponse> { ... }
    // async getShopById(id: string): Promise<Shop> { ... }
    // async searchShops(query: string): Promise<Shop[]> { ... }
}

export const shopService = new ShopService();