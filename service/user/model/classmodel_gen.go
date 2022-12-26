// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	classFieldNames          = builder.RawFieldNames(&Class{})
	classRows                = strings.Join(classFieldNames, ",")
	classRowsExpectAutoSet   = strings.Join(stringx.Remove(classFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	classRowsWithPlaceHolder = strings.Join(stringx.Remove(classFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	classModel interface {
		Insert(ctx context.Context, data *Class) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Class, error)
		FindOneByCode(ctx context.Context, code string) (*Class, error)
		Update(ctx context.Context, data *Class) error
		Delete(ctx context.Context, id int64) error
		FindList(ctx context.Context,  param *Class) ([]Class, error)
	}

	defaultClassModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Class struct {
		Id         int64     `db:"id"`
		Num        int64     `db:"num"`  // 人数
		Name       string    `db:"name"` // 班级名称
		Code       string    `db:"code"` // 班级编号
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newClassModel(conn sqlx.SqlConn) *defaultClassModel {
	return &defaultClassModel{
		conn:  conn,
		table: "`class`",
	}
}

func (m *defaultClassModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultClassModel) FindOne(ctx context.Context, id int64) (*Class, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", classRows, m.table)
	var resp Class
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultClassModel) FindOneByCode(ctx context.Context, code string) (*Class, error) {
	var resp Class
	query := fmt.Sprintf("select %s from %s where `code` = ? limit 1", classRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, code)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultClassModel) Insert(ctx context.Context, data *Class) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, classRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Num, data.Name, data.Code)
	return ret, err
}

func (m *defaultClassModel) Update(ctx context.Context, newData *Class) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, classRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Num, newData.Name, newData.Code, newData.Id)
	return err
}


func (m *defaultClassModel) FindList(ctx context.Context, param *Class) ([]Class, error) {
	query := fmt.Sprintf("select %s from %s where `name` like ? limit 1", classRows, m.table)
	var resp =make([]Class,10)
	err := m.conn.QueryRowCtx(ctx, &resp, query,fmt.Sprintf("`%%%s%%`",param.Name))
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	} 
}

func (m *defaultClassModel) tableName() string {
	return m.table
}
