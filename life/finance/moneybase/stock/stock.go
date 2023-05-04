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
		Tencent    TencentConfig    `sc:"tencent"`
		Tushare    TushareConfig    `sc:"tushare"`
	} `sc:"keys"`
}

type MairuiClubConfig struct {
	Key     string `sc:"key"`
	HSLTURL string `sc:"hslt-url"` // 沪深股票基础信息列表接口
}

func (m *MairuiClubConfig) getHSLTURL() string {
	return fmt.Sprintf("%s%s", m.HSLTURL, m.Key)
}

type TencentConfig struct {
	GubenUrl         string `sc:"guben-url"`
	MinuteURL        string `sc:"minute-url"`
	ZiJinLiuxiangUrl string `sc:"zijin-liuxiang-url"`
}

type TushareConfig struct {
	BaseUri string `sc:"base-uri"`
	Token   string `sc:"token"`
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

	syncerTencent := NewSyncerTencent()
	err = syncerTencent.Init(WithDao(dao))
	if err != nil {
		return err
	}

	syncerTushare := NewSyncerTushare()
	err = syncerTushare.Init(WithDao(dao))
	if err != nil {
		return err
	}

	// 定时器
	initCron(ctx)

	return nil
}
