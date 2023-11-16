package models

type Point struct {
	Id        uint    `json:"id"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Activated bool    `json:"activated"`
	Materials bool    `json:"materials"`
	Days      uint32  `json:"days"`
	Approved  uint32  `json:"approved"`
	Cards     uint32  `json:"cards"`
}
