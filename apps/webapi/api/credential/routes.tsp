import "@typespec/http";
import "../models.tsp";

using TypeSpec.Http;

@useAuth(BearerAuth)
interface Credential {
  @route("/organization/{organizationId}/credentials")
  list(organizationId: string): CredentialDetails[];

  @post()
  @route("/organization/{organizationId}/credential")
  create(organizationId: string, @body req: { name: string }): CredentialDetails;

  @delete()
  @route("/organization/{organizationId}/credential/{credentialId}")
  delete(organizationId: string, credentialId: string): void;
}
