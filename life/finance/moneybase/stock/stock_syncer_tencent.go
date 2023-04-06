package stock

import (
	"encoding/json"
	"net/http"
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

func (s *SyncerTencent) SyncAllStockBases() error {
	return common.ErrorStockUnimplementedMethod
}

func (s *SyncerTencent) MethodSupported(methodName string) (supported bool, err error) {
	switch methodName {
	case "sync-single-stock-guben", "sync-all-stock-guben":
		return true, nil
	default:
		return false, nil
	}
}

func (s *SyncerTencent) Name() string {
	return "tencent"
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

func (s *SyncerTencent) Sync() error {
	return common.ErrorStockUnimplementedMethod
}

func (s *SyncerTencent) SyncSingleStockGuBen(code string) error {
	url := c.Keys.Tencent.GubenUrl + code
	log.Errorf("[SyncSingleStockGuBen] 腾讯股本接口：%s", url)

	resp, err := http.Get(c.Keys.Tencent.GubenUrl + code)
	if err != nil {
		panic(err)
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
