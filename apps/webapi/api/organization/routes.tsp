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
}
