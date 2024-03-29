package invitation

import (
	"time"

	"provar.se/webapi/lib/database"
	"provar.se/webapi/lib/types"
)

// Invitation struct represents the invitation table in the database
type Invitation struct {
	ID             string         `db:"id" json:"id"`
	OrganizationID string         `db:"organization_id" json:"organizationId"`
	CreatedAt      time.Time      `db:"created_at" json:"createdAt"`
	CreatedBy      string         `db:"created_by" json:"createdBy"`
	ExpiresAt      time.Time      `db:"expires_at" json:"expiresAt"`
	AcceptedAt     types.NullTime `db:"accepted_at" json:"acceptedAt"`
	Secret         string         `db:"secret" json:"secret"`
	Name           string         `db:"name" json:"name"`
	Email          string         `db:"email" json:"email"`
}

// FindPendingByOrganizationID returns all pending invitations for an organization
func FindPendingByOrganizationID(id string) ([]*Invitation, error) {
	invites := []*Invitation{}
	query := `
		SELECT *
		FROM private.invitation
		WHERE
			organization_id = $1 AND
			accepted_at IS NULL AND
			expires_at > NOW()
	`
	err := database.DB().Select(&invites, query, id)
	if err != nil {
		return nil, err
	}
	return invites, nil
}
