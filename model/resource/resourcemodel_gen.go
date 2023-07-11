// Code generated by goctl. DO NOT EDIT.

package resource

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	resourceFieldNames          = builder.RawFieldNames(&Resource{})
	resourceRows                = strings.Join(resourceFieldNames, ",")
	resourceRowsExpectAutoSet   = strings.Join(stringx.Remove(resourceFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	resourceRowsWithPlaceHolder = strings.Join(stringx.Remove(resourceFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheResourceIdPrefix = "cache:resource:id:"
)

type (
	resourceModel interface {
		Insert(ctx context.Context, data *Resource) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Resource, error)
		Update(ctx context.Context, data *Resource) error
		Delete(ctx context.Context, id int64) error
	}

	defaultResourceModel struct {
		sqlc.CachedConn
		table string
	}

	Resource struct {
		Id          int64          `db:"id"`
		CreatedAt   sql.NullTime   `db:"created_at"`
		UpdatedAt   sql.NullTime   `db:"updated_at"`
		DeletedAt   sql.NullTime   `db:"deleted_at"`
		Vid         sql.NullInt64  `db:"vid"`          // 所属视频
		Uid         sql.NullInt64  `db:"uid"`          // 所属用户
		Title       sql.NullString `db:"title"`        // 分P使用的标题
		Url         sql.NullString `db:"url"`          // 视频链接
		Duration    float64        `db:"duration"`     // 视频时长
		Status      int64          `db:"status"`       // 审核状态
		Quality     sql.NullInt64  `db:"quality"`      // 视频最大质量
		OriginalUrl sql.NullString `db:"original_url"` // 原始mp4链接
	}
)

func newResourceModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultResourceModel {
	return &defaultResourceModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`resource`",
	}
}

func (m *defaultResourceModel) Delete(ctx context.Context, id int64) error {
	resourceIdKey := fmt.Sprintf("%s%v", cacheResourceIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, resourceIdKey)
	return err
}

func (m *defaultResourceModel) FindOne(ctx context.Context, id int64) (*Resource, error) {
	resourceIdKey := fmt.Sprintf("%s%v", cacheResourceIdPrefix, id)
	var resp Resource
	err := m.QueryRowCtx(ctx, &resp, resourceIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", resourceRows, m.table)
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

func (m *defaultResourceModel) Insert(ctx context.Context, data *Resource) (sql.Result, error) {
	resourceIdKey := fmt.Sprintf("%s%v", cacheResourceIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, resourceRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.Vid, data.Uid, data.Title, data.Url, data.Duration, data.Status, data.Quality, data.OriginalUrl)
	}, resourceIdKey)
	return ret, err
}

func (m *defaultResourceModel) Update(ctx context.Context, data *Resource) error {
	resourceIdKey := fmt.Sprintf("%s%v", cacheResourceIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, resourceRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.Vid, data.Uid, data.Title, data.Url, data.Duration, data.Status, data.Quality, data.OriginalUrl, data.Id)
	}, resourceIdKey)
	return err
}

func (m *defaultResourceModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheResourceIdPrefix, primary)
}

func (m *defaultResourceModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", resourceRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultResourceModel) tableName() string {
	return m.table
}
