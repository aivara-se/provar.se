/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
	'/auth/email/confirm': {
		post: operations['EmailAuthentication_confirm'];
	};
	'/auth/email/prepare': {
		post: operations['EmailAuthentication_prepare'];
	};
	'/auth/whoami': {
		get: operations['Authentication_whoami'];
	};
	'/organization': {
		post: operations['Organization_create'];
	};
	'/organization/list': {
		get: operations['Organization_list'];
	};
	'/organization/{organizationId}/credential': {
		post: operations['Credential_create'];
	};
	'/organization/{organizationId}/credential/list': {
		get: operations['Organization_credentials'];
	};
	'/organization/{organizationId}/details': {
		get: operations['Organization_details'];
		patch: operations['Organization_updateDetails'];
	};
	'/organization/{organizationId}/invitation/list': {
		get: operations['Organization_invitations'];
	};
	'/organization/{organizationId}/member/list': {
		get: operations['Organization_members'];
	};
	'/organization/{organizationId}/settings': {
		get: operations['Organization_settings'];
	};
	'/ping': {
		get: operations['HealthCheck_basic'];
	};
	'/ping/secure': {
		get: operations['HealthCheck_secure'];
	};
}

export type webhooks = Record<string, never>;

export interface components {
	schemas: {
		CredentialDetails: {
			id: string;
			createdAt: string;
			modifiedAt: string;
			lastUsedAt: string;
			organizationId: string;
			name: string;
			secret: string;
		};
		InvitationDetails: {
			id: string;
			organizationId: string;
			createdAt: string;
			createdBy: string;
			expiresAt: string;
			acceptedAt?: string;
			secret: string;
			name: string;
			email: string;
		};
		OrganizationDetails: {
			id: string;
			createdAt: string;
			createdBy: string;
			modifiedAt: string;
			name: string;
			slug: string;
			description: string;
		};
		UserDetails: {
			id: string;
			createdAt: string;
			modifiedAt: string;
			avatar: string;
			email: string;
			emailVerifiedAt?: string;
			name: string;
		};
	};
	responses: never;
	parameters: never;
	requestBodies: never;
	headers: never;
	pathItems: never;
}

export type $defs = Record<string, never>;

export type external = Record<string, never>;

export interface operations {
	EmailAuthentication_confirm: {
		requestBody: {
			content: {
				'application/json': {
					token: string;
				};
			};
		};
		responses: {
			/** @description The request has succeeded. */
			200: {
				content: {
					'application/json': {
						accessToken: string;
					};
				};
			};
		};
	};
	EmailAuthentication_prepare: {
		requestBody: {
			content: {
				'application/json': {
					email: string;
				};
			};
		};
		responses: {
			/** @description There is no content to send for this request, but the headers may be useful. */
			204: {
				content: never;
			};
		};
	};
	Authentication_whoami: {
		responses: {
			/** @description The request has succeeded. */
			200: {
				content: {
					'application/json': {
						/** @enum {string} */
						type: 'user' | 'credential';
						user?: components['schemas']['UserDetails'];
						credential?: components['schemas']['CredentialDetails'];
					};
				};
			};
		};
	};
	Organization_create: {
		requestBody: {
			content: {
				'application/json': {
					name: string;
					slug: string;
					description: string;
				};
			};
		};
		responses: {
			/** @description The request has succeeded. */
			200: {
				content: {
					'application/json': components['schemas']['OrganizationDetails'];
				};
			};
		};
	};
	Organization_list: {
		responses: {
			/** @description The request has succeeded. */
			200: {
				content: {
					'application/json': components['schemas']['OrganizationDetails'][];
				};
			};
		};
	};
	Credential_create: {
		parameters: {
			path: {
				organizationId: string;
			};
		};
		requestBody: {
			content: {
				'application/json': {
					name: string;
				};
			};
		};
		responses: {
			/** @description The request has succeeded. */
			200: {
				content: {
					'application/json': components['schemas']['OrganizationDetails'];
				};
			};
		};
	};
	Organization_credentials: {
		parameters: {
			path: {
				organizationId: string;
			};
		};
		responses: {
			/** @description The request has succeeded. */
			200: {
				content: {
					'application/json': components['schemas']['CredentialDetails'][];
				};
			};
		};
	};
	Organization_details: {
		parameters: {
			path: {
				organizationId: string;
			};
		};
		responses: {
			/** @description The request has succeeded. */
			200: {
				content: {
					'application/json': components['schemas']['OrganizationDetails'];
				};
			};
		};
	};
	Organization_updateDetails: {
		parameters: {
			path: {
				organizationId: string;
			};
		};
		requestBody: {
			content: {
				'application/json': {
					name?: string;
					slug?: string;
					description?: string;
				};
			};
		};
		responses: {
			/** @description There is no content to send for this request, but the headers may be useful. */
			204: {
				content: never;
			};
		};
	};
	Organization_invitations: {
		parameters: {
			path: {
				organizationId: string;
			};
		};
		responses: {
			/** @description The request has succeeded. */
			200: {
				content: {
					'application/json': components['schemas']['InvitationDetails'][];
				};
			};
		};
	};
	Organization_members: {
		parameters: {
			path: {
				organizationId: string;
			};
		};
		responses: {
			/** @description The request has succeeded. */
			200: {
				content: {
					'application/json': components['schemas']['UserDetails'][];
				};
			};
		};
	};
	Organization_settings: {
		parameters: {
			path: {
				organizationId: string;
			};
		};
		responses: {
			/** @description The request has succeeded. */
			200: {
				content: {
					'application/json': {
						[key: string]: string;
					};
				};
			};
		};
	};
	HealthCheck_basic: {
		responses: {
			/** @description There is no content to send for this request, but the headers may be useful. */
			204: {
				content: never;
			};
		};
	};
	HealthCheck_secure: {
		responses: {
			/** @description There is no content to send for this request, but the headers may be useful. */
			204: {
				content: never;
			};
		};
	};
}
