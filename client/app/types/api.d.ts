// Declaring Nuxt has a global property $api and it behaves like $fetch

declare module '#app' {
    interface NuxtApp {
        $api: typeof $fetch
    }
}

declare module 'vue' {
    interface ComponentCustomProperties {
        $api: typeof $fetch
    }
}

export {}
