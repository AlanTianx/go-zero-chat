package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UserTokensModel = (*customUserTokensModel)(nil)

type (
	// UserTokensModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserTokensModel.
	UserTokensModel interface {
		userTokensModel
		Save(ctx context.Context, data *UserTokens) error
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

func (m *defaultUserTokensModel) Save(ctx context.Context, data *UserTokens) error {
	userTokensIdKey := fmt.Sprintf("%s%v", cacheUserTokensKeyPrefix, data.TokenKey)
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Select("status").Save(data).Error
	}, userTokensIdKey)
	return err
}
