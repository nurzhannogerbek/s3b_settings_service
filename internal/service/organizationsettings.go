package service

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository"
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/tool/uuid"
)

type OrganizationSettingsService struct {
	repo repository.Organization
}

func NewOrganizationSettingsService(repo repository.Organization) *OrganizationSettingsService {
	return &OrganizationSettingsService{
		repo: repo,
	}
}

func (orgSetSer *OrganizationSettingsService) Create(organizationSettings *common.OrganizationSettings) error {
	if err := organizationSettings.Validate(); err != nil {
		return err
	}

	if err := orgSetSer.repo.Create(organizationSettings); err != nil {
		return err
	}

	return nil
}

func (orgSetSer *OrganizationSettingsService) Delete(organizationID *string) error {
	if err := uuid.Validate(organizationID); err != nil {
		return err
	}

	if err := orgSetSer.repo.Delete(organizationID); err != nil {
		return err
	}

	return nil
}

func (orgSetSer *OrganizationSettingsService) Get(organizationID *string) (*common.OrganizationSettings, error) {
	if err := uuid.Validate(organizationID); err != nil {
		return nil, err
	}

	organizationSettings, err := orgSetSer.repo.Get(organizationID)
	if err != nil {
		return nil, err
	}
	return organizationSettings, nil
}

func (orgSetSer *OrganizationSettingsService) Update(newOrganizationSettings *common.OrganizationSettings) (*common.OrganizationSettings, error) {
	if err := newOrganizationSettings.Validate(); err != nil {
		return nil, err
	}

	organizationSettings, err := orgSetSer.repo.Update(newOrganizationSettings)
	if err != nil {
		return nil, err
	}

	return organizationSettings, nil
}