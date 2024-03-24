import "@typespec/http";
import "../models.tsp";

using TypeSpec.Http;

@route("/organization")
@useAuth(BearerAuth)
interface Organization {
  @post()
  create(@body req: { name: string; slug: string }): OrganizationDetails;

  @route("/list")
  list(): OrganizationDetails[];

  @route("/{id}/details")
  details(id: string): OrganizationDetails;

  @route("/{id}/settings")
  settings(id: string): Record<string>;
}
