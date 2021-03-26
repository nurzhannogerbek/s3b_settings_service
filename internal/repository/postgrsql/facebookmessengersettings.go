package postgrsql

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"

	"github.com/jmoiron/sqlx"
)

// FacebookMessengerSettingsRepository
// Contains information about facebook messenger settings repository.
type FacebookMessengerSettingsRepository struct {
	db *sqlx.DB
}

// NewOrganizationSettingsRepository
// Creates new FacebookMessengerSettingsRepository.
func NewFacebookMessengerSettingsRepository(db *sqlx.DB) *FacebookMessengerSettingsRepository {
	return &FacebookMessengerSettingsRepository{
		db: db,
	}
}

// Create
// Creates new facebook messenger settings record in database.
func (fmsr *FacebookMessengerSettingsRepository) Create(fms *common.FacebookMessengerSettings) error {
	rows, err := fmsr.db.NamedQuery(`
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
	`, *fms)
	if err != nil {
		return err
	}

	var lastInsertedID string
	for rows.Next() {
		if err := rows.Scan(&lastInsertedID); err != nil {
			return err
		}
	}

	fms.ChannelId = &lastInsertedID

	return nil
}
