import { useNuxtApp } from '#app'

export default defineNuxtPlugin((nuxtApp) => {
    const config = useRuntimeConfig()
    const token = useCookie('token')
    const { $toast } = useNuxtApp()
    const { parseError } = useErrorParser()

    const api = $fetch.create({
        baseURL: config.public.apiBase,

        onRequest({ options }) {
            if (token.value) {
                // Ensure headers is a Headers object
                const headers = new Headers(options.headers)
                headers.set('Authorization', `Bearer ${token.value}`)
                options.headers = headers
            }
        },
        
        onResponseError({ response }) {
            if (response.status === 401) {
                token.value = null
                $toast?.error(parseError("Your session has expired. Please log in again."))
                navigateTo('/auth/login')
            }
        }
    })

    return {
        provide: {
            api
        }
    }
})