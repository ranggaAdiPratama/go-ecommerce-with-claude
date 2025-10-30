import type { ApiResponse } from "./api.types";

export interface Shop {
    id: string;
    name: string;
    logo: string;
    rank: 'bronze' | 'silver' | 'gold' | 'platinum';
}



export type ShopsResponse = ApiResponse<Shop[]>;