package postgrsql

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/repository"
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/database/postgresql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewRepositories(newDB *sqlx.DB) *repository.Repositories {
	return &repository.Repositories{
		Organization: NewOrganizationSettingsRepo(newDB),
	}
}

type OrganizationSettingsRepo struct {
	db *sqlx.DB
}

func NewOrganizationSettingsRepo(newDB *sqlx.DB) *OrganizationSettingsRepo {
	return &OrganizationSettingsRepo{
		db: newDB,
	}
}

func (osr *OrganizationSettingsRepo) Create(organization *common.OrganizationSettings) error {
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

func (osr OrganizationSettingsRepo) Delete(organizationID *string) error {
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

func (osr OrganizationSettingsRepo) Get(organizationID *string) (*common.OrganizationSettings, error) {
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

func (osr OrganizationSettingsRepo) Update(organization *common.OrganizationSettings) (*common.OrganizationSettings, error) {
	updateCondition := postgresql.UpdateConditionFromStruct(organization)
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

	rows, err := osr.db.NamedQuery(queryString, *organization)
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
			&organizationSettings.LocationID); err != nil {
			return nil, err
		}
	}

	return &organizationSettings, err
}
