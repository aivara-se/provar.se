package invitation

import (
	"time"

	"github.com/automattic/go-gravatar"
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
	Avatar         string         `db:"avatar" json:"avatar"`
}

// Link returns the link to the application page where it is possible to accept
// the invitation
func (i *Invitation) Link() string {
	cfg := config.Get()
	return cfg.Provar.AppURL + "/auth/accept?organizationId=" + i.OrganizationID + "&invitationId=" + i.ID + "&secret=" + i.Secret
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

// Delete deletes an invitation with the given id
func (i *Invitation) Delete() error {
	return DeleteByID(i.ID, i.OrganizationID)
}

// Create creates a new invitation in the database
func Create(orgID, name, email, createdBy string) (*Invitation, error) {
	avatar := gravatar.NewGravatarFromEmail(email)
	avatar.Default = "mp"
	invite := &Invitation{
		ID:             database.NewID(),
		OrganizationID: orgID,
		CreatedAt:      time.Now(),
		CreatedBy:      createdBy,
		ExpiresAt:      time.Now().Add(InvitationTTL),
		Secret:         random.String(64),
		Name:           name,
		Email:          email,
		Avatar:         avatar.GetURL(),
	}
	query := `
		INSERT INTO private.invitation (id, organization_id, created_at, created_by, expires_at, secret, name, email, avatar)
		VALUES (:id, :organization_id, :created_at, :created_by, :expires_at, :secret, :name, :email, :avatar)
	`
	_, err := database.DB().NamedExec(query, invite)
	return invite, err
}

// FindByID returns an invitation by its id
func FindByID(id string, organizationID string) (*Invitation, error) {
	invite := &Invitation{}
	query := `
		SELECT *
		FROM private.invitation
		WHERE
			id = $1 AND
			organization_id = $2
	`
	err := database.DB().Get(invite, query, id, organizationID)
	if err != nil {
		return nil, err
	}
	return invite, nil
}

// FindBySecret returns an invitation by its secret
func FindBySecret(secret string, organizationID string) (*Invitation, error) {
	invite := &Invitation{}
	query := `
		SELECT *
		FROM private.invitation
		WHERE
			secret = $1 AND
			organization_id = $2
	`
	err := database.DB().Get(invite, query, secret, organizationID)
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

// DeleteByID deletes an invitation with the given id
func DeleteByID(id string, organizationID string) error {
	query := `
		DELETE
		FROM private.invitation
		WHERE
			id = $1 AND
			organization_id = $2
	`
	_, err := database.DB().Exec(query, id, organizationID)
	return err
}
