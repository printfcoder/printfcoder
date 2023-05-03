package stock

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/printfcoder/printfcoder/life/finance/moneybase/common"
	log "github.com/stack-labs/stack/logger"
)

func NewSyncerTencent() *SyncerTencent {
	return &SyncerTencent{}
}

type SyncerTencent struct {
	Options *SyncerOptions
}

func (s *SyncerTencent) Init(opts ...SyncerOption) error {
	s.Options = &SyncerOptions{}
	for _, o := range opts {
		o(s.Options)
	}

	if s.Options.Dao == nil {
		log.Errorf("DB 未传入")
		return common.ErrorDBNil
	}

	syncerMux.Lock()
	syncers[s.Name()] = s
	syncerMux.Unlock()

	return nil
}

func (s *SyncerTencent) GetStockQT(symbols ...string) ([]StockQTDataTencent, error) {
	wg := sync.WaitGroup{}

	var ret []StockQTDataTencent
	for i := 0; i < len(symbols); i++ {
		wg.Add(1)
		symbol := symbols[i]
		url := c.Keys.Tencent.ZiJinLiuxiangUrl + symbol
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Errorf("[GetStockQT] 解析腾讯[%s]股盘 panic：%s", symbol, err)
				}
				wg.Done()
			}()
			resp, err := http.Get(url)
			if err != nil {
				log.Errorf("[GetStockQT] 读取腾讯[%s]股盘异常：%s", symbol, err)
				return
			}
			defer resp.Body.Close()

			b, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Errorf("[GetStockQT] 解析腾讯[%s]股盘body异常：%s", symbol, err)
				return
			}

			str, err := common.GbkToUtf8(b)
			if err != nil {
				log.Errorf("[GetStockQT] 解析腾讯[%s]股盘body-str异常：%s", symbol, err)
				return
			}
			data := parseCurrentBodyString(symbol, string(str))
			ret = append(ret, data)
		}()
	}

	wg.Wait()

	return ret, nil
}

func (s *SyncerTencent) SyncAllStockBases() error {
	return common.ErrorStockUnimplementedMethod
}

func (s *SyncerTencent) MethodSupported(methodName string) (supported bool, err error) {
	switch methodName {
	case "sync-single-stock-guben", "sync-all-stock-guben",
		"get-current-value", "write-single-qt-daily", "write-qt-daily":
		return true, nil
	default:
		return false, nil
	}
}

func (s *SyncerTencent) Sync() error {
	return common.ErrorStockUnimplementedMethod
}

func (s *SyncerTencent) WriteSingleStockQTDaily(symbol string) error {
	qts, err := s.GetStockQT(symbol)
	if err != nil {
		log.Errorf("[WriteSingleStockQTDaily] 获取腾讯QT接口异常：%s", err)
		return common.ErrorStockQTDailyReadError
	}

	for i := 0; i < len(qts); i++ {
		qtTencent := qts[i]
		qt := qtTencent.ToStockQT()

		err = s.Options.Dao.WriteStockQTDaily(qt)
		if err != nil {
			log.Errorf("[WriteSingleStockQTDaily] 写入腾讯QT异常：%s", err)
			return common.ErrorStockQTDailyWriteError
		}
	}

	return nil
}

func (s *SyncerTencent) WriteStockQTDaily() error {
	sbs, err := s.Options.Dao.ReadAllAStockBases()
	if err != nil {
		log.Errorf("[WriteStockQTDaily] 读取所有股票基本信息异常。err: %s", err)
		return common.ErrorStockSyncAllGuBenToDB
	}

	for i, v := range sbs {
		log.Infof("[WriteStockQTDaily] sync [%d-%s-%s-%s]", i+1, v.JYS, v.DM, v.MC)
		symbol := v.JYS + v.DM
		qts, err := s.GetStockQT(symbol)
		if err != nil {
			log.Errorf("[WriteStockQTDaily] 获取腾讯QT接口异常：%s", err)
			return common.ErrorStockQTDailyReadError
		}

		for i := 0; i < len(qts); i++ {
			qtTencent := qts[i]
			qt := qtTencent.ToStockQT()

			err = s.Options.Dao.WriteStockQTDaily(qt)
			if err != nil {
				log.Errorf("[WriteStockQTDaily] 写入腾讯QT异常：%s", err)
				return common.ErrorStockQTDailyWriteError
			}
		}

		// 睡50ms，免得被封了
		time.Sleep(50 * time.Millisecond)
	}

	return nil
}

func (s *SyncerTencent) SyncSingleStockGuBen(code string) error {
	url := c.Keys.Tencent.GubenUrl + code
	log.Infof("[SyncSingleStockGuBen] 腾讯股本接口：%s", url)

	resp, err := http.Get(c.Keys.Tencent.GubenUrl + code)
	if err != nil {
		log.Errorf("[SyncSingleStockGuBen] 读取腾讯股本接口异常：%s", err)
		return common.ErrorStockVendorGuBenReadError
	}
	defer resp.Body.Close()

	var gubenRsp GuBenTencentRsp
	err = json.NewDecoder(resp.Body).Decode(&gubenRsp)
	if err != nil {
		log.Errorf("[SyncSingleStockGuBen] 获取A股所有股票基本信息异常。err: %s", err)
		return common.ErrorStockVendorGuBenInvalidStruct
	}

	if len(gubenRsp.Data.GuBen) == 0 || gubenRsp.Data.GuBen[0].GPDM == "" {
		log.Errorf("[SyncSingleStockGuBen] 获取A股所有股票基本信息异常。err: %s", err)
		return common.ErrorStockVendorGuBenIsNil
	}

	err = s.Options.Dao.WriteAStockGuBen(gubenRsp.Data.ToGuBenInfo())
	if err != nil {
		log.Errorf("[SyncSingleStockGuBen] 同步股票股本到数据库异常。err: %s", err)
		return common.ErrorStockVendorGuBenInvalidStruct
	}

	return nil
}

func (s *SyncerTencent) SyncAllStockGuBen() error {
	sbs, err := s.Options.Dao.ReadAllAStockBases()
	if err != nil {
		log.Errorf("[SyncAllStockGuBen] 读取所有股票基本信息异常。err: %s", err)
		return common.ErrorStockSyncAllGuBenToDB
	}

	for i, v := range sbs {
		log.Infof("[SyncAllStockGuBen] sync [%d-%s-%s-%s]", i+1, v.JYS, v.DM, v.MC)
		code := v.JYS + v.DM
		err = s.SyncSingleStockGuBen(code)
		if err != nil {
			log.Errorf("[SyncAllStockGuBen] 同步股票[%s-%s-%s]股本到数据库异常。err: %s", v.JYS, v.DM, v.MC, err)
			continue
		}
		// 睡50ms，免得被封了
		time.Sleep(50 * time.Millisecond)
	}

	return nil
}

func (s *SyncerTencent) Name() string {
	return "tencent"
}

func parseCurrentBodyString(symbol, input string) StockQTDataTencent {
	/**
	0: 未知
	      1: 股票名字
	      2: 股票代码
	      3: 当前价格
	      4: 昨收
	      5: 今开
	      6: 成交量（手）
	      7: 外盘
	      8: 内盘
	      9: 买一
	     10: 买一量（手）
	     11-18: 买二 买五
	     19: 卖一
	     20: 卖一量
	     21-28: 卖二 卖五
	     29: 最近逐笔成交
	     30: 时间
	     31: 涨跌
	     32: 涨跌%
	     33: 最高
	     34: 最低
	     35: 价格/成交量（手）/成交额
	     36: 成交量（手）
	     37: 成交额（万）
	     38: 换手率
	     39: 市盈率
	     40:
	     41: 最高
	     42: 最低
	     43: 振幅
	     44: 流通市值
	     45: 总市值
	     46: 市净率
	     47: 涨停价
	     48: 跌停价
	*/
	// v_sz000610="xxxx"，取xxx部分
	input = input[len(symbol)+2 : len(input)-3]
	datas := strings.Split(input, "~")
	data := StockQTDataTencent{
		MC:                                datas[1],
		DaiMa:                             datas[2],
		JiaGeChengJiaoLiangShouChengJiaoE: datas[35],
	}

	data.DangQianJiaGe = toFloat(datas, 3)
	data.ZuoShou = toFloat(datas, 4)
	data.JinKai = toFloat(datas, 5)
	data.ChengJiaoLiangShou = toInt(datas, 6)
	data.WaiPan = toInt(datas, 7)
	data.NeiPan = toInt(datas, 8)
	data.Mai3Yi = toFloat(datas, 9)
	data.Mai3YiShou = toInt(datas, 10)
	data.Mai4Yi = toFloat(datas, 19)
	data.Mai4YiShou = toInt(datas, 20)
	data.ZuiJinZhuBiChengJiao = toInt(datas, 29)
	data.ShiJian = toInt(datas, 30)
	data.ZhangDie = toFloat(datas, 31)
	data.ZhangDiePercent = toFloat(datas, 32)
	data.Max = toFloat(datas, 33)
	data.Min = toFloat(datas, 34)
	data.ChengJiaoLiangShou2 = toInt(datas, 36)
	data.ChengJiaoEWan = toFloat(datas, 37)
	data.HuanShouLv = toFloat(datas, 38)
	data.ShiYingLv = toFloat(datas, 39)
	data.ZuiGao2 = toFloat(datas, 41)
	data.ZuiDi2 = toFloat(datas, 42)
	data.ZhenFu = toFloat(datas, 43)
	data.LiuTongShiZhi = toFloat(datas, 44)
	data.ZongShiZhi = toFloat(datas, 45)
	data.ShiJingLv = toFloat(datas, 46)
	data.ZhangTingJia = toFloat(datas, 47)
	data.DieTingJia = toFloat(datas, 48)

	return data
}

func toFloat(input []string, idx int) float64 {
	val, err := strconv.ParseFloat(input[idx], 64)
	if err != nil {
		log.Errorf("[toFloat]转换字符串[%d]到浮点异常，%s", idx, err)
	}

	return val
}

func toInt(input []string, idx int) int64 {
	val, err := strconv.ParseInt(input[idx], 10, 64)
	if err != nil {
		log.Errorf("[toFloat]转换字符串[%d]到整形异常，%s", idx, err)
	}

	return val
}
