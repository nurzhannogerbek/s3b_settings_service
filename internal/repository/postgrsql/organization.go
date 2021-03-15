package postgrsql

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/service"
	"github.com/jmoiron/sqlx"
)

type OrganizationRepository struct {
	db *sqlx.DB
}

func (or *OrganizationRepository) Create(organization *service.Organization) error {
	rows, err := or.db.NamedQuery(`
		insert into organizations_settings (
			country_id,
			location_id,
			organization_setting_address,
			organization_setting_postal_code,
			organization_setting_work_time,
			organization_setting_privacy,
			timezone_id
		)
		values (
			:country_id,
			:location_id,
			:organization_setting_address,
			:organization_setting_postal_code
			:organization_setting_work_time,
			:organization_setting_privacy
			:timezone_id
		)
		returning
			organization_id;`, *organization)
	if err != nil {
		return err
	}

	var lastInsertedID string
	for rows.Next() {
		if err := rows.Scan(&lastInsertedID); err != nil {
			return err
		}
	}

	organization.OrganizationID = &lastInsertedID

	return nil
}
