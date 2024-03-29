import "@typespec/http";

using TypeSpec.Http;

interface HealthCheck {
  @route("/ping")
  basic(): void;

  @route("/ping/secure")
  @useAuth(BearerAuth)
  secure(): void;
}
