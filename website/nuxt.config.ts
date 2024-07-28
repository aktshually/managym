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
		"@sidebase/nuxt-auth",
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
		google: {
			clientId: process.env.NUXT_GOOGLE_CLIENT_ID,
			clientSecret: process.env.NUXT_GOOGLE_CLIENT_SECRET,
		},
		auth: {
			isEnabled: true,
			disableServerSideAuth: false,
			baseURL: "http://localhost:3000/api/auth",
			provider: {
				type: "authjs",
				trustHost: false,
				defaultProvider: "google",
				addDefaultCallbackUrl: true,
			},
			sessionRefresh: {
				enablePeriodically: true,
				enableOnWindowFocus: true,
			},
		},
	},
})
