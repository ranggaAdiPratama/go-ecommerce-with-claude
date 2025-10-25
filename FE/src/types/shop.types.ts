export interface Shop {
    id: string;
    name: string;
    logo: string;
    rank: 'bronze' | 'silver' | 'gold' | 'platinum';
}

export interface MetaData {
    code: number;
    message: string;
}

export interface ShopsResponse {
    meta_data: MetaData;
    data: Shop[];
}