package shared

type AccountDataList struct {
	Data []AccountDataListItem `json:"data"`
}

type AccountDataListItem struct {
	Uid                 string `json:"uid"`
	Name                string `json:"name"`
	Email               string `json:"email"`
	Mobile              string `json:"mobile"`
	IsEmailVerify       bool   `json:"is_email_verify"`
	Realname            string `json:"realname"`
	IdentityCard        string `json:"identity_card"`
	Token               string `json:"token"`
	IsGuest             bool   `json:"is_guest"`
	GuestId             string `json:"guest_id"`
	SafeMobile          string `json:"safe_mobile"`
	Account             string `json:"account"`
	IsLogin             bool   `json:"is_login"`
	LoginType           int    `json:"login_type"`
	Payload             string `json:"payload"`
	ChannelId           int    `json:"channel_id"`
	AsteriskName        string `json:"asterisk_name"`
	AccessToken         string `json:"accessToken"`
	DeviceId            string `json:"deviceId"`
	Country             string `json:"country"`
	AreaCode            string `json:"area_code"`
	ThirdLoginTimestamp int    `json:"thirdLoginTimestamp"`
}
