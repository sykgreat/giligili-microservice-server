// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserIdPrefix = "cache:user:id:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByEmail(ctx context.Context, email string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id int64) error

		TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id         int64          `db:"id"`
		CreatedAt  sql.NullTime   `db:"created_at"`
		UpdatedAt  sql.NullTime   `db:"updated_at"`
		DeletedAt  sql.NullTime   `db:"deleted_at"`
		Username   sql.NullString `db:"username"`    // 用户名
		Email      string         `db:"email"`       // 邮箱
		Password   sql.NullString `db:"password"`    // 密码
		Avatar     sql.NullString `db:"avatar"`      // 头像
		SpaceCover sql.NullString `db:"space_cover"` // 空间封面
		Gender     int64          `db:"gender"`      // 性别:0未知、1男、3女
		Birthday   time.Time      `db:"birthday"`    // 生日
		Sign       string         `db:"sign"`        // 个性签名
		ClientIp   sql.NullString `db:"client_ip"`   // 客户端IP
		Status     string         `db:"status"`
		Role       int64          `db:"role"`
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userIdKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, userIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultUserModel) FindOneByEmail(ctx context.Context, email string) (*User, error) {
	var resp User
	err := m.QueryRowCtx(ctx, &resp, email, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, email)
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

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.DeletedAt, data.Username, data.Email, data.Password, data.Avatar, data.SpaceCover, data.Gender, data.Birthday, data.Sign, data.ClientIp, data.Status, data.Role)
	}, userIdKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.Username, data.Email, data.Password, data.Avatar, data.SpaceCover, data.Gender, data.Birthday, data.Sign, data.ClientIp, data.Status, data.Role, data.Id)
	}, userIdKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
