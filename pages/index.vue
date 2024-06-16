<template>
	<main class="h-[85vh] flex items-center">
		<div class="flex flex-col gap-5 items-center w-screen h-full justify-center z-10">
			<div class="text-center">
				<h1>It always feels good to</h1>
				<h1 class="bg-clip-text text-transparent bg-gradient-to-r from-sky-800 to-[#39a5d3]">
					make progress
				</h1>
			</div>
			<p class="text-center w-2/5">
				Managym is not just a way of organizing your workouts. It's a way of unifying workout management, health care & more.
			</p>
			<NuxtLink
				to="/workouts"
				class="primary"
			>
				Create a workout
			</NuxtLink>
		</div>
	</main>
</template>
