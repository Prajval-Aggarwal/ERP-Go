package request

type CreateChannelRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Type        string `json:"type"`
}

type AddMemeber struct {
	ChannelId string   `json:"channel_id"`
	EmailIds  []string `json:"email_id"`
}

type DeleteInactiveMemeber struct {
	UserEmail string `json:"user_email"`
}
