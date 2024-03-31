import "@typespec/http";
import "../models.tsp";

using TypeSpec.Http;

@useAuth(BearerAuth)
interface Organization {
  @route("/organizations")
  list(): OrganizationDetails[];

  @post()
  @route("/organization")
  create(@body req: { name: string; slug: string, description: string }): OrganizationDetails;

  @route("/organization/{organizationId}")
  details(organizationId: string): OrganizationDetails;

  @route("/organization/{organizationId}/public")
  public(organizationId: string): PublicOrganizationDetails;

  @patch()
  @route("/organization/{organizationId}")
  update(organizationId: string, @body req: { name: string; slug: string, description: string }): void;

  @delete()
  @route("/organization/{organizationId}")
  delete(organizationId: string): void;

  @route("/organization/{organizationId}/settings")
  settings(organizationId: string): Record<string>;

  @route("/organization/{organizationId}/members")
  members(organizationId: string): UserDetails[];

  @delete()
  @route("/organization/{organizationId}/member/{userId}")
  removeMember(organizationId: string, userId: string): void;
}
