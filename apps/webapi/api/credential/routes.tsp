import "@typespec/http";
import "../models.tsp";

using TypeSpec.Http;

@useAuth(BearerAuth)
interface Credential {
  @post()
  @route("/organization/{organizationId}/credential")
  create(organizationId: string, @body req: { name: string }): CredentialDetails;

  @route("/organization/{organizationId}/credential/list")
  list(organizationId: string): CredentialDetails[];
}
