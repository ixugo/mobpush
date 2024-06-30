package mobpush

type PushForward struct {
	NextType       int       `json:"nextType"`
	URL            string    `json:"url"`
	Scheme         string    `json:"scheme"`
	SchemeDataList []PushMap `json:"schemeDataList"`
}

type PushFactoryExtra struct {
	XiaomiExtra XiaomiExtra `json:"xiaomiExtra"`
}

type XiaomiExtra struct {
	ChannelID string `json:"channelId"`
}
