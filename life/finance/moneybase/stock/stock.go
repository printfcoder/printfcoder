package stock

import (
	"context"
	"fmt"
	"sync"

	"github.com/printfcoder/printfcoder/life/finance/moneybase/db"
	"github.com/stack-labs/stack/config"
)

func init() {
	config.RegisterOptions(&c)
}

var (
	c KeysConfig

	s sync.Mutex
)

type KeysConfig struct {
	Keys struct {
		MairuiClub MairuiClubConfig `sc:"mairui-club"`
	} `sc:"keys"`
}

type MairuiClubConfig struct {
	Key     string `sc:"key"`
	HSLTURL string `sc:"hslt-url"` // 沪深股票基础信息列表接口
}

func (m *MairuiClubConfig) getHSLTURL() string {
	return fmt.Sprintf("%s%s", m.HSLTURL, m.Key)
}

func Init(ctx context.Context) error {
	// 初始化Dao，目前只支持Postgre
	err := db.Init(ctx)
	if err != nil {
		return err
	}

	dao := &DaoPostgre{
		db: db.DB(),
	}

	// mairui 插件
	syncerMairui := NewSyncerMairui()
	err = syncerMairui.Init(WithDao(dao))
	if err != nil {
		return err
	}

	return nil
}
