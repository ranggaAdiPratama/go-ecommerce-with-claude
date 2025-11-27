import React, { useEffect, useRef, useState, type FormEvent } from "react"
import { usePageTitle } from "../hooks/usePageTitle"
import type { RegisterRequest } from "../types/auth.types"
import Swal from "sweetalert2"
import { authService } from "../services/auth.service"
import { Eye, EyeOff, Lock, Mail, Store, User, UserPlus } from "lucide-react"
import { useNavigate } from "react-router-dom"
import { useSettings } from "../hooks/useSettings"

export const RegisterPage = () => {
    usePageTitle('Register')

    const navigate = useNavigate();

    const inputRef = useRef<HTMLInputElement>(null)

    const { settings, loading } = useSettings()

    const [formData, setFormData] = useState<RegisterRequest>({
        name: '',
        email: '',
        username: '',
        password: '',
        role: 'user',
    })

    const [showPassword, setShowPassword] = useState(false)
    const [loadingState, setLoadingState] = useState(false)

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value
        })
    }

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault()

        if (!formData.name || !formData.email || !formData.username || !formData.password) {
            Swal.fire({
                icon: 'warning',
                title: 'Incomplete Form',
                text: 'Please fill in all fields',
                confirmButtonColor: '#4F46E5',
            })

            return
        }

        if (formData.password.length < 8) {
            Swal.fire({
                icon: 'warning',
                title: 'Weak Password',
                text: 'Password must be at least 8 characters',
                confirmButtonColor: '#4F46E5',
            })

            return
        }

        try {
            setLoadingState(true)

            await authService.register(formData)

            Swal.fire({
                icon: 'success',
                title: 'Registration Successful!',
                text: 'Your account has been created successfully',
                confirmButtonColor: '#4F46E5',
                confirmButtonText: 'Login Now!'
            }).then(() => {
                navigate('/auth/login')
            });
        } catch (error) {
            Swal.fire({
                icon: 'error',
                title: 'Registration Failed',
                text: error instanceof Error ? error.message : 'An error occurred during registration',
                confirmButtonColor: '#4F46E5',
            });
        } finally {
            setLoadingState(false);
        }
    }

    const logoEmptyState = (<Store className="w-20 h-20 text-indigo-600" />)

    useEffect(() => {
        inputRef.current?.focus();
    }, [])

    return (
        <div className="min-h-screen bg-radient-to-br from-blue-50 to-indigo-100 flex items-center justify-center p-4">
            <div className="bg-white rounded-2xl shadow-xl p-8 w-full max-w-md">
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
                    <h1 className="text-3xl font-bold text-gray-900 mb-2">Create Account</h1>
                    <p className="text-gray-600">Sign up to get started</p>
                </div>

                <form onSubmit={handleSubmit} className="space-y-4">
                    <div>
                        <label htmlFor="name" className="block text-sm font-medium text-gray-700 mb-2">
                            Full Name
                        </label>
                        <div className="relative">
                            <User className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
                            <input
                                ref={inputRef}
                                type="text"
                                id="name"
                                name="name"
                                value={formData.name}
                                onChange={handleChange}
                                className="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-600 focus:border-transparent outline-none transition"
                                placeholder="Ngaran Aing"
                                required
                            />
                        </div>
                    </div>
                    <div>
                        <label htmlFor="email" className="block text-sm font-medium text-gray-700 mb-2">
                            Email Address
                        </label>
                        <div className="relative">
                            <Mail className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 w-5 h-5" />
                            <input
                                type="email"
                                id="email"
                                name="email"
                                value={formData.email}
                                onChange={handleChange}
                                className="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-600 focus:border-transparent outline-none transition"
                                placeholder="aing@email.com"
                                required
                            />
                        </div>
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
                                    placeholder="usernameaing"
                                    required
                                />
                            </div>
                        </div>
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
                                    minLength={8}
                                />
                                <button
                                    type="button"
                                    onClick={() => setShowPassword(!showPassword)}
                                    className="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600"
                                >
                                    {showPassword ? <EyeOff className="w-5 h-5" /> : <Eye className="w-5 h-5" />}
                                </button>
                            </div>
                            <p className="text-xs text-gray-500 mt-1">At least 8 characters</p>
                        </div>
                        <br></br>
                        <button
                            type="submit"
                            disabled={loadingState}
                            className="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-semibold py-3 px-4 rounded-lg transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center gap-2"
                        >
                            {
                                loadingState ? (
                                    <>
                                        <div className="w-5 h-5 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
                                        Creating Account...
                                    </>
                                ) : (
                                    <>
                                        <UserPlus className="w-5 h-5" />
                                        Create Account
                                    </>
                                )
                            }
                        </button>
                    </div>
                </form>
                <div className="mt-6 text-center">
                    <p className="text-sm text-gray-600">
                        Already have an account?
                        <a href="/auth/login" className="text-indigo-600 hover:text-indigo-700 font-semibold">
                            Sign In
                        </a>
                    </p>
                </div>
            </div>
        </div>
    )
}
