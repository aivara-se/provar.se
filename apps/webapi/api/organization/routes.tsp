import "@typespec/http";
import "../models.tsp";

using TypeSpec.Http;

@route("/organization")
@useAuth(BearerAuth)
interface Organization {
  @post()
  create(@body req: { name: string; slug: string, description: string }): OrganizationDetails;

  @route("/list")
  list(): OrganizationDetails[];

  @route("/{id}/details")
  details(id: string): OrganizationDetails;

  @route("/{id}/settings")
  settings(id: string): Record<string>;

  @route("/{id}/member/list")
  members(id: string): UserDetails[];

  @route("/{id}/credential/list")
  credentials(id: string): CredentialDetails[];

  @route("/{id}/invitation/list")
  invitations(id: string): InvitationDetails[];

  @patch()
  @route("/{id}/details")
  updateDetails(id: string, @body req: { name: string; slug: string, description: string }): void;
}
