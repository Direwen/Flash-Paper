export default defineNuxtPlugin(() => {
    const config = useRuntimeConfig()
    // get a reactive ref of browser cookie "token"
    const token = useCookie<string | null>('token')

    const api = $fetch.create({
        baseURL: config.public.apiBase as string,
        onRequest({options}) {
            if (token.value) {
                options.headers = new Headers(options.headers)
                options.headers.set('Authorization', `Bearer ${token.value}`)
            }
        },
        onResponseError({response}) {
            if (response.status === 401) {
                // remove the token (reactive ref) and redirect to login
                token.value = null
                navigateTo("/auth/login")
            }
        }
    })

    return {
        provide: {api}
    }
})