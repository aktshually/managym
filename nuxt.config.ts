// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	devtools: { enabled: true },
	css: ["~/assets/css/main.css"],
	modules: ["@pinia/nuxt", "@nuxt/eslint", "nuxt-icon", "@nuxt/image", "@nuxt/ui"],
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
		dir: "assets",
	},
})
