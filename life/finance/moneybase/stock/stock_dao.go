package stock

import "database/sql"

type Dao interface {
	// SetDB 设置DB
	SetDB(db *sql.DB)

	// WriteAllAStocks 写入所有A股信息到数据库
	WriteAllAStocks(aStockBases ...AStockBase) error
}
