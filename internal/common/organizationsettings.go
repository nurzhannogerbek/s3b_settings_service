package common

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/tool/uuid"
	"encoding/json"
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

type WeekWorkTime struct {
	Monday    *DayWorkTime `json:"monday"`
	Tuesday   *DayWorkTime `json:"tuesday"`
	Wednesday *DayWorkTime `json:"wednesday"`
	Friday    *DayWorkTime `json:"friday"`
	Saturday  *DayWorkTime `json:"saturday"`
	Sunday    *DayWorkTime `json:"sunday"`
}

type DayWorkTime struct {
	BeginTime *int     `json:"beginTime"`
	EndTime   *int     `json:"endTime"`
	Break     *[]Break `json:"break"`
}

type Break struct {
	BeginTime *int `json:"beginTime"`
	EndTime   *int `json:"endTime"`
}

type Privacy struct {
	DataTransferAndStorage *bool `json:"dataTransferAndStorage"`
	SidePersonDataTransfer *bool `json:"sidePersonDataTransfer"`
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

	if orgSet.OrganizationSettingWorkTime != nil {
		if err := orgSet.ValidateWorkTime(); err != nil {
			return err
		}
	}

	if orgSet.OrganizationSettingPrivacy != nil {
		if err := orgSet.ValidatePrivacy(); err != nil {
			return err
		}
	}

	if orgSet.TimezoneID != nil {
		if err := uuid.Validate(orgSet.TimezoneID); err != nil {
			return err
		}
	}

	return nil
}

func (orgSet OrganizationSettings) ValidateWorkTime() error {
	var weekWorkTime WeekWorkTime
	if err := json.Unmarshal([]byte(*orgSet.OrganizationSettingWorkTime), &weekWorkTime); err != nil {
		return err
	}

	return nil
}

func (orgSet OrganizationSettings) ValidatePrivacy() error {
	var privacy Privacy
	if err := json.Unmarshal([]byte(*orgSet.OrganizationSettingPrivacy), &privacy); err != nil {
		return err
	}

	return nil
}
