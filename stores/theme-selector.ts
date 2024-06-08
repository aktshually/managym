import { defineStore } from "pinia"
import { useStorage } from "@vueuse/core"

export const useThemeSelectorStore = defineStore("themeSelector", () => {
	const theme = useStorage("theme", 1)
	if (!theme.value) {
		localStorage.setItem("theme", "1")
	}

	const toggle = () => {
		theme.value = !theme.value ? 1 : 0
	}

	const get = () => {
		return theme.value
	}

	return { theme, toggle, get }
})
