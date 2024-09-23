package domain

type Permissions struct {
	ID     int    `json:"id"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status int    `json:"status"`
}
