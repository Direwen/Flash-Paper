<script setup lang="ts">
import { MazFire, MazBars3, MazXMark, MazLockClosed } from '@maz-ui/icons'

// 1. Mobile Menu State
const isMobileMenuOpen = ref(false)

// 2. Navigation Links Configuration
const navLinks = [
    { label: 'How it Works', to: '/#how-it-works', isButton: false },
    { label: 'Dashboard', to: '/dashboard', isButton: false },
    { label: 'Sign In', to: '/auth/login', isButton: false },
    { label: 'Get Started', to: '/auth/register', isButton: true }, // Special styling
]

// Helper to close menu when a link is clicked
const closeMenu = () => {
    isMobileMenuOpen.value = false
}
</script>

<template>
    <div class="min-h-screen flex flex-col bg-background text-foreground font-base selection:bg-primary selection:text-black">

        <header class="fixed top-0 left-0 right-0 z-50 border-b border-white/5 bg-background/80 backdrop-blur-md">
            <div class="max-w-7xl mx-auto px-6 h-20 flex items-center justify-between">
            
            <NuxtLink to="/" class="flex items-center gap-2 group z-50" @click="closeMenu">
                <span class="text-2xl"></span>
                <MazFire class="w-8 h-8 text-primary group-hover:text-primary/80 transition-colors" />
                <span class="font-bold text-xl tracking-tight group-hover:text-primary transition-colors">
                FlashPaper
                </span>
            </NuxtLink>

            <nav class="hidden md:flex items-center gap-8">
                <template v-for="link in navLinks" :key="link.label">
                
                <MazBtn 
                    v-if="link.isButton" 
                    :to="link.to"
                    color="primary"
                    size="sm"
                    class="font-bold text-black"
                >
                    {{ link.label }}
                </MazBtn>

                <NuxtLink 
                    v-else 
                    :to="link.to" 
                    class="text-sm font-medium text-white/70 hover:text-primary transition-colors"
                >
                    {{ link.label }}
                </NuxtLink>

                </template>
            </nav>

            <button class="md:hidden text-white z-50 p-2" @click="isMobileMenuOpen = !isMobileMenuOpen">
                <MazXMark v-if="isMobileMenuOpen" class="w-6 h-6" />
                <MazBars3 v-else class="w-6 h-6" />
            </button>
            </div>

            <transition
            enter-active-class="transition duration-200 ease-out"
            enter-from-class="opacity-0 -translate-y-2"
            enter-to-class="opacity-100 translate-y-0"
            leave-active-class="transition duration-150 ease-in"
            leave-from-class="opacity-100 translate-y-0"
            leave-to-class="opacity-0 -translate-y-2"
            >
            <div v-show="isMobileMenuOpen" class="absolute top-20 left-0 w-full bg-secondary border-b border-white/5 shadow-2xl md:hidden flex flex-col p-6 gap-4">
                <template v-for="link in navLinks" :key="link.label">
                <NuxtLink 
                    :to="link.to"
                    @click="closeMenu"
                    class="text-lg font-medium py-2 border-b border-white/5 last:border-0"
                    :class="link.isButton ? 'text-primary' : 'text-white/80'"
                >
                    {{ link.label }}
                </NuxtLink>
                </template>
            </div>
            </transition>
        </header>

        <main class="flex-grow pt-20 flex flex-col relative">
            <slot />
        </main>

        <footer class="border-t border-white/5 bg-secondary/30 py-12 mt-auto">
            <div class="max-w-7xl mx-auto px-6 grid grid-cols-1 md:grid-cols-4 gap-8">
            
            <div class="col-span-1 md:col-span-1">
                <div class="flex items-center gap-2 mb-4">
                    <MazFire class="w-8 h-8 text-primary group-hover:text-primary/80 transition-colors" />
                    <span class="font-bold text-lg">FlashPaper</span>
                </div>
                <p class="text-sm text-white/40 leading-relaxed">
                Zero-knowledge encrypted message sharing. Your secrets burn after reading, leaving no trace.
                </p>
                <div class="flex gap-2 mt-6 text-xs text-primary/80">
                    <MazLockClosed class="w-4 h-4" />
                    <span class="flex items-center gap-1">End-to-End Encrypted</span>
                </div>
            </div>

            <div>
                <h4 class="font-bold text-white mb-4">Product</h4>
                <ul class="space-y-2 text-sm text-white/50">
                <li><NuxtLink to="/#how-it-works" class="hover:text-primary transition-colors">How it Works</NuxtLink></li>
                <li><NuxtLink to="/dashboard" class="hover:text-primary transition-colors">Dashboard</NuxtLink></li>
                <li><NuxtLink to="/api-docs" class="hover:text-primary transition-colors">API Docs</NuxtLink></li>
                </ul>
            </div>

            <div>
                <h4 class="font-bold text-white mb-4">Legal</h4>
                <ul class="space-y-2 text-sm text-white/50">
                <li><NuxtLink to="/privacy" class="hover:text-primary transition-colors">Privacy Policy</NuxtLink></li>
                <li><NuxtLink to="/terms" class="hover:text-primary transition-colors">Terms of Service</NuxtLink></li>
                <li><a href="#" class="hover:text-primary transition-colors">Open Source</a></li>
                </ul>
            </div>

            <div class="text-right flex flex-col justify-end">
                <p class="text-xs text-white/20">
                v1.0.0 · Built with paranoia
                </p>
                <p class="text-xs text-white/20 mt-2">
                &copy; {{ new Date().getFullYear() }} FlashPaper.
                </p>
            </div>
            </div>
        </footer>

    </div>
</template>