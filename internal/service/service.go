/*
 * Package service contains all business logic.
 */

package service

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository"
)

// OrganizationSettings
// Organization settings service interface.
type OrganizationSettings interface {
	CreateOrganizationSettings(organization *common.OrganizationSettings) error
	DeleteOrganizationSettings(organizationID *string) error
	GetOrganizationSettingsByID(organizationID *string) (*common.OrganizationSettings, error)
	UpdateOrganizationSettings(organization *common.OrganizationSettings) (*common.OrganizationSettings, error)
	RestoreDeletedOrganizationSettings(organizationID *string) error
}

// Channel
// Channel service interface.
type Channel interface {
	CreateChannel(channel *common.Channel) (*common.Channel, error)
	UpdateChannel(channel *common.Channel) (*common.Channel, error)
	GetChannels(organizationId *string) (*[]common.Channel, error)
	GetChannel(channelId *string) (*common.Channel, error)
}

// Organization
// Organization service interface.
type Organization interface {
	CreateOrganization(newOrganization common.OrganizationCreateInput) (*common.Organization, error)
	CreateOrganizationDepartment(organization common.OrganizationDepartmentCreateInput) (*common.Organization, error)
	DeleteOrganizations(organizationsIDs *[]string) error
	GetOrganizationByID(organizationID *string) (*common.Organization, error)
	GetOrganizationsByIDs(organizationsIDs *[]string) (*[]common.Organization, error)
	UpdateOrganizationName(organization common.OrganizationNameUpdateInput) (*common.Organization, error)
	RestoreDeletedOrganizations(organizationsIDs *[]string) error
	GetOrganizationDepartmentsByID(parentOrganizationID *string) (*[]common.Organization, error)
	GetAllOrganizationDepartments(rootOrganizationID *string) (*[]common.Organization, error)
	GetArchivedOrganizationDepartmentsByID(parentOrganizationID *string) (*[]common.Organization, error)
	GetAllArchivedOrganizationDepartments(rootOrganizationID *string) (*[]common.Organization, error)
}

// Services
// Contains all services available in the package.
type Services struct {
	OrganizationSettings OrganizationSettings
	Channel              Channel
	Organization         Organization
}

// Dependencies
// Contains dependencies for creating services.
type Dependencies struct {
	Repositories *repository.Repositories
}

// NewServices
// Creates new services based on Dependencies.
func NewServices(d Dependencies) *Services {
	organizationSettingsService := NewOrganizationSettingsService(d.Repositories.OrganizationSettings)
	channelService := NewChannelService(d.Repositories.Channel)
	organizationService := NewOrganizationService(d.Repositories.Organization)

	return &Services{
		OrganizationSettings: organizationSettingsService,
		Channel:              channelService,
		Organization:         organizationService,
	}
}
