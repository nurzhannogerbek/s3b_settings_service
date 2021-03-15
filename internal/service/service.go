package service

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository"
)

type OrganizationSettings interface {
	Create(organization *common.OrganizationSettings) error
	Delete(organizationID *string) error
	Get(organizationID *string) (*common.OrganizationSettings, error)
	Update(organization *common.OrganizationSettings) (*common.OrganizationSettings, error)
}

type Services struct {
	OrganizationSettings OrganizationSettings
}

type Dependencies struct {
	Repositories *repository.Repositories
}

func NewServices(deps Dependencies) *Services {
	organizationSettingsService := NewOrganizationSettingsService(deps.Repositories.Organization)

	return &Services{
		OrganizationSettings: organizationSettingsService,
	}
}