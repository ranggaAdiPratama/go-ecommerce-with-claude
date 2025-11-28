import { usePageTitle } from '../hooks/usePageTitle';
import { DashboardHeader } from '../components/dashboard/DashboardHeader';
import { DashboardSidebar } from '../components/dashboard/DashboardSidebar';

export const DashboardPage = () => {
    usePageTitle('Dashboard');

    return (
        <div className="min-h-screen bg-gray-50">
            <DashboardHeader />

            <div className="flex">
                <DashboardSidebar />

                <main className="flex-1 p-8">
                    <div className="max-w-7xl mx-auto">
                        <h1 className="text-3xl font-bold text-gray-900 mb-8">Dashboard</h1>

                        {/* Dashboard Content */}
                        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
                            {/* Stats Cards */}
                            <div className="bg-white rounded-lg shadow p-6">
                                <p className="text-sm text-gray-600 mb-1">Total Sales</p>
                                <p className="text-2xl font-bold text-gray-900">Rp 10,250,000</p>
                                <p className="text-sm text-green-600 mt-2">+12% from last month</p>
                            </div>

                            <div className="bg-white rounded-lg shadow p-6">
                                <p className="text-sm text-gray-600 mb-1">Orders</p>
                                <p className="text-2xl font-bold text-gray-900">324</p>
                                <p className="text-sm text-green-600 mt-2">+8% from last month</p>
                            </div>

                            <div className="bg-white rounded-lg shadow p-6">
                                <p className="text-sm text-gray-600 mb-1">Products</p>
                                <p className="text-2xl font-bold text-gray-900">156</p>
                                <p className="text-sm text-gray-500 mt-2">Active products</p>
                            </div>

                            <div className="bg-white rounded-lg shadow p-6">
                                <p className="text-sm text-gray-600 mb-1">Customers</p>
                                <p className="text-2xl font-bold text-gray-900">1,234</p>
                                <p className="text-sm text-green-600 mt-2">+15% from last month</p>
                            </div>
                        </div>

                        {/* Recent Activity */}
                        <div className="mt-8 bg-white rounded-lg shadow">
                            <div className="p-6 border-b border-gray-200">
                                <h2 className="text-xl font-semibold text-gray-900">Recent Activity</h2>
                            </div>
                            <div className="p-6">
                                <p className="text-gray-500 text-center py-8">No recent activity</p>
                            </div>
                        </div>
                    </div>
                </main>
            </div>
        </div>
    );
};