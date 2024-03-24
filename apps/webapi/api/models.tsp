model UserDetails {
  id: string;
  createdAt: float64;
  modifiedAt: float64;
  avatar: string;
  email: string;
  emailVerifiedAt?: float64;
  name: string;
}

model OrganizationDetails {
  id: string;
  createdAt: float64;
  createdBy: string;
  modifiedAt: float64;
  name: string;
  slug: string;
  description: string;
};

model CredentialDetails {
  id: string;
  createdAt: float64;
  modifiedAt: float64;
  lastUsedAt: float64;
  organizationId: string;
  name: string;
  secret: string;
}
