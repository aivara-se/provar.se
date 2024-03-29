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

  @route("/organization/{organizationId}/invitation/{secret}")
  details(organizationId: string, secret: string): InvitationDetails;

  @post()
  @route("/organization/{organizationId}/invitation/{secret}/accept")
  accept(organizationId: string, secret: string): void;
}
