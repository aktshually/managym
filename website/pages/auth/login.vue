<script setup lang="ts">
const { signIn, signOut, status, data } = useAuth()
console.log(data.value)

if (data.value) {
	const splittedName = data.value.user?.name?.split(" ") as string[]
	const firstName = splittedName[0]

	let lastName = ""
	if (splittedName.length > 0) {
		lastName = splittedName[-1]
	}

	await useFetch("http://localhost:8080/users", {
		method: "POST",
		body: {
			first_name: firstName,
			last_name: lastName,
			password: data.value.user.id,
		},
	})
}
</script>

<template>
	<div>
		<button @click="signIn('google')">
			Sign in
		</button>
		<button @click="signOut()">
			Sign out
		</button>
		<p>{{ status }}</p>
		<p>{{ data?.user?.name }}</p>
		<p>{{ data?.user?.email }}</p>
		<NuxtImg :src="data?.user?.image || ''" />
	</div>
</template>
