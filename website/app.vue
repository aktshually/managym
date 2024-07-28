<script setup lang="ts">
const { data } = useAuth()
</script>

<template>
	<header class="flex items-center w-full h-[15vh] m-0 p-0 drop-shadow-sm px-5">
		<NuxtImg
			v-if="data?.user?.image"
			class="rounded-full"
			:src="data?.user?.image"
		/>
		<h1 class="flex items-center w-full justify-center text-diesel-950 dark:text-leaf-100">
			Managym
		</h1>
		<ThemeToggle />
	</header>
	<NuxtPage />
	<UNotifications />
</template>
