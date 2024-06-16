import type { Exercise } from "./exercise"

type Set = {
	repetitions: number
	weight?: number
	type: "normal" | "warmup" | "drop" | "failure"
}

export type WorkoutExercise = Exercise & {
	sets: Set[]
}

export type Workout = {
	name: string
	exercises: WorkoutExercise[]
}
