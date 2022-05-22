package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"time"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	err := dao()
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
	}
	return
}

func dao() error {
	id := 123
	var username string
	var created time.Time
	var query = "SELECT username, created_at FROM users WHERE id=?"
	err := db.QueryRowContext(ctx, query, id).Scan(&username, &created)
	//此处wrap错误后返回上层
	//因为：1.查询错误后 此层无法正常推进后续工作 属于功能错误
	//2.此处的doa层为applications类型的，不属于最高可重用性的包，且报错来自于协作的其他库，故使用wrap报装出错的sql的信息，返回给调用者，让调用者自己处理
	if err != nil {
		return errors.Wrapf(err, `failed to query %q`, query)
	}
	return nil
}
