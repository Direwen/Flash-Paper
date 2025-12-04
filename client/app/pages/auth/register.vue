<script setup lang="ts">
import { MazEnvelope, MazLockClosed, MazFire } from '@maz-ui/icons'
import { ref } from 'vue'

const form = ref({
    email: '',
    password: ''
})

const isLoading = ref(false)
const authStore = useAuthStore()
const router = useRouter()
const { $toast } = useNuxtApp()

const handleRegister = async () => {
    isLoading.value = true
    try {
        await authStore.register(form.value)
        $toast?.success("Created Successfully")
        router.push('/')
    } catch (error) {
        console.error(error)
        $toast?.error("Failed to Register")
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
                    <h1 class="text-2xl font-bold text-white mb-2">Create Account</h1>
                    <p class="text-sm text-white/40">Join purely for the dashboard. No tracking.</p>
                </div>

                <form @submit.prevent="handleRegister" class="space-y-6">
                
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

                <MazInput
                    v-model="form.password"
                    label="Password"
                    type="password"
                    color="primary"
                    class="w-full"
                    hint="Minimum 6 characters"
                >
                    <template #left-icon>
                        <MazLockClosed class="w-5 h-5 text-white/40" />
                    </template>
                </MazInput>

                <MazBtn
                    type="submit"
                    color="primary" 
                    block
                    :loading="isLoading"
                    class="font-bold tracking-wide"
                >
                    CREATE ACCOUNT
                </MazBtn>

                <p class="text-[12px] text-center text-white/30 leading-tight">
                    By signing up, you agree to our 
                    <a href="#" class="hover:text-white/50 underline">Terms</a> and 
                    <a href="#" class="hover:text-white/50 underline">Privacy Policy</a>.
                </p>

                </form>

                <div class="mt-8 text-center text-sm text-white/40">
                Already have an account? 
                <NuxtLink to="/auth/login" class="text-primary/80 hover:text-primary font-medium transition-colors">
                    Sign in
                </NuxtLink>
                </div>

            </div>
        </div>

    </div>
</template>