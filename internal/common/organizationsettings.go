package common

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/tool/uuid"
	"errors"
	"unicode/utf8"
)

type OrganizationSettingPrivacy int

const (
	DataTransferAndStorage OrganizationSettingPrivacy = iota
	SidePersonDataTransfer
)

func (orgSetPri OrganizationSettingPrivacy) String() string {
	switch orgSetPri {
	case DataTransferAndStorage:
		return "dataTransferAndStorage"
	case SidePersonDataTransfer:
		return "SidePersonDataTransfer"
	default:
		return ""
	}
}

type OrganizationSettings struct {
	OrganizationID                *string `db:"organization_id" json:"organizationId"`
	CountryID                     *string `db:"country_id" json:"countryId"`
	LocationID                    *string `db:"location_id" json:"locationId"`
	OrganizationSettingAddress    *string `db:"organization_setting_address" json:"organizationSettingAddress"`
	OrganizationSettingPostalCode *string `db:"organization_setting_postal_code" json:"organizationSettingPostalCode"`
	OrganizationSettingWorkTime   *string `db:"organization_setting_work_time" json:"organizationSettingWorkTime"`
	OrganizationSettingPrivacy    *string `db:"organization_setting_privacy" json:"organizationSettingPrivacy"`
	TimezoneID                    *string `db:"timezone_id" json:"timezoneId"`
}

func (orgSet OrganizationSettings) Validate() error {
	if err := uuid.Validate(orgSet.OrganizationID); err != nil {
		return err
	}

	if orgSet.CountryID != nil {
		if err := uuid.Validate(orgSet.CountryID); err != nil {
			return err
		}
	}

	if orgSet.CountryID != nil {
		if err := uuid.Validate(orgSet.LocationID); err != nil {
			return err
		}
	}

	if orgSet.OrganizationSettingAddress != nil {
		orgSettAdd := utf8.RuneCountInString(*orgSet.OrganizationSettingAddress)
		if orgSettAdd < 50 || orgSettAdd < 5 {
			return errors.New("organizationSettingAddress length should be less than 50 and greater than 5")
		}
	}

	if orgSet.OrganizationSettingPostalCode != nil {
		orgSetPosCod := utf8.RuneCountInString(*orgSet.OrganizationSettingPostalCode)
		if orgSetPosCod > 50 || orgSetPosCod < 5 {
			return errors.New("organizationSettingPostalCode length should be less than 50 and greater than 5")
		}
	}

	if err := orgSet.ValidateWorkTime(); err != nil {
		return err
	}

	if err := orgSet.ValidatePrivacy(); err != nil {
		return err
	}

	if err := uuid.Validate(orgSet.TimezoneID); err != nil {
		return err
	}

	return nil
}

func (orgSet OrganizationSettings) ValidateWorkTime() error {
	return nil
}

func (orgSet OrganizationSettings) ValidatePrivacy() error {
	return nil
}
