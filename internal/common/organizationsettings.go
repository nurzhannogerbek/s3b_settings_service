package common

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/pkg/tool/uuid"
	"encoding/json"
	"errors"
	"unicode/utf8"
)

// OrganizationSettings
// Contains information about organization settings.
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

// WeekWorkTime
// Contains information about organization week work time.
type WeekWorkTime struct {
	Monday    DayWorkTime `json:"monday"`
	Tuesday   DayWorkTime `json:"tuesday"`
	Wednesday DayWorkTime `json:"wednesday"`
	Thursday  DayWorkTime `json:"thursday"`
	Friday    DayWorkTime `json:"friday"`
	Saturday  DayWorkTime `json:"saturday"`
	Sunday    DayWorkTime `json:"sunday"`
}

// Validate
// Validates WeekWorkTime struct.
func (wwt *WeekWorkTime) Validate() error {
	if err := wwt.Monday.Validate(); err != nil {
		return err
	}

	if err := wwt.Tuesday.Validate(); err != nil {
		return err
	}

	if err := wwt.Wednesday.Validate(); err != nil {
		return err
	}

	if err := wwt.Thursday.Validate(); err != nil {
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

// DayWorkTime
// Contains information about day work time of the week (WeekWorkTime).
type DayWorkTime struct {
	BeginTime int         `json:"beginTime"`
	EndTime   int         `json:"endTime"`
	BreakTime []BreakTime `json:"breakTime"`
	IsActive  bool        `json:"isActive"`
}

// Validate
// Validates DayWorkTime struct.
func (dwt *DayWorkTime) Validate() error {
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

// BreakTime
// Contains information about break time of the day (DayWorkTime).
type BreakTime struct {
	BeginTime int `json:"beginTime"`
	EndTime   int `json:"endTime"`
}

// Validate
// Validates BreakTime struct.
func (bt *BreakTime) Validate() error {
	if bt.BeginTime < 0 || bt.BeginTime > 86399 {
		return errors.New("beginTime should be in range 0 - 86399")
	}

	if bt.EndTime < 0 || bt.EndTime > 86399 {
		return errors.New("endTime should be in range 0 - 86399")
	}

	return nil
}

// Privacy
// Contains information about privacy of the organization (OrganizationSettings).
type Privacy struct {
	DataTransferAndStorage bool `json:"dataTransferAndStorage"`
	SidePersonDataTransfer bool `json:"sidePersonDataTransfer"`
}

// Validate
// Validates OrganizationSettings struct.
func (os *OrganizationSettings) Validate() error {
	if err := uuid.Validate(os.OrganizationID); err != nil {
		return err
	}

	if os.CountryID != nil {
		if *os.CountryID != "null" {
			if err := uuid.Validate(os.CountryID); err != nil {
				return err
			}
		}
	}

	if os.LocationID != nil {
		if *os.LocationID != "null" {
			if err := uuid.Validate(os.LocationID); err != nil {
				return err
			}
		}
	}

	if os.OrganizationSettingAddress != nil {
		organizationSettingAddress := utf8.RuneCountInString(*os.OrganizationSettingAddress)
		if organizationSettingAddress > 50 || organizationSettingAddress < 5 {
			return errors.New("organizationSettingAddress length should be less than 50 and greater than 5")
		}
	}

	if os.OrganizationSettingPostalCode != nil {
		organizationSettingPostalCode := utf8.RuneCountInString(*os.OrganizationSettingPostalCode)
		if organizationSettingPostalCode > 50 || organizationSettingPostalCode < 5 {
			return errors.New("organizationSettingPostalCode length should be less than 50 and greater than 5")
		}
	}

	if os.OrganizationSettingWorkTime != nil {
		if err := os.ValidateWorkTime(); err != nil {
			return err
		}
	}

	if os.OrganizationSettingPrivacy != nil {
		if err := os.ValidatePrivacy(); err != nil {
			return err
		}
	}

	if os.TimezoneID != nil {
		if *os.TimezoneID != "null" {
			if err := uuid.Validate(os.TimezoneID); err != nil {
				return err
			}
		}
	}

	return nil
}

// ValidateWorkTime
// Validates work time of the organization (OrganizationSettings).
func (os *OrganizationSettings) ValidateWorkTime() error {
	var breakTimes []BreakTime
	var breakTime BreakTime
	breakTimes = append(breakTimes, breakTime)
	weekWorkTime := WeekWorkTime{
		Monday: DayWorkTime{
			BeginTime: 0,
			EndTime:   86399,
			BreakTime: breakTimes,
			IsActive:  true,
		},
		Tuesday: DayWorkTime{
			BeginTime: 0,
			EndTime:   86399,
			BreakTime: breakTimes,
			IsActive:  true,
		},
		Wednesday: DayWorkTime{
			BeginTime: 0,
			EndTime:   86399,
			BreakTime: breakTimes,
			IsActive:  true,
		},
		Thursday: DayWorkTime{
			BeginTime: 0,
			EndTime:   86399,
			BreakTime: breakTimes,
			IsActive:  true,
		},
		Friday: DayWorkTime{
			BeginTime: 0,
			EndTime:   86399,
			BreakTime: breakTimes,
			IsActive:  true,
		},
		Saturday: DayWorkTime{
			BeginTime: 0,
			EndTime:   86399,
			BreakTime: breakTimes,
			IsActive:  true,
		},
		Sunday: DayWorkTime{
			BeginTime: 0,
			EndTime:   86399,
			BreakTime: breakTimes,
			IsActive:  true,
		},
	}

	if err := json.Unmarshal([]byte(*os.OrganizationSettingWorkTime), &weekWorkTime); err != nil {
		return err
	}

	if err := weekWorkTime.Validate(); err != nil {
		return err
	}

	b, err := json.Marshal(weekWorkTime)
	if err != nil {
		return err
	}

	*os.OrganizationSettingWorkTime = string(b)

	return nil
}

// ValidatePrivacy
// Validates privacy of the organization (OrganizationSettings).
func (os *OrganizationSettings) ValidatePrivacy() error {
	var privacy Privacy
	if err := json.Unmarshal([]byte(*os.OrganizationSettingPrivacy), &privacy); err != nil {
		return err
	}

	return nil
}
