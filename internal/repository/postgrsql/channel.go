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

// Create
// Creates new channel record in database.
func (cr *ChannelRepository) Create(c *common.Channel) error {
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
