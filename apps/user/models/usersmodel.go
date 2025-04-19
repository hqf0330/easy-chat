package models

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		FindByPhone(ctx context.Context, phone string) (*Users, error)
		ListByName(ctx context.Context, name string) ([]*Users, error)
		ListByIds(ctx context.Context, ids []string) ([]*Users, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

func (c customUsersModel) ListByName(ctx context.Context, name string) ([]*Users, error) {
	query := fmt.Sprintf("select %s from %s where `nickname` like ? ", usersRows, c.table)
	var resp []*Users
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, fmt.Sprint("%", name, "%"))
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (c customUsersModel) ListByIds(ctx context.Context, ids []string) ([]*Users, error) {
	query := fmt.Sprintf("select %s from %s where `id` in ('%s') ", usersRows, c.table, strings.Join(ids, "','"))

	var resp []*Users
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (c customUsersModel) FindByPhone(ctx context.Context, phone string) (*Users, error) {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, phone)
	var resp Users
	err := c.QueryRowCtx(ctx, &resp, usersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `phone` = ? limit 1", usersRows, c.table)
		return conn.QueryRowCtx(ctx, v, query, phone)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c, opts...),
	}
}
