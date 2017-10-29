package dao

import (
	"allen-service/model"
	"sync"
	"testing"
)

var (
	once sync.Once
	d    *Dao
)

func startDao() {
	d = New()
}
func testInsert(t *testing.T) {
	u := &model.UserInfo{}
	d.Insert(u)
}
