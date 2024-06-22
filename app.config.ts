export default defineAppConfig({
	ui: {
		icons: {
			dynamic: true,
		},
		notification: {
			background: "bg-brown-200 dark:bg-shark-800",
			title: "font-bold",
			progress: {
				background: "bg-astral-500 text-astral-500 dark:text-astral-600 dark:bg-astral-600",
			},
			icon: {
				base: "border-transparent",
			},
			default: {
				actionButton: {
					color: "gray",
				},
				closeButton: {
					color: "gray",
					variant: "outline",
				},
			},
		},
		button: {
			base: "border text-diesel-950 rounded-lg transition-all dark:text-leaf-100",
			padding: "px-7 py-4",
			color: {
				gray: {
					solid: "bg-brown-200 border-astral-500 dark:border-astral-600 dark:bg-shark-800 hover:brightness-90",
					outline: "border-transparent",
				},
				shark: {
					solid: "bg-astral-500 px-6 text-leaf-100 py-3 border-transparent hover:brightness-90 dark:bg-astral-600",
					outline: "bg-transparent px-6 py-3 border border-astral-500 hover:bg-brown-200 hover:dark:bg-shark-800 dark:border-astral-600",
					ghost: "text-diesel-950 border-none bg-transparent p-1 hover:bg-zinc-400 hover:text-leaf-100 hover:dark:bg-zinc-900 dark:text-leaf-100",
				},
			},
		},
		dropdown: {
			background: "bg-brown-200 dark:bg-shark-800",
			ring: "ring-0",
			item: {
				base: "transition-all",
				active: "brightness-75 bg-brown-200 dark:bg-shark-800",
				icon: "opacity-75 dark:text-leaf-100",
			},

		},
	},
})
