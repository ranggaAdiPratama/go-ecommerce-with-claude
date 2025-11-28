import { LayoutDashboard, Package, Settings, ShoppingBag, Users } from "lucide-react";
import { useLocation, Link } from 'react-router-dom';

export const DashboardSidebar = () => {
    const location = useLocation();

    const menuItems = [
        { icon: LayoutDashboard, label: 'Dashboard', path: '/dashboard' },
        { icon: ShoppingBag, label: 'Orders', path: '/dashboard/orders' },
        { icon: Package, label: 'Products', path: '/dashboard/products' },
        { icon: Users, label: 'Customers', path: '/dashboard/customers' },
        { icon: Settings, label: 'Settings', path: '/dashboard/settings' },
    ];

    return (
        <aside className="w-64 bg-white border-r border-gray-200 min-h-screen">
            <nav className="p-4 space-y-2">
                {menuItems.map((item) => {
                    const Icon = item.icon;
                    const isActive = location.pathname === item.path;

                    return (
                        <Link
                            key={item.path}
                            to={item.path}
                            className={`flex items-center gap-3 px-4 py-3 rounded-lg transition-colors ${isActive
                                ? 'bg-indigo-600 text-white'
                                : 'text-gray-700 hover:bg-gray-100'
                                }`}
                        >
                            <Icon className="w-5 h-5" />
                            <span className="font-medium">{item.label}</span>
                        </Link>
                    );
                })}
            </nav>
        </aside>
    );
}