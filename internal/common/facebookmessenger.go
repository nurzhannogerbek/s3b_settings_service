package common

type CategoryList struct {
	Id   *string `json:"id"`
	Name *string `json:"name"`
}

type Data struct {
	AccessToken  *string         `json:"access_token"`
	Category     *string         `json:"category"`
	CategoryList []*CategoryList `json:"category_list"`
	Name         *string         `json:"name"`
	Id           *string         `json:"id"`
	Tasks        []*string       `json:"tasks"`
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

// FacebookMessengerUser
// The facebook messenger user information with 'Long Lived User Access Token'.
type FacebookMessengerUser struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
