/** @type {import('tailwindcss').Config} */
export default {
	content: ["*.vue", "./pages/**/*.vue", "./components/**/*.vue"],
	theme: {
		colors: {
			black: "#000000",
			neutral: {
				200: "#E2E8CE",
				300: "#E6E1C5",
				800: "#262626",
				950: "#120309",
			},
			sky: {
				800: "#348AAE",
				900: "#247BA0",
			},
			indigo: {
				300: "#B7C3F3",
			},
		},
		extend: {
			fontFamily: {
				sans: "Urbanist",
			},
		},
	},
	darkMode: "class",
	plugins: [],
}
