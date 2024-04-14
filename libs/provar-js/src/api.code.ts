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

export type OAuth2Authentication = {
	confirm: {
		RequestBody: paths['/auth/oauth2/{provider}/confirm']['post']['requestBody']['content']['application/json'];
		ResponseBody: paths['/auth/oauth2/{provider}/confirm']['post']['responses'][200]['content']['application/json'];
	};
	login: {
		RequestBody: void;
		ResponseBody: paths['/auth/oauth2/{provider}/login']['get']['responses'][204]['content']['application/json'];
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
	details: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}']['get']['responses'][200]['content']['application/json'];
	};
	update: {
		RequestBody: paths['/organization/{organizationId}']['patch']['requestBody']['content']['application/json'];
		ResponseBody: paths['/organization/{organizationId}']['patch']['responses'][204]['content']['application/json'];
	};
	delete: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}']['delete']['responses'][204]['content']['application/json'];
	};
	removeMember: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/member/{userId}']['delete']['responses'][204]['content']['application/json'];
	};
	members: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/members']['get']['responses'][200]['content']['application/json'];
	};
	public: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/public']['get']['responses'][200]['content']['application/json'];
	};
	settings: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/settings']['get']['responses'][200]['content']['application/json'];
	};
	list: {
		RequestBody: void;
		ResponseBody: paths['/organizations']['get']['responses'][200]['content']['application/json'];
	};
};

export type Credential = {
	create: {
		RequestBody: paths['/organization/{organizationId}/credential']['post']['requestBody']['content']['application/json'];
		ResponseBody: paths['/organization/{organizationId}/credential']['post']['responses'][200]['content']['application/json'];
	};
	delete: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/credential/{credentialId}']['delete']['responses'][204]['content']['application/json'];
	};
	list: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/credentials']['get']['responses'][200]['content']['application/json'];
	};
};

export type Feedback = {
	create: {
		RequestBody: paths['/organization/{organizationId}/feedback']['post']['requestBody']['content']['application/json'];
		ResponseBody: paths['/organization/{organizationId}/feedback']['post']['responses'][204]['content']['application/json'];
	};
	details: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/feedback/{feedbackId}']['get']['responses'][200]['content']['application/json'];
	};
	delete: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/feedback/{feedbackId}']['delete']['responses'][204]['content']['application/json'];
	};
	search: {
		RequestBody: paths['/organization/{organizationId}/feedbacks']['post']['requestBody']['content']['application/json'];
		ResponseBody: paths['/organization/{organizationId}/feedbacks']['post']['responses'][200]['content']['application/json'];
	};
	count: {
		RequestBody: paths['/organization/{organizationId}/feedbacks/count']['post']['requestBody']['content']['application/json'];
		ResponseBody: paths['/organization/{organizationId}/feedbacks/count']['post']['responses'][200]['content']['application/json'];
	};
};

export type Invitation = {
	create: {
		RequestBody: paths['/organization/{organizationId}/invitation']['post']['requestBody']['content']['application/json'];
		ResponseBody: paths['/organization/{organizationId}/invitation']['post']['responses'][204]['content']['application/json'];
	};
	details: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/invitation/{invitationId}']['get']['responses'][200]['content']['application/json'];
	};
	delete: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/invitation/{invitationId}']['delete']['responses'][204]['content']['application/json'];
	};
	accept: {
		RequestBody: paths['/organization/{organizationId}/invitation/{invitationId}/accept']['post']['requestBody']['content']['application/json'];
		ResponseBody: paths['/organization/{organizationId}/invitation/{invitationId}/accept']['post']['responses'][204]['content']['application/json'];
	};
	list: {
		RequestBody: void;
		ResponseBody: paths['/organization/{organizationId}/invitations']['get']['responses'][200]['content']['application/json'];
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

export const createOAuth2AuthenticationEndpoints = (f: Fetcher) => {
	return {
		confirm: async (
			provider: string,
			body: OAuth2Authentication['confirm']['RequestBody']
		): Promise<OAuth2Authentication['confirm']['ResponseBody']> => {
			return f.fetch<
				OAuth2Authentication['confirm']['ResponseBody'],
				OAuth2Authentication['confirm']['RequestBody']
			>('POST', `/auth/oauth2/${provider}/confirm`, body);
		},
		login: async (provider: string): Promise<OAuth2Authentication['login']['ResponseBody']> => {
			return f.fetch<OAuth2Authentication['login']['ResponseBody']>(
				'GET',
				`/auth/oauth2/${provider}/login`
			);
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
		details: async (organizationId: string): Promise<Organization['details']['ResponseBody']> => {
			return f.fetch<Organization['details']['ResponseBody']>(
				'GET',
				`/organization/${organizationId}`
			);
		},
		update: async (
			organizationId: string,
			body: Organization['update']['RequestBody']
		): Promise<Organization['update']['ResponseBody']> => {
			return f.fetch<Organization['update']['ResponseBody'], Organization['update']['RequestBody']>(
				'PATCH',
				`/organization/${organizationId}`,
				body
			);
		},
		delete: async (organizationId: string): Promise<Organization['delete']['ResponseBody']> => {
			return f.fetch<Organization['delete']['ResponseBody']>(
				'DELETE',
				`/organization/${organizationId}`
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
		members: async (organizationId: string): Promise<Organization['members']['ResponseBody']> => {
			return f.fetch<Organization['members']['ResponseBody']>(
				'GET',
				`/organization/${organizationId}/members`
			);
		},
		public: async (organizationId: string): Promise<Organization['public']['ResponseBody']> => {
			return f.fetch<Organization['public']['ResponseBody']>(
				'GET',
				`/organization/${organizationId}/public`
			);
		},
		settings: async (organizationId: string): Promise<Organization['settings']['ResponseBody']> => {
			return f.fetch<Organization['settings']['ResponseBody']>(
				'GET',
				`/organization/${organizationId}/settings`
			);
		},
		list: async (): Promise<Organization['list']['ResponseBody']> => {
			return f.fetch<Organization['list']['ResponseBody']>('GET', '/organizations');
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
		delete: async (
			organizationId: string,
			credentialId: string
		): Promise<Credential['delete']['ResponseBody']> => {
			return f.fetch<Credential['delete']['ResponseBody']>(
				'DELETE',
				`/organization/${organizationId}/credential/${credentialId}`
			);
		},
		list: async (organizationId: string): Promise<Credential['list']['ResponseBody']> => {
			return f.fetch<Credential['list']['ResponseBody']>(
				'GET',
				`/organization/${organizationId}/credentials`
			);
		}
	};
};

export const createFeedbackEndpoints = (f: Fetcher) => {
	return {
		create: async (
			organizationId: string,
			body: Feedback['create']['RequestBody']
		): Promise<Feedback['create']['ResponseBody']> => {
			return f.fetch<Feedback['create']['ResponseBody'], Feedback['create']['RequestBody']>(
				'POST',
				`/organization/${organizationId}/feedback`,
				body
			);
		},
		details: async (
			organizationId: string,
			feedbackId: string
		): Promise<Feedback['details']['ResponseBody']> => {
			return f.fetch<Feedback['details']['ResponseBody']>(
				'GET',
				`/organization/${organizationId}/feedback/${feedbackId}`
			);
		},
		delete: async (
			organizationId: string,
			feedbackId: string
		): Promise<Feedback['delete']['ResponseBody']> => {
			return f.fetch<Feedback['delete']['ResponseBody']>(
				'DELETE',
				`/organization/${organizationId}/feedback/${feedbackId}`
			);
		},
		search: async (
			organizationId: string,
			body: Feedback['search']['RequestBody']
		): Promise<Feedback['search']['ResponseBody']> => {
			return f.fetch<Feedback['search']['ResponseBody'], Feedback['search']['RequestBody']>(
				'POST',
				`/organization/${organizationId}/feedbacks`,
				body
			);
		},
		count: async (
			organizationId: string,
			body: Feedback['count']['RequestBody']
		): Promise<Feedback['count']['ResponseBody']> => {
			return f.fetch<Feedback['count']['ResponseBody'], Feedback['count']['RequestBody']>(
				'POST',
				`/organization/${organizationId}/feedbacks/count`,
				body
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
		details: async (
			organizationId: string,
			invitationId: string
		): Promise<Invitation['details']['ResponseBody']> => {
			return f.fetch<Invitation['details']['ResponseBody']>(
				'GET',
				`/organization/${organizationId}/invitation/${invitationId}`
			);
		},
		delete: async (
			organizationId: string,
			invitationId: string
		): Promise<Invitation['delete']['ResponseBody']> => {
			return f.fetch<Invitation['delete']['ResponseBody']>(
				'DELETE',
				`/organization/${organizationId}/invitation/${invitationId}`
			);
		},
		accept: async (
			organizationId: string,
			invitationId: string,
			body: Invitation['accept']['RequestBody']
		): Promise<Invitation['accept']['ResponseBody']> => {
			return f.fetch<Invitation['accept']['ResponseBody'], Invitation['accept']['RequestBody']>(
				'POST',
				`/organization/${organizationId}/invitation/${invitationId}/accept`,
				body
			);
		},
		list: async (organizationId: string): Promise<Invitation['list']['ResponseBody']> => {
			return f.fetch<Invitation['list']['ResponseBody']>(
				'GET',
				`/organization/${organizationId}/invitations`
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
