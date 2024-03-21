import { createResolver } from '@nuxt/kit'
const { resolve } = createResolver(import.meta.url)

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: true,
  css: [
    "~/assets/app.scss",
  ],
  app: {
    head: {
      title: 'Sistem Informasi Kepegawaian',
      link: [
        {
          rel: "stylesheet",
          href: "https://fonts.googleapis.com/css2?family=Oxygen:wght@300;400;700&display=swap",
        },
      ],
    },
  },
  modules: ['@pinia/nuxt', '@nuxt/ui', "nuxt-lodash", '@pinia/nuxt'],
  plugins: ["~/plugins/preline.client.ts"],
  pinia: {
    storesDirs: ['./stores/**'],
  },
  vite: {
    define: {
      "process.env.DEBUG": false,
    },
  },
  ui: {
    notifications: {
      // Show toasts at the top right of the screen
      position: 'top-0 bottom-auto'
    },
    icons: {
      dynamic: true
    }
  },
  lodash: {
    prefix: "_",
    prefixSkip: ["string"],
    upperAfterPrefix: false,
    exclude: ["map"],
    alias: [
      ["camelCase", "stringToCamelCase"], // => stringToCamelCase
      ["kebabCase", "stringToKebab"], // => stringToKebab
      ["isDate", "isLodashDate"], // => _isLodashDate
    ],
  },
  devtools: { enabled: true },
  runtimeConfig: {
    apiSecret: '', // can be overridden by NUXT_API_SECRET environment variable
    public: {
      apiBase: '', // can be overridden by NUXT_PUBLIC_API_BASE environment variable
    }
  },
})