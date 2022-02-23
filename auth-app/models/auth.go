package models

type Auth struct {
	BaseModel
	Phone    string `gorm:"uniqueIndex"`
	Name     string
	Password string
	Role     string
}

var Roles = []string{
	"admin", "staff",
}
