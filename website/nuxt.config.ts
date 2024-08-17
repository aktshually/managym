// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	devtools: { enabled: true },
	css: ["~/assets/css/main.css"],
	modules: [
		"@pinia/nuxt",
		"@nuxt/eslint",
		"nuxt-icon",
		"@nuxt/image",
		"@nuxt/ui",
		"nuxt-auth-utils",
	],
	postcss: {
		plugins: {
			tailwindcss: {},
			autoprefixer: {},
		},
	},
	eslint: {
		config: {
			stylistic: {
				indent: "tab",
				quotes: "double",
			},
		},
	},
	components: [
		"~/components",
	],
	image: {
		dir: "./assets",
	},
	runtimeConfig: {
		oauth: {
			google: {
				clientId: process.env.NUXT_GOOGLE_CLIENT_ID,
				clientSecret: process.env.NUXT_GOOGLE_CLIENT_SECRET,
			},
		},
		public: {
			baseURL: process.env.API_BASE_URL,
		},
	},
})
