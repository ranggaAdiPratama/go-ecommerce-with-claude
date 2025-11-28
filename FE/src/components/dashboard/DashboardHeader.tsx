/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState, useEffect } from 'react';
import { User, LogOut, ChevronDown, Store } from 'lucide-react';
import { useNavigate } from 'react-router-dom';
import { useSettings } from '../../hooks/useSettings';
import { authService } from '../../services/auth.service';

export const DashboardHeader = () => {
    const { settings, loading } = useSettings();
    const navigate = useNavigate();
    const [showDropdown, setShowDropdown] = useState(false);
    const [user, setUser] = useState<any>(null);

    useEffect(() => {
        const userData = authService.getUser();
        setUser(userData);
    }, []);

    const handleLogout = async () => {
        setShowDropdown(false);

        await authService.logout();

        setUser(null);

        navigate('/');
    };

    const logoEmptyState = (<Store className="w-20 h-20 text-indigo-600" />)

    return (
        <header className="bg-white shadow-sm border-b border-gray-200">
            <div className="px-6 py-4">
                <div className="flex items-center justify-between">
                    {/* Logo + Toggle Button */}
                    <div className="flex items-center gap-4">
                        {/* Logo */}
                        <div className="flex items-center gap-3">
                            {loading ? (
                                <div className="w-10 h-10 bg-gray-200 rounded-full animate-pulse"></div>
                            ) : settings?.logo ? (
                                <img
                                    src={settings.logo}
                                    alt={settings.name}
                                    className="h-10 w-auto object-contain"
                                />
                            ) : (
                                logoEmptyState
                            )}
                            <span className="text-xl font-bold text-gray-900">
                                {loading ? '...' : settings?.name || 'Warunk Aink'}
                            </span>
                        </div>
                    </div>

                    {/* User Dropdown */}
                    <div className="relative">
                        <button
                            onClick={() => setShowDropdown(!showDropdown)}
                            className="flex items-center gap-3 px-4 py-2 hover:bg-gray-100 rounded-lg transition-colors"
                        >
                            <div className="w-9 h-9 bg-indigo-100 rounded-full flex items-center justify-center">
                                <User className="w-5 h-5 text-indigo-600" />
                            </div>
                            <div className="text-left">
                                <p className="text-sm font-semibold text-gray-900">{user?.name || 'User'}</p>
                                <p className="text-xs text-gray-500">{user?.role || 'user'}</p>
                            </div>
                            <ChevronDown className="w-4 h-4 text-gray-500" />
                        </button>

                        {/* Dropdown Menu */}
                        {showDropdown && (
                            <>
                                <div
                                    className="fixed inset-0 z-10"
                                    onClick={() => setShowDropdown(false)}
                                ></div>
                                <div className="absolute right-0 mt-2 w-56 bg-white rounded-lg shadow-xl border border-gray-200 z-20">
                                    <div className="p-3 border-b border-gray-200">
                                        <p className="text-sm font-semibold text-gray-900">{user?.name}</p>
                                        <p className="text-xs text-gray-500">{user?.email}</p>
                                    </div>
                                    <div className="py-2">
                                        <button
                                            onClick={() => {
                                                setShowDropdown(false);
                                                // TODO: Navigate to profile page
                                                console.log('Profile clicked');
                                            }}
                                            className="w-full flex items-center gap-3 px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 transition-colors"
                                        >
                                            <User className="w-4 h-4" />
                                            Profile
                                        </button>
                                        <button
                                            onClick={handleLogout}
                                            className="w-full flex items-center gap-3 px-4 py-2 text-sm text-red-600 hover:bg-red-50 transition-colors"
                                        >
                                            <LogOut className="w-4 h-4" />
                                            Logout
                                        </button>
                                    </div>
                                </div>
                            </>
                        )}
                    </div>
                </div>
            </div>
        </header>
    );
};