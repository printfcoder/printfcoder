package stock

import "context"

type Syncer interface {
	// Init 初始化
	Init(opts ...SyncerOption) error
	// Name 同步插件名字
	Name() string
	// Sync 同步汇总
	Sync() error
	// SyncAllStockBases 同步所有股票基本信息
	SyncAllStockBases() error
	// SyncGuBen 同步股本
	SyncGuBen() error
}

type SyncerOption func(o *SyncerOptions)

type SyncerOptions struct {
	Dao Dao
}

func WithDao(db Dao) SyncerOption {
	return func(o *SyncerOptions) {
		o.Dao = db
	}
}

func SyncerAdapter(ctx context.Context) Syncer {

}
