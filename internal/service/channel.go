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

// CreateChannel
// Validates and Creates new channel record in database.
func (cs *ChannelService) CreateChannel(c *common.Channel) error {
	if err := c.Validate(); err != nil {
		return err
	}

	if err := cs.repository.CreateChannel(c); err != nil {
		return err
	}

	return nil
}

// GetChannels
// Get the list of all channels of the specific organization.
func (cs *ChannelService) GetChannels(rootOrganizationId *string) (*[]common.Channel, error) {
	return nil, nil
}

// GetChannels
// Get the information about the specific channel.
func (cs *ChannelService) GetChannel(channelId *string) (*common.Channel, error) {
	return nil, nil
}
