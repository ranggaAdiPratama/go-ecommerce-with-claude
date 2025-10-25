import { Store } from 'lucide-react';

export const Header = () => {
    return (
        <header className="bg-white shadow-sm">
            <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
                <div className="text-center">
                    <div className="flex items-center justify-center gap-3 mb-2">
                        <Store className="w-10 h-10 text-indigo-600" />
                        <h1 className="text-4xl font-bold text-gray-900">Warunk Aink</h1>
                    </div>
                    <h2 className="text-2xl font-semibold text-gray-800 mb-2">Featured Shops</h2>
                    <p className="text-sm text-gray-600">Discover our latest partner stores</p>
                </div>
            </div>
        </header>
    );
};