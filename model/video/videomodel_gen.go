// Code generated by goctl. DO NOT EDIT.

package video

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
	videoFieldNames          = builder.RawFieldNames(&Video{})
	videoRows                = strings.Join(videoFieldNames, ",")
	videoRowsExpectAutoSet   = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	videoRowsWithPlaceHolder = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheVideoIdPrefix = "cache:video:id:"
)

type (
	videoModel interface {
		Insert(ctx context.Context, data *Video) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Video, error)
		Update(ctx context.Context, data *Video) error
		Delete(ctx context.Context, id int64) error
	}

	defaultVideoModel struct {
		sqlc.CachedConn
		table string
	}

	Video struct {
		Id          int64         `db:"id"`
		CreatedAt   sql.NullTime  `db:"created_at"`
		UpdatedAt   sql.NullTime  `db:"updated_at"`
		DeletedAt   sql.NullTime  `db:"deleted_at"`
		Title       string        `db:"title"` // 标题
		Cover       string        `db:"cover"`
		Desc        string        `db:"desc"`         // 视频简介
		Uid         int64         `db:"uid"`          // 用户ID
		Copyright   int64         `db:"copyright"`    // 是否为原创
		Clicks      int64         `db:"clicks"`       // 点击量
		Status      int64         `db:"status"`       // 审核状态
		PartitionId sql.NullInt64 `db:"partition_id"` // 分区ID
	}
)

func newVideoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultVideoModel {
	return &defaultVideoModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`video`",
	}
}

func (m *defaultVideoModel) Delete(ctx context.Context, id int64) error {
	videoIdKey := fmt.Sprintf("%s%v", cacheVideoIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, videoIdKey)
	return err
}

func (m *defaultVideoModel) FindOne(ctx context.Context, id int64) (*Video, error) {
	videoIdKey := fmt.Sprintf("%s%v", cacheVideoIdPrefix, id)
	var resp Video
	err := m.QueryRowCtx(ctx, &resp, videoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
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

func (m *defaultVideoModel) Insert(ctx context.Context, data *Video) (sql.Result, error) {
	videoIdKey := fmt.Sprintf("%s%v", cacheVideoIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, videoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.Title, data.Cover, data.Desc, data.Uid, data.Copyright, data.Clicks, data.Status, data.PartitionId)
	}, videoIdKey)
	return ret, err
}

func (m *defaultVideoModel) Update(ctx context.Context, data *Video) error {
	videoIdKey := fmt.Sprintf("%s%v", cacheVideoIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, videoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.Title, data.Cover, data.Desc, data.Uid, data.Copyright, data.Clicks, data.Status, data.PartitionId, data.Id)
	}, videoIdKey)
	return err
}

func (m *defaultVideoModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheVideoIdPrefix, primary)
}

func (m *defaultVideoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultVideoModel) tableName() string {
	return m.table
}
