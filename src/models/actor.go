package models

type Actor struct {
	First_name string `gorm:"column:first_name"`
	Last_name  string `gorm:"column:last_name"`
}

func (Actor) TableName() string {
	return "actor"
}
