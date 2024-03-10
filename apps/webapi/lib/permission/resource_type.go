package permission

type ResourceType string

const (
	ResourceTypeCredential   ResourceType = "credential"
	ResourceTypeOrganization ResourceType = "organization"
	ResourceTypeProject      ResourceType = "project"
	ResourceTypeUser         ResourceType = "user"
)
