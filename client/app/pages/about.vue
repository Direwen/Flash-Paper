<script setup lang="ts">
import { 
    MazLockClosed, MazEye, MazFire, MazKey, 
    MazServer, MazShieldCheck, MazArrowRight
} from '@maz-ui/icons'

const faqs = [
    {
        question: "How does the encryption work?",
        answer: "Think of it like putting your message in an unbreakable safe. Your browser uses AES-256 encryption (the same standard banks and governments use) to lock your message before sending it to us. We only store the encrypted data—never the key. The key stays with you, hidden in the link."
    },
    {
        question: "What does 'zero-knowledge' mean?",
        answer: "It means we literally cannot read your secrets, even if we tried. The decryption key is stored in the URL fragment (the part after the '#'), which your browser never sends to our servers. It's like mailing a locked box where only you and your recipient have the key."
    },
    {
        question: "How does self-destruction work?",
        answer: "The moment someone opens your secret (or the expiration time passes), we permanently delete it from our database. A background cleanup process double-checks that nothing remains. Once erased, the data is cryptographically unrecoverable—gone forever."
    },
    {
        question: "Can I trust FlashPaper with sensitive data?",
        answer: "Yes. All encryption happens client-side (on your device) using the Web Crypto API before anything reaches our servers. You're not trusting us—you're trusting proven cryptographic standards. Our code is open source, we store no logs, and we don't track IP addresses."
    },
    {
        question: "What happens if I lose the link?",
        answer: "The secret is permanently lost. Since the decryption key only exists in that link and we never store it, there's no way for anyone—including us—to recover your message. That's the trade-off for true end-to-end privacy."
    }
]

const flowSteps = [
    { icon: MazCodeBracket, label: "Your Secret", step: 1 },
    { icon: MazKey, label: "Generate Key", step: 2 },
    { icon: MazLockClosed, label: "Encrypt (Browser)", step: 3 },
    { icon: MazServer, label: "Store Cipher", step: 4 },
    { icon: MazFire, label: "View & Burn", step: 5 },
]

import { MazCodeBracket } from '@maz-ui/icons';
</script>

<template>
    <div class="w-full mx-auto px-6 py-16 space-y-24">

        <div class="text-center max-w-3xl mx-auto space-y-6">
            <div class="inline-flex p-4 bg-primary/10 rounded-full text-primary mb-2">
                <MazFire class="w-12 h-12" />
            </div>
            
            <h1 class="text-4xl md:text-5xl font-bold text-white">
                How <span class="text-primary">FlashPaper</span> Works
            </h1>
            <p class="text-lg text-white/50 leading-relaxed">
                Understanding our zero-knowledge encryption and self-destructing messages.
            </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-6 max-w-7xl mx-auto">
            <div class="bg-secondary border border-white/5 p-8 rounded-xl text-center space-y-4 hover:border-primary/30 transition-colors group">
                <div class="mx-auto w-12 h-12 bg-secondary border border-white/10 rounded-lg flex items-center justify-center text-primary group-hover:scale-110 transition-transform">
                    <MazLockClosed class="w-6 h-6" />
                </div>
                <h3 class="text-xl font-bold text-white">AES-256-GCM</h3>
                <p class="text-sm text-white/40">Military-grade encryption algorithm used by governments and banks worldwide.</p>
            </div>

            <div class="bg-secondary border border-white/5 p-8 rounded-xl text-center space-y-4 hover:border-primary/30 transition-colors group">
                <div class="mx-auto w-12 h-12 bg-secondary border border-white/10 rounded-lg flex items-center justify-center text-primary group-hover:scale-110 transition-transform">
                    <MazServer class="w-6 h-6" />
                </div>
                <h3 class="text-xl font-bold text-white">Zero Knowledge</h3>
                <p class="text-sm text-white/40">We never see your unencrypted data. The key stays in your URL hash.</p>
            </div>

            <div class="bg-secondary border border-white/5 p-8 rounded-xl text-center space-y-4 hover:border-primary/30 transition-colors group">
                <div class="mx-auto w-12 h-12 bg-secondary border border-white/10 rounded-lg flex items-center justify-center text-primary group-hover:scale-110 transition-transform">
                    <MazFire class="w-6 h-6" />
                </div>
                <h3 class="text-xl font-bold text-white">Auto-Destruct</h3>
                <p class="text-sm text-white/40">Secrets are permanently deleted after viewing or expiration.</p>
            </div>
        </div>

        <div class="max-w-7xl mx-auto">
            <div class="bg-secondary/30 border border-white/5 rounded-2xl p-12">
                <h3 class="text-center text-lg font-bold text-white mb-12">Encryption Flow</h3>
                
                <div class="flex flex-wrap items-center justify-center gap-8 md:gap-12 relative">
                    <template v-for="(step, index) in flowSteps" :key="step.step">
                        
                        <div class="flex flex-col items-center gap-4 relative group z-10">
                            <div class="relative w-16 h-16 bg-secondary border border-primary/20 rounded-full flex items-center justify-center shadow-[0_0_15px_rgba(251,191,36,0.1)] group-hover:border-primary transition-colors">
                                <component :is="step.icon" class="w-7 h-7 text-primary" />
                                
                                <div class="absolute -top-1 -right-1 w-6 h-6 bg-primary text-black text-xs font-bold rounded-full flex items-center justify-center">
                                    {{ step.step }}
                                </div>
                            </div>
                            <span class="text-xs font-mono text-white/60 uppercase tracking-wider">{{ step.label }}</span>
                        </div>

                        <div v-if="index < flowSteps.length - 1" class="hidden md:block text-white/10">
                            <MazArrowRight class="w-6 h-6" />
                        </div>

                    </template>
                </div>
            </div>
        </div>

        <div class="max-w-3xl mx-auto">
            <h3 class="text-center text-2xl font-bold text-white mb-8">Frequently Asked Questions</h3>
            
            <div class="space-y-4">
                <MazAccordion>
                    <template v-for="(faq, i) in faqs" :key="i" #[`title-${i+1}`]>
                        {{ faq.question }}
                    </template>
                    <template v-for="(faq, i) in faqs" :key="i" #[`content-${i+1}`]>
                        {{ faq.answer }}
                    </template>
                </MazAccordion>
            </div>
        </div>

        <div class="flex justify-center gap-8 pt-12 border-t border-white/5 max-w-lg mx-auto">
            <div class="flex items-center gap-2 text-primary/80 text-xs font-medium uppercase tracking-widest">
                <MazCodeBracket class="w-4 h-4" /> Open Source
            </div>
            <div class="flex items-center gap-2 text-primary/80 text-xs font-medium uppercase tracking-widest">
                <MazShieldCheck class="w-4 h-4" /> Audited
            </div>
            <div class="flex items-center gap-2 text-primary/80 text-xs font-medium uppercase tracking-widest">
                <MazEye class="w-4 h-4" /> No Tracking
            </div>
        </div>

    </div>
</template>

<style scoped>
:deep(.m-accordion-item) {
    background-color: var(--color-secondary);
    border: 1px solid rgba(255, 255, 255, 0.05);
    border-radius: 0.75rem;
    margin-bottom: 1rem;
    overflow: hidden;
    transition: border-color 0.2s;
}
:deep(.m-accordion-item:hover) {
    border-color: rgba(255, 255, 255, 0.1);
}
:deep(.m-accordion-item__header) {
    background-color: transparent !important;
}
</style>