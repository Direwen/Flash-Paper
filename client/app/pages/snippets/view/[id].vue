<script setup lang="ts">
import { 
    MazLockClosed, MazFire, MazEye, MazClock, 
    MazDocumentDuplicate, MazShieldCheck 
} from '@maz-ui/icons'
import hljs from 'highlight.js'
import 'highlight.js/styles/atom-one-dark.css'
import type { ApiResponse } from '~/types/dashboard'

definePageMeta({
    layout: 'default'
})

const route = useRoute()
const { $api, $toast } = useNuxtApp()
const authStore = useAuthStore()

// State
const id = route.params.id as string
const step = ref<'locked' | 'revealed' | 'burnt'>('locked')
const isLoading = ref(false)
const errorState = ref<string>("This secret has been burnt, expired, or never existed.")

// Data Containers
const secretContent = ref<string>('')
const metadata = ref<{
    title?: string,
    language?: string,
    owner_id: string | null,
    views_left: number,
    expires_at: string
} | null>(null)

// Fetch Metadata Immediately WITHOUT burning a view.
const { data: metaResponse, error: metaError } = await useAsyncData<ApiResponse<any>>(
    `snippet-meta-${id}`,
    () => $api(`/snippets/${id}/meta`),
    { server: false }
)

// Handle Metadata Error
if (metaError.value) {
    step.value = 'burnt'
} else if (metaResponse.value?.data) {
    metadata.value = metaResponse.value.data
}

const isOwner = computed(() => {
    if (!authStore.user || !metadata.value?.owner_id) return false
    return authStore.user.id === metadata.value.owner_id
})

const isLastView = computed(() => {
    return metadata.value?.views_left === 1
})

const revealSecret = async () => {
    isLoading.value = true

    try {
        const response: any = await $api(`/snippets/${id}`)
        
        if (response.success) {
            secretContent.value = response.data.content
            // Merge the full details (Title/Language) into metadata ref
            metadata.value = { ...metadata.value, ...response.data }
            step.value = 'revealed'
            
            nextTick(() => {
                hljs.highlightAll()
            })
        }
    } catch (error: any) {
        step.value = 'burnt'
        const status = error.response?.status
        errorState.value = "This secret has been burnt, expired, or never existed."
    } finally {
        isLoading.value = false
    }
}

const copyContent = () => {
    navigator.clipboard.writeText(secretContent.value)
    $toast.success("Secret copied to clipboard")
}
</script>

<template>
    <div class="min-h-[80vh] flex items-center justify-center p-4">
        
        <div v-if="step === 'locked'" class="max-w-md w-full text-center space-y-8 animate-in fade-in zoom-in duration-500">
            
            <div class="relative inline-block">
                <div class="absolute inset-0 bg-primary/20 blur-xl rounded-full"></div>
                <div class="relative bg-secondary p-6 rounded-2xl border border-white/10 shadow-2xl">
                    <MazLockClosed class="w-16 h-16 text-primary" />
                </div>
            </div>

            <div class="space-y-4">
                <h1 class="text-4xl font-bold text-white">Secure Transmission</h1>
                
                <div v-if="isOwner" class="bg-warning/10 border border-warning/20 rounded-lg p-4 text-left">
                    <div class="flex items-start gap-3">
                        <MazFire class="w-5 h-5 text-warning shrink-0 mt-0.5" />
                        <div>
                            <h4 class="font-bold text-warning text-sm">Creator Warning</h4>
                            <p class="text-xs text-warning/80 mt-1">
                                You created this. Revealing it counts as a view.
                            </p>
                        </div>
                    </div>
                </div>

                <div v-else-if="isLastView" class="bg-danger/10 border border-danger/20 rounded-lg p-4 text-left">
                    <div class="flex items-start gap-3">
                        <MazFire class="w-5 h-5 text-danger shrink-0 mt-0.5" />
                        <div>
                            <h4 class="font-bold text-danger text-sm">Final View</h4>
                            <p class="text-xs text-danger/80 mt-1">
                                This secret has 1 view left. It will be permanently destroyed after you read it.
                            </p>
                        </div>
                    </div>
                </div>

                <p v-else class="text-white/60 leading-relaxed">
                    You have received a secure message.
                    <br />
                    <span class="text-warning">Warning:</span> Reading this message will decrease its view count.
                </p>
            </div>

            <MazBtn 
                size="xl" 
                color="primary" 
                :loading="isLoading" 
                @click="revealSecret"
                class="font-bold tracking-wide shadow-lg shadow-primary/20 w-full"
            >
                REVEAL SECRET
            </MazBtn>
            
            <p class="text-xs text-white/20 mt-4">
                Do not share this link. Bots may burn the view.
            </p>
        </div>

        <div v-else-if="step === 'revealed'" class="max-w-3xl w-full animate-in fade-in slide-in-from-bottom-4 duration-500">
            
            <div class="flex items-center justify-between mb-6">
                <div>
                    <h2 class="text-2xl font-bold text-white">{{ metadata?.title || 'Untitled Secret' }}</h2>
                    <p class="text-sm text-white/40 font-mono mt-1">{{ id }}</p>
                </div>
                <MazBtn size="sm" color="secondary" @click="copyContent">
                    <template #left-icon><MazDocumentDuplicate class="w-4 h-4"/></template>
                </MazBtn>
            </div>

            <div class="relative group">
                <div class="absolute -inset-0.5 bg-gradient-to-r from-primary/20 to-purple-600/20 rounded-xl blur opacity-75"></div>
                
                <div class="relative bg-[#0d1117] rounded-xl border border-white/10 overflow-hidden shadow-2xl">
                    <div class="flex items-center gap-2 px-4 py-3 bg-white/5 border-b border-white/5">
                        <div class="flex gap-1.5">
                            <div class="w-3 h-3 rounded-full bg-red-500/20"></div>
                            <div class="w-3 h-3 rounded-full bg-yellow-500/20"></div>
                            <div class="w-3 h-3 rounded-full bg-green-500/20"></div>
                        </div>
                        <span class="ml-2 text-xs font-mono text-white/30 uppercase">{{ metadata?.language || 'text' }}</span>
                    </div>

                    <pre class="p-6 overflow-x-auto text-sm font-mono text-gray-300 leading-relaxed"><code :class="`language-${metadata?.language || 'text'}`">{{ secretContent }}</code></pre>
                </div>
            </div>

            <div class="mt-8 text-center">
                <div class="inline-flex items-center gap-2 px-4 py-2 bg-danger/10 text-danger rounded-full text-sm font-medium animate-pulse">
                    <MazFire class="w-4 h-4" />
                    This message is burning...
                </div>
                <div class="mt-4">
                    <MazBtn to="/" color="primary" outline>Create Your Own</MazBtn>
                </div>
            </div>
        </div>

        <div v-else class="max-w-md w-full text-center space-y-6">
            <div class="inline-flex p-6 bg-danger/10 rounded-full text-danger mb-4">
                <MazFire class="w-16 h-16" />
            </div>
            
            <h1 class="text-3xl font-bold text-white">Gone Up in Smoke</h1>
            
            <p class="text-lg text-white/50 leading-relaxed">
                {{ errorState }}
            </p>

            <div class="pt-6">
                <MazBtn to="/" color="primary" size="lg">Create New Secret</MazBtn>
            </div>
        </div>

    </div>
</template>