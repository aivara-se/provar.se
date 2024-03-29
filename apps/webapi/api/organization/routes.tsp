import "@typespec/http";
import "../models.tsp";

using TypeSpec.Http;

@useAuth(BearerAuth)
interface Organization {
  @post()
  @route("/organization")
  create(@body req: { name: string; slug: string, description: string }): OrganizationDetails;

  @route("/organization/list")
  list(): OrganizationDetails[];

  @route("/organization/{organizationId}/details")
  details(organizationId: string): OrganizationDetails;

  @route("/organization/{organizationId}/settings")
  settings(organizationId: string): Record<string>;

  @route("/organization/{organizationId}/member/list")
  members(organizationId: string): UserDetails[];

  @patch()
  @route("/organization/{organizationId}/details")
  update(organizationId: string, @body req: { name: string; slug: string, description: string }): void;

  @delete()
  @route("/organization/{organizationId}")
  delete(organizationId: string): void;

  @delete()
  @route("/organization/{organizationId}/member/{userId}")
  removeMember(organizationId: string, userId: string): void;
}
