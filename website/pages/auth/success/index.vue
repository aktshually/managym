<script setup lang="ts">
const { user } = useUserSession()
const runtimeConfig = useRuntimeConfig()

if (user.value) {
	const [firstName, lastName] = splitGivenName(user.value.name)

	const res = await $fetch.raw(`/users`, {
		method: "POST",
		body: {
			first_name: firstName,
			last_name: lastName,
			email: user.value.email,
			password: "",
		},
		baseURL: runtimeConfig.public.baseURL,
	})
	if (res.status === 201) {
		const router = useRouter()
		router.push("/")
	}
}
</script>

<template>
	<p>Logging in...</p>
</template>
