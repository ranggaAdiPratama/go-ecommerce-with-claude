/* eslint-disable @typescript-eslint/no-explicit-any */
import { LayoutDashboard, LogOut, Search, ShoppingCart, Store, User, X } from "lucide-react";
import { useSettings } from "../../hooks/useSettings"
import { Link } from "react-router-dom";
import { useState } from "react";
import { authService } from "../../services/auth.service";

const cartItems: any[] = [
    // Mock cart items - uncomment to test
    // { id: '1', name: 'Product 1', price: 25000, quantity: 2, image: 'https://placehold.co/80x80' },
    // { id: '2', name: 'Product 2', price: 50000, quantity: 1, image: 'https://placehold.co/80x80' },
    // { id: '3', name: 'Product 3', price: 15000, quantity: 3, image: 'https://placehold.co/80x80' },
];

export const Navbar = () => {
    const { settings, loading } = useSettings()

    const [isLoggedIn, setIsLoggedIn] = useState(authService.isAuthenticated())
    const [userName, setUserName] = useState(authService.getUser()?.name)
    const [showCartDropdown, setShowCartDropdown] = useState(false);
    const [showAccountDropdown, setShowAccountDropdown] = useState(false);

    const cartCount = cartItems.length > 0 ? cartItems.reduce((sum, item) => sum + item.quantity, 0) : 0;
    const cartTotal = cartItems.reduce((sum, item) => sum + (item.price * item.quantity), 0);

    const handleLogout = async () => {
        await authService.logout()

        setIsLoggedIn(authService.isAuthenticated())
        setUserName('')

        setShowAccountDropdown(false);
    };

    const logoEmptyState = (<Store className="w-20 h-20 text-indigo-600" />)

    return (
        <nav className="bg-white shadow-md sticky top-0 z-50">
            <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <div className="flex items-center justify-between h-16">
                    {/* Logo */}
                    <Link to="/" className="flex items-center gap-3 hover:opacity-80 transition-opacity">
                        {loading ? (
                            <div className="w-10 h-10 bg-gray-200 rounded-full animate-pulse"></div>
                        ) : settings?.logo ? (
                            <img
                                src={settings.logo}
                                alt={settings.name}
                                className="h-10 w-auto object-contain"
                            />
                        ) : (
                            <div className="w-10 h-10 bg-indigo-100 rounded-full flex items-center justify-center">
                                {logoEmptyState}
                            </div>
                        )}
                        <span className="text-xl font-bold text-gray-900 hidden sm:block">
                            {loading ? '...' : settings?.name || 'Warunk Aink'}
                        </span>
                    </Link>

                    {/* Search Bar */}
                    <div className="flex-1 max-w-2xl mx-8 hidden md:block">
                        <div className="relative">
                            <input
                                type="text"
                                placeholder="Search products..."
                                className="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-600 focus:border-transparent outline-none"
                            />
                            <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
                        </div>
                    </div>

                    {/* Right Icons */}
                    <div className="flex items-center gap-4">
                        {/* Mobile Search Icon */}
                        <button className="md:hidden p-2 hover:bg-gray-100 rounded-lg transition-colors">
                            <Search className="w-6 h-6 text-gray-700" />
                        </button>

                        {/* Cart Dropdown */}
                        <div className="relative">
                            <button
                                onClick={() => setShowCartDropdown(!showCartDropdown)}
                                className="relative p-2 hover:bg-gray-100 rounded-lg transition-colors"
                            >
                                <ShoppingCart className="w-6 h-6 text-gray-700" />
                                {cartCount > 0 && (
                                    <span className="absolute -top-1 -right-1 bg-red-500 text-white text-xs font-bold rounded-full w-5 h-5 flex items-center justify-center">
                                        {cartCount}
                                    </span>
                                )}
                            </button>

                            {/* Cart Dropdown Menu */}
                            {showCartDropdown && (
                                <>
                                    <div
                                        className="fixed inset-0 z-10"
                                        onClick={() => setShowCartDropdown(false)}
                                    ></div>
                                    <div className="absolute right-0 mt-2 w-80 bg-white rounded-lg shadow-xl border border-gray-200 z-20">
                                        <div className="p-4 border-b border-gray-200">
                                            <div className="flex items-center justify-between">
                                                <h3 className="text-lg font-semibold text-gray-900">Shopping Cart</h3>
                                                <button
                                                    onClick={() => setShowCartDropdown(false)}
                                                    className="text-gray-400 hover:text-gray-600"
                                                >
                                                    <X className="w-5 h-5" />
                                                </button>
                                            </div>
                                        </div>

                                        {cartItems.length === 0 ? (
                                            <div className="p-8 text-center">
                                                <ShoppingCart className="w-12 h-12 text-gray-300 mx-auto mb-3" />
                                                <p className="text-gray-600">Your cart is empty</p>
                                            </div>
                                        ) : (
                                            <>
                                                <div className="max-h-64 overflow-y-auto">
                                                    {cartItems.slice(0, 3).map((item) => (
                                                        <div key={item.id} className="p-4 border-b border-gray-100 hover:bg-gray-50">
                                                            <div className="flex gap-3">
                                                                <img
                                                                    src={item.image}
                                                                    alt={item.name}
                                                                    className="w-16 h-16 object-cover rounded"
                                                                />
                                                                <div className="flex-1">
                                                                    <h4 className="text-sm font-medium text-gray-900">{item.name}</h4>
                                                                    <p className="text-xs text-gray-500 mt-1">Qty: {item.quantity}</p>
                                                                    <p className="text-sm font-semibold text-indigo-600 mt-1">
                                                                        Rp {item.price.toLocaleString('id-ID')}
                                                                    </p>
                                                                </div>
                                                            </div>
                                                        </div>
                                                    ))}
                                                </div>
                                                <div className="p-4 border-t border-gray-200">
                                                    <div className="flex justify-between mb-3">
                                                        <span className="font-semibold text-gray-900">Total:</span>
                                                        <span className="font-bold text-indigo-600">
                                                            Rp {cartTotal.toLocaleString('id-ID')}
                                                        </span>
                                                    </div>
                                                    <Link
                                                        to="/cart"
                                                        onClick={() => setShowCartDropdown(false)}
                                                        className="block w-full bg-indigo-600 hover:bg-indigo-700 text-white text-center font-medium py-2 rounded-lg transition-colors"
                                                    >
                                                        View Cart
                                                    </Link>
                                                </div>
                                            </>
                                        )}
                                    </div>
                                </>
                            )}
                        </div>

                        {/* Account Dropdown */}
                        <div className="relative">
                            {isLoggedIn ? (
                                <>
                                    <button
                                        onClick={() => setShowAccountDropdown(!showAccountDropdown)}
                                        className="flex items-center gap-2 p-2 hover:bg-gray-100 rounded-lg transition-colors"
                                    >
                                        <User className="w-6 h-6 text-gray-700" />
                                        <span className="hidden lg:block text-sm font-medium text-gray-700">
                                            {userName}
                                        </span>
                                    </button>

                                    {/* Account Dropdown Menu */}
                                    {showAccountDropdown && (
                                        <>
                                            <div
                                                className="fixed inset-0 z-10"
                                                onClick={() => setShowAccountDropdown(false)}
                                            ></div>
                                            <div className="absolute right-0 mt-2 w-56 bg-white rounded-lg shadow-xl border border-gray-200 z-20">
                                                <div className="p-3 border-b border-gray-200">
                                                    <p className="text-sm font-semibold text-gray-900">{userName}</p>
                                                    <p className="text-xs text-gray-500">user@example.com</p>
                                                </div>
                                                <div className="py-2">
                                                    <Link
                                                        to="/dashboard"
                                                        onClick={() => setShowAccountDropdown(false)}
                                                        className="flex items-center gap-3 px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 transition-colors"
                                                    >
                                                        <LayoutDashboard className="w-4 h-4" />
                                                        Dashboard
                                                    </Link>
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
                                </>
                            ) : (
                                <Link
                                    to="/auth/login"
                                    className="flex items-center gap-2 px-4 py-2 bg-indigo-600 hover:bg-indigo-700 text-white rounded-lg transition-colors font-medium text-sm"
                                >
                                    <User className="w-5 h-5" />
                                    <span className="hidden lg:block">Login</span>
                                </Link>
                            )}
                        </div>
                    </div>
                </div>
            </div>
        </nav>
    );
}