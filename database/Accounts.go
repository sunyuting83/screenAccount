package database

type Accounts struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Account   string `gorm:"index"`
	Password  string
	Phone     string `gorm:"index"`
	Gold      int
	Price     int
	User      User
	CreatedAt int64 `gorm:"autoUpdateTime:milli"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
}

// Get Count
func (accounts *Accounts) GetCount(UserID string) (count int64, err error) {
	if err = sqlDB.Model(&accounts).Where("user_id = ?", UserID).Count(&count).Error; err != nil {
		return
	}
	return
}

// Get Count
func (accounts *Accounts) FindOnePhone(phone string) (account Accounts, err error) {
	if err = sqlDB.Joins("User", "id = (?)", account.UserID).First(&account, "phone = ? ", phone).Error; err != nil {
		return
	}
	return
}
func (accounts *Accounts) FindOneAccount(acc string) (account Accounts, err error) {
	if err = sqlDB.Joins("User", "id = (?)", account.UserID).First(&account, "account = ?", acc).Error; err != nil {
		return
	}
	return
}

func AccountBatches(accounts []Accounts) {
	sqlDB.Create(&accounts)
}

// Account List
func GetAccountList(page int, UserID string, Limit int) (accounts *[]Accounts, err error) {
	p := makePage(page, Limit)
	if err = sqlDB.
		Where("user_id = ?", UserID).
		Order("accounts.created_at desc").
		Limit(Limit).Offset(p).
		Joins("User", "id = (?)", UserID).
		Find(&accounts).Error; err != nil {
		return
	}
	return
}

// makePage make page
func makePage(p, Limit int) int {
	p = p - 1
	if p <= 0 {
		p = 0
	}
	page := p * Limit
	return page
}
