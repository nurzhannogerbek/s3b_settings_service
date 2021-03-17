package common

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/tool/uuid"
	"encoding/json"
	"errors"
	"unicode/utf8"
)

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
	Monday    DayWorkTime `json:"monday"`
	Tuesday   DayWorkTime `json:"tuesday"`
	Wednesday DayWorkTime `json:"wednesday"`
	Friday    DayWorkTime `json:"friday"`
	Saturday  DayWorkTime `json:"saturday"`
	Sunday    DayWorkTime `json:"sunday"`
}

func (wwt WeekWorkTime) Validate() error {
	if err := wwt.Monday.Validate(); err != nil {
		return err
	}

	if err := wwt.Tuesday.Validate(); err != nil {
		return err
	}

	if err := wwt.Wednesday.Validate(); err != nil {
		return err
	}

	if err := wwt.Friday.Validate(); err != nil {
		return err
	}

	if err := wwt.Saturday.Validate(); err != nil {
		return err
	}

	if err := wwt.Sunday.Validate(); err != nil {
		return err
	}

	return nil
}

type DayWorkTime struct {
	BeginTime int         `json:"beginTime"`
	EndTime   int         `json:"endTime"`
	BreakTime []BreakTime `json:"breakTime"`
}

func (dwt DayWorkTime) Validate() error {
	if dwt.BeginTime < 0 || dwt.BeginTime > 86399 {
		return errors.New("beginTime should be in range 0 - 86399")
	}

	if dwt.EndTime < 0 || dwt.EndTime > 86399 {
		return errors.New("endTime should be in range 0 - 86399")
	}


	for _, i := range dwt.BreakTime {
		if err := i.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type BreakTime struct {
	BeginTime int `json:"beginTime"`
	EndTime   int `json:"endTime"`
}

func (bt BreakTime) Validate() error {
	if bt.BeginTime < 0 || bt.BeginTime > 86399 {
		return errors.New("beginTime should be in range 0 - 86399")
	}

	if bt.EndTime < 0 || bt.EndTime > 86399 {
		return errors.New("endTime should be in range 0 - 86399")
	}

	return nil
}

type Privacy struct {
	DataTransferAndStorage bool `json:"dataTransferAndStorage"`
	SidePersonDataTransfer bool `json:"sidePersonDataTransfer"`
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

	if err := weekWorkTime.Validate(); err != nil {
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
