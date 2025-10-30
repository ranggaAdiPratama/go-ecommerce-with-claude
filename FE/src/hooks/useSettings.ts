import { useEffect, useState } from "react";
import type { Settings } from "../types/setting.types";
import { settingService } from "../services/setting.service";

interface UseSettingsResult {
    settings: Settings | null;
    loading: boolean;
    error: string | null;
}

export const useSettings = (): UseSettingsResult => {
    const [settings, setSettings] = useState<Settings | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchSettings = async () => {
            try {
                setLoading(true);
                setError(null);
                const data = await settingService.getSettings();

                setSettings(data);

                if (data.name) document.title = data.name

                if (data.logo) {
                    const favicon = document.querySelector<HTMLLinkElement>("link[rel='icon']")

                    if (favicon) {
                        favicon.href = data.logo
                    } else {
                        const newFavicon = document.createElement('link')

                        newFavicon.rel = 'icon'
                        newFavicon.href = data.logo

                        document.head.appendChild(newFavicon)
                    }
                }
            } catch (err) {
                setError(err instanceof Error ? err.message : 'Failed to fetch settings');
                console.error('Error fetching settings:', err);
            } finally {
                setLoading(false);
            }
        };

        fetchSettings();
    }, []);

    return { settings, loading, error };
}