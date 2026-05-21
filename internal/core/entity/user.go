package entity

type User struct {
	BaseEntity
	Username string `gorm:"unique;not null;type:varchar(20)"`
	Email    string `gorm:"unique;not null;type:varchar(100)"`
	Password string `gorm:"not null;type:text"`
	Role     string `gorm:"type:varchar(20);default:'user';check:role IN ('admin', 'user')"`
}
