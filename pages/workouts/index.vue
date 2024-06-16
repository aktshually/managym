<script setup lang="ts">
import { useStorage } from "@vueuse/core"
import type { Workout } from "~/types/workout"

const workouts = useStorage<Workout[]>("workouts", [])
</script>

<template>
	<main class="w-full h-full">
		<div
			v-if="workouts.length === 0"
			class="flex items-center justify-center h-[85vh]"
		>
			<h2 class="w-1/2 text-center opacity-75">
				You do not have any workouts created on this device. Click on the "+" button to create one
			</h2>
		</div>
		<div
			v-for="workout in workouts"
			v-else
			:key="workout.name"
		>
			{{ workout.name }}
		</div>
		<div>
			<NuxtLink
				class="!px-6 !py-5 !m-0 primary rounded-full absolute bottom-8 right-8"
				to="/workouts/create"
			>
				<Icon
					name="ic:round-plus"
					class="w-6 h-6"
				/>
			</NuxtLink>
		</div>
	</main>
</template>
