import type { Exercise } from "./exercise"

export type SetType = "normal" | "warmup" | "drop" | "failure"

export type Set = {
	repetitions: number
	weight?: number
	type: SetType
}

export type WorkoutExercise = Exercise & {
	sets: Set[]
}

export type Workout = {
	name: string
	exercises: WorkoutExercise[]
}
