package models

type User struct {
	Id             uint    `json:"id"`
	Name           string  `json:"name"`
	Email          string  `json:"email" gorm:"unique"`
	Password       []byte  `json:"-"`
	Longitude      float64 `json:"longitude"`
	Latitude       float64 `json:"latitude"`
	Grade          string  `json:"grade"`
	AvailableHours float64 `json:"hours"`
}
