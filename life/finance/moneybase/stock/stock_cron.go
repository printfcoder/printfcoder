package stock

import (
	"context"

	"github.com/robfig/cron/v3"
	log "github.com/stack-labs/stack/logger"
)

func initCron(ctx context.Context) {
	go func() {
		c := cron.New(cron.WithSeconds())
		id, err := c.AddFunc("1 2 18 ? * *", func() {
			ctxB := context.WithValue(context.Background(), methodWrapperKey{}, "write-qt-daily")
			err := WriteStockQTDaily(ctxB)
			if err != nil {
				log.Errorf("[WriteStockQTDaily] cron query err: %s", err)
			}
			log.Infof("running cron %v", 123)
		})
		if err != nil {
			log.Errorf("add cron err: %s", err)
		}

		log.Infof("running cron %v", id)

		c.Start()
	}()
}
