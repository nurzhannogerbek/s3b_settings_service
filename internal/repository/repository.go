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

// Channel
// Channel repository interface.
type Channel interface {
	CreateChannel(channel *common.Channel) error
	GetChannel(channelId *string) (*common.Channel, error)
	GetChannels(rootOrganizationId *string) (*[]common.Channel, error)
}

// Organization
// Organization repository interface.
type Organization interface {
	CreateOrganization(organization *common.OrganizationCreateInput) (*common.Organization, error)
	CreateOrganizationDepartment(department *common.OrganizationCreateDepartmentInput) (*common.Organization, error)
	DeleteOrganizations(organizationsIDs *[]string) error
	GetOrganizationByID(organizationID *string) (*common.Organization, error)
	GetOrganizationsByIDs(organizationsIDs *[]string) (*[]common.Organization, error)
	UpdateOrganization(organization *common.OrganizationUpdateInput) (*common.Organization, error)
	RestoreDeletedOrganizations(organizationsIDs *[]string) error
	GetOrganizationDepartmentsByID(parentOrganizationID *string) (*[]common.Organization, error)
	GetAllOrganizationDepartments(rootOrganizationID *string) (*[]common.Organization, error)
	GetArchivedOrganizationDepartmentsByID(parentOrganizationID *string) (*[]common.Organization, error)
	GetAllArchivedOrganizationDepartments(rootOrganizationID *string) (*[]common.Organization, error)
}

// Repositories
// Contains all repositories available in the package.
type Repositories struct {
	OrganizationSettings OrganizationSettings
	Channel              Channel
	Organization         Organization
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
	r.Channel = postgresql.NewChannelRepository(db)
	r.Organization = postgresql.NewOrganizationRepository(db)
}
