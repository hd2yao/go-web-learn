package table

type User struct {
	Id        int64  `gorm:"column:id;primary_key"`
	UserName  string `gorm:"column:username"`
	Password  string `gorm:"column:password;type:varchar(1000)"`
	CreatedAt string `gorm:"column:created_at"`
	//UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (u *User) TableName() string {
	return "users"
}
