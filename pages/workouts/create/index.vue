<script setup lang="ts">
import { useStorage } from "@vueuse/core"
import type { Exercise } from "~/types/exercise"
import type { Workout } from "~/types/workout"
import exercisesData from "~/public/exercises.json"

const workouts = useStorage<Workout[]>("workouts", [])
let createdWorkoutNumber = 0

if (workouts.value.length > 0) {
	for (let i = 0; i < workouts.value.length; i++) {
		const workout = workouts.value[i]
		if (workout.name.includes("New workout") && !workout.name.includes("#")) {
			createdWorkoutNumber = 1
		}
		else {
			createdWorkoutNumber = Number(workout.name.at(13)) + 1
		}
	}
}

const createdWorkoutName = ref(createdWorkoutNumber > 0 ? `New workout #${createdWorkoutNumber}` : "New workout")
const selectedExercises = ref<Exercise[]>([])
const workout = ref<Workout>({
	name: createdWorkoutName.value,
	exercises: [],
})
const exercises: Exercise[] = exercisesData
const searchQuery = ref("")

for (let i = 0; i <= exercises.length; i++) {
	const exercise = exercises[i]
	if (exercise.name?.includes(searchQuery.value) || exercise.equipment?.includes(searchQuery.value) || exercise.primaryMuscles?.includes(searchQuery.value)) {
		selectedExercises.value.push(exercise)
	}

	if (selectedExercises.value.length === 6) {
		break
	}
}

const onExerciseButtonClick = (exercise: Exercise) => {
	for (let i = 0; i < workout.value.exercises.length; i++) {
		if (workout.value.exercises[i].name === exercise.name) {
			return
		}
	}

	workout.value.exercises.push({
		...exercise,
		sets: [{
			repetitions: 1,
			weight: 0,
			type: "normal",
		}],
	})
}
</script>

<template>
	<main class="w-full flex flex-col-reverse gap-4 md:flex-row md:justify-between pb-8">
		<section
			class="shadow-md md:w-2/3 border flex flex-col gap-8 border-neutral-950 border-opacity-25 rounded-md p-4 dark:border-neutral-200 dark:border-opacity-25"
			:class="workout.exercises.length === 0 ? 'h-[80vh]' : 'h-fit'"
		>
			<div class="flex gap-3 items-center">
				<Icon
					name="mingcute:pencil-line"
					class="text-sky-900 dark:text-sky-900 w-6 h-6"
				/>
				<input
					v-model="createdWorkoutName"
					type="text"
					class="bg-transparent outline-none text-ellipsis text-2xl border-b-2 border-b-transparent font-bold focus:border-b-sky-800 dark:focus:border-b-sky-900"
				>
			</div>
			<div
				v-if="workout.exercises.length === 0"
				class="flex items-center justify-center h-full"
			>
				<p class="text-center !opacity-75">
					You haven't added any exercises yet. Click on any exercise in the search area to add one.
				</p>
			</div>
			<div
				v-for="exercise in workout.exercises"
				v-else
				:key="exercise.name as string"
			>
				<WorkoutExercise :exercise="exercise" />
			</div>
		</section>
		<section class="flex flex-col gap-3 shadow-md max-h-[80vh] md:w-1/3 border border-neutral-950 border-opacity-25 rounded-md p-4 dark:border-neutral-200 dark:border-opacity-25">
			<div class="flex gap-2 items-center border border-sky-800 dark:border-sky-900 rounded-md p-2">
				<Icon
					name="ion:search-outline"
					class="text-sky-900 dark:text-sky-900 w-6 h-6"
				/>
				<input
					v-model="searchQuery"
					type="text"
					class="outline-none bg-transparent"
					placeholder="Try writing an exercise"
				>
			</div>
			<div
				v-for="exercise in selectedExercises"
				:key="exercise.name as string"
				class="max-h-[80vh]"
			>
				<ExerciseCard
					:exercise="exercise"
					:add-to-workout="onExerciseButtonClick"
				/>
			</div>
		</section>
	</main>
</template>
