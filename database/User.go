package database

type User struct {
	ID        uint `gorm:"primaryKey"`
	UserName  string
	Accounts  []Accounts
	CreatedAt int64 `gorm:"autoUpdateTime:milli"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
}

func (user *User) Insert() (err error) {
	sqlDB.Create(&user)
	return nil
}
func CheckUserName(username string) (user *User, err error) {
	if err = sqlDB.First(&user, "user_name = ? ", username).Error; err != nil {
		return
	}
	return
}

// Account List
func GetUserList() (user *[]User, err error) {
	if err = sqlDB.
		Select("id, user_name, created_at, updated_at").
		Order("updated_at desc").
		Find(&user).Error; err != nil {
		return
	}
	return
}
