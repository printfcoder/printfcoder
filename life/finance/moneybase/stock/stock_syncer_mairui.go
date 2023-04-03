package stock

import (
	"encoding/json"
	"github.com/printfcoder/printfcoder/life/finance/moneybase/common"
	log "github.com/stack-labs/stack/logger"
	"net/http"
)

func NewSyncerMairui() *SyncerMairui {
	return &SyncerMairui{}
}

type SyncerMairui struct {
	Options *SyncerOptions
}

func (s *SyncerMairui) MethodSupported(methodName string) (supported bool, err error) {
	switch methodName {
	case "sync-stock-base":
		return true, nil
	default:
		return false, nil
	}
}

func (s *SyncerMairui) Name() string {
	return "mairui"
}

func (s *SyncerMairui) Init(opts ...SyncerOption) error {
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

func (s *SyncerMairui) Sync() error {
	return common.ErrorStockUnimplementedMethod
}

func (s *SyncerMairui) SyncAllStockBases() error {
	resp, err := http.Get(c.Keys.MairuiClub.getHSLTURL())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// {"dm":"300418","mc":"昆仑万维","jys":"sz"}
	var aStockBases []AStockBase
	err = json.NewDecoder(resp.Body).Decode(&aStockBases)
	if err != nil {
		log.Errorf("获取A股所有股票基本信息异常。err: %s", err)
		return err
	}

	err = s.Options.Dao.WriteAllAStocks(aStockBases...)
	if err != nil {
		log.Errorf("mairui同步股票基本信息到数据库异常。err: %s", err)
		return err
	}

	return nil
}

func (s *SyncerMairui) SyncGuBen() error {
	return common.ErrorStockUnimplementedMethod
}
