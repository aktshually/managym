const runtimeConfig = useRuntimeConfig()

export default oauthGoogleEventHandler({
	config: {
		clientId: runtimeConfig.oauth.google.clientId,
		clientSecret: runtimeConfig.oauth.google.clientSecret,
		scope: ["email", "openid", "profile"],
	},
	async onSuccess(event, { user }) {
		await setUserSession(event, {
			user: {
				id: user.sub,
				name: user.name,
				email: user.email,
				image: user.picture,
			},
		})
		return sendRedirect(event, "/auth/success")
	},
	onError(event, error) {
		console.error("Google OAuth error:", error)
		return sendRedirect(event, "/")
	},
})
