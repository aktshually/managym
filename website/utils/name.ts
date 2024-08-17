export const splitGivenName = (name: string) => {
	const splittedName = name.split(" ") as string[]
	const firstName = splittedName[0]

	let lastName = ""
	if (splittedName.length > 1) {
		lastName = splittedName[splittedName.length - 1]
	}

	return [firstName, lastName]
}
