import tailwindcss from "@tailwindcss/vite";

const mazUiConfig = {
  injectComponents: true, // Automatically imports components
  theme: {
    preset: 'maz-ui',
    strategy: 'hybrid', // Recommended: Critical CSS inlined, rest async
    darkClass: 'dark',  // Use 'dark' class to trigger dark mode
    defaultMode: 'dark', // Default to Dark Mode to match your design immediately
    
    overrides: {
      // Font matching your design (Install 'Inter' or 'Fira Code' in main.css)
      fontFamily: {
        base: 'Inter, sans-serif',
        title: 'Inter, sans-serif',
      },
      
      colors: {
        // 🌑 DARK MODE (The "FlashPaper" Look)
        dark: {
          // The "Neon Amber" from the 'Get Started' button
          // Hex approx: #FBBF24 -> HSL: 38 92% 50%
          primary: '38 92% 50%', 
          
          // The Card Backgrounds (slightly lighter than main bg)
          // Hex approx: #18181b -> HSL: 240 5% 10%
          secondary: '240 5% 10%', 
          
          // The Main Page Background (Deep Charcoal)
          // Hex approx: #0f0f11 -> HSL: 240 6% 7%
          background: '240 6% 7%',
          
          // Text Color (Off-white for readability)
          foreground: '240 5% 96%',
          
          // Status Colors (Adjusted for dark mode visibility)
          success: '142 76% 36%', // Green
          warning: '38 92% 50%',  // Amber
          danger: '0 84% 60%',    // Red (Burn color)
          info: '217 91% 60%',    // Blue
        },
        
        // ☀️ LIGHT MODE (Inverted, just in case)
        light: {
          primary: '38 92% 45%',    // Slightly darker amber for contrast on white
          secondary: '240 5% 96%',  // Light gray surface
          background: '0 0% 100%',  // Pure white
          foreground: '240 10% 4%', // Almost black text
          
          success: '142 72% 29%',
          warning: '38 92% 45%',
          danger: '0 72% 51%',
          info: '217 91% 51%',
        },
      },
      
      // You can also round the corners to match the mockup's soft edges
      borderRadius: {
        DEFAULT: '0.5rem', // 8px (Matches the input box and cards)
        lg: '0.75rem',     // 12px
        xl: '1rem',        // 16px
      }
    },
  },
  composables: {
    useToast: true,
  },
}

export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },
  css: ['./app/assets/css/main.css'],
  vite: {
    plugins: [
      tailwindcss(),
    ],
  },
  modules: ['@maz-ui/nuxt', '@pinia/nuxt'],
  mazUi: mazUiConfig,
  runtimeConfig: {
    public: {
      apiBase: "http://localhost:8080"
    }
  },
} as any)