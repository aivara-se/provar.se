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
			>('post', '/auth/email/confirm', body);
		},
		prepare: async (
			body: EmailAuthentication['prepare']['RequestBody']
		): Promise<EmailAuthentication['prepare']['ResponseBody']> => {
			return f.fetch<
				EmailAuthentication['prepare']['ResponseBody'],
				EmailAuthentication['prepare']['RequestBody']
			>('post', '/auth/email/prepare', body);
		}
	};
};

export const createAuthenticationEndpoints = (f: Fetcher) => {
	return {
		whoami: async (): Promise<Authentication['whoami']['ResponseBody']> => {
			return f.fetch<Authentication['whoami']['ResponseBody']>('get', '/auth/whoami');
		}
	};
};

export const createOrganizationEndpoints = (f: Fetcher) => {
	return {
		create: async (
			body: Organization['create']['RequestBody']
		): Promise<Organization['create']['ResponseBody']> => {
			return f.fetch<Organization['create']['ResponseBody'], Organization['create']['RequestBody']>(
				'post',
				'/organization',
				body
			);
		}
	};
};

export const createHealthCheckEndpoints = (f: Fetcher) => {
	return {
		basic: async (): Promise<HealthCheck['basic']['ResponseBody']> => {
			return f.fetch<HealthCheck['basic']['ResponseBody']>('get', '/ping');
		},
		secure: async (): Promise<HealthCheck['secure']['ResponseBody']> => {
			return f.fetch<HealthCheck['secure']['ResponseBody']>('get', '/ping/secure');
		}
	};
};
