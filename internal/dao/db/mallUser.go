package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type MallUser struct {
	ID        int       `gorm:"primaryKey;column:id"`
	NickName  string    `gorm:"column:nickname"`
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

func (dao *UserDbDao) GetRawSql(ctx context.Context) (string, error) {
	var sql string
	db := dao.Db.Table(TABLE_MALL_USER).WithContext((ctx)).Where("account = ?", "test").Where("password = ?", "xxxxafdasfdfx")
	stmt := db.Session((&gorm.Session{DryRun: true})).Find(&MallUser{}).Statement
	if stmt.Error != nil {
		return "", stmt.Error
	}
	sql = stmt.SQL.String()
	return sql, nil
}
