import "@typespec/http";
import "../models.tsp";

using TypeSpec.Http;

@useAuth(BearerAuth)
interface Invitation {
  @route("/organization/{organizationId}/invitation/list")
  list(organizationId: string): InvitationDetails[];
}
