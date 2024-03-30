import "@typespec/http";
import "../models.tsp";

using TypeSpec.Http;

@useAuth(BearerAuth)
interface Invitation {
  @post()
  @route("/organization/{organizationId}/invitation")
  create(organizationId: string, @body req: { name: string; email: string }): void;

  @route("/organization/{organizationId}/invitation/list")
  list(organizationId: string): InvitationDetails[];

  @route("/organization/{organizationId}/invitation/{invitationId}")
  details(organizationId: string, invitationId: string): InvitationDetails;

  @post()
  @route("/organization/{organizationId}/invitation/{invitationId}/accept")
  accept(organizationId: string, invitationId: string, @body req: { secret: string }): void;

  @delete()
  @route("/organization/{organizationId}/invitation/{invitationId}")
  delete(organizationId: string, invitationId: string): void;
}
