package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func Test_MYSQL(t *testing.T) {
	t.Log(GetMysqlClient())
}
