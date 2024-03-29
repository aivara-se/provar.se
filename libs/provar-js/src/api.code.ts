import type { paths } from './api.type.ts';
import type { Fetcher } from './fetcher.ts';

export type EmailAuthentication = {
	confirm: {
		RequestBody: paths['/auth/email/confirm']['post']['requestBody']['content']['application/json'];
		ResponseBody: paths['/auth/email/confirm']['post']['responses'][200]['content']['application/json'];
	};
	prepare: {
		RequestBody: paths['/auth/email/prepare']['post']['requestBody']['content']['application/json'];
		ResponseBody: paths['/auth/email/prepare']['post']['responses'][204]['content']['application/json'];
	};
};

export type Authentication = {
	whoami: {
		RequestBody: void;
		ResponseBody: paths['/auth/whoami']['get']['responses'][200]['content']['application/json'];
	};
};

export type Organization = {
	create: {
		RequestBody: paths['/organization']['post']['requestBody']['content']['application/json'];
		ResponseBody: paths['/organization']['post']['responses'][200]['content']['application/json'];
	};
	list: {
		RequestBody: void;
		ResponseBody: paths['/organization/list']['get']['responses'][200]['content']['application/json'];
	};
	delete: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}']['delete']['responses'][204]['content']['application/json'];
	};
	details: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/details']['get']['responses'][200]['content']['application/json'];
	};
	update: {
		RequestBody: paths['/organization/{organizationId}/details']['patch']['requestBody']['content']['application/json'];
		ResponseBody: paths['/organization/{organizationId}/details']['patch']['responses'][204]['content']['application/json'];
	};
	members: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/member/list']['get']['responses'][200]['content']['application/json'];
	};
	removeMember: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/member/{userId}']['delete']['responses'][204]['content']['application/json'];
	};
	settings: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/settings']['get']['responses'][200]['content']['application/json'];
	};
};

export type Credential = {
	create: {
		RequestBody: paths['/organization/{organizationId}/credential']['post']['requestBody']['content']['application/json'];
		ResponseBody: paths['/organization/{organizationId}/credential']['post']['responses'][200]['content']['application/json'];
	};
	list: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/credential/list']['get']['responses'][200]['content']['application/json'];
	};
};

export type Invitation = {
	create: {
		RequestBody: paths['/organization/{organizationId}/invitation']['post']['requestBody']['content']['application/json'];
		ResponseBody: paths['/organization/{organizationId}/invitation']['post']['responses'][204]['content']['application/json'];
	};
	list: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/invitation/list']['get']['responses'][200]['content']['application/json'];
	};
	details: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/invitation/{secret}']['get']['responses'][200]['content']['application/json'];
	};
	accept: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/invitation/{secret}/accept']['post']['responses'][204]['content']['application/json'];
	};
};

export type HealthCheck = {
	basic: {
		RequestBody: void;
		ResponseBody: paths['/ping']['get']['responses'][204]['content']['application/json'];
	};
	secure: {
		RequestBody: void;
		ResponseBody: paths['/ping/secure']['get']['responses'][204]['content']['application/json'];
	};
};

export const createEmailAuthenticationEndpoints = (f: Fetcher) => {
	return {
		confirm: async (
			body: EmailAuthentication['confirm']['RequestBody']
		): Promise<EmailAuthentication['confirm']['ResponseBody']> => {
			return f.fetch<
				EmailAuthentication['confirm']['ResponseBody'],
				EmailAuthentication['confirm']['RequestBody']
			>('POST', '/auth/email/confirm', body);
		},
		prepare: async (
			body: EmailAuthentication['prepare']['RequestBody']
		): Promise<EmailAuthentication['prepare']['ResponseBody']> => {
			return f.fetch<
				EmailAuthentication['prepare']['ResponseBody'],
				EmailAuthentication['prepare']['RequestBody']
			>('POST', '/auth/email/prepare', body);
		}
	};
};

export const createAuthenticationEndpoints = (f: Fetcher) => {
	return {
		whoami: async (): Promise<Authentication['whoami']['ResponseBody']> => {
			return f.fetch<Authentication['whoami']['ResponseBody']>('GET', '/auth/whoami');
		}
	};
};

export const createOrganizationEndpoints = (f: Fetcher) => {
	return {
		create: async (
			body: Organization['create']['RequestBody']
		): Promise<Organization['create']['ResponseBody']> => {
			return f.fetch<Organization['create']['ResponseBody'], Organization['create']['RequestBody']>(
				'POST',
				'/organization',
				body
			);
		},
		list: async (): Promise<Organization['list']['ResponseBody']> => {
			return f.fetch<Organization['list']['ResponseBody']>('GET', '/organization/list');
		},
		delete: async (organizationId: string): Promise<Organization['delete']['ResponseBody']> => {
			return f.fetch<Organization['delete']['ResponseBody']>(
				'DELETE',
				`/organization/${organizationId}`
			);
		},
		details: async (organizationId: string): Promise<Organization['details']['ResponseBody']> => {
			return f.fetch<Organization['details']['ResponseBody']>(
				'GET',
				`/organization/${organizationId}/details`
			);
		},
		update: async (
			organizationId: string,
			body: Organization['update']['RequestBody']
		): Promise<Organization['update']['ResponseBody']> => {
			return f.fetch<Organization['update']['ResponseBody'], Organization['update']['RequestBody']>(
				'PATCH',
				`/organization/${organizationId}/details`,
				body
			);
		},
		members: async (organizationId: string): Promise<Organization['members']['ResponseBody']> => {
			return f.fetch<Organization['members']['ResponseBody']>(
				'GET',
				`/organization/${organizationId}/member/list`
			);
		},
		removeMember: async (
			organizationId: string,
			userId: string
		): Promise<Organization['removeMember']['ResponseBody']> => {
			return f.fetch<Organization['removeMember']['ResponseBody']>(
				'DELETE',
				`/organization/${organizationId}/member/${userId}`
			);
		},
		settings: async (organizationId: string): Promise<Organization['settings']['ResponseBody']> => {
			return f.fetch<Organization['settings']['ResponseBody']>(
				'GET',
				`/organization/${organizationId}/settings`
			);
		}
	};
};

export const createCredentialEndpoints = (f: Fetcher) => {
	return {
		create: async (
			organizationId: string,
			body: Credential['create']['RequestBody']
		): Promise<Credential['create']['ResponseBody']> => {
			return f.fetch<Credential['create']['ResponseBody'], Credential['create']['RequestBody']>(
				'POST',
				`/organization/${organizationId}/credential`,
				body
			);
		},
		list: async (organizationId: string): Promise<Credential['list']['ResponseBody']> => {
			return f.fetch<Credential['list']['ResponseBody']>(
				'GET',
				`/organization/${organizationId}/credential/list`
			);
		}
	};
};

export const createInvitationEndpoints = (f: Fetcher) => {
	return {
		create: async (
			organizationId: string,
			body: Invitation['create']['RequestBody']
		): Promise<Invitation['create']['ResponseBody']> => {
			return f.fetch<Invitation['create']['ResponseBody'], Invitation['create']['RequestBody']>(
				'POST',
				`/organization/${organizationId}/invitation`,
				body
			);
		},
		list: async (organizationId: string): Promise<Invitation['list']['ResponseBody']> => {
			return f.fetch<Invitation['list']['ResponseBody']>(
				'GET',
				`/organization/${organizationId}/invitation/list`
			);
		},
		details: async (
			organizationId: string,
			secret: string
		): Promise<Invitation['details']['ResponseBody']> => {
			return f.fetch<Invitation['details']['ResponseBody']>(
				'GET',
				`/organization/${organizationId}/invitation/${secret}`
			);
		},
		accept: async (
			organizationId: string,
			secret: string
		): Promise<Invitation['accept']['ResponseBody']> => {
			return f.fetch<Invitation['accept']['ResponseBody']>(
				'POST',
				`/organization/${organizationId}/invitation/${secret}/accept`
			);
		}
	};
};

export const createHealthCheckEndpoints = (f: Fetcher) => {
	return {
		basic: async (): Promise<HealthCheck['basic']['ResponseBody']> => {
			return f.fetch<HealthCheck['basic']['ResponseBody']>('GET', '/ping');
		},
		secure: async (): Promise<HealthCheck['secure']['ResponseBody']> => {
			return f.fetch<HealthCheck['secure']['ResponseBody']>('GET', '/ping/secure');
		}
	};
};
