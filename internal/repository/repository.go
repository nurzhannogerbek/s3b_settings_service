package repository

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
)

type Organization interface {
	Create(organization *common.OrganizationSettings) error
	Delete(organizationID *string) error
	Get(organizationID *string) (*common.OrganizationSettings, error)
	Update(organization *common.OrganizationSettings) (*common.OrganizationSettings, error)
}

type Repositories struct {
	Organization Organization
}

