import "@typespec/http";
import "../models.tsp";

using TypeSpec.Http;

@route("/auth")
@useAuth(BearerAuth)
interface Authentication {
  @route("whoami")
  whoami(): {
    type: "user" | "credential";
    user?: User;
    credential?: Credential;
  };
}

@route("/auth/email")
interface EmailAuthentication {
  @post()
  @route("prepare")
  prepare(@body req: { email: string }): void;

  @post()
  @route("confirm")
  confirm(@body req: { token: string }): {
    accessToken: string;
  };
}
