package permission

import (
	"time"

	"provar.se/webapi/lib/database"
)

// Permission struct represents the permission table in the database
type Permission struct {
	ID             string         `db:"id"`
	CreatedAt      time.Time      `db:"created_at"`
	OrganizationID string         `db:"organization_id"`
	PrincipalType  PrincipalType  `db:"principal_type"`
	Principal      string         `db:"principal"`
	ResourceType   ResourceType   `db:"resource_type"`
	Resources      string         `db:"resources"`
	Permission     PermissionType `db:"permission"`
}

// PermissionQuery struct represents a query to check if a principal has a permission
type PermissionQuery struct {
	OrganizationID string
	PrincipalType  PrincipalType
	PrincipalID    string
	ResourceType   ResourceType
	ResourceID     string
	Permission     PermissionType
}

// HasPermission checks if the principal has the given permission for the given resources
func HasPermission(q PermissionQuery) bool {
	permissions := []*Permission{}
	query := `
		SELECT * FROM private.permission
		WHERE
			(organization_id = $1 OR organization_id = '*') AND
			(principal_type = $2 OR principal_type = '*') AND
			(principal = $3 OR principal = '*') AND
			(resourceType = $4 OR resourceType = '*') AND
			(resources = $5 OR resources = '*') AND
			permission = $6
		LIMIT 1
	`
	err := database.DB().Get(permissions, query, q.OrganizationID, q.PrincipalType, q.PrincipalID, q.ResourceType, q.ResourceID, q.Permission)
	return err != nil
}

// FindByPrincipal returns permissions for the given principal
func FindByPrincipal(principalID string) ([]*Permission, error) {
	permissions := []*Permission{}
	query := "SELECT * FROM private.permission WHERE principal = $1"
	err := database.DB().Get(permissions, query, principalID)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

// DeleteByID deletes a permission with the given id
func DeleteByID(id string) error {
	query := "DELETE FROM private.permission WHERE id = $1"
	_, err := database.DB().Exec(query, id)
	return err
}

// DeleteByID deletes a permission with the given id
func (p *Permission) Delete() error {
	return DeleteByID(p.ID)
}
