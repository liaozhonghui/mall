package db

import (
	"context"
	"mall/internal/entity"
	"time"

	"gorm.io/gorm"
)

type MallUser struct {
	ID        int       `gorm:"primaryKey;column:id"`
	NickName  string    `gorm:"column:nick_name"`
	Account   string    `gorm:"column:account"`
	Password  string    `gorm:"column:password"`
	Icon      string    `gorm:"column:icon"`
	Gender    int       `gorm:"column:gender"`
	Status    int       `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

const TABLE_MALL_USER = "mall_user"

type UserDbDao struct {
	Db *gorm.DB
}

func NewUserDbDao() UserDbDao {
	r := UserDbDao{
		Db: GetDbInstance("default"),
	}
	return r
}
func (dao UserDbDao) WithDBInstance(db *gorm.DB) UserDbDao {
	dao.Db = db
	return dao
}

func (dao *UserDbDao) CreateUser(ctx context.Context, user entity.User) (int, error) {
	mallUser := entityToDbUser(user)
	db := dao.Db.Table(TABLE_MALL_USER).Create(&mallUser)
	if db.Error != nil {
		return 0, db.Error
	}
	return mallUser.ID, nil
}

func (dao *UserDbDao) GetUserByAccount(ctx context.Context, account string, password string) (user entity.User, err error) {
	var mallUser MallUser
	db := dao.Db.Table(TABLE_MALL_USER).WithContext((ctx)).Where("account = ?", account).Where("password = ?", password).First(&mallUser)
	if db.Error != nil {
		return entity.User{}, db.Error
	}
	return dbUserToEntity(mallUser), nil
}

func (dao *UserDbDao) FindUserById(ctx context.Context, id int) (user entity.User, err error) {
	var mu MallUser
	db := dao.Db.Table(TABLE_MALL_USER).WithContext((ctx)).Where("id = ?", id).First(&mu)
	if db.Error != nil {
		return entity.User{}, db.Error
	}
	return dbUserToEntity(mu), nil
}

func entityToDbUser(user entity.User) MallUser {
	return MallUser{
		ID:        user.Id,
		NickName:  user.NickName,
		Account:   user.Account,
		Password:  user.Password,
		Icon:      user.Icon,
		Gender:    user.Gender,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
func dbUserToEntity(mu MallUser) entity.User {
	return entity.User{
		Id:        mu.ID,
		NickName:  mu.NickName,
		Account:   mu.Account,
		Password:  mu.Password,
		Icon:      mu.Icon,
		Gender:    mu.Gender,
		Status:    mu.Status,
		CreatedAt: mu.CreatedAt,
		UpdatedAt: mu.UpdatedAt,
	}
}
