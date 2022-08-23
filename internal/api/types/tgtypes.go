package tgtypes

type TGMethod string
type GetBody map[string]string

// API METHODS
// getMe
type GetMeT struct {
	Ok     bool         `json:"ok"`
	Result GetMeResultT `json:"result"`
}
type GetMeResultT struct {
	Id        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

// getUpdates
type GetUpdatesT struct {
	Ok     bool                `json:"ok"`
	Result []GetUpdatesResultT `json:"result"`
}
type GetUpdatesResultT struct {
	UpdateID int                `json:"update_id"`
	Message  GetUpdatesMessageT `json:"message,omitempty"`
}
type GetUpdatesMessageT struct {
	MessageID int `json:"message_id"`
	From      struct {
		ID           int    `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Username     string `json:"username"`
		LanguageCode string `json:"language_code"`
	} `json:"from"`
	Chat struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Type      string `json:"type"`
	} `json:"chat"`
	Date     int    `json:"date"`
	Text     string `json:"text"`
	Entities []struct {
		Offset int    `json:"offset"`
		Length int    `json:"length"`
		Type   string `json:"type"`
	} `json:"entities,omitempty"`
	WebAppData []struct {
		Data       string `json:"data"`
		ButtonText string `json:"button_text"`
	} `json:"web_app_data,omitempty"`
}
