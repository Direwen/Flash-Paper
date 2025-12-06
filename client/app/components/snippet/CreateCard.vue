<script setup lang="ts">
import { 
    MazLockClosed, MazFire, MazClock, MazEye, 
    MazDocumentDuplicate, MazArrowPath
} from '@maz-ui/icons'


// Props & Emits
const props = withDefaults(defineProps<{
    showBackground?: boolean
}>(), {
    showBackground: true
})
const emit = defineEmits(['created']) // Let parent know something happened (e.g. refresh list)

// Services
const { $api, $toast } = useNuxtApp()
const authStore = useAuthStore()
const isGuest = computed(() => !authStore.user)

// State
const isLoading = ref(false)
const createdSnippet = ref<any>(null) // Stores the result after success

// Form Data
const form = ref({
    content: '',
    title: '',
    language: 'text',
    maxViews: 1,
    expiresIn: 60, // Default 1 hour
})

// Options
const languageOptions = [
    { label: 'Plain Text', value: 'text' },
    { label: 'Go', value: 'go' },
    { label: 'Python', value: 'python' },
    { label: 'JavaScript', value: 'javascript' },
    { label: 'TypeScript', value: 'typescript' },
    { label: 'HTML', value: 'html' },
    { label: 'SQL', value: 'sql' },
]

const expiryOptions = [
    { label: '10 Minutes', value: 10 },
    { label: '1 Hour', value: 60 },
    { label: '24 Hours', value: 1440 },
    { label: '3 Days', value: 4320 },
    { label: '7 Days', value: 10080 },
]

const viewOptions = [
    { label: '1 View', value: 1 },
    { label: '5 Views', value: 5 },
    { label: '10 Views', value: 10 },
    { label: '50 Views', value: 50 },
]

// Actions
const handleSubmit = async () => {
    if (!form.value.content) {
        $toast.error("Content is required")
        return
    }

    isLoading.value = true
    try {
        const response: any = await $api('/snippets', {
            method: 'POST',
            body: {
                content: form.value.content,
                title: form.value.title,
                language: form.value.language,
                max_views: form.value.maxViews,
                expires_in: form.value.expiresIn,
            }
        })
        
        if (response.success) {
            createdSnippet.value = response.data
            emit('created')
            $toast.success("Secret created successfully")
        }
    } catch (e) {
        $toast.error("Failed to create secret")
    } finally {
        isLoading.value = false
    }
}

const copyLink = () => {
    if (!createdSnippet.value) return
    const id = createdSnippet.value.id
    const fullLink = `${window.location.origin}/snippets/view/${id}`
    
    navigator.clipboard.writeText(fullLink)
    $toast.success("Link copied to clipboard")
}

const resetForm = () => {
    createdSnippet.value = null
    form.value.content = ''
    form.value.title = ''
    form.value.language = 'text'
}

const shareLink = computed(() => {
    if (!createdSnippet.value) return ''
    return `${window.location.origin}/snippets/view/${createdSnippet.value.id}`
})
</script>

<template>
    <div class="relative w-full max-w-lg mx-auto">
        <div v-if="showBackground && !createdSnippet" class="absolute -inset-0.5 bg-gradient-to-br from-primary/30 to-purple-600/30 rounded-2xl blur opacity-75 transition duration-500"></div>

        <div 
            class="relative rounded-xl p-6 lg:p-8 shadow-2xl overflow-hidden"
            :class="{'bg-secondary border border-white/5': showBackground}"
        >
            
            <div v-if="createdSnippet" class="text-center space-y-6 py-4 animate-in fade-in zoom-in duration-300">
                <div class="inline-flex p-4 bg-success/10 rounded-full text-success mb-2">
                    <MazLockClosed class="w-12 h-12" />
                </div>
                
                <div>
                    <h3 class="text-2xl font-bold text-white">Secret Armed</h3>
                    <p class="text-white/40 text-sm mt-1">This link will self-destruct after use.</p>
                </div>

                <div class="bg-black/30 border border-primary/30 rounded-lg p-4 flex items-center gap-3">
                    <div class="flex-grow font-mono text-primary text-sm truncate">
                        {{ shareLink }}
                    </div>
                    <MazBtn size="sm" color="primary" @click="copyLink">
                        <MazDocumentDuplicate class="w-4 h-4" />
                    </MazBtn>
                </div>

                <div class="pt-4">
                    <MazBtn color="primary" block outline @click="resetForm">
                        <template #left-icon><MazArrowPath class="w-4 h-4"/></template>
                        Create Another
                    </MazBtn>
                </div>
            </div>

            <div v-else class="space-y-5">
                <div class="flex items-center gap-3 mb-4">
                    <div class="p-2 bg-background rounded-lg text-primary">
                        <MazLockClosed class="w-6 h-6" />
                    </div>
                    <div>
                        <h3 class="font-bold text-lg text-white">New Secret</h3>
                        <p class="text-xs text-white/40">AES-256-GCM Encrypted</p>
                    </div>
                </div>

                <MazInput 
                    v-model="form.title" 
                    placeholder="Title" 
                    color="primary"
                    size="sm"
                    class="w-full"
                />

                <div class="relative">
                    <MazTextarea
                        v-model="form.content"
                        color="primary"
                        placeholder="Paste sensitive data here..."
                        :rows="5"
                        class="w-full font-mono text-sm [&_textarea]:!max-h-[10rem] [&_textarea]:!overflow-y-auto"
                    />
                    <div class="absolute bottom-2 right-2 w-32">
                        <MazSelect 
                            v-model="form.language" 
                            :options="languageOptions" 
                            color="primary" 
                            size="xs"
                            list-position="top"
                        />
                    </div>
                </div>

                <div class="grid grid-cols-2 gap-4">
                    <MazSelect
                        v-model="form.maxViews"
                        color="primary"
                        :options="viewOptions"
                        label="Self-destruct after"
                        list-position="top"
                    >
                        <template #left-icon><MazEye class="w-4 h-4"/></template>
                    </MazSelect>

                    <MazSelect
                        v-model="form.expiresIn"
                        color="primary"
                        :options="expiryOptions"
                        label="Expires in"
                        list-position="top"
                    >
                        <template #left-icon><MazClock class="w-4 h-4"/></template>
                    </MazSelect>
                </div>

                <div class="pt-2">
                    <NuxtLink v-if="isGuest" to="/auth/login">
                        <MazBtn 
                            block 
                            color="primary" 
                            size="lg"
                            class="font-bold tracking-wide shadow-lg shadow-primary/10"
                        >
                            <template #left-icon><MazLockClosed class="w-5 h-5"/></template>
                            LOGIN TO CREATE
                        </MazBtn>
                    </NuxtLink>
                    <MazBtn 
                        v-else
                        block 
                        color="primary" 
                        size="lg" 
                        :loading="isLoading" 
                        @click="handleSubmit"
                        class="font-bold tracking-wide shadow-lg shadow-primary/10"
                    >
                        <template #left-icon><MazFire class="w-5 h-5"/></template>
                        CREATE SECURE LINK
                    </MazBtn>
                </div>
            </div>

        </div>
    </div>
</template>