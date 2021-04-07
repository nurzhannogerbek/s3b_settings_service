package postgrsql

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
			channel_id::text;
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
func (cr *ChannelRepository) GetChannels(organizationId *string) (*[]common.Channel, error) {
	var channels []common.Channel

	rows, err := cr.db.Query(`
		select
			channels.channel_id::text,
			channels.channel_name::text,
			channels.channel_description::text,
			channels.channel_type_id::text,
			channels.channel_technical_id::text,
			channels.channel_status_id::text,
			array_agg(channels_organizations_relationship.organization_id)::text[] organizations_ids
		from
			channels
		left join channels_organizations_relationship on
			channels.channel_id = channels_organizations_relationship.channel_id
		left join organizations on
			channels_organizations_relationship.organization_id = organizations.organization_id
		where
			organizations.tree_organization_id like concat('%', '\', $1::text, '%')
		group by
			channels.channel_id;`, *organizationId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var channel common.Channel

		if err = rows.Scan(&channel.ChannelId,
			&channel.ChannelName,
			&channel.ChannelDescription,
			&channel.ChannelTypeId,
			&channel.ChannelTechnicalId,
			&channel.ChannelStatusId,
			pq.Array(&channel.OrganizationsIds)); err != nil {
			return nil, err
		}

		channels = append(channels, channel)
	}

	return &channels, nil
}

// GetChannel
// Get the information about the specific channel.
func (cr *ChannelRepository) GetChannel(channelId *string) (*common.Channel, error) {
	var channel common.Channel

	row := cr.db.QueryRow(`
		select
			channels.channel_id::text,
			channels.channel_name::text,
			channels.channel_description::text,
			channels.channel_type_id::text,
			channels.channel_technical_id::text,
			channels.channel_status_id::text,
			array_agg(distinct channels_organizations_relationship.organization_id)::text[] organizations_ids
		from
			channels
		left join channels_organizations_relationship on
			channels.channel_id = channels_organizations_relationship.channel_id
		where
			channels.channel_id = $1
		group by
    		channels.channel_id
		limit 1;`, *channelId)

	if err := row.Scan(&channel.ChannelId,
		&channel.ChannelName,
		&channel.ChannelDescription,
		&channel.ChannelTypeId,
		&channel.ChannelTechnicalId,
		&channel.ChannelStatusId,
		pq.Array(&channel.OrganizationsIds)); err != nil {
		return nil, err
	}

	return &channel, nil
}

