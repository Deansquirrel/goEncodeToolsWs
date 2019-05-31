package object

//================================================================
type VersionResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Version string `json:"version"`
}

//================================================================
type MessageRequest struct {
	RequestText string `json:"requettext"`
	RequestKey  string `json:"requestkey"`
	OprType     int    `json:"oprtype"` //1-加密，2-解密
}

type MessageResponse struct {
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
	ResponseText string `json:"responsetext"`
}

//================================================================
