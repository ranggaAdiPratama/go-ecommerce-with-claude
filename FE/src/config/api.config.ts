export const apiConfig = {
    baseUrl: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080',
    endpoints: {
        shops: '/api/shops',
    },
    timeout: 10000,
}

export default apiConfig