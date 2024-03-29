import "@typespec/http";
import "../models.tsp";

using TypeSpec.Http;

@useAuth(BearerAuth)
interface Authentication {
  @route("/auth/whoami")
  whoami(): {
    type: "user" | "credential";
    user?: UserDetails;
    credential?: CredentialDetails;
  };
}

interface EmailAuthentication {
  @post()
  @route("/auth/email/prepare")
  prepare(@body req: { email: string }): void;

  @post()
  @route("/auth/email/confirm")
  confirm(@body req: { token: string }): {
    accessToken: string;
  };
}
