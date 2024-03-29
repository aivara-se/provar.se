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
}
