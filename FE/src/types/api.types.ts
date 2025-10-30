export interface MetaData {
    code: number;
    message: string;
}

export interface ApiResponse<T> {
    meta_data: MetaData;
    data: T;
}