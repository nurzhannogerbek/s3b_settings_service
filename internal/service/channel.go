package service

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository"
)

// ChannelService
// Contains all dependencies for channel service.
type ChannelService struct {
	repository repository.Channel
}

// NewChannelService
// Creates new channel service (Channel).
func NewChannelService(c repository.Channel) *ChannelService {
	return &ChannelService{
		repository: c,
	}
}

// Create
// Validates and Creates new channel record in database.
func (cs *ChannelService) Create(c *common.Channel) error {
	if err := c.Validate(); err != nil {
		return err
	}

	if err := cs.repository.Create(c); err != nil {
		return err
	}

	return nil
}
