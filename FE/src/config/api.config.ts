export const apiConfig = {
    baseUrl: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080',
    endpoints: {
        categories: '/api/categories',
        login: '/api/auth/login',
        logout: '/api/auth/logout',
        register: '/api/auth/register',
        settings: '/api/settings',
        shops: '/api/shops',
    },
    timeout: 10000,
}

export default apiConfig