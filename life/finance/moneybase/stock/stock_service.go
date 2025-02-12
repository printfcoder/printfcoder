package stock

import "context"

// SyncAllStockBases 从A股中同步所有股票基本信息
func SyncAllStockBases(ctx context.Context) error {
	return SyncerAdapter(ctx).SyncAllStockBases()
}

// SyncSingleStockGuBen 从A股中同步单个股票股本
func SyncSingleStockGuBen(ctx context.Context, code string) error {
	return SyncerAdapter(ctx).SyncSingleStockGuBen(code)
}

// SyncAllStockGuBen 从A股中同步所有股票股本
func SyncAllStockGuBen(ctx context.Context) error {
	return SyncerAdapter(ctx).SyncAllStockGuBen()
}

// GetStockQT 从A股获取股票QT信息
func GetStockQT(ctx context.Context, symbols ...string) ([]StockQTData, error) {
	return SyncerAdapter(ctx).GetStockQT(symbols...)
}
