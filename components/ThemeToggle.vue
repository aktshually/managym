<script setup lang="ts">
import { useStorage } from "@vueuse/core"
import { useThemeSelectorStore } from "~/stores/theme-selector"

const theme = useThemeSelectorStore()
const onClick = () => {
	theme.toggle()
	useStorage("theme", theme.get())
	document.body.className = theme.get() === 1 ? "" : "dark"
}
</script>

<template>
	<Icon
		v-if="theme.get() === 1"
		name="iconamoon:mode-dark-thin"
		class="cursor-pointer w-6 h-6"
		@click="onClick"
	/>
	<Icon
		v-else
		name="material-symbols-light:light-mode-outline"
		class="cursor-pointer w-6 h-6 text-neutral-200"
		@click="onClick"
	/>
</template>
