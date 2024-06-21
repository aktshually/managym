<script setup lang="ts">
import { EQUIPMENT_FOR_BODYWEIGHT_EXERCISES } from "~/const/equipment"
import type { SetType, WorkoutExercise } from "~/types/workout"

const { exercise, remove, undo } = defineProps<{
	exercise: WorkoutExercise
	remove: VoidFunction
	undo: VoidFunction
}>()
const toast = useToast()

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
	type: SetType
}

const setsTableRows = ref<RowExercise[]>([])
for (let i = 0; i < sets.length; i++) {
	const { weight, repetitions, type } = sets[i]

	const objectToPush: RowExercise = {
		id: i,
		repetitions,
		type,
	}

	if (!EQUIPMENT_FOR_BODYWEIGHT_EXERCISES.includes(equipment as string)) {
		objectToPush.weight = weight
	}

	setsTableRows.value.push(objectToPush)
}

const addSet = () => {
	const { repetitions, weight, type } = setsTableRows.value[setsTableRows.value.length - 1]
	setsTableRows.value.push({
		id: setsTableRows.value.length,
		repetitions,
		weight,
		type,
	})
}

const onChangeSetType = (index: number, newType: "normal" | "warmup" | "drop" | "failure") => {
	setsTableRows.value[index].type = newType
}

const deleteSet = (index: number) => {
	setsTableRows.value.splice(index, 1)
	for (let i = 0; i < setsTableRows.value.length; i++) {
		setsTableRows.value[i].id = i
	}

	if (setsTableRows.value.length === 0) {
		deleteExercise()
	}
}

const dropdownItems = (row: RowExercise) => [
	[
		{
			label: "Normal",
			click: () => onChangeSetType(row.id, "normal"),
		},
		{
			label: "Warmup",
			click: () => onChangeSetType(row.id, "warmup"),
		},
		{
			label: "Drop set",
			click: () => onChangeSetType(row.id, "drop"),
		},
		{
			label: "Failure set",
			click: () => onChangeSetType(row.id, "failure"),
		},
	],
	[
		{
			label: "Delete",
			icon: "i-heroicons-trash-20-solid",
			click: () => deleteSet(row.id),
		},
	],
]

const validateRepetitionsNumber = (reps: number) => {
	if (reps < 1) return 1
	else if (reps > 100) return 100
	return reps
}

const deletionNotificationActions = ref([
	{
		label: "Undo",
		click: undo,
	},
])

const deleteExercise = () => {
	remove()
	toast.add({
		title: `"${exercise.name}" was deleted.`,
		actions: deletionNotificationActions.value,
	})
}
</script>

<template>
	<div>
		<div class="flex justify-between items-center">
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
			<div>
				<UButton
					color="shark"
					variant="ghost"
					class="p-2"
					@click="deleteExercise()"
				>
					<Icon
						name="i-heroicons-trash-20-solid"
						class="w-5 h-5 text-diesel-950 dark:text-brown-100"
					/>
				</UButton>
			</div>
		</div>
		<UTable
			:columns="setsTableColumns"
			:rows="setsTableRows"
		>
			<template #id-data="{ row }">
				<p
					v-if="setsTableRows[row.id].type === 'drop'"
					class="w-6 h-6 text-blue-500 dark:!text-sky-600"
				>
					D
				</p>
				<p
					v-else-if="setsTableRows[row.id].type === 'failure'"
					class="w-6 h-6 text-red-800 dark:!text-red-500"
				>
					F
				</p>
				<p
					v-else-if="setsTableRows[row.id].type === 'warmup'"
					class="w-6 h-6 text-yellow-600 dark:!text-yellow-300"
				>
					W
				</p>
				<p
					v-else
				>
					{{ row.id+1-(setsTableRows.slice(0, row.id).filter(set => set.type === "warmup").length) }}
				</p>
			</template>

			<template #repetitions-data="{ row }">
				<input
					v-model="setsTableRows[row.id].repetitions"
					type="number"
					class="minimalist outline-none bg-transparent p-0 m-0 w-12 border-none text-diesel-950 dark:text-brown-100"
					@change="setsTableRows[row.id].repetitions = validateRepetitionsNumber(setsTableRows[row.id].repetitions)"
				>
			</template>

			<template #weight-data="{ row }">
				<input
					v-model="setsTableRows[row.id].weight"
					type="number"
					class="minimalist outline-none bg-transparent p-0 m-0 w-12 border-none text-diesel-950 dark:text-brown-100"
				>
			</template>

			<template #actions-data="{ row }">
				<div class="flex items-end justify-end">
					<UDropdown
						:items="dropdownItems(row)"
						class="m-0 p-0 !bg-brown-100 dark:!bg-shark-950 hover:bg-astral-500 dark:hover:bg-astral-600 hover:text-leaf-100"
					>
						<UButton
							color="shark"
							variant="ghost"
						>
							<Icon
								name="i-heroicons-ellipsis-horizontal-20-solid"
								class="w-6 h-6 text-diesel-950 dark:text-brown-100"
							/>
						</UButton>
					</UDropdown>
				</div>
			</template>
		</UTable>
		<div class="flex flex-col gap-4">
			<UButton
				color="shark"
				variant="solid"
				class="!p-2 !rounded-md primary w-full flex items-center justify-center gap-2"
				@click="addSet()"
			>
				<Icon name="ic:round-plus" />
				Add set
			</UButton>
		</div>
	</div>
</template>
