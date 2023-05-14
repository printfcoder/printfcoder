package stock

import (
	"database/sql"
	"fmt"
	"strconv"
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

func (d *DaoPostgre) WriteStockQTDaily(qt StockQTData) error {
	if d.db == nil {
		return common.ErrorDBNil
	}

	tx, err := d.db.Begin()
	if err != nil {
		log.Errorf("[WriteStockQTDaily]，开启事务异常。err: %s", err)
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 20230101235929 -> 20230101000001 20230101235959，转成一天的开始和结束
	shijianPrefix := fmt.Sprintf("%d", qt.ShiJian)[0:8]
	shijianBegin, _ := strconv.ParseInt(shijianPrefix+"000001", 10, 64)
	shijianEnd, _ := strconv.ParseInt(shijianPrefix+"235959", 10, 64)

	// 查出一天内有相同写过的
	row, err := tx.Query(`SELECT 1 FROM stock_qt_daily WHERE dai_ma = $1 AND shi_jian <= $2 AND shi_jian >=  $3`, qt.DaiMa, shijianEnd, shijianBegin)
	if err != nil {
		log.Errorf("[WriteStockQTDaily]，查询QT[%s-%s-%d]信息异常，err: %s", qt.MC, qt.DaiMa, qt.ShiJian, err)
		return err
	}
	defer func() {
		if err != nil {
			if err = row.Close(); err != nil {
				log.Errorf("[WriteStockQTDaily]，row[%s-%d]关闭异常。err: %s", qt.DaiMa, qt.ShiJian, err)
			}
		}
	}()

	// 不存在，则写入
	if !row.Next() {
		log.Infof("[WriteStockQTDaily] 查询QT[%s-%s-%d]信息异常", qt.MC, qt.DaiMa, qt.ShiJian)
		_, err = tx.Exec(`INSERT INTO stock_qt_daily (
                            mc, dai_ma, dang_qian_jia_ge, zuo_shou, jin_kai, cheng_jiao_liang, wai_pan, nei_pan,
                            mai_3_yi, mai_3_yi_shou, mai_4_yi, mai_4_yi_shou, zui_jin_zhu_bi_cheng_jiao, shi_jian,
                            zhang_die, zhang_die_percent, max_, min_, jia_ge_cheng_jiao_liang_shou_cheng_jiao_e,
                            cheng_jiao_liang_shou_2, cheng_jiao_e_wan, huan_shou_lv, shi_ying_lv, zui_gao_2, zui_di_2,
                            zhen_fu, liu_tong_shi_zhi, zong_shi_zhi, shi_jing_lv, zhang_ting_jia, die_ting_jia)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, 
					        $9, $10, $11, $12, $13, $14,
					        $15, $16, $17, $18, $19, $20,
					        $21, $22, $23, $24, $25, $26,$27,
					        $28, $29, $30, $31
					        )`,
			qt.MC, qt.DaiMa, qt.DangQianJiaGe, qt.ZuoShou, qt.JinKai, qt.ChengJiaoLiangShou, qt.WaiPan, qt.NeiPan,
			qt.Mai3Yi, qt.Mai3YiShou, qt.Mai4Yi, qt.Mai4YiShou, qt.ZuiJinZhuBiChengJiao, qt.ShiJian, qt.ZhangDie,
			qt.ZhangDiePercent, qt.Max, qt.Min, qt.JiaGeChengJiaoLiangShouChengJiaoE, qt.ChengJiaoLiangShou2,
			qt.ChengJiaoEWan, qt.HuanShouLv, qt.ShiYingLv, qt.ZuiGao2, qt.ZuiDi2, qt.ZhenFu, qt.LiuTongShiZhi,
			qt.ZongShiZhi, qt.ShiJingLv, qt.ZhangTingJia, qt.DieTingJia,
		)
		if err != nil {
			log.Errorf("[WriteStockQTDaily]，插入QT[%s-%d]信息异常。err: %s", qt.DaiMa, qt.ShiJian, err)
			return err
		}
	} else { // 有则更新
		err = row.Close()
		if err != nil {
			log.Errorf("[WriteStockQTDaily]，row[%s-%d]关闭异常。err: %s", qt.DaiMa, qt.ShiJian, err)
			return common.ErrorDBRowClose
		}

		log.Infof("[WriteAStockGuBen] 查询到今天QT信息 %s-%d", qt.DaiMa, qt.ShiJian)
		_, err = tx.Exec(`UPDATE stock_qt_daily SET dang_qian_jia_ge = $3, cheng_jiao_liang = $4, wai_pan = $5, nei_pan = $6,
                            mai_3_yi = $7, mai_3_yi_shou = $8, mai_4_yi = $9, mai_4_yi_shou = $10, zui_jin_zhu_bi_cheng_jiao = $11, shi_jian = $12,
                            zhang_die = $13 , zhang_die_percent = $14, max_ = $15, min_ = $16, jia_ge_cheng_jiao_liang_shou_cheng_jiao_e = $17,
                            cheng_jiao_liang_shou_2 = $18 , cheng_jiao_e_wan = $19 , huan_shou_lv = $20, shi_ying_lv = $21 , zui_gao_2 = $22 , zui_di_2 = $23,
                            zhen_fu = $24, liu_tong_shi_zhi = $25, zong_shi_zhi = $26, shi_jing_lv = $27, zhang_ting_jia = $28, die_ting_jia = $29
					WHERE dai_ma = $1 AND shi_jian <= $31 AND shi_jian >=  $30 AND $2=$2`, qt.DaiMa, qt.ShiJian, qt.DangQianJiaGe, qt.ChengJiaoLiangShou, qt.WaiPan,
			qt.NeiPan, qt.Mai3Yi, qt.Mai3YiShou, qt.Mai4Yi, qt.Mai4YiShou, qt.ZuiJinZhuBiChengJiao,
			qt.ShiJian, qt.ZhangDie, qt.ZhangDiePercent, qt.Max, qt.Min, qt.JiaGeChengJiaoLiangShouChengJiaoE,
			qt.ChengJiaoLiangShou2, qt.ChengJiaoEWan, qt.HuanShouLv, qt.ShiJingLv, qt.ZuiGao2, qt.ZuiDi2,
			qt.ZhenFu, qt.LiuTongShiZhi, qt.ZongShiZhi, qt.ShiJingLv, qt.ZhangTingJia, qt.DieTingJia, shijianBegin, shijianEnd,
		)
		if err != nil {
			log.Errorf("[WriteStockQTDaily]，更新QT[%s-%d]信息异常。err: %s", qt.DaiMa, qt.ShiJian, err)
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Errorf("[WriteStockQTDaily]，更新QT事务提交异常。err: %s", err)
		return common.ErrorDBCommit
	}

	return nil
}

func (d *DaoPostgre) QueryXiaDieTrend(daiMa []StockXiaDie, dateStart, dateEnd int64) (trend []StockXiaDie, err error) {
	for _, dm := range daiMa {
		rows, errIn := d.db.Query(`SELECT mc, dang_qian_jia_ge, shi_jian FROM stock_qt_daily WHERE dai_ma = $1 AND shi_jian >= $2 AND shi_jian <= $3 AND zhang_die <= 0`, dm.DaiMa, dateStart, dateEnd)
		if errIn != nil {
			err = fmt.Errorf("[QueryXiaDieTrend] query db error: %s. dateStart[%d], dateEnd[%d]", errIn, dateStart, dateEnd)
			log.Error(err)
			return nil, err
		}
		s := StockXiaDie{
			DaiMa: dm.DaiMa,
		}
		for rows.Next() {
			ab := StockDailyJiaGe{}
			err = rows.Scan(&ab.MC, &ab.JiaGe, &ab.Day)
			if err != nil {
				log.Errorf("[QueryXiaDieTrend] 获取下跌数据异常。err: %s", err)
				return nil, common.ErrorDBQueryScan
			}

			s.Trend = append(s.Trend, ab)
		}
		rows.Close()
		if len(s.Trend) > 0 {
			trend = append(trend, s)
		}
	}
	return
}

func (d *DaoPostgre) QueryAllXiaDieTrendByDay(dateStart, dateEnd int64) (trend []StockXiaDie, err error) {
	rows, errIn := d.db.Query(`SELECT mc, dai_ma FROM stock_qt_daily WHERE shi_jian >= $1 AND shi_jian <= $2 AND jin_kai !=0 AND dang_qian_jia_ge != 0 AND zhang_die <= 0`, dateStart, dateEnd)
	if errIn != nil {
		err = fmt.Errorf("[QueryXiaDieTrend] query db error: %s. dateStart[%d], dateEnd[%d]", err, dateStart, dateEnd)
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		ab := StockXiaDie{}
		err = rows.Scan(&ab.MC, &ab.DaiMa)
		if err != nil {
			log.Errorf("[QueryXiaDieTrend] 获取所有下跌数据异常。err: %s", err)
			return nil, common.ErrorDBQueryScan
		}

		trend = append(trend, ab)
	}

	return
}

func (d *DaoPostgre) QueryBenchmarkTradeDays(daiMa string, benchmarkDay int64, days int) (tradeDays []int64, err error) {
	rows, err := d.db.Query(`SELECT shi_jian FROM stock_qt_daily WHERE dai_ma = $1 AND shi_jian <= $2 ORDER BY shi_jian DESC LIMIT $3`, daiMa, benchmarkDay, days)
	if err != nil {
		err = fmt.Errorf("[QueryBenchmarkTradeDays] query db error: %s. daiMa[%s], days[%d]", err, daiMa, days)
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tradeDay int64
		err = rows.Scan(&tradeDay)
		if err != nil {
			log.Errorf("[QueryBenchmarkTradeDays] 获取基准时间数据异常。err: %s", err)
			return nil, common.ErrorDBQueryScan
		}

		tradeDays = append(tradeDays, tradeDay)
	}

	return

}
