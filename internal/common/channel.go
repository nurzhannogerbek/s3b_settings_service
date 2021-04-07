package common

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/tool/uuid"

	"github.com/lib/pq"
)

// Channel
// Contains information about channel.
type Channel struct {
	ChannelId          *string           `db:"channel_id" json:"channelId"`
	ChannelName        *string           `db:"channel_name" json:"channelName"`
	ChannelDescription *string           `db:"channel_description" json:"channelDescription"`
	ChannelTypeId      *string           `db:"channel_type_id" json:"channelTypeId"`
	ChannelTechnicalId *string           `db:"channel_technical_id" json:"channelTechnicalId"`
	ChannelStatusId    *string           `db:"channel_status_id" json:"channelStatusId"`
	OrganizationIds    pq.StringArray    `db:"organization_ids" json:"organizationIds"`
}

// Validate
// Validates Channel struct.
func (c *Channel) Validate() error {
	if c.ChannelId != nil {
		if err := uuid.Validate(c.ChannelId); err != nil {
			return err
		}
	}

	if c.ChannelTypeId != nil {
		if err := uuid.Validate(c.ChannelTypeId); err != nil {
			return err
		}
	}

	if c.ChannelStatusId != nil {
		if err := uuid.Validate(c.ChannelStatusId); err != nil {
			return err
		}
	}

	return nil
}
