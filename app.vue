<script setup lang="ts">
// import { useThemeSelectorStore } from "~/stores/theme-selector"

// // const theme = useThemeSelectorStore()
// // // useHead({
// // // 	bodyAttrs: {
// // // 		class: theme ? "" : "dark",
// // // 	},
// // // })
</script>

<template>
	<header class="flex items-center w-full h-[15vh] m-0 p-0 drop-shadow-sm px-5">
		<h1 class="flex items-center w-full justify-center text-neutral-950 dark:text-neutral-200">
			Managym
		</h1>
		<ThemeToggle />
	</header>
	<NuxtPage />
</template>
