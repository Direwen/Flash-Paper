<script setup lang="ts">
import { 
    MazFire, MazEye, MazTrash, MazDocumentDuplicate, 
    MazClock, MazArrowTopRightOnSquare // <--- New Icon for "Open"
} from '@maz-ui/icons'
import type { DashboardStats, OverviewSnippet, PaginatedList, ApiResponse } from './types/dashboard'

definePageMeta({
    middleware: 'auth',
})

const { $api } = useNuxtApp()
const { $toast } = useNuxtApp()
const token = useCookie('token')

// State
const page = ref(1)
const limit = ref(10)

const [{ data: statsResponse, error: statsError }, { data: listResponse, refresh: refreshList, error: listError, pending: loadingList }] = await Promise.all([
    useAsyncData<ApiResponse<DashboardStats>>(
        'dashboard-stats',
        () => $api('/dashboard'),
        { server: false }
    ),
    useAsyncData<ApiResponse<PaginatedList<OverviewSnippet>>>(
        'dashboard-snippets',
        () => $api('/snippets', {
            query: { page: page.value, limit: limit.value }
        }),
        { watch: [page, limit], server: false }
    )
])

const stats = computed(() => statsResponse.value?.data || { active_snippets: 0, active_burnt_snippets: 0, total_views: 0 })
const snippets = computed(() => listResponse.value?.data.data || [])
const meta = computed(() => listResponse.value?.data.meta)

const copyLink = (id: string) => {
    const link = `${window.location.origin}/view/${id}`
    navigator.clipboard.writeText(link)
    $toast?.success('Copied to clipboard')
}

const openLink = (id: string) => {
    const link = `/view/${id}`
    window.open(link, '_blank')
}

const deleteSnippet = async (id: string) => {
    if(!confirm("Burn this secret immediately?")) return
    try {
        await $api(`/snippets/${id}`, { method: 'DELETE' })
        $toast?.success('Secret burned')
        refreshList()
    } catch (e) {
        $toast?.error('Failed to delete')
    }
}

// VISUAL HELPERS
// Truncate UUID: "db420be3..." -> "db420..."
const truncateId = (id: string) => id.substring(0, 8) + '...'

// Progress Bar Math
const getProgressWidth = (current: number, max: number) => {
    if (max === 0) return 0
    return (current / max) * 100
}

const formatRelativeTime = (dateStr: string) => {
    const date = new Date(dateStr)
    const now = new Date()
    const diffMs = now.getTime() - date.getTime()
    const diffMins = Math.floor(diffMs / 60000)
    
    if (diffMins < 60) return `${diffMins}m ago`
    const diffHours = Math.floor(diffMs / 3600000)
    if (diffHours < 24) return `${diffHours}h ago`
    return date.toLocaleDateString()
}

const formatExpiresIn = (dateStr: string) => {
    const date = new Date(dateStr)
    const now = new Date()
    const diffMs = date.getTime() - now.getTime()
    if (diffMs <= 0) return 'Expired'
    
    const diffMins = Math.floor(diffMs / 60000)
    const diffHours = Math.floor(diffMs / 3600000)
    if (diffMins < 60) return `${diffMins}m`
    if (diffHours < 24) return `${diffHours}h`
    const diffDays = Math.floor(diffMs / 86400000)
    return `${diffDays}d`
}
</script>

<template>
    <div class="w-full mx-auto px-6 py-12">
        
        <div class="flex flex-col md:flex-row md:items-center justify-between gap-6 mb-12">
            <div>
                <p class="text-3xl font-semibold text-white/40">Manage your <span class="text-primary">Active Secrets</span></p>
            </div>
            <MazBtn to="/" color="primary" left-icon="plus" class="!text-background">
                <template #left-icon><MazFire class="w-5 h-5"/></template>
                New Secret
            </MazBtn>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-12">
            <div class="bg-secondary border border-white/5 p-6 rounded-xl relative overflow-hidden group">
                <div class="absolute right-0 top-0 p-6 opacity-10 group-hover:opacity-20 transition-opacity">
                    <MazFire class="w-16 h-16 text-primary" />
                </div>
                <p class="text-white/40 text-sm font-medium uppercase tracking-wider mb-2">Active Secrets</p>
                <p class="text-4xl font-bold text-white">{{ stats.active_snippets }}</p>
            </div>
            <div class="bg-secondary border border-white/5 p-6 rounded-xl relative overflow-hidden group">
                <div class="absolute right-0 top-0 p-6 opacity-10 group-hover:opacity-20 transition-opacity">
                    <MazEye class="w-16 h-16 text-primary" />
                </div>
                <p class="text-white/40 text-sm font-medium uppercase tracking-wider mb-2">Total Views</p>
                <p class="text-4xl font-bold text-white">{{ stats.total_views }}</p>
            </div>
            <div class="bg-secondary border border-white/5 p-6 rounded-xl relative overflow-hidden group">
                <div class="absolute right-0 top-0 p-6 opacity-10 group-hover:opacity-20 transition-opacity">
                    <MazClock class="w-16 h-16 text-warning" />
                </div>
                <p class="text-white/40 text-sm font-medium uppercase tracking-wider mb-2">Burnt (Active)</p>
                <p class="text-4xl font-bold text-white">{{ stats.active_burnt_snippets }}</p>
            </div>
        </div>

        <div class="bg-secondary border border-white/5 rounded-xl overflow-hidden shadow-2xl">

            <div v-if="snippets.length === 0 && !loadingList" class="p-12 text-center text-white/30">
                <MazFire class="w-12 h-12 mx-auto mb-4 opacity-50" />
                <p>No active secrets found.</p>
            </div>

            <div v-else class="overflow-x-auto">
                <table class="w-full text-left text-sm">
                    <thead class="bg-white/5 text-white/50 uppercase text-xs tracking-wider">
                        <tr>
                            <th class="px-6 py-4 font-medium">ID / Title</th>
                            <th class="px-6 py-4 font-medium">Created</th>
                            <th class="px-6 py-4 font-medium">Expires In</th>
                            <th class="px-6 py-4 font-medium">Views Left</th>
                            <th class="px-6 py-4 font-medium text-right">Actions</th>
                        </tr>
                    </thead>

                    <tbody class="divide-y divide-white/5">
                        <tr v-for="item in snippets" :key="item.id" class="group hover:bg-white/[0.02] transition-colors">
                            
                            <td class="px-6 py-4">
                                <div class="font-medium text-white mb-0.5">{{ item.title || 'Untitled Secret' }}</div>
                                <div class="text-xs font-mono text-white/30 group-hover:text-primary/70 transition-colors">
                                    {{ truncateId(item.id) }}
                                </div>
                            </td>

                            <td class="px-6 py-4 text-white/60 font-mono text-xs">
                                {{ formatRelativeTime(item.created_at) }}
                            </td>

                            <td class="px-6 py-4">
                                <div class="flex items-center gap-2 text-white/80">
                                    <MazClock class="w-4 h-4 text-warning" />
                                    <span class="font-mono text-xs">{{ formatExpiresIn(item.expires_at) }}</span>
                                </div>
                            </td>

                            <td class="px-6 py-4 w-48">
                                <div class="flex items-center gap-3">
                                    <div class="flex-grow h-1.5 bg-white/10 rounded-full overflow-hidden">
                                        <div 
                                            class="h-full bg-primary transition-all duration-500 shadow-[0_0_10px_rgba(251,191,36,0.5)]" 
                                            :style="{ width: `${getProgressWidth(item.current_views, item.max_views)}%` }"
                                        ></div>
                                    </div>
                                    <span class="text-xs font-mono text-white/60 whitespace-nowrap">
                                        <span class="text-white">{{ item.max_views - item.current_views }}</span>/{{ item.max_views }}
                                    </span>
                                </div>
                            </td>

                            <td class="px-6 py-4 text-right">
                                <div class="flex items-center justify-end gap-1 opacity-60 group-hover:opacity-100 transition-opacity">
                                    
                                    <button @click="copyLink(item.id)" class="p-2 hover:text-white hover:bg-white/10 rounded-lg transition-colors" title="Copy Link">
                                        <MazDocumentDuplicate class="w-4 h-4"/>
                                    </button>
                                    
                                    <button @click="openLink(item.id)" class="p-2 hover:text-primary hover:bg-primary/10 rounded-lg transition-colors" title="View Secret">
                                        <MazArrowTopRightOnSquare class="w-4 h-4"/>
                                    </button>

                                    <button @click="deleteSnippet(item.id)" class="p-2 hover:text-danger hover:bg-danger/10 rounded-lg transition-colors" title="Burn Secret">
                                        <MazTrash class="w-4 h-4"/>
                                    </button>
                                </div>
                            </td>

                        </tr>
                    </tbody>
                </table>
            </div>

            <div class="flex justify-center items-center py-4">
                <span class="text-xs lg:text-sm text-white/40" v-if="meta">
                    Page {{ meta.current_page }} of {{ meta.total_pages }}
                </span>
            </div>

            <div class="p-4 border-t border-white/5 flex justify-center gap-2" v-if="meta && meta.total_pages > 1">
                <MazBtn size="sm" color="secondary" :disabled="page <= 1" @click="page--">Previous</MazBtn>
                <span class="px-4 py-1 flex items-center text-sm font-mono bg-black/20 rounded">{{ page }}</span>
                <MazBtn size="sm" color="secondary" :disabled="page >= meta.total_pages" @click="page++">Next</MazBtn>
            </div>
        </div>
    </div>
</template>