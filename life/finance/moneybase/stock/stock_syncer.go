package stock

import (
	"context"
	"sync"

	log "github.com/stack-labs/stack/logger"
)

var (
	syncers   = map[string]Syncer{}
	syncerMux sync.Mutex
)

type Syncer interface {
	// Init 初始化
	Init(opts ...SyncerOption) error

	// Name 同步插件名字
	Name() string

	// Sync 同步汇总
	Sync() error

	// SyncAllStockBases 同步所有股票基本信息
	SyncAllStockBases() error

	// SyncAllStockGuBen 同步所有股票股本
	SyncAllStockGuBen() error

	// SyncSingleStockGuBen 同步单支股本
	SyncSingleStockGuBen(symbol string) error

	// GetStockQT 获取股票当前价值
	GetStockQT(date string, symbol ...string) ([]StockQTDataTencent, error)

	// GetStockTop10GuDong 获取 Top10 股东
	GetStockTop10GuDong(symbol string, startDate string, endDate string, guDongType int) (data []StockTop10GuDong, err error)

	// WriteSingleStockQTDaily 写入单个每天QT
	WriteSingleStockQTDaily(date string, symbol string) error

	// WriteStockQTDaily 写入每天QT
	WriteStockQTDaily(date string) error

	// MethodSupported 是否支持该方法
	MethodSupported(methodName string) (supported bool, err error)

	// SyncTop10GuDong 写入 Top10 股东
	SyncTop10GuDong(startDate string, endDate string, guDongType int) error
}

type SyncerOption func(o *SyncerOptions)

type SyncerOptions struct {
	Dao  Dao
	Name string
}

func WithDao(db Dao) SyncerOption {
	return func(o *SyncerOptions) {
		o.Dao = db
	}
}

func SyncerAdapter(ctx context.Context) Syncer {
	methodName, syncerName, err := getMethodNameFromHTTP(ctx)
	if err != nil {
		// todo 优化
	}

	if methodName == "" {
		// todo 优化
	}

	for _, syncer := range syncers {
		if ok, err := syncer.MethodSupported(methodName); ok {
			if syncerName != "" && syncer.Name() != syncerName {
				log.Errorf("[SyncerAdapter] syncer[%s] is not the one [%s]", syncer.Name(), syncerName)
				continue
			}
			return syncer
		} else {
			log.Errorf("[SyncerAdapter] check syncer[%s] method support err: %s", syncer.Name(), err)
			continue
		}
	}

	panic("no syncer")
}
