package service

type Organization struct {
	OrganizationID              *string `db:"organization_id" json:"organizationId"`
	CountryID                   *string `db:"country_id" json:"countryId"`
	LocationID                  *string `db:"location_id" json:"locationId"`
	OrganizationSettingAddress  *string `db:"organization_setting_address" json:"organizationSettingAddress"`
	OrganizationSettingWorkTime *string `db:"organization_setting_work_time" json:"organizationSettingWorkTime"`
	OrganizationSettingPrivacy  *string `db:"organization_setting_privacy" json:"organizationSettingPrivacy"`
	TimezoneID                  *string `db:"timezone_id" json:"timezoneId"`
}

func CreateOrganization(org Organization) (*Organization, error) {
	return nil, nil
}
