export const useErrorParser = () => {
    const parseError = (error: any, fallback = "An unexpected error occurred") => {
        // 1. Check if it's a Nuxt/Fetch error with response data
        if (error.response?._data?.error) {
            return error.response._data.error
        }
        
        // 2. Check if it's a standard JS Error
        if (error.message) {
            return error.message
        }

        // 3. Return fallback
        return fallback
    }

    return { parseError }
}