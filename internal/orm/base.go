package orm

import (
	"context"
	"go-base/internal/pkg/common"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DatetimeFormat 时间格式
var DatetimeFormat = "2006-01-02 15:04:05"
var TimeZero = "0000-00-00 00:00:00"

type IdModel struct {
	Id int64 `json:"id" gorm:"column:id;type:bigint(20);not null;primaryKey;autoIncrement"`
}

type TimestampModel struct {
	CreatedAt int64 `json:"created_at" gorm:"column:created_at;type:int(10);not null;default:0"`
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at;type:int(10);not null;default:0"`
}

type DeletedAtModel struct {
	DeletedAt int64 `json:"deleted_at" gorm:"column:deleted_at;type:int(10);not null;default:0"`
}

type GormConnWithContext func(ctx context.Context, db *gorm.DB) *gorm.DB

type GormConn func(db *gorm.DB) *gorm.DB

// GetConn 获取查询连接
func GetConnWithContext(table schema.Tabler) GormConnWithContext {
	return func(ctx context.Context, db *gorm.DB) *gorm.DB {
		return instance.WithContext(ctx).Table(table.TableName())
	}
}

func GetConn(table schema.Tabler) GormConn {
	return func(db *gorm.DB) *gorm.DB {
		return instance.Table(table.TableName())
	}
}

// GetTableConnWithCtx 获取带table的连接
func GetTableConnWithCtx(ctx context.Context, table schema.Tabler) *gorm.DB {
	return instance.WithContext(ctx).Table(table.TableName())
}

// GetTableConn 获取带table的连接
func GetTableConn(table schema.Tabler) *gorm.DB {
	return instance.Table(table.TableName())
}

// GetTableTrans 获取带table的事务连接
func GetTableTrans(table schema.Tabler, tx *gorm.DB) *gorm.DB {
	return tx.Table(table.TableName())
}

// ExecTrans 执行事务
func ExecTrans(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return instance.WithContext(ctx).Transaction(fn)
}

// WithPagination 翻页查询
func WithPagination(pg common.Pagination) GormConn {
	pg.HandleOffset()
	return func(db *gorm.DB) *gorm.DB {
		if pg.Size > 0 {
			return db.Offset(pg.Offset).Limit(pg.Size)
		}
		return db
	}
}

func WithNotDeleted() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("deleted_at = 0")
	}
}
