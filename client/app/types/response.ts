export type ApiResponse<T> = 
    | { success: true; data: T; message?: never; error?: never }
    | { success: false; error: string; data?: never }