package proto

type PubInfo struct {
	MediaInfo
	Info        ClientUserInfo `json:"info"`
	Tracks      TrackMap       `json:"tracks"`
	Description string         `json:"description,omitempty"`
}

type GetPubResp struct {
	RoomInfo
	Pubs []PubInfo `json:"pubs,omitempty"`
}

type GetMediaParams struct {
	RID RID
	MID MID
}

type FindServiceParams struct {
	Service string
	MID     MID
	RID     RID
}

type GetSFURPCParams struct {
	RPCID   string
	EventID string
	ID      string
	Name    string
	Service string
}

type UserInfoResp struct {
	ClientUserInfo
	UserInfo
}

type RoomIdParams struct {
	RID RID `json:"rid"`
}
