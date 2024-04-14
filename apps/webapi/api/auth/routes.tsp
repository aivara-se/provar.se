import "@typespec/http";
import "../models.tsp";

using TypeSpec.Http;

model OAuth2Confirmation {
  state: string;
  code: string;
}

@useAuth(BearerAuth)
interface Authentication {
  @route("/auth/whoami")
  whoami(): PrincipalDetails;
}

interface EmailAuthentication {
  @post()
  @route("/auth/email/prepare")
  prepare(@body req: { email: string }): void;

  @post()
  @route("/auth/email/confirm")
  confirm(@body req: { token: string }): { accessToken: string; };
}

interface OAuth2Authentication {
  @route("/auth/oauth2/{provider}/login")
  login(provider: string): void;

  @post()
  @route("/auth/oauth2/{provider}/confirm")
  confirm(provider: string, @body req: OAuth2Confirmation): { accessToken: string; };
}
