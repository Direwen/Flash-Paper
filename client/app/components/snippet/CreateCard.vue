<script setup lang="ts">
import { MazLockClosed } from '@maz-ui/icons'
</script>

<template>
    <div class="relative group w-full max-w-lg">
        <div class="absolute -inset-0.5 bg-gradient-to-br from-primary/30 to-purple-600/30 rounded-2xl blur opacity-75 group-hover:opacity-100 transition duration-500"></div>
        
        <div class="relative bg-secondary rounded-xl border border-white/5 p-6 lg:p-8 shadow-2xl">
        
        <div class="flex items-center gap-3 mb-6">
            <div class="p-2 bg-background rounded-lg text-primary">
            <MazLockClosed class="w-6 h-6" />
            </div>
            <div>
            <h3 class="font-bold text-lg text-white">Create Secure Secret</h3>
            <p class="text-xs text-white/40">End-to-end encrypted in your browser</p>
            </div>
        </div>

        <div class="h-64 border-2 border-dashed border-white/10 rounded-lg flex flex-col items-center justify-center text-white/30 gap-2">
            <span class="text-2xl">🏗️</span>
            <span class="text-sm font-mono">Form Component Slot</span>
        </div>

        <div class="mt-6 flex items-center justify-center gap-4 text-[10px] text-white/20 uppercase tracking-widest">
            <span>🔒 Zero Knowledge</span>
            <span>🔥 Auto-Destruct</span>
        </div>

        </div>
    </div>
</template>