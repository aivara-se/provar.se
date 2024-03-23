model User {
  id: string;
  email: string;
  name: string;
  avatar: string;
}

model Credential {
  id: string;
  organizationId: string;
  name: string;
}
