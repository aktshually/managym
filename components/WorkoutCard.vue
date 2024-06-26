<script setup lang="ts">
import type { WorkoutExercise } from "~/types/workout"

const { name, exercises } = defineProps<{
	name: string
	exercises: WorkoutExercise[]
}>()

const data = new Blob([JSON.stringify({ name, exercises })], { type: "application/json" })
const downloadURL = window.URL.createObjectURL(data)
</script>

<template>
	<div class="border border-diesel-950 flex flex-col gap-4 rounded-lg p-4 dark:border-leaf-100">
		<div>
			<h2>{{ name }}</h2>
			<p class="!opacity-60">
				{{ exercises.map((exercise) => exercise.name).join(", ") }}
			</p>
		</div>
		<div class="flex gap-2">
			<UButton
				color="shark"
				class="py-2"
			>
				Start
			</UButton>
			<UButton
				color="shark"
				class="py-2"
				variant="outline"
				:to="downloadURL"
				:download="`${name}.json`"
			>
				Export
			</UButton>
		</div>
	</div>
</template>
