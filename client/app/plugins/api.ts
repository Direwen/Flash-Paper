// plugins/api.ts
export default defineNuxtPlugin((nuxtApp) => {
    const config = useRuntimeConfig()
    const token = useCookie('token')

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