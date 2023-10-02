export async function getOrganizations() {
	return [
		{ id: 'org-a', name: 'Normative', environment: 'production' },
		{ id: 'org-b', name: 'Normative - dev', environment: 'staging' }
	];
}
