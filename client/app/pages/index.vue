<script setup lang="ts">
import { MazLockClosed, MazShieldCheck, MazClock } from '@maz-ui/icons'

const authStore = useAuthStore()
</script>

<template>
    <div class="flex-grow flex items-center justify-center relative overflow-hidden py-12 lg:py-0">

        <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[400px] h-[400px] bg-primary/10 rounded-full blur-[120px] pointer-events-none"></div>

        <div class="max-w-7xl md:py-6 lg:py-12 w-full mx-auto px-6 grid lg:grid-cols-2 gap-12 items-center">
            
            <!-- Text Section -->
            <div class="space-y-8 relative z-10">
            
                <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-primary/10 border border-primary/20 text-primary text-xs font-bold tracking-wider uppercase animate-fade-in">
                    <MazLockClosed class="w-3 h-3" />
                    AES-256-GCM Encrypted
                </div>
                
                <h1 class="text-5xl lg:text-7xl font-bold leading-tight tracking-tight">
                    Share Secrets That <br />
                    <span class="text-transparent bg-clip-text bg-gradient-to-r from-primary to-orange-500 animate-pulse">
                    Self-Destruct
                    </span>
                </h1>
                
                <p class="text-lg text-white/60 max-w-lg leading-relaxed">
                    Send passwords, API keys, and private data that burns after reading. 
                    We never store your keys, so we can never see your data.
                </p>

                <div class="flex flex-wrap items-center gap-6 text-sm font-medium text-white/40 pt-2">
                    <div class="flex items-center gap-2">
                    <MazShieldCheck class="w-5 h-5 text-primary" />
                    <span>No Server Storage</span>
                    </div>
                    <div class="flex items-center gap-2">
                    <MazClock class="w-5 h-5 text-warning" />
                    <span>Time-Based Expiry</span>
                    </div>
                </div>
            </div>

            <!-- Create Card Section -->
            <div class="flex justify-center lg:justify-end relative z-10">
                <span>{{ authStore.user ?? "No User" }}</span>
                <SnippetCreateCard />
            </div>

        </div>
    </div>
</template>

<style scoped>
.animate-fade-in {
    animation: fadeIn 1s ease-out;
}
@keyframes fadeIn {
    from { opacity: 0; transform: translateY(10px); }
    to { opacity: 1; transform: translateY(0); }
}
</style>