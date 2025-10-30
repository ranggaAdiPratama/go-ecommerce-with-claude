import { Store } from 'lucide-react';
import { useSettings } from '../../hooks/useSettings';

export const Header = () => {
    const { settings, loading } = useSettings()

    const logoEmptyState = (<Store className="w-20 h-20 text-indigo-600" />)

    return (
        <header className="bg-white shadow-sm">
            <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
                <div className="text-center">
                    <div className="flex items-center justify-center gap-3 mb-2">
                        {
                            loading
                                ? logoEmptyState
                                : settings?.logo
                                    ? (
                                        <img
                                            src={settings.logo}
                                            alt={settings.name}
                                            className="h-20 w-auto object-contain"
                                        />
                                    )
                                    : logoEmptyState
                        }
                    </div>
                    <h1 className="text-2xl font-semibold text-gray-800 mb-2">{loading ? 'Loading...' : settings?.name || 'Shop'}</h1>
                    <p className="text-sm text-gray-600">Discover our latest partner stores</p>
                </div>
            </div>
        </header>
    );
};