package test

import (
	"mall/internal/core"
	"mall/internal/dao/db"
	"testing"
)

func TestRawSql(t *testing.T) {
	t.Log("TestMallUser")

	if err := core.InitConfig("../configs/config.yaml"); err != nil {
		t.Error(err)
	}

	dao := db.NewUserDbDao()

	sql, err := dao.GetRawSql(t.Context())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("sql:", sql)
}
