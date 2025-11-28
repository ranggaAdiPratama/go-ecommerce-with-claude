import { Link, useNavigate } from 'react-router-dom';
import { usePageTitle } from '../hooks/usePageTitle';
import { useState, type FormEvent } from 'react';
import type { LoginRequest } from '../types/auth.types';
import Swal from 'sweetalert2';
import { authService } from '../services/auth.service';
import { Eye, EyeOff, Lock, LogIn, Store, User } from 'lucide-react';
import { useSettings } from '../hooks/useSettings';

export const LoginPage = () => {
    usePageTitle("Login");

    const { settings } = useSettings()

    const navigate = useNavigate();

    const [formData, setFormData] = useState<LoginRequest>({
        username: '',
        password: '',
    });

    const [showPassword, setShowPassword] = useState(false);
    const [loading, setLoading] = useState(false);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value,
        });
    };

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault();

        if (!formData.username || !formData.password) {
            Swal.fire({
                icon: 'warning',
                title: 'Incomplete Form',
                text: 'Please fill in all fields',
                confirmButtonColor: '#4F46E5',
            });
            return;
        }

        try {
            setLoading(true);

            const data = await authService.login(formData);

            if (typeof data === 'object' && data !== null && 'user' in data) {
                navigate('/dashboard');
            }
        } catch (error) {
            Swal.fire({
                icon: 'error',
                title: 'Login Failed',
                text: error instanceof Error ? error.message : 'Invalid username or password',
                confirmButtonColor: '#4F46E5',
            });
        } finally {
            setLoading(false);
        }
    };

    const logoEmptyState = (<Store className="w-20 h-20 text-indigo-600" />)

    return (
        <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center p-4">
            <div className="bg-white rounded-2xl shadow-xl p-8 w-full max-w-md">
                {/* Header */}
                <div className="text-center mb-8">
                    <div className="inline-flex items-center justify-center w-16 h-16 bg-indigo-100 rounded-full mb-4">
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
                    <h1 className="text-3xl font-bold text-gray-900 mb-2">Welcome Back</h1>
                    <p className="text-gray-600">Sign in to your account</p>
                </div>

                {/* Form */}
                <form onSubmit={handleSubmit} className="space-y-4">
                    {/* Username */}
                    <div>
                        <label htmlFor="username" className="block text-sm font-medium text-gray-700 mb-2">
                            Username
                        </label>
                        <div className="relative">
                            <User className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
                            <input
                                type="text"
                                id="username"
                                name="username"
                                value={formData.username}
                                onChange={handleChange}
                                className="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-600 focus:border-transparent outline-none transition"
                                placeholder="john@example.com"
                                required
                            />
                        </div>
                    </div>

                    {/* Password */}
                    <div>
                        <label htmlFor="password" className="block text-sm font-medium text-gray-700 mb-2">
                            Password
                        </label>
                        <div className="relative">
                            <Lock className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
                            <input
                                type={showPassword ? 'text' : 'password'}
                                id="password"
                                name="password"
                                value={formData.password}
                                onChange={handleChange}
                                className="w-full pl-10 pr-12 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-600 focus:border-transparent outline-none transition"
                                placeholder="••••••••"
                                required
                            />
                            <button
                                type="button"
                                onClick={() => setShowPassword(!showPassword)}
                                className="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600"
                            >
                                {showPassword ? <EyeOff className="w-5 h-5" /> : <Eye className="w-5 h-5" />}
                            </button>
                        </div>
                    </div>

                    {/* Forgot Password */}
                    <div className="text-right">
                        <a href="#" className="text-sm text-indigo-600 hover:text-indigo-700">
                            Forgot password?
                        </a>
                    </div>

                    {/* Submit Button */}
                    <button
                        type="submit"
                        disabled={loading}
                        className="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-semibold py-3 px-4 rounded-lg transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center gap-2"
                    >
                        {loading ? (
                            <>
                                <div className="w-5 h-5 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
                                Signing in...
                            </>
                        ) : (
                            <>
                                <LogIn className="w-5 h-5" />
                                Sign In
                            </>
                        )}
                    </button>
                </form>

                {/* Footer */}
                <div className="mt-6 text-center">
                    <p className="text-sm text-gray-600">
                        Don't have an account?{' '}
                        <Link to="/auth/register" className="text-indigo-600 hover:text-indigo-700 font-semibold">
                            Sign Up
                        </Link>
                    </p>
                </div>
            </div>
        </div>
    );
};