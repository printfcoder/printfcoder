package stock

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/printfcoder/printfcoder/life/finance/moneybase/common"
	log "github.com/stack-labs/stack/logger"
)

func NewSyncerTushare() *SyncerTushare {
	return &SyncerTushare{}
}

type SyncerTushare struct {
	Options *SyncerOptions
}

func (s *SyncerTushare) Init(opts ...SyncerOption) error {
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

func (s *SyncerTushare) GetStockQT(date string, symbols ...string) ([]StockQTDataTencent, error) {
	wg := sync.WaitGroup{}

	var ret []StockQTDataTencent
	for i := 0; i < len(symbols); i += 599 {
		wg.Add(1)
		var symbol []string
		if len(symbols)-i < 599 {
			symbol = symbols[i:]
		} else {
			symbol = symbols[i : i+599]
		}
		url := c.Keys.Tushare.BaseUri
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Errorf("[GetStockQT] 解析Tushare[%s]股盘 panic：%s", symbol, err)
				}
				wg.Done()
			}()

			body := tushareReqBody{
				APIName: "daily",
				Token:   c.Keys.Tushare.Token,
				Params: &tushareDailyParams{
					TsCode:    strings.Join(symbol, ","),
					TradeDate: date,
					StartDate: date,
					EndDate:   date,
				},
			}

			bodyStr, _ := json.Marshal(body)
			payload := bytes.NewBuffer(bodyStr)
			resp, err := http.Post(url, "application/json", payload)
			if err != nil {
				log.Errorf("[GetStockQT] 读取Tushare[%s]股盘异常：%s", symbol, err)
				return
			}
			defer resp.Body.Close()

			b, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Errorf("[GetStockQT] 解析Tushare[%s]股盘body异常：%s", symbol, err)
				return
			}

			rsp := &tushareRsp{}
			err = json.Unmarshal(b, rsp)
			if err != nil {
				log.Errorf("[GetStockQT] 解析Tushare[%s]股盘json异常：%s", symbol, err)
				return
			}

			if rsp.Code != 0 {
				log.Errorf("[GetStockQT] rsp返回异常：%d-%s-[%v]", symbol, rsp.Code, rsp.Msg, rsp)
				return
			}
			ret = append(ret, rsp.ToStdQTData()...)
		}()
	}

	wg.Wait()

	return ret, nil
}

func (s *SyncerTushare) GetStockTop10GuDong(symbol string, startDate string, endDate string, guDongType int) (data []StockTop10GuDong, err error) {
	url := c.Keys.Tushare.BaseUri
	var apiName string
	if guDongType == GuDongTypeFloat {
		apiName = "top10_floatholders"
	} else {
		apiName = "top10_holders"
	}

	body := tushareReqBody{
		APIName: apiName,
		Token:   c.Keys.Tushare.Token,
		Params: &tushareTop10GuDongParams{
			TsCode:    symbol,
			StartDate: startDate,
			EndDate:   endDate,
		},
	}

	bodyStr, _ := json.Marshal(body)
	payload := bytes.NewBuffer(bodyStr)
	resp, err := http.Post(url, "application/json", payload)
	if err != nil {
		log.Errorf("[GetStockTop10GuDong] 读取Tushare[%s]Top10股东异常：%s", symbol, err)
		return
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("[GetStockTop10GuDong] 解析Tushare[%s]Top10 body异常：%s", symbol, err)
		return
	}

	rsp := &tushareRsp{}
	err = json.Unmarshal(b, rsp)
	if err != nil {
		log.Errorf("[GetStockTop10GuDong] 解析Tushare[%s]Top10 json异常：%s", symbol, err)
		return
	}

	if rsp.Code != 0 {
		log.Errorf("[GetStockTop10GuDong] rsp返回异常：%d-%s-[%v]", symbol, rsp.Code, rsp.Msg, rsp)
		return
	}

	data = rsp.ToTop10GuDong(guDongType)

	return
}

func (s *SyncerTushare) SyncAllStockBases() error {
	return common.ErrorStockUnimplementedMethod
}

func (s *SyncerTushare) MethodSupported(methodName string) (supported bool, err error) {
	switch methodName {
	case "write-single-qt-daily", "get-current-value", "write-qt-daily", "get-top10-gudong",
		"sync-top10-gudong":
		return true, nil
	default:
		return false, nil
	}
}

func (s *SyncerTushare) Sync() error {
	return common.ErrorStockUnimplementedMethod
}

func (s *SyncerTushare) WriteSingleStockQTDaily(date string, symbol string) error {
	qts, err := s.GetStockQT(date, symbol)
	if err != nil {
		log.Errorf("[WriteSingleStockQTDaily] 获取Tushare daily接口异常：%s", err)
		return common.ErrorStockQTDailyReadError
	}

	for i := 0; i < len(qts); i++ {
		log.Infof("[WriteSingleStockQTDaily] 执行第%d个，共有%d个", i+1, len(qts))
		qtTencent := qts[i]
		qt := qtTencent.ToStockQT()

		// 价格为0不用再查了
		if qt.DangQianJiaGe == 0 {
			si := StockIgnore{
				DM: qt.DaiMa,
			}
			si.SetType(IgnoreTuiShi)
			log.Errorf("[WriteSingleStockQTDaily] 忽略该股票，[%+v]", si)
			AddNewIgnore(si)
		}

		err = s.Options.Dao.WriteStockQTDaily(qt)
		if err != nil {
			log.Errorf("[WriteSingleStockQTDaily] 写入Tushare QT异常：%s", err)
			return common.ErrorStockQTDailyWriteError
		}
	}

	return nil
}

func (s *SyncerTushare) WriteStockQTDaily(date string) error {
	sbs, err := s.Options.Dao.ReadAllAStockBases()
	if err != nil {
		log.Errorf("[WriteStockQTDaily] 读取所有股票基本信息异常。err: %s", err)
		return common.ErrorStockSyncAllGuBenToDB
	}

	for i := 0; i < len(sbs); i += 500 {
		log.Infof("[WriteStockQTDaily] sync [%d-%d]", i+1, i+500)

		var symbols []string
		for j := 0; j < i+500; j++ {
			if j < len(sbs) {
				symbols = append(symbols, sbs[j].DM+"."+strings.ToUpper(sbs[j].JYS))
			}
		}

		qts, err := s.GetStockQT(date, symbols...)
		if err != nil {
			log.Errorf("[WriteStockQTDaily] 获取Tushare daily接口异常：%s", err)
			return common.ErrorStockQTDailyReadError
		}

		for j := 0; j < len(qts); j++ {
			qtTencent := qts[j]
			qt := qtTencent.ToStockQT()
			if qt.ShiJian == 0 {
				// todo 查询原因
				log.Errorf("有空数据，%v-[%d]", qtTencent, j)
				continue
			}

			// 价格为0不用再查了
			if qt.DangQianJiaGe == 0 {
				si := StockIgnore{
					DM: qt.DaiMa,
				}
				si.SetType(IgnoreTuiShi)
				log.Errorf("[WriteStockQTDaily] 忽略该股票，[%+v]", si)
				AddNewIgnore(si)
			}

			err = s.Options.Dao.WriteStockQTDaily(qt)
			if err != nil {
				log.Errorf("[WriteStockQTDaily] 写入腾讯Tushare daily异常：%s", err)
				return common.ErrorStockQTDailyWriteError
			}
		}

		// 睡50ms，免得被封了
		time.Sleep(50 * time.Millisecond)
	}

	return nil
}

func (s *SyncerTushare) SyncTop10GuDong(startDate, endDate string, guDongType int) error {
	sbs, err := s.Options.Dao.ReadAllAStockBasesForSyncGuDong(guDongType)
	if err != nil {
		log.Errorf("[SyncTop10GuDong] 读取所有股票基本信息异常。err: %s", err)
		return common.ErrorStockSyncAllGuBenToDB
	}

	for i := 0; i < len(sbs); i++ {
		symbol := sbs[i].DM + "." + strings.ToUpper(sbs[i].JYS)
		log.Infof("[SyncTop10GuDong] 执行第%d个[%s]，共有%d个", i+1, symbol, len(sbs))
		top10s, err := s.GetStockTop10GuDong(symbol, startDate, endDate, guDongType)
		if err != nil {
			log.Errorf("[SyncTop10GuDong] 获取Tushare Top10 股东异常：%s", err)
			time.Sleep(10 * time.Second)
			continue
		}

		// 空的下次就不用再查了
		if top10s == nil {
			si := StockIgnore{
				DM: sbs[i].DM,
			}
			si.SetType(guDongType)
			log.Errorf("[SyncTop10GuDong] 忽略该股票，[%+v]", si)
			AddNewIgnore(si)
		}

		err = s.Options.Dao.WriteStockTop10GuDong(top10s)
		if err != nil {
			log.Errorf("[WriteStockQTDaily] 写入Tushare Top10 股东异常：%s", err)
			return common.ErrorStockTop10GuDongWriteError
		}

		// 睡50ms，免得被封了，普通账号1min只能调10次
		time.Sleep(10 * time.Second)
	}

	return nil
}

func (s *SyncerTushare) SyncSingleStockGuBen(code string) error {
	return common.ErrorStockUnimplementedMethod
}

func (s *SyncerTushare) SyncAllStockGuBen() error {
	return common.ErrorStockUnimplementedMethod
}

func (s *SyncerTushare) Name() string {
	return "Tushare"
}
