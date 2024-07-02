<script setup lang="ts">
const duration = ref({
	hours: 0,
	minutes: 0,
	seconds: 0,
})

onMounted(() => {
	setInterval(() => {
		duration.value.seconds++
		if (duration.value.seconds === 60) {
			duration.value.seconds = 0
			duration.value.minutes++

			if (duration.value.minutes === 60) {
				duration.value.minutes = 0
				duration.value.hours++
			}
		}
	}, 1000)
})
</script>

<template>
	<div>
		<time>
			{{ `${duration.hours > 0 ? duration.hours.toString().padStart(2, "0") + ":" : ""}${duration.minutes.toString().padStart(2, "0")}:${duration.seconds.toString().padStart(2, "0")}` }}
		</time>
		<UProgress :value="20" />
	</div>
</template>
