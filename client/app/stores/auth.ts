import { defineStore } from 'pinia'

interface User {
    id: string
    email: string
    created_at: string
}

export const useAuthStore = defineStore('auth', () => {
    // Injections
    const { $api } = useNuxtApp()
    const token = useCookie('token')
    // State
    const user = ref<User | null>(null)
    const isLoading = ref(false)
    // 3. Actions
    const register = async (credentials: any) => {
        // $api automatically uses baseURL
        await $api('/auth/register', {
        method: 'POST',
        body: credentials
        })
    }

    const login = async (credentials: any) => {
        const response = await $api<{ success: boolean, data: { token: string } }>('/auth/login', {
        method: 'POST',
        body: credentials
        })
        
        // Set the cookie. The plugin will detect this change for future requests!
        if (response.success && response.data.token) {
        token.value = response.data.token
        await fetchUser() // Load profile immediately
        }
    }

    const fetchUser = async () => {
        if (!token.value) return

        try {
        const response = await $api<{ success: boolean, data: User }>('/api/me')
        if (response.success) {
            user.value = response.data
        }
        } catch (error) {
        // If /me fails (token expired), log them out
        token.value = null
        user.value = null
        }
    }

    const logout = () => {
        token.value = null
        user.value = null
        navigateTo('/auth/login')
    }

    // 4. Initialize (Run this when app starts)
    const initAuth = async () => {
        if (token.value && !user.value) {
        await fetchUser()
        }
    }

    return {
        user,
        isLoading,
        register,
        login,
        logout,
        fetchUser,
        initAuth
    }
})