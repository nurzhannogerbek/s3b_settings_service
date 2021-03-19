package postgrsql

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/database/postgresql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// OrganizationSettingsRepository
// Contains information about organization settings repository.
type OrganizationSettingsRepository struct {
	db *sqlx.DB
}

// NewOrganizationSettingsRepository
// Creates new OrganizationSettingsRepository.
func NewOrganizationSettingsRepository(db *sqlx.DB) *OrganizationSettingsRepository {
	return &OrganizationSettingsRepository{
		db: db,
	}
}

// Create
// Creates new organization settings record in database.
func (osr *OrganizationSettingsRepository) Create(os *common.OrganizationSettings) error {
	rows, err := osr.db.NamedQuery(`
		insert into organizations_settings (
			organization_id,
		    country_id,
			location_id,
			organization_setting_address,
			organization_setting_postal_code,
			organization_setting_work_time,
			organization_setting_privacy,
			timezone_id
		)
		values (
		    :organization_id,
			:country_id,
			:location_id,
			:organization_setting_address,
			:organization_setting_postal_code,
			:organization_setting_work_time,
			:organization_setting_privacy,
			:timezone_id
		)
		returning
			organization_id;`, *os)
	if err != nil {
		return err
	}

	var lastInsertedID string
	for rows.Next() {
		if err := rows.Scan(&lastInsertedID); err != nil {
			return err
		}
	}

	os.OrganizationID = &lastInsertedID

	return nil
}

// Delete
// Deletes organization settings from database by ID.
func (osr *OrganizationSettingsRepository) Delete(organizationID *string) error {
	_, err := osr.db.Query(`
		update 
		    organizations_settings
		set
			entry_deleted_date_time = now()
		where 
			organization_id = $1;`, *organizationID)

	if err != nil {
		return err
	}

	return nil
}

// GetByID
// Queries organization settings by ID from database.
func (osr *OrganizationSettingsRepository) GetByID(organizationID *string) (*common.OrganizationSettings, error) {
	var organizationSettings common.OrganizationSettings
	err := osr.db.Get(&organizationSettings, `
		select 
		    organization_id,
		    country_id,
		    location_id,
		    organization_setting_address,
		    organization_setting_postal_code,
		    organization_setting_work_time,
		    organization_setting_privacy,
		    timezone_id
		from 
		    organizations_settings
		where 
			organization_id = $1;`, *organizationID)

	if err != nil {
		return nil, err
	}

	return &organizationSettings, nil
}

// Update
// Updates organization settings record in database.
func (osr *OrganizationSettingsRepository) Update(os *common.OrganizationSettings) (*common.OrganizationSettings, error) {
	updateCondition := postgresql.UpdateConditionFromStruct(os)
	queryString := fmt.Sprintf(`
		update 
			organizations_settings
		set
			entry_updated_date_time = now(),
			%s
		where
			organization_id = :organization_id
		returning
			organization_id,
		    country_id,
		    location_id,
		    organization_setting_address,
		    organization_setting_postal_code,
		    organization_setting_work_time,
		    organization_setting_privacy,
		    timezone_id;`, updateCondition)

	rows, err := osr.db.NamedQuery(queryString, *os)
	if err != nil {
		return nil, err
	}

	var organizationSettings common.OrganizationSettings
	for rows.Next() {
		if err := rows.Scan(&organizationSettings.OrganizationID,
			&organizationSettings.CountryID,
			&organizationSettings.LocationID,
			&organizationSettings.OrganizationSettingAddress,
			&organizationSettings.OrganizationSettingPostalCode,
			&organizationSettings.OrganizationSettingWorkTime,
			&organizationSettings.OrganizationSettingPrivacy,
			&organizationSettings.TimezoneID); err != nil {
			return nil, err
		}
	}

	return &organizationSettings, err
}
