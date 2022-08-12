package response

import "time"

type LogList struct {
	Id          int       `json:"id"`
	Username    string    `json:"username" xorm:"'username'"`
	Method      string    `json:"method"`
	Action      string    `json:"action"`
	Ip          string    `json:"ip"`
	StatusCode  int       `json:"statusCode"`
	Params      string    `json:"params"`
	Results     string    `json:"results"`
	CreatedTime time.Time `json:"createdTime"`
}
