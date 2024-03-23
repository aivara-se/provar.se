import "@typespec/http";

using TypeSpec.Http;

model OrganizationDetails {
  id: string;
  name: string;
  slug: string;
  createdAt: float64;
  createdBy: string;
  modifiedAt: float64;
};

@route("/organization")
@useAuth(BearerAuth)
interface Organization {
  @post()
  create(@body req: { name: string; slug: string }): OrganizationDetails;

  @route("/list")
  list(): OrganizationDetails[];
}
