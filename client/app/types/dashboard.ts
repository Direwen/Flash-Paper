export interface OverviewSnippet {
    id: string
    title: string
    language: string
    max_views: number
    current_views: number
    expires_at: string
    created_at: string
}

export interface DashboardStats {
    active_snippets: number
    active_burnt_snippets: number
    total_views: number
}

export interface PaginatedList<T> {
    data: T[]
    meta: {
        current_page: number
        per_page: number
        total_items: number
        total_pages: number
    }
}