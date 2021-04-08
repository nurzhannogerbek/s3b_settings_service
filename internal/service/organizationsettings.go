package service

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository"
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/tool/uuid"
)

// OrganizationSettingsService
// Contains all dependencies for organization settings service.
type OrganizationSettingsService struct {
	repository repository.OrganizationSettings
}

// NewOrganizationSettingsService
// Creates new organization settings service (OrganizationSettings).
func NewOrganizationSettingsService(os repository.OrganizationSettings) *OrganizationSettingsService {
	return &OrganizationSettingsService{
		repository: os,
	}
}

// CreateOrganizationSettings
// Validates and Creates new organization settings record in database.
func (oss *OrganizationSettingsService) CreateOrganizationSettings(os *common.OrganizationSettings) error {
	if err := os.Validate(); err != nil {
		return err
	}

	if err := oss.repository.CreateOrganizationSettings(*os); err != nil {
		return err
	}

	return nil
}

// DeleteOrganizationSettings
// Validates and Deletes organization settings record by ID in database.
func (oss *OrganizationSettingsService) DeleteOrganizationSettings(organizationID *string) error {
	if err := uuid.Validate(organizationID); err != nil {
		return err
	}

	if err := oss.repository.DeleteOrganizationSettings(*organizationID); err != nil {
		return err
	}

	return nil
}

// GetOrganizationSettingsById
// Validates and Queries organization settings record by ID from database.
func (oss *OrganizationSettingsService) GetOrganizationSettingsByID(organizationID *string) (*common.OrganizationSettings, error) {
	if err := uuid.Validate(organizationID); err != nil {
		return nil, err
	}

	organizationSettings, err := oss.repository.GetOrganizationSettingsByID(*organizationID)
	if err != nil {
		return nil, err
	}

	return organizationSettings, nil
}

// UpdateOrganizationSettings
// Validates and updates organization settings record by ID in database.
func (oss *OrganizationSettingsService) UpdateOrganizationSettings(os *common.OrganizationSettings) (*common.OrganizationSettings, error) {
	if err := os.Validate(); err != nil {
		return nil, err
	}

	organizationSettings, err := oss.repository.UpdateOrganizationSettings(*os)
	if err != nil {
		return nil, err
	}

	return organizationSettings, nil
}

// RestoreDeletedOrganizationSettings
// Validates and Restores deleted organization settings record by ID in database.
func (oss *OrganizationSettingsService) RestoreDeletedOrganizationSettings(organizationID *string) error {
	if err := uuid.Validate(organizationID); err != nil {
		return err
	}

	if err := oss.repository.RestoreDeletedOrganizationSettings(*organizationID); err != nil {
		return err
	}

	return nil
}
