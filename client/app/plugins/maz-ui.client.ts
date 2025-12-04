import { ToastPlugin } from 'maz-ui/plugins'
import { useToast } from 'maz-ui/composables'

export default defineNuxtPlugin((nuxtApp) => {

    nuxtApp.vueApp.use(ToastPlugin)

    return {
        provide: {
            toast: useToast()
        }
    }
})