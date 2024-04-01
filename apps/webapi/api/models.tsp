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

model PublicOrganizationDetails {
  id: string;
  name: string;
};

model CredentialDetails {
  id: string;
  createdAt: string;
  createdBy: string;
  modifiedAt: string;
  lastUsedAt: string;
  organizationId: string;
  name: string;
  secret: string;
}

model InvitationDetails {
  id: string;
  organizationId: string;
  createdAt: string;
  createdBy: string;
  expiresAt: string;
  acceptedAt?: string;
  secret: string;
  name: string;
  email: string;
  avatar: string;
}

model PrincipalDetails {
  type: "user" | "credential";
  user?: UserDetails;
  credential?: CredentialDetails;
}

enum FeedbackType {
  text: "text",
  cnps: "cnps",
  csat: "csat",
}

model FeedbackDetails {
  id: string;
  organizationId: string;
  createdAt: string;
  questionType: string;
  feedbackTime: string;
  feedbackType: FeedbackType;
  feedbackData: Record<string>;
  feedbackTags: Record<string>;
  feedbackMeta: Record<string>;
  feedbackUser: Record<string>;
}
