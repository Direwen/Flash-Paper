<script setup lang="ts">
import { MazEnvelope, MazLockClosed, MazFire } from '@maz-ui/icons'
import { ref } from 'vue'
import { useAuthStore } from '~/stores/auth'
import { useNuxtApp } from '#app'


const form = ref({
    email: '',
    password: ''
})

const isLoading = ref(false)
const authStore = useAuthStore()
const router = useRouter()
const { $toast } = useNuxtApp()
const { parseError } = useErrorParser()

const handleLogin = async () => {
    isLoading.value = true
    try {
        await authStore.login(form.value)
        $toast?.success(`Welcome back ${authStore.user?.email}`)
        router.push("/")
    } catch (error) {
        $toast?.error(parseError(error, "Failed to log in"))
    } finally {
        isLoading.value = false
    }
}
</script>

<template>
    <div class="flex-grow flex items-center justify-center relative py-12 px-4">
        
        <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[500px] h-[500px] bg-primary/5 rounded-full blur-[100px] pointer-events-none"></div>

        <div class="relative w-full max-w-md">
            <div class="absolute -inset-0.5 bg-gradient-to-b from-primary/20 to-transparent rounded-2xl blur-sm opacity-50"></div>
            
            <div class="relative bg-secondary rounded-xl border border-white/5 p-8 shadow-2xl backdrop-blur-sm">
                
                <div class="text-center mb-8">
                <div class="inline-flex items-center justify-center w-12 h-12 rounded-full bg-primary/10 text-primary mb-4">
                    <MazFire class="w-12 h-12 text-primary" />
                </div>
                <h1 class="text-2xl font-bold text-white mb-2">Welcome Back</h1>
                <p class="text-sm text-white/40">Sign in to manage your encrypted secrets</p>
                </div>

                <form @submit.prevent="handleLogin" class="space-y-6">
                
                <MazInput
                    v-model="form.email"
                    label="Email Address"
                    type="email"
                    color="primary"
                    class="w-full"
                    auto-focus
                >
                    <template #left-icon>
                        <MazEnvelope class="w-5 h-5 text-white/40" />
                    </template>
                </MazInput>

                <div class="space-y-1">
                    <MazInput
                    v-model="form.password"
                    label="Password"
                    type="password"
                    color="primary"
                    class="w-full"
                    >
                        <template #left-icon>
                            <MazLockClosed class="w-5 h-5 text-white/40" />
                        </template>
                    </MazInput>

                    <div class="flex justify-end">
                        <NuxtLink to="/auth/forgot-password" class="text-xs text-primary/80 hover:text-primary hover:underline transition-all">
                            Forgot password?
                        </NuxtLink>
                    </div>
                </div>

                <MazBtn
                    type="submit"
                    color="primary"
                    block
                    :loading="isLoading"
                    class="font-bold tracking-wide"
                >
                    SIGN IN
                </MazBtn>

                </form>

                <div class="mt-8 text-center text-sm text-white/40">
                    Don't have an account? 
                    <NuxtLink to="/auth/register" class="text-primary hover:text-primary/80 font-medium transition-colors">
                        Sign up
                    </NuxtLink>
                </div>

            </div>
        </div>

    </div>
</template>