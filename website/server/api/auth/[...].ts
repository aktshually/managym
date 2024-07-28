import { NuxtAuthHandler } from "#auth"

const runtimeConfig = useRuntimeConfig()

export default NuxtAuthHandler({
	secret: process.env.NUXT_AUTH_SECRET,
	providers: [
		{
			id: "google",
			name: "Google",
			type: "oauth",
			wellKnown: "https://accounts.google.com/.well-known/openid-configuration",
			authorization: { params: { scope: "openid email profile" } },
			idToken: true,
			checks: ["pkce", "state"],
			profile(profile) {
				return {
					id: profile.sub,
					name: profile.name,
					email: profile.email,
					image: profile.picture,
				}
			},
			clientId: runtimeConfig.google.clientId,
			clientSecret: runtimeConfig.google.clientSecret,
		},
	],
	callbacks: {
		async signIn(_params) {
			console.log("xxxxxx")
			return true
		},
		async session({ user, session, token }) {
			console.log(user)
			console.log(token)

			return session
		},
	},
	session: {
		strategy: "database",
	},
})
