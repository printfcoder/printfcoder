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

func (d *DaoPostgre) ReadAllAStockBases() (ret []AStockBase, err error) {
	if d.db == nil {
		return nil, common.ErrorDBNil
	}

	rows, err := d.db.Query("SELECT DISTINCT dm, mc, jys FROM stock_base")
	if err != nil {
		log.Errorf("[ReadAllAStockBases] 读取股票基本信息异常。err: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		ab := AStockBase{}
		err = rows.Scan(&ab.DM, &ab.MC, &ab.JYS)
		if err != nil {
			log.Errorf("[ReadAllAStockBases] scan所有股票基本信息异常。err: %s", err)
			return nil, common.ErrorDBQueryScan
		}

		ret = append(ret, ab)
	}

	return
}

func (d *DaoPostgre) WriteAllAStockBases(aStockBases ...AStockBase) error {
	if d.db == nil {
		return common.ErrorDBNil
	}

	log.Infof("get stock base: %+v", aStockBases)

	tx, err := d.db.Begin()
	if err != nil {
		log.Errorf("[WriteAllAStockBases] 写入股所有股票基本信息，开启事务异常。err: %s", err)
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	stmt, err := tx.Prepare("INSERT INTO stock_base(dm, mc, jys, update_time) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Errorf("[WriteAllAStockBases] 写入股所有股票基本信息，准备批量插入异常。err: %s", err)
		return err
	}

	version := time.Now()

	// 执行插入
	for _, row := range aStockBases {
		_, err = stmt.Exec(row.DM, row.MC, row.JYS, version)
		if err != nil {
			log.Errorf("[WriteAllAStockBases] 写入股所有股票基本信息，插入数据库异常。err: %s", err)
			return err
		}
	}

	result, err := tx.Exec("DELETE FROM stock_base WHERE update_time < $1 ", version)
	if err != nil {
		log.Errorf("[WriteAllAStockBases] 写入股所有股票基本信息，删除老版本异常。err: %s", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	log.Infof("[WriteAllAStockBases] 写入股所有股票基本信息，插入新数据[%d]条，清空老数据[%d]条", len(aStockBases), rowsAffected)

	// 提交事务
	err = tx.Commit()
	if err != nil {
		log.Errorf("[WriteAllAStockBases] 写入股所有股票基本信息，事务提交异常。err: %s", err)
		return err
	}

	return nil
}

func (d *DaoPostgre) WriteAStockGuBen(guBenInfo GuBenInfo) error {
	if d.db == nil {
		return common.ErrorDBNil
	}

	tx, err := d.db.Begin()
	if err != nil {
		log.Errorf("[WriteAStockGuBen]，开启事务异常。err: %s", err)
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 插入股本，如果变动日期没有变动，则不用插入
	for _, v := range guBenInfo.GuBen {
		row, err := tx.Query(`SELECT 1 FROM stock_guben WHERE dai_ma = $1 AND bian_dong_ri_qi = $2`, v.GPDM, v.BDRQ)
		if err != nil {
			log.Errorf("[WriteAStockGuBen]，查询股本信息异常。err: %s", err)
			return err
		}

		// 不存在，则写入
		if !row.Next() {
			log.Infof("[WriteAStockGuBen] 未查询到股本信息 %s-%s-%s", v.GPDM, v.BDRQ)
			_, err = tx.Exec("INSERT INTO stock_guben (dai_ma, zong_gu_ben, liu_tong_gu_ben, ren_jun_chi_gu, gu_dong_ren_shu, bian_dong_ri_qi, g_g_ri_qi, shang_qi_gu_dong_ren_shu) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
				v.GPDM, v.ZGB, v.LTGB, v.RJCG, v.GDRS, v.BDRQ, v.GGRQ, v.SQ_GDRS,
			)

			if err != nil {
				log.Errorf("[WriteAStockGuBen]，插入股本信息异常。err: %s", err)
				return err
			}
		}
		err = row.Close()
		if err != nil {
			log.Errorf("[WriteAStockGuBen]，row关闭异常。err: %s", err)
			return err
		}
	}

	// 插入机构持股，如果变动日期没有变动，则不用插入
	for _, v := range guBenInfo.FundChiGu {
		row, err := tx.Query(`SELECT 1 FROM stock_fund_chi_gu WHERE jjdm = $1 AND zqdm = $2 AND update_time = $3`, v.JJDM, v.ZQDM, v.UPDATETIME)
		if err != nil {
			log.Errorf("[WriteAStockGuBen]，查询机构持股信息异常。参数：[%s-%s-%v]，err: %s", v.JJDM, v.ZQDM, v.UPDATETIME, err)
			return err
		}

		// 不存在，则写入
		if !row.Next() {
			log.Infof("[WriteAStockGuBen] 未查询到机构持股信息 %s-%s-%s", v.JJDM, v.ZQDM, v.UPDATETIME)
			_, err = tx.Exec(`INSERT INTO stock_fund_chi_gu (jjdm, zqmc, zqsclx, ccsz, jzbl, bcbgrq, bcccgs,
                               scbgrq, scccgs, ccbd, update_time, jjmc, zqdm)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,
				v.JJDM, v.ZQMC, v.ZQSCLX, v.CCSZ,
				v.JZBL, v.BCBGRQ, v.BCCCGS, v.SCBGRQ,
				v.SCCCGS, v.CCBD, v.UPDATETIME, v.JJMC, v.ZQDM,
			)
			if err != nil {
				log.Errorf("[WriteAStockGuBen]，插入机构持股信息异常。err: %s", err)
				return err
			}
		}

		err = row.Close()
		if err != nil {
			log.Errorf("[WriteAStockGuBen]，row关闭异常。err: %s", err)
			return err
		}
	}

	// 插入股东信息，如果变动日期没有变动，则不用插入
	for _, v := range guBenInfo.GuDong {
		row, err := tx.Query(`SELECT 1 FROM stock_gu_dong WHERE symbol = $1 AND id = $2 AND company_code = $3 AND report_date = $4`,
			v.Symbol, v.Id, v.CompanyCode, v.ReportDate)
		if err != nil {
			log.Errorf("[WriteAStockGuBen]，查询股东信息异常。err: %s", err)
			return err
		}

		// 不存在，则写入
		if !row.Next() {
			log.Infof("[WriteAStockGuBen] 未查询到股东信息 %s-%s-%s-%+v", v.Symbol, v.Id, v.CompanyCode, v.ReportDate)
			_, err = tx.Exec(`INSERT INTO stock_gu_dong (symbol, id, company_code, rank, shareholder_name, shareholder_type, shareholder_num,
								shareholder_percent, report_date, shareholder_change)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
				v.Symbol, v.Id, v.CompanyCode, v.Rank,
				v.ShareholderName, v.ShareholderType, v.ShareholderNum, v.ShareholderPercent,
				v.ReportDate, v.ShareholderChange,
			)
			if err != nil {
				log.Errorf("[WriteAStockGuBen]，插入股东信息异常。err: %s", err)
				return err
			}
		}
		err = row.Close()
		if err != nil {
			log.Errorf("[WriteAStockGuBen]，row关闭异常。err: %s", err)
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Errorf("[WriteAStockGuBen]，插入股东事务提交异常。err: %s", err)
		return common.ErrorDBCommit
	}

	return nil
}
