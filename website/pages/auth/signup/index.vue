<script setup lang="ts">
import type { InferType } from "yup"
import { object, string } from "yup"
import type { FormSubmitEvent } from "#ui/types"

const loginProgress = ref(1)
const user = reactive<{
	email: string
	password: string
	name: string
}>({
	email: "",
	password: "",
	name: "",
})

const strongPasswordValidation = /^(?=.*[A-Z]{1,})(?=.*[!@#$&*]{1,})(?=.*[0-9]{1,}.*)(?=.*[a-z]{1,}).{5,}$/g

const schema = object({
	email: string().email("Invalid email address").required("This field is required").max(100, "Email length must be less than 100 characters"),
	password: string().matches(strongPasswordValidation, "Must contain at least one capital letter, lowercase letter, symbol and number").required("This field is required").max(100, "Name length must be less than 100 characters"),
	name: string().min(4, "This field must be at least 4 characters long").required("This field is required").max(100, "Password length must be less than 100 characters"),
})

type Schema = InferType<typeof schema>

const orderedFieldsToValidate = ["name", "email", "password"]

const runtimeConfig = useRuntimeConfig()

const onSubmit = async (event: FormSubmitEvent<Schema>) => {
	event.preventDefault()

	const toValidate = orderedFieldsToValidate[loginProgress.value - 1] as keyof typeof user

	try {
		await schema.validateAt(toValidate, user)
	}
	catch (e) {
		return
	}

	loginProgress.value++

	if (loginProgress.value === 4) {
		const [firstName, lastName] = splitGivenName(user.name)
		const res = await $fetch.raw(`/users`, {
			method: "POST",
			body: {
				first_name: firstName,
				last_name: lastName,
				email: user.email,
				password: user.password,
			},
			baseURL: runtimeConfig.public.baseURL,
		})
		if (res.status === 201) {
			const router = useRouter()
			router.push("/")
		}
	}
}
</script>

<template>
	<main class="flex items-center justify-center w-screen gap-5 flex-col">
		<div class="w-full my-10">
			<UProgress
				:value="loginProgress"
				size="lg"
				color="shark"
				:max="3"
			/>
			<div class="flex justify-between relative bottom-[1.2rem]">
				<div />
				<div
					v-for="i in 3"
					:key="i"
					class="bg-astral-500 rounded-full w-5 h-5 flex items-center justify-center p-3 dark:bg-astral-600"
				>
					{{ i }}
				</div>
			</div>
		</div>
		<div class="flex items-center flex-col w-full mt-5">
			<UForm
				:schema="schema"
				:state="user"
				class="flex w-full items-center justify-center *:flex *:flex-col *:w-5/6 *:text-center md:*:w-1/2"
			>
				<div
					v-if="loginProgress === 1"
					class="flex flex-col gap-1"
				>
					<UFormGroup
						name="name"
					>
						<div class="flex flex-col gap-4">
							<h1>How do you want to be called?</h1>
							<UInput
								v-model="user.name"
								color="shark"
								placeholder="Ex: John"
								size="lg"
								class="w-full"
								type="text"
								name="name"
								autocomplete="given-name"
							>
								<template #trailing>
									<UButton
										icon="material-symbols-light:arrow-right-alt"
										class="py-[0.6rem] px-2 absolute right-0 rounded-l-none text-sm"
										color="shark"
										variant="solid"
										:padded="false"
										type="submit"
										@click="onSubmit"
									/>
								</template>
							</UInput>
						</div>
					</UFormGroup>
				</div>
				<div
					v-else-if="loginProgress === 2"
					class="flex flex-col gap-1"
				>
					<UFormGroup
						name="email"
					>
						<div class="flex flex-col gap-4">
							<h1>Insert your email address</h1>
							<UInput
								v-model="user.email"
								type="email"
								color="shark"
								placeholder="example@provider.com"
								size="lg"
								class="w-full"
								name="email"
								autocomplete="email"
							>
								<template #trailing>
									<UButton
										icon="material-symbols-light:arrow-right-alt"
										class="py-[0.6rem] px-2 absolute right-0 rounded-l-none text-sm"
										color="shark"
										variant="solid"
										:padded="false"
										@click="onSubmit"
									/>
								</template>
							</UInput>
						</div>
					</UFormGroup>
				</div>
				<div
					v-else
					class="flex flex-col gap-1"
				>
					<UFormGroup
						class="flex flex-col gap-4"
						name="password"
					>
						<div class="flex flex-col gap-4">
							<h1>Insert your password</h1>
							<UInput
								v-model="user.password"
								color="shark"
								placeholder="Use symbols, capital and lowercase letters"
								size="lg"
								class="w-full"
								name="password"
								autocomplete="off"
							>
								<template #trailing>
									<UButton
										icon="material-symbols-light:arrow-right-alt"
										class="py-[0.6rem] px-2 absolute right-0 rounded-l-none text-sm"
										color="shark"
										variant="solid"
										:padded="false"
										:loading="loginProgress === 4"
										@click="onSubmit"
									/>
								</template>
							</UInput>
						</div>
					</UFormGroup>
				</div>
			</UForm>
		</div>
		<div class="flex items-center justify-center w-full gap-2 flex-col">
			<UDivider
				label="OR"
				class="w-5/6 md:w-1/2"
			/>
			<a
				href="/auth/google"
				class="hover:opacity-75 transition-opacity"
				aria-label="Sign in with Google"
			>
				<Icon
					name="material-symbols:g-mobiledata-badge-sharp"
					class="text-red-500 w-6 h-6"
				/>
			</a>
		</div>
	</main>
</template>
