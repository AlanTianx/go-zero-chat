// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"strings"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stringx"
	"gorm.io/gorm"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUsersIdPrefix   = "cache:users:id:"
	cacheUsersUuidPrefix = "cache:users:uuid:"
)

type (
	usersModel interface {
		Insert(ctx context.Context, data *Users) error
		FindOne(ctx context.Context, id int64) (*Users, error)
		FindOneByUuid(ctx context.Context, uuid string) (*Users, error)
		FindOneByPhone(ctx context.Context, phone string) (*Users, error)
		FirstOrCreateByPhone(ctx context.Context, u Users) (*Users, error)
		Update(ctx context.Context, data *Users) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUsersModel struct {
		gormc.CachedConn
		table string
	}

	Users struct {
		Id        int64        `gorm:"column:id"`
		Uuid      string       `gorm:"column:uuid"`
		Name      string       `gorm:"column:name"`  // 用户名
		Phone     string       `gorm:"column:phone"` // 手机号
		CreatedAt sql.NullTime `gorm:"column:created_at"`
		UpdatedAt sql.NullTime `gorm:"column:updated_at"`
	}
)

func newUsersModel(conn *gorm.DB, c cache.CacheConf) *defaultUsersModel {
	return &defaultUsersModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`users`",
	}
}

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) error {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)
	usersUuidKey := fmt.Sprintf("%s%v", cacheUsersUuidPrefix, data.Uuid)
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Save(&data).Error
	}, usersIdKey, usersUuidKey)
	return err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id int64) (*Users, error) {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, id)
	var resp Users
	err := m.QueryCtx(ctx, &resp, usersIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Users{}).Where("`id` = ?", id).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByUuid(ctx context.Context, uuid string) (*Users, error) {
	usersUuidKey := fmt.Sprintf("%s%v", cacheUsersUuidPrefix, uuid)
	var resp Users
	err := m.QueryRowIndexCtx(ctx, &resp, usersUuidKey, m.formatPrimary, func(conn *gorm.DB, v interface{}) (interface{}, error) {
		if err := conn.Model(&Users{}).Where("`uuid` = ?", uuid).Take(&resp).Error; err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByPhone(ctx context.Context, phone string) (*Users, error) {
	usersPhoneKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, phone)
	var resp Users
	err := m.QueryCtx(ctx, &resp, usersPhoneKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Users{}).Where("`phone` = ?", phone).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FirstOrCreateByPhone(ctx context.Context, u Users) (*Users, error) {
	//var resp Users
	if u.Phone == "" {
		return nil, errors.New("参数错误")
	}
	err := m.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
		return conn.Model(&Users{}).Where("phone = ?", u.Phone).FirstOrCreate(&u).Error
	})
	switch err {
	case nil:
		return &u, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Update(ctx context.Context, data *Users) error {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, data.Id)
	usersUuidKey := fmt.Sprintf("%s%v", cacheUsersUuidPrefix, data.Uuid)
	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Select("name").Save(data).Error
	}, usersIdKey, usersUuidKey)
	return err
}

func (m *defaultUsersModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, id)
	usersUuidKey := fmt.Sprintf("%s%v", cacheUsersUuidPrefix, data.Uuid)
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Delete(&Users{}, id).Error
	}, usersIdKey, usersUuidKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUsersIdPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&Users{}).Where("`id` = ?", primary).Take(v).Error
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}

func (Users) TableName() string {
	return "users"
}
