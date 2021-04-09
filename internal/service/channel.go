package service

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository"
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/tool/uuid"
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
func (cs *ChannelService) CreateChannel(c *common.Channel) (*common.Channel, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}

	channel, err := cs.repository.CreateChannel(c)
	if err != nil {
		return nil, err
	}

	return channel, nil
}

// GetChannels
// Get the list of all channels of the specific organization.
func (cs *ChannelService) GetChannels(organizationId *string) (*[]common.Channel, error) {
	if err := uuid.Validate(organizationId); err != nil {
		return nil, err
	}

	channels, err := cs.repository.GetChannels(organizationId)
	if err != nil {
		return nil, err
	}

	return channels, nil
}

// GetChannels
// Get the information about the specific channel.
func (cs *ChannelService) GetChannel(channelId *string) (*common.Channel, error) {
	if err := uuid.Validate(channelId); err != nil {
		return nil, err
	}

	channel, err := cs.repository.GetChannel(channelId)
	if err != nil {
		return nil, err
	}

	return channel, nil
}
