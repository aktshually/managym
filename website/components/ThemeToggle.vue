<script setup lang="ts">
import { useToggle } from "@vueuse/shared"
import { useStorage } from "@vueuse/core"

const toggleDark = useToggle(isDark)
isDark.value = useStorage("vueuse-color-scheme", "auto").value === "auto"
</script>

<template>
	<Icon
		v-if="!isDark"
		name="iconamoon:mode-dark-thin"
		class="cursor-pointer w-6 h-6 text-diesel-950"
		@click="toggleDark()"
	/>
	<Icon
		v-else
		name="material-symbols-light:light-mode-outline"
		class="cursor-pointer w-6 h-6 dark:text-leaf-100"
		@click="toggleDark()"
	/>
</template>
