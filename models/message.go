package models

type Message struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
