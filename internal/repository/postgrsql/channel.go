package postgrsql

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"

	"github.com/jmoiron/sqlx"
)

// ChannelRepository
// Contains information about channel repository.
type ChannelRepository struct {
	db *sqlx.DB
}

// NewChannelRepository
// Creates new ChannelRepository.
func NewChannelRepository(db *sqlx.DB) *ChannelRepository {
	return &ChannelRepository{
		db: db,
	}
}

// CreateChannel
// Creates new channel record in database.
func (cr *ChannelRepository) CreateChannel(c *common.Channel) error {
	rows, err := cr.db.NamedQuery(`
		insert into channels (
			channel_name,
			channel_description,
			channel_type_id,
			channel_technical_id,
			channel_status_id
		)
		values (
			:channel_name,
			:channel_description,
			:channel_type_id,
			:channel_technical_id,
			:channel_status_id
		)
		returning
			channel_id;
	`, *c)
	if err != nil {
		return err
	}

	var lastInsertedID string
	for rows.Next() {
		if err := rows.Scan(&lastInsertedID); err != nil {
			return err
		}
	}

	c.ChannelId = &lastInsertedID

	return nil
}

// GetChannels
// Get the list of all channels of the specific organization.
func (cr *ChannelRepository) GetChannels(rootOrganizationId *string) (*[]common.Channel, error) {
	return nil, nil
}

// GetChannels
// Get the information about the specific channel.
func (cr *ChannelRepository) GetChannel(channelId *string) (*common.Channel, error) {
	return nil, nil
}

