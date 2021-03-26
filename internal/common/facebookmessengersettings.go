package common

import "bitbucket.org/3beep-workspace/3beep_settings_service/pkg/tool/uuid"

type CategoryList struct {
	Id   *string `json:"id"`
	Name *string `json:"name"`
}
type Data struct {
	AccessToken  *string `json:"accessToken"`
	Category     *string `json:"category"`
	CategoryList []*CategoryList `json:"categoryList"`
	Name         *string `json:"name"`
	Id           *string `json:"id"`
	Tasks        []*string
}

type Paging struct {
	Cursors *Cursors `json:"cursors"`
}

type Cursors struct {
	Before *string `json:"before"`
	After  *string `json:"after"`
}

// FacebookPages
// Contains information about facebook pages.
type FacebookPages struct {
	Data   []*Data `json:"data"`
	Paging *Paging `json:"paging"`
}

// FacebookMessengerSettings
// Contains information about facebook messenger settings.
type FacebookMessengerSettings struct {
	ChannelId          *string `db:"channel_id" json:"channelId"`
	ChannelName        *string `db:"channel_name" json:"channelName"`
	ChannelDescription *string `db:"channel_description" json:"channelDescription"`
	ChannelTypeId      *string `db:"channel_type_id" json:"channelTypeId"`
	ChannelTechnicalId *string `db:"channel_technical_id" json:"channelTechnicalId"`
	ChannelStatusId    *string `db:"channel_status_id" json:"channelStatusId"`
}

// Validate
// Validates FacebookMessengerSettings struct.
func (fms *FacebookMessengerSettings) Validate() error {
	if fms.ChannelId != nil {
		if err := uuid.Validate(fms.ChannelId); err != nil {
			return err
		}
	}

	if err := uuid.Validate(fms.ChannelTypeId); err != nil {
		return err
	}

	if err := uuid.Validate(fms.ChannelStatusId); err != nil {
		return err
	}

	return nil
}
