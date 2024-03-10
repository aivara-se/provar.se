package permission

type PermissionType string

const (
	PermissionTypeCredentialAdmin   PermissionType = "credential.admin"
	PermissionTypeOrganizationAdmin PermissionType = "organization.admin"
	PermissionTypeUserAdmin         PermissionType = "user.admin"
)
