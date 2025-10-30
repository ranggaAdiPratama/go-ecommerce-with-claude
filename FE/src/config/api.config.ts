export const apiConfig = {
    baseUrl: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080',
    endpoints: {
        settings: '/api/settings',
        shops: '/api/shops',
    },
    timeout: 10000,
}

export default apiConfig