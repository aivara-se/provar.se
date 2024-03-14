import "@typespec/http";

using TypeSpec.Http;

@route("/auth")
@useAuth(BearerAuth)
interface Authentication {
  @route("whoami")
  whoami(): void;
}

@route("/auth/email")
interface EmailAuthentication {
  @post()
  @route("prepare")
  prepare(@body req: { email: string }): void;

  @post()
  @route("confirm")
  confirm(@body req: { token: string }): void;
}
