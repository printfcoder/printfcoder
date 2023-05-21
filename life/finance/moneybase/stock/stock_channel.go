package stock

import (
	"context"
	"time"

	log "github.com/stack-labs/stack/logger"
)

var (
	ignoreChannel = make(chan StockIgnore, 10)
)

func initChannels(ctx context.Context) {
	log.Infof("initChannels")
	go func() {
		for {
			select {
			case ignore := <-ignoreChannel:
				log.Infof("[initChannels] receive new ignore: %s-%d", ignore.DM, ignore.IgnoreType)
				if err := dao.InsertIgnoreStock(ignore.DM, ignore.IgnoreType); err != nil {
					log.Errorf("[initChannels] receive new ignore: %s-%d err: %s", ignore.DM, ignore.IgnoreType, err)
				}
			}
		}
	}()
}

func AddNewIgnore(si StockIgnore) {
	go func() {
		select {
		case ignoreChannel <- si:
			// do nothing
		case <-time.After(10 * time.Second):
			log.Errorf("[AddNewIgnore] 管道阻塞，重试: %s-%d。", si.DM, si.IgnoreType)
			AddNewIgnore(si)
		}
	}()
}
