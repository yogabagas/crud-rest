package models

type Item struct {
	ID   int    `json:"tblItemID"`
	Code string `json:"ItemCode"`
	Nama string `json:"ItemName"`
}
