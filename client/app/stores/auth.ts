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
        console.log(response)
        // Set the cookie. The plugin will detect this change for future requests!
        if (response.success && response.data.token) {
            token.value = response.data.token
            console.log("Fetched", token.value)
            await fetchUser()
        }
    }

    const fetchUser = async () => {
        if (!token.value) return

        try {
            const response = await $api<{ success: boolean, data: User }>('/me')
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
    }

    // 4. Initialize (Run this when app starts)
    const initAuth = async () => {
        if (token.value && !user.value) {
            await fetchUser()
        }
    }

    return {
        token,
        user,
        isLoading,
        register,
        login,
        logout,
        fetchUser,
        initAuth
    }
})