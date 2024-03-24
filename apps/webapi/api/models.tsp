model UserDetails {
  id: string;
  createdAt: string;
  modifiedAt: string;
  avatar: string;
  email: string;
  emailVerifiedAt?: string;
  name: string;
}

model OrganizationDetails {
  id: string;
  createdAt: string;
  createdBy: string;
  modifiedAt: string;
  name: string;
  slug: string;
  description: string;
};

model CredentialDetails {
  id: string;
  createdAt: string;
  modifiedAt: string;
  lastUsedAt: string;
  organizationId: string;
  name: string;
  secret: string;
}
