import type { ApiResponse } from "./api.types";

export interface Settings {
    name: string;
    logo: string;
}

export type SettingsResponse = ApiResponse<Settings>;