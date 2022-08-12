package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UserTokensModel = (*customUserTokensModel)(nil)

type (
	// UserTokensModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserTokensModel.
	UserTokensModel interface {
		userTokensModel
	}

	customUserTokensModel struct {
		*defaultUserTokensModel
	}
)

// NewUserTokensModel returns a model for the database table.
func NewUserTokensModel(conn *gorm.DB, c cache.CacheConf) UserTokensModel {
	return &customUserTokensModel{
		defaultUserTokensModel: newUserTokensModel(conn, c),
	}
}
