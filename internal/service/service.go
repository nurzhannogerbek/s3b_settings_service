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
	Create(organization *common.OrganizationSettings) error
	Delete(organizationID *string) error
	GetByID(organizationID *string) (*common.OrganizationSettings, error)
	Update(organization *common.OrganizationSettings) (*common.OrganizationSettings, error)
	RestoreDeleted(organizationID *string) error
}

// Channel
// Channel service interface.
type Channel interface {
	Create(channel *common.Channel) error
}

// Services
// Contains all services available in the package.
type Services struct {
	OrganizationSettings OrganizationSettings
	Channel Channel
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

	return &Services{
		OrganizationSettings: organizationSettingsService,
		Channel: channelService,
	}
}
