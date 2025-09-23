package test

import (
	"fmt"
	"mall/internal/core"
	"mall/internal/dao/db"
	"mall/internal/entity"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	config := "../configs/config.yaml"
	if err := core.InitConfig(config); err != nil {
		fmt.Printf("init config failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("init config success")
	code := m.Run()

	os.Exit(code)
}

func TestUserCreateSql(t *testing.T) {
	t.Log("TestUserCreateSql")

	dao := db.NewUserDbDao()
	user := &entity.User{
		NickName:  "test",
		Account:   "test",
		Password:  "test",
		Icon:      "test",
		Gender:    1,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	id, err := dao.CreateUser(t.Context(), *user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("id:", id)
}
