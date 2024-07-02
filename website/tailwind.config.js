/** @type {import('tailwindcss').Config} */
export default {
	content: ["*.{ts,vue}", "./pages/**/*.vue", "./components/**/*.vue", "./app.config.ts"],
	theme: {
		extend: {
			colors: {
				diesel: {
					950: "#120309",
				},
				shark: {
					800: "#454545",
					950: "#262626",
				},
				astral: {
					500: "#348AAE",
					600: "#247BA0",
				},
				leaf: {
					100: "#E2E8CE",
				},
				brown: {
					100: "#E6E1C5",
					200: "#dbd4ac",
				},
				sky: {
					800: "#348AAE",
					900: "#247BA0",
				},
				indigo: {
					300: "#B7C3F3",
				},
			},
			fontFamily: {
				sans: "Urbanist",
			},
		},
	},
	darkMode: "class",
	plugins: [],
}
