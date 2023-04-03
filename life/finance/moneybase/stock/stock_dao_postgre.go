package stock

import (
	"database/sql"
	"time"

	"github.com/printfcoder/printfcoder/life/finance/moneybase/common"
	log "github.com/stack-labs/stack/logger"
)

type DaoPostgre struct {
	db *sql.DB
}

func (d *DaoPostgre) SetDB(db *sql.DB) {
	d.db = db
}

func (d *DaoPostgre) WriteAllAStocks(aStockBases ...AStockBase) error {
	if d.db == nil {
		return common.ErrorDBNil
	}

	log.Infof("get stock base: %+v", aStockBases)

	tx, err := d.db.Begin()
	if err != nil {
		log.Errorf("写入股所有股票基本信息，开启事务异常。err: %s", err)
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	stmt, err := tx.Prepare("INSERT INTO stock_base(dm, mc, jys, update_time) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Errorf("写入股所有股票基本信息，准备批量插入异常。err: %s", err)
		return err
	}

	version := time.Now()

	// 执行插入
	for _, row := range aStockBases {
		_, err = stmt.Exec(row.DM, row.MC, row.JYS, version)
		if err != nil {
			log.Errorf("写入股所有股票基本信息，插入数据库异常。err: %s", err)
			return err
		}
	}

	result, err := tx.Exec("DELETE FROM stock_base WHERE update_time < $1 ", version)
	if err != nil {
		log.Errorf("写入股所有股票基本信息，删除老版本异常。err: %s", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	log.Infof("写入股所有股票基本信息，插入新数据[%d]条，清空老数据[%d]条", len(aStockBases), rowsAffected)

	// 提交事务
	err = tx.Commit()
	if err != nil {
		log.Errorf("写入股所有股票基本信息，事务提交异常。err: %s", err)
		return err
	}

	return nil
}
