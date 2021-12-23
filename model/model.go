package model

type User struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Jk         uint   `json:"jk"`
	Role       uint   `json:"role"`
	Domisili   uint   `json:"domisili"`
	Kota_pelak uint   `json:"kota_pelak"`
	Email      string `json:"email" gorm:"unique"`
	Password   []byte `json:"-"`
}

type JawabTest struct {
	Id      uint `json:"id"`
	Soal    uint `json:"soal"`
	Jawaban uint `json:"jawaban"`
}
