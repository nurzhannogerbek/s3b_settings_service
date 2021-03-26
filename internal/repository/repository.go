/*
 * Package repository contains all repositories (database) and their queries.
 */

package repository

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	postgresql "bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository/postgrsql"

	"github.com/jmoiron/sqlx"
)

// OrganizationSettings
// Organization settings repository interface.
type OrganizationSettings interface {
	Create(organization *common.OrganizationSettings) error
	Delete(organizationID *string) error
	GetByID(organizationID *string) (*common.OrganizationSettings, error)
	Update(organization *common.OrganizationSettings) (*common.OrganizationSettings, error)
	RestoreDeleted(organizationID *string) error
}

// FacebookMessengerSettings
// Facebook messenger settings repository interface.
type FacebookMessengerSettings interface {
	Create(facebookMessengerSettings *common.FacebookMessengerSettings) error
}

// Repositories
// Contains all repositories available in the package.
type Repositories struct {
	OrganizationSettings OrganizationSettings
	FacebookMessengerSettings FacebookMessengerSettings
}

// NewRepositories
// Creates new repositories.
func NewRepositories() *Repositories {
	return &Repositories{}
}

// SetPostgresqlRepositories
// Sets postgresql repositories in Repositories struct.
func (r *Repositories) SetPostgresqlRepositories(db *sqlx.DB) {
	r.OrganizationSettings = postgresql.NewOrganizationSettingsRepository(db)
	r.FacebookMessengerSettings = postgresql.NewFacebookMessengerSettingsRepository(db)
}
