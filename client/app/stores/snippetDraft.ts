export const useSnippetDraftStore = defineStore('snippetDraft', () => {
    const draft = ref({
        content: '',
        title: '',
        language: 'text',
        maxViews: 1,
        expiresIn: 60,
    })

    const hasDraft = computed(() => !!draft.value.content)

    function clear() {
        draft.value.content = ''
        draft.value.title = ''
        draft.value.language = 'text'
        draft.value.maxViews = 1
        draft.value.expiresIn = 60
    }

    return { 
        draft, 
        hasDraft, 
        clear 
    }
})
