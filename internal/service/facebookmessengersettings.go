package service

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository"
)

// FacebookMessengerSettingsService
// Contains all dependencies for facebook messenger settings service.
type FacebookMessengerSettingsService struct {
	repository repository.FacebookMessengerSettings
}

// NewFacebookMessengerSettingsService
// Creates new facebook messenger settings service (FacebookMessengerSettings).
func NewFacebookMessengerSettingsService(fms repository.FacebookMessengerSettings) *FacebookMessengerSettingsService {
	return &FacebookMessengerSettingsService{
		repository: fms,
	}
}

// Create
// Validates and Creates new facebook messenger channel settings record in database.
func (fmss *FacebookMessengerSettingsService) Create(fms *common.FacebookMessengerSettings) error {
	if err := fms.Validate(); err != nil {
		return err
	}

	if err := fmss.repository.Create(fms); err != nil {
		return err
	}

	return nil
}
