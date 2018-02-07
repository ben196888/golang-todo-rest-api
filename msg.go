package main

type jsonMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
}
