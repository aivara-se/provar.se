package invitation

import (
	"time"

	"provar.se/webapi/lib/config"
	"provar.se/webapi/lib/database"
	"provar.se/webapi/lib/random"
	"provar.se/webapi/lib/types"
)

const (
	// InvitationTTL is the time an invitation is valid for
	InvitationTTL = time.Hour * 24 * 7
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

// Create creates a new invitation in the database
func Create(orgID, name, email, createdBy string) (*Invitation, error) {
	invite := &Invitation{
		ID:             database.NewID(),
		OrganizationID: orgID,
		CreatedAt:      time.Now(),
		CreatedBy:      createdBy,
		ExpiresAt:      time.Now().Add(InvitationTTL),
		Secret:         random.String(64),
		Name:           name,
		Email:          email,
	}
	query := `
		INSERT INTO private.invitation (id, organization_id, created_at, created_by, expires_at, secret, name, email)
		VALUES (:id, :organization_id, :created_at, :created_by, :expires_at, :secret, :name, :email)
	`
	_, err := database.DB().NamedExec(query, invite)
	return invite, err
}

// FindBySecret returns an invitation by its secret
func FindBySecret(secret string) (*Invitation, error) {
	invite := &Invitation{}
	query := `SELECT * FROM private.invitation WHERE secret = $1`
	err := database.DB().Get(invite, query, secret)
	if err != nil {
		return nil, err
	}
	return invite, nil
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

// Link returns the link to the application page where it is possible to accept
// the invitation
func (i *Invitation) Link() string {
	cfg := config.Get()
	return cfg.Provar.AppURL + "/auth/invitation/accept/" + i.Secret
}

// IsAcceptable returns true if the invitation can be accepted by the user
func (i *Invitation) IsAcceptable() bool {
	if i.AcceptedAt.Valid {
		return false
	}
	if i.ExpiresAt.Before(time.Now()) {
		return false
	}
	return true
}
