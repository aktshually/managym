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
			v-else
			class="flex flex-col gap-4 pb-8"
		>
			<WorkoutCard
				v-for="workout in workouts"
				:key="workout.name"
				:name="workout.name"
				:exercises="workout.exercises"
			/>
		</div>
		<div>
			<UButton
				icon="ic:round-plus"
				color="shark"
				variant="solid"
				to="/workouts/create"
				class="!text-leaf-100 fixed bottom-5 right-5"
			/>
		</div>
	</main>
</template>
