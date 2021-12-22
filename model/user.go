package model

type User struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Phone_Number string `json:"phone_number"`
	GenderID     uint   `json:"gender_id"`
	RoleID       uint   `json:"role_id"`
	CityID       uint   `json:"city_id"`
	City_EventID uint   `json:"city_event_id"`
	Email        string `json:"email" gorm:"unique"`
	Password     []byte `json:"-"`
}
