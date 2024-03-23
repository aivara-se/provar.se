model UserDetails {
  id: string;
  email: string;
  name: string;
  avatar: string;
  createdAt: float64;
  modifiedAt: float64;
}

model OrganizationDetails {
  id: string;
  name: string;
  slug: string;
  createdAt: float64;
  createdBy: string;
  modifiedAt: float64;
};

model CredentialDetails {
  id: string;
  organizationId: string;
  name: string;
}
