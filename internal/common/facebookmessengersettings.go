package common

type CategoryList struct {
	Id *string `json:"id"`
	Name *string `json:"name"`
}
type Data struct {
	AccessToken *string `json:"access_token"`
	Category *string `json:"category"`
	CategoryList []*CategoryList `json:"category_list"`
	Name *string `json:"name"`
	Id *string `json:"id"`
	Tasks []*string
}

type Paging struct {
	Cursors *Cursors `json:"cursors"`
}

type Cursors struct {
	Before *string `json:"before"`
	After *string `json:"after"`
}

// FacebookPages
// Contains information about facebook pages.
type FacebookPages struct {
	Data []*Data `json:"data"`
	Paging *Paging `json:"paging"`
}
