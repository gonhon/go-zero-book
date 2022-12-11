package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ClassModel = (*customClassModel)(nil)

type (
	// ClassModel is an interface to be customized, add more methods here,
	// and implement the added methods in customClassModel.
	ClassModel interface {
		classModel
	}

	customClassModel struct {
		*defaultClassModel
	}
)

// NewClassModel returns a model for the database table.
func NewClassModel(conn sqlx.SqlConn) ClassModel {
	return &customClassModel{
		defaultClassModel: newClassModel(conn),
	}
}
