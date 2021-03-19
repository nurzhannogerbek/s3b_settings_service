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

// Create
// Validates and Creates new organization settings record in database.
func (oss *OrganizationSettingsService) Create(os *common.OrganizationSettings) error {
	if err := os.Validate(); err != nil {
		return err
	}

	if err := oss.repository.Create(os); err != nil {
		return err
	}

	return nil
}

// Delete
// Validates and Deletes organization settings record by ID in database.
func (oss *OrganizationSettingsService) Delete(organizationID *string) error {
	if err := uuid.Validate(organizationID); err != nil {
		return err
	}

	if err := oss.repository.Delete(organizationID); err != nil {
		return err
	}

	return nil
}

// GetById
// Validates and Queries organization settings record by ID from database.
func (oss *OrganizationSettingsService) GetByID(organizationID *string) (*common.OrganizationSettings, error) {
	if err := uuid.Validate(organizationID); err != nil {
		return nil, err
	}

	organizationSettings, err := oss.repository.GetByID(organizationID)
	if err != nil {
		return nil, err
	}

	return organizationSettings, nil
}

// Update
// Validates and updates organization settings record by ID in database.
func (oss *OrganizationSettingsService) Update(os *common.OrganizationSettings) (*common.OrganizationSettings, error) {
	if err := os.Validate(); err != nil {
		return nil, err
	}

	organizationSettings, err := oss.repository.Update(os)
	if err != nil {
		return nil, err
	}

	return organizationSettings, nil
}
