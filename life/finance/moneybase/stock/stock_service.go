package stock

import "context"

// SyncAllStockBases 从A股中同步所有股票
func SyncAllStockBases(ctx context.Context) error {
	return SyncerAdapter(ctx).SyncAllStockBases()
}
