import apiConfig from "../config/api.config";
import type { Settings, SettingsResponse } from "../types/setting.types";
import { apiService } from "./api.service";

class SettingService {
    async getSettings(): Promise<Settings> {
        const response = await apiService.get<SettingsResponse>(apiConfig.endpoints.settings);

        if (response.meta_data.code !== 200) {
            throw new Error(response.meta_data.message || 'Failed to fetch settings');
        }

        return response.data;
    }
}

export const settingService = new SettingService()