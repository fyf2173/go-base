package orm

import (
	"context"
	"go-base/internal/pkg/sqlparser"
	"go-base/resources"
	"sort"
	"sync"
	"time"

	"github.com/fyf2173/ysdk-go/util"
	"github.com/fyf2173/ysdk-go/xdb"
	"github.com/fyf2173/ysdk-go/xlog"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var instance *gorm.DB

var once sync.Once

func InitConn() {
	once.Do(func() {
		if viper.GetString("data_driver") == "sqlite3" {
			db, err := gorm.Open(sqlite.Open(viper.GetString("sqliteDsn")), &gorm.Config{})
			if err != nil {
				panic(err)
			}
			instance = db
		}
		if viper.GetString("data_driver") == "mysql" {
			db, err := xdb.NewGorm(viper.GetString("env"), util.ViperMustGetNode("db", xdb.DbConfig{}))
			if err != nil {
				panic(err)
			}
			instance = db
		}
	})
	return
}

func DB() *gorm.DB {
	return instance
}

type GooseMigration struct {
	ID        int    `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Version   string `gorm:"column:version;NOT NULL"`
	CreatedAt int64  `gorm:"column:created_at;default:0;NOT NULL"`
	DeletedAt int64  `gorm:"column:deleted_at;default:0;NOT NULL"`
}

func (m *GooseMigration) TableName() string {
	return "gomig_migration"
}

func AutoMigrate(ctx context.Context) error {
	var migrateBlocks []sqlparser.VersionBlock

	// read all the sql files from the embed
	embedSQLFiles, err := resources.InstallationResource.ReadDir(".")
	if err != nil {
		xlog.Warn(ctx, "failed to read embed InstallSQL")
		return err
	}
	for _, file := range embedSQLFiles {
		sql, err := resources.InstallationResource.ReadFile(file.Name())
		if err != nil {
			xlog.Warn(ctx, "failed to read "+file.Name())
			return err
		}
		blocks := sqlparser.ParseSQLMigration(string(sql))
		migrateBlocks = append(migrateBlocks, blocks...)
	}

	// get all the migration history
	var histories []*GooseMigration
	instance.WithContext(ctx).Model(&GooseMigration{}).Where("deleted_at = 0").Order("created_at desc").Find(&histories)
	historyMap := util.SliceFieldFilteredMapWithKey(histories, func(val *GooseMigration) (string, string, error) { return val.Version, val.Version, nil })

	// sort by version asc and apply the migration
	sort.SliceStable(migrateBlocks, func(i, j int) bool { return migrateBlocks[i].Version < migrateBlocks[j].Version })
	for _, mig := range migrateBlocks {
		// skip if the migration has been applied
		if _, ok := historyMap[mig.Version]; ok {
			continue
		}
		// apply the migration
		if err := instance.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			for _, sql := range mig.Up {
				if err := tx.Exec(sql).Error; err != nil {
					return err
				}
			}
			for _, sql := range mig.Down {
				if err := tx.Exec(sql).Error; err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			return err
		}
		// record the migration history
		instance.WithContext(ctx).Create(&GooseMigration{
			Version:   mig.Version,
			CreatedAt: time.Now().Unix(),
		})
	}

	return nil
}
