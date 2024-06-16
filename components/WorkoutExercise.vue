<script setup lang="ts">
import { EQUIPMENT_FOR_BODYWEIGHT_EXERCISES } from "~/const/equipment"
import type { WorkoutExercise } from "~/types/workout"

const { exercise } = defineProps<{
	exercise: WorkoutExercise
}>()

const { images, name, primaryMuscles, secondaryMuscles, sets, equipment } = exercise

const setsTableColumns = [
	{
		key: "id",
		label: "Set",
	},
	{
		key: "repetitions",
		label: "Repetitions",
	},
]

if (!EQUIPMENT_FOR_BODYWEIGHT_EXERCISES.includes(equipment as string)) {
	setsTableColumns.push({
		key: "weight",
		label: "Weight (Kg)",
	})
}

setsTableColumns.push({
	key: "actions",
	label: "",
})

type RowExercise = {
	id: number
	repetitions: number
	weight?: number
}

const setsTableRows = ref<RowExercise[]>([])
for (let i = 0; i < sets.length; i++) {
	const { weight, repetitions } = sets[i]

	const objectToPush: RowExercise = {
		id: i + 1,
		repetitions,
	}

	if (!EQUIPMENT_FOR_BODYWEIGHT_EXERCISES.includes(equipment as string)) {
		objectToPush.weight = weight
	}

	setsTableRows.value.push(objectToPush)
}

const addSet = () => {
	const { repetitions, weight, type } = sets[sets.length - 1]
	setsTableRows.value.push({
		id: setsTableRows.value.length + 1,
		repetitions,
		weight,
	})
	sets.push({
		repetitions,
		weight,
		type,
	})
}

const onChangeSetType = (index: number, newType: "normal" | "warmup" | "drop" | "failure") => {
	sets[index].type = newType
}

const dropdownItems = (row: RowExercise) => [
	[
		{
			label: "Normal",
			click: () => onChangeSetType(row.id - 1, "normal"),
		},
		{
			label: "Warmup",
			click: () => onChangeSetType(row.id - 1, "warmup"),
		},
		{
			label: "Drop set",
			click: () => onChangeSetType(row.id - 1, "drop"),
		},
		{
			label: "Failure set",
			click: () => onChangeSetType(row.id - 1, "failure"),
		},
	],
	[
		{
			label: "Delete",
			icon: "i-heroicons-trash-20-solid",
		},
	],
]

const validateRepetitionsNumber = (reps: number) => {
	if (reps < 1) return 1
	else if (reps > 100) return 100
	return reps
}
</script>

<template>
	<div>
		<div class="flex gap-4 items-center">
			<NuxtImg
				:src="`/exercises/${(images as string[])[1]}`"
				class="rounded-full w-14 h-14"
			/>
			<div>
				<h2>{{ name }}</h2>
				<p class="capitalize text-ellipsis">
					{{ primaryMuscles?.concat(...secondaryMuscles as string[]).join(", ") }}
				</p>
			</div>
		</div>
		<UTable
			:columns="setsTableColumns"
			:rows="setsTableRows"
		>
			<template #id-data="{ row }">
				<p
					v-if="sets[row.id-1].type === 'drop'"
					class="w-6 h-6 text-blue-500 dark:!text-sky-600"
				>
					D
				</p>
				<p
					v-else-if="sets[row.id-1].type === 'failure'"
					class="w-6 h-6 text-red-800 dark:!text-red-500"
				>
					F
				</p>
				<p
					v-else-if="sets[row.id-1].type === 'warmup'"
					class="w-6 h-6 text-yellow-600 dark:!text-yellow-300"
				>
					W
				</p>
				<p
					v-else
				>
					{{ row.id-(sets.slice(0, row.id-1).filter(set => set.type === "warmup").length) }}
				</p>
			</template>

			<template #repetitions-data="{ row }">
				<input
					v-model="sets[row.id-1].repetitions"
					type="number"
					class="minimalist outline-none bg-transparent p-0 m-0 w-12 border-none text-neutral-950 dark:text-neutral-300"
					@change="sets[row.id-1].repetitions = validateRepetitionsNumber(sets[row.id-1].repetitions)"
				>
			</template>

			<template #weight-data="{ row }">
				<input
					v-model="sets[row.id-1].weight"
					type="number"
					class="minimalist outline-none bg-transparent p-0 m-0 w-12 border-none text-neutral-950 dark:text-neutral-300"
				>
			</template>

			<template #actions-data="{ row }">
				<div class="flex items-end justify-end">
					<UDropdown
						:items="dropdownItems(row)"
						class="m-0 p-0 !bg-neutral-300 dark:!bg-neutral-800 hover:bg-sky-800 dark:hover:bg-sky-900 hover:text-neutral-200"
					>
						<button
							class="bg-neutral-300 max-w-fit rounded-md p-1 transition-colors dark:bg-neutral-800 hover:dark:bg-zinc-900"
						>
							<Icon
								name="i-heroicons-ellipsis-horizontal-20-solid"
								class="w-6 h-6 text-neutral-950 dark:text-neutral-300"
							/>
						</button>
					</UDropdown>
				</div>
			</template>
		</UTable>
		<div class="flex flex-col gap-4">
			<button
				class="!p-2 !rounded-md primary w-full flex items-center justify-center gap-2"
				@click="addSet()"
			>
				<Icon name="ic:round-plus" />
				Add set
			</button>
		</div>
	</div>
</template>
