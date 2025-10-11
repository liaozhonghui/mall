package pg

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Email     string    `gorm:"column:email" json:"email"`
	Name      string    `gorm:"column:name" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

const TABLE_NAME = "public.employee"

type UserDbDao struct {
	Db *gorm.DB
}

func NewUserDao() *UserDbDao {
	r := &UserDbDao{
		Db: GetInstance(),
	}
	return r
}

func (dao *UserDbDao) FindUserByEmail(ctx context.Context, email string) (userId int, err error) {
	var user UserModel
	result := dao.Db.Table(TABLE_NAME).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}
