package stock

import "database/sql"

type Dao interface {
	// SetDB 设置DB
	SetDB(db *sql.DB)

	// ReadAllAStockBases 读取所有A股股本信息到数据库
	ReadAllAStockBases() (list []AStockBase, err error)

	// WriteAllAStockBases 写入所有A股信息到数据库
	WriteAllAStockBases(aStockBases ...AStockBase) error

	// WriteAStockGuBen 写入所有A股股本信息到数据库
	WriteAStockGuBen(guBenInfo GuBenInfo) error
}
