package models

type CodeMessage struct {
	Code int    `json:"code" required:"true" description:"code"`
	Msg  string `json:"message" required:"true" description:"code description"`
}
