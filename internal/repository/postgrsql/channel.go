package postgrsql

import (
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/common"
	"bitbucket.org/3beep-workspace/3beep_settings_service/internal/environment"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// ChannelRepository
// Contains information about channel repository.
type ChannelRepository struct {
	db *sqlx.DB
}

// NewChannelRepository
// Creates new ChannelRepository.
func NewChannelRepository(db *sqlx.DB) *ChannelRepository {
	return &ChannelRepository{
		db: db,
	}
}

// SetWebhookToTelegram
// Set webhook to the telegram chat bot.
func SetWebhookToTelegram (channelName string, channelTechnicalId string) error {
	client := &http.Client{}

	request, err := http.NewRequest("GET", fmt.Sprintf("https://api.telegram.org/bot%s/setWebhook", channelTechnicalId), nil)
	if err != nil {
		return err
	}

	query := request.URL.Query()
	query.Add("url", fmt.Sprintf("%ssend_message_from_telegram/%s", environment.TelegramBotURL, channelName))
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("response status code: %q", response.StatusCode)
	}

	return nil
}

// SetWebhookToFacebookMessenger
// Set webhook to the facebook messenger chat bot.
func SetWebhookToFacebookMessenger () error {
	client := &http.Client{}

	request, err := http.NewRequest("GET", "https://graph.facebook.com/oauth/access_token", nil)
	if err != nil {
		return err
	}

	query := request.URL.Query()
	query.Add("client_id", environment.FacebookAppId)
	query.Add("client_secret", environment.FacebookAppSecret)
	query.Add("grant_type", "client_credentials")
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var facebookAppToken struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
	}
	if err = json.Unmarshal(body, &facebookAppToken); err != nil {
		return err
	}

	request, err = http.NewRequest("POST", fmt.Sprintf("https://graph.facebook.com/v9.0/%s/subscriptions", environment.FacebookAppId), nil)
	if err != nil {
		return err
	}

	query = request.URL.Query()
	query.Add("access_token", facebookAppToken.AccessToken)
	query.Add("object", "page")
	query.Add("callback_url", environment.FacebookMessengerBotURL)
	query.Add("fields", "messages,messaging_postbacks,messaging_optins,message_deliveries,message_reads,messaging_payments,messaging_pre_checkouts,messaging_checkout_updates,messaging_account_linking,messaging_referrals,message_echoes,messaging_game_plays,standby,messaging_handovers,messaging_policy_enforcement,message_reactions,inbox_labels")
	query.Add("include_values", "true")
	query.Add("verify_token", environment.FacebookMessengerBotVerifyToken)
	request.URL.RawQuery = query.Encode()

	response, err = client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("response status code: %q", response.StatusCode)
	}

	return nil
}

// SetInstagramPrivateSession
// Set session of the private instagram chat bot.
// Fire and forget.
func (cr *ChannelRepository) SetInstagramPrivateSession (channelId string, channelStatusId string) error {
	var channelStatusName string
	err := cr.db.QueryRowx(`
	select
		channel_status_name
	from
		channel_statuses
	where
		channel_status_id = $1
	limit 1;`, &channelStatusId).Scan(&channelStatusName)
	if err != nil {
		return err
	}

	var action string
	switch channelStatusName {
	case "active":
		action = "true"
	default:
		action = "false"
	}

	client := &http.Client{}

	data := url.Values{}
	data.Set("channel_id", channelId)
	data.Set("action", action)

	urlAddress := fmt.Sprintf("%stoggle_session_for_instagram_private_channel", environment.InstagramPrivateBotURL)

	request, err := http.NewRequest("POST", urlAddress, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}

func stringExistsInArray(originalArray []string, originalString string) bool {
	for _, value := range originalArray {
		if value == originalString {
			return true
		}
	}
	return false
}

// CreateChannel
// Creates new channel record in database.
func (cr *ChannelRepository) CreateChannel(c *common.Channel) (*common.Channel, error) {
	var channelTypeName string
	err := cr.db.QueryRowx(`
		select
			lower(channel_type_name) as channel_type_name
		from
			channel_types
		where
			channel_type_id = $1
		limit 1;`, &c.ChannelTypeId).Scan(&channelTypeName)
	if err != nil {
		return nil, fmt.Errorf("unexpected channel type, err: %q", err.Error())
	}

	availableChannelTypes := []string{"telegram", "facebook_messenger", "whatsapp", "instagram_private"}
	if !stringExistsInArray(availableChannelTypes, channelTypeName) {
		return nil, fmt.Errorf("creating a channel for the %q type is currently not possible", channelTypeName)
	}

	err = cr.db.QueryRowx(`
		insert into channels (
			channel_name,
			channel_description,
			channel_type_id,
			channel_technical_id,
			channel_status_id
		)
		values (
			$1,
			$2,
			$3,
			$4,
			$5
		)
		returning
			channel_id::text;`,
		&c.ChannelName,
		&c.ChannelDescription,
		&c.ChannelTypeId,
		&c.ChannelTechnicalId,
		&c.ChannelStatusId).Scan(&c.ChannelId)
	if err != nil {
		return nil, fmt.Errorf("failed to create a channel, err: %q", err.Error())
	}

	_, err = cr.db.Exec(`
		insert into channels_organizations_relationship(
			channel_id,
			organization_id
		)
		select
			$1::uuid channel_id,
			organizations_ids
		from
			unnest($2::uuid[]) organizations_ids;`,
		&c.ChannelId,
		&c.OrganizationsIds)
	if err != nil {
		return nil, fmt.Errorf("сouldn't link channel to organizations, err: %q", err.Error())
	}

	switch channelTypeName {
	case "telegram":
		_, err = cr.db.Exec(`
			insert into telegram_business_accounts (
				business_account,
				channel_id
			)
			values (
				$1,
				$2
			);`,
			&c.ChannelName,
			&c.ChannelId)
		if err != nil {
			return nil, fmt.Errorf("сouldn't insert new value in the 'telegram_business_accounts' table, err: %q", err.Error())
		}

		err = SetWebhookToTelegram(*c.ChannelName, *c.ChannelTechnicalId)
		if err != nil {
			return nil, fmt.Errorf("couldn't set webhook via telegram api, err: %q", err.Error())
		}
	case "facebook_messenger":
		_, err = cr.db.Exec(`
			insert into facebook_messenger_business_accounts (
				business_account,
				channel_id
			)
			values (
				$1,
				$2
			);`,
			&c.ChannelName,
			&c.ChannelId)
		if err != nil {
			return nil, fmt.Errorf("сouldn't insert new value in the 'facebook_messenger_business_accounts' table, err: %q", err.Error())
		}

		err = SetWebhookToFacebookMessenger()
		if err != nil {
			return nil, fmt.Errorf("couldn't set webhook via facebook graph api, err: %q", err.Error())
		}
	case "instagram_private":
		_, err = cr.db.Exec(`
			insert into instagram_private_business_accounts (
				business_account,
				channel_id
			)
			values (
				$1,
				$2
			);`,
			&c.ChannelName,
			&c.ChannelId)
		if err != nil {
			return nil, fmt.Errorf("сouldn't insert new value in the 'instagram_private_business_accounts' table, err: %q", err.Error())
		}

		err = cr.SetInstagramPrivateSession(*c.ChannelId, *c.ChannelStatusId)
		if err != nil {
			return nil, fmt.Errorf("couldn't set sessions in the private instagram, err: %q", err.Error())
		}
	case "whatsapp":
		_, err = cr.db.Exec(`
			insert into whatsapp_business_accounts (
				business_account,
				channel_id
			)
			values (
				concat('+', $1::text),
				$2::uuid
			);`,
			&c.ChannelName,
			&c.ChannelId)
		if err != nil {
			return nil, fmt.Errorf("сouldn't insert new value in the 'whatsapp_business_accounts' table, err: %q", err.Error())
		}
	default:
		//
	}

	return c, nil
}

// UpdateChannel
// Updates the specific channel information in database.
func (cr *ChannelRepository) UpdateChannel(c *common.Channel) (*common.Channel, error) {
	var channelTypeName string
	err := cr.db.QueryRowx(`
		select
			lower(channel_type_name) as channel_type_name
		from
			channel_types
		where
			channel_type_id = $1
		limit 1;`, &c.ChannelTypeId).Scan(&channelTypeName)
	if err != nil {
		return nil, fmt.Errorf("unexpected channel type, err: %q", err.Error())
	}

	availableChannelTypes := []string{"telegram", "facebook_messenger", "whatsapp", "instagram_private"}
	if !stringExistsInArray(availableChannelTypes, channelTypeName) {
		return nil, fmt.Errorf("updating the channel for the %q type is currently not possible", channelTypeName)
	}

	_, err = cr.db.Exec(`
		update
	    	channels
		set
	    	channel_name = $2,
			channel_description = $3,
			channel_type_id = $4,
			channel_technical_id = $5,
			channel_status_id = $6
		where
			channel_id = $1;`,
			&c.ChannelId,
			&c.ChannelName,
			&c.ChannelDescription,
			&c.ChannelTypeId,
			&c.ChannelTechnicalId,
			&c.ChannelStatusId)
	if err != nil {
		return nil, fmt.Errorf("failed to update the channel information, err: %q", err.Error())
	}

	_, err = cr.db.Exec(`delete from channels_organizations_relationship where channel_id = $1;`, &c.ChannelId)
	if err != nil {
		return nil, fmt.Errorf("failed to delete the link between channel and organizations, err: %q", err.Error())
	}

	_, err = cr.db.Exec(`
		insert into channels_organizations_relationship(
			channel_id,
			organization_id
		)
		select
			$1::uuid channel_id,
			organizations_ids
		from
			unnest($2::uuid[]) organizations_ids;`,
			&c.ChannelId,
			&c.OrganizationsIds)
	if err != nil {
		return nil, fmt.Errorf("сouldn't link channel to organizations, err: %q", err.Error())
	}

	switch channelTypeName {
	case "telegram":
		_, err = cr.db.Exec(`delete from telegram_business_accounts where channel_id = $1;`, &c.ChannelId)
		if err != nil {
			return nil, fmt.Errorf("сouldn't insert new value in the 'telegram_business_accounts' table, err: %q", err.Error())
		}

		_, err = cr.db.Exec(`
			insert into telegram_business_accounts (
				business_account,
				channel_id
			)
			values (
				$1,
				$2
			);`,
			&c.ChannelName,
			&c.ChannelId)
		if err != nil {
			return nil, fmt.Errorf("сouldn't insert new value in the 'telegram_business_accounts' table, err: %q", err.Error())
		}

		err = SetWebhookToTelegram(*c.ChannelName, *c.ChannelTechnicalId)
		if err != nil {
			return nil, fmt.Errorf("couldn't set webhook via telegram api, err: %q", err.Error())
		}
	case "facebook_messenger":
		_, err = cr.db.Exec(`delete from facebook_messenger_business_accounts where channel_id = $1;`, &c.ChannelId)
		if err != nil {
			return nil, fmt.Errorf("сouldn't insert new value in the 'telegram_business_accounts' table, err: %q", err.Error())
		}

		_, err = cr.db.Exec(`
			insert into facebook_messenger_business_accounts (
				business_account,
				channel_id
			)
			values (
				$1,
				$2
			);`,
			&c.ChannelName,
			&c.ChannelId)
		if err != nil {
			return nil, fmt.Errorf("сouldn't insert new value in the 'facebook_messenger_business_accounts' table, err: %q", err.Error())
		}

		err = SetWebhookToFacebookMessenger()
		if err != nil {
			return nil, fmt.Errorf("couldn't set webhook via facebook graph api, err: %q", err.Error())
		}
	case "instagram_private":
		_, err = cr.db.Exec(`delete from instagram_private_business_accounts where channel_id = $1;`, &c.ChannelId)
		if err != nil {
			return nil, fmt.Errorf("сouldn't insert new value in the 'telegram_business_accounts' table, err: %q", err.Error())
		}

		_, err = cr.db.Exec(`
			insert into instagram_private_business_accounts (
				business_account,
				channel_id
			)
			values (
				$1,
				$2
			);`,
			&c.ChannelName,
			&c.ChannelId)
		if err != nil {
			return nil, fmt.Errorf("сouldn't insert new value in the 'instagram_private_business_accounts' table, err: %q", err.Error())
		}

		err = cr.SetInstagramPrivateSession(*c.ChannelId, *c.ChannelStatusId)
		if err != nil {
			return nil, fmt.Errorf("couldn't set sessions in the private instagram, err: %q", err.Error())
		}
	case "whatsapp":
		_, err = cr.db.Exec(`delete from whatsapp_business_accounts where channel_id = $1;`, &c.ChannelId)
		if err != nil {
			return nil, fmt.Errorf("сouldn't insert new value in the 'telegram_business_accounts' table, err: %q", err.Error())
		}

		_, err = cr.db.Exec(`
			insert into whatsapp_business_accounts (
				business_account,
				channel_id
			)
			values (
				concat('+', $1),
				$2
			);`,
			&c.ChannelName,
			&c.ChannelId)
		if err != nil {
			return nil, fmt.Errorf("сouldn't insert new value in the 'whatsapp_business_accounts' table, err: %q", err.Error())
		}
	default:
		//
	}

	return c, nil
}

// GetChannels
// Get the list of all channels of the specific organization.
func (cr *ChannelRepository) GetChannels(organizationId *string) (*[]common.Channel, error) {
	var channels []common.Channel

	rows, err := cr.db.Query(`
		select
			channels.channel_id::text,
			channels.channel_name::text,
			channels.channel_description::text,
			channels.channel_type_id::text,
			channels.channel_technical_id::text,
			channels.channel_status_id::text,
		    array_agg(distinct channels_organizations_relationship.organization_id) filter (where channels_organizations_relationship.organization_id is not null)
		from
			channels
		left join channels_organizations_relationship on
			channels.channel_id = channels_organizations_relationship.channel_id
		left join organizations on
			channels_organizations_relationship.organization_id = organizations.organization_id
		where
			organizations.tree_organization_id like concat('%', '\', $1::text, '%')
		group by
			channels.channel_id;`, *organizationId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var channel common.Channel

		if err = rows.Scan(&channel.ChannelId,
			&channel.ChannelName,
			&channel.ChannelDescription,
			&channel.ChannelTypeId,
			&channel.ChannelTechnicalId,
			&channel.ChannelStatusId,
			&channel.OrganizationsIds); err != nil {
			return nil, err
		}

		channels = append(channels, channel)
	}

	return &channels, nil
}

// GetChannel
// Get the information about the specific channel.
func (cr *ChannelRepository) GetChannel(channelId *string) (*common.Channel, error) {
	var channel common.Channel

	row := cr.db.QueryRow(`
		select
			channels.channel_id::text,
			channels.channel_name::text,
			channels.channel_description::text,
			channels.channel_type_id::text,
			channels.channel_technical_id::text,
			channels.channel_status_id::text,
			array_agg(distinct channels_organizations_relationship.organization_id) filter (where channels_organizations_relationship.organization_id is not null)
		from
			channels
		left join channels_organizations_relationship on
			channels.channel_id = channels_organizations_relationship.channel_id
		where
			channels.channel_id = $1
		group by
    		channels.channel_id
		limit 1;`, *channelId)

	if err := row.Scan(&channel.ChannelId,
		&channel.ChannelName,
		&channel.ChannelDescription,
		&channel.ChannelTypeId,
		&channel.ChannelTechnicalId,
		&channel.ChannelStatusId,
		&channel.OrganizationsIds); err != nil {
		return nil, err
	}

	return &channel, nil
}

