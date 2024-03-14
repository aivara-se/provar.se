import "@typespec/http";

using TypeSpec.Http;

@route("/ping")
interface HealthCheck {
  @route("/")
  basic(): void;

  @route("/secure")
  @useAuth(BearerAuth)
  secure(): void;
}
