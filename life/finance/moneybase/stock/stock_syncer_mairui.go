package stock

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/printfcoder/printfcoder/life/finance/moneybase/common"
	log "github.com/stack-labs/stack/logger"
)

func NewSyncerMairui() *SyncerMairui {
	return &SyncerMairui{}
}

// AStockBaseMairui A股股票基本数据结构
type AStockBaseMairui struct {
	DM         string    `json:"dm"`          // 代码
	MC         string    `json:"mc"`          // 名称
	JYS        string    `json:"jys"`         // 交易所
	UpdateTime time.Time `json:"update_time"` // 更新时间
}

type SyncerMairui struct {
	Options *SyncerOptions
}

func (s *SyncerMairui) Name() string {
	return "mairui"
}

func (s *SyncerMairui) Init(opts ...SyncerOption) error {
	options := &SyncerOptions{}
	for _, o := range opts {
		o(s.Options)
	}

	if options.Dao == nil {
		log.Errorf("DB 未传入")
		return common.ErrorDBNil
	}

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

	err = s.Options.Dao.WriteAllAStocks(aStockBases)
	if err != nil {
		log.Errorf("mairui同步股票基本信息到数据库异常。err: %s", err)
		return err
	}

	return nil
}

func (s *SyncerMairui) SyncGuBen() error {
	return common.ErrorStockUnimplementedMethod
}
