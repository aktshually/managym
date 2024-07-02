<script setup lang="ts">
import type { Exercise } from "~/types/exercise"

const { exercise, addToWorkout } = defineProps<{
	exercise: Exercise
	addToWorkout: (exercise: Exercise) => void
}>()
const { images, name, primaryMuscles, secondaryMuscles } = exercise
</script>

<template>
	<button
		class="flex gap-4 items-center w-full text-left transition-opacity hover:opacity-75"
		@click="addToWorkout(exercise)"
	>
		<div>
			<NuxtImg
				:src="`/exercises/${(images as string[])[1]}`"
				class="rounded-full w-14 h-14"
			/>
		</div>
		<div>
			<h2>{{ name }}</h2>
			<p class="capitalize">
				{{ primaryMuscles?.concat(...secondaryMuscles as string[]).join(", ") }}
			</p>
		</div>
	</button>
</template>
