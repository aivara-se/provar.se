import "@typespec/http";

using TypeSpec.Http;

model OrganizationDetails {
  id: string;
  name: string;
  slug: string;
};

@route("/organization")
@useAuth(BearerAuth)
interface Organization {
  @post()
  create(@body req: { name: string; slug: string }): OrganizationDetails;
}
