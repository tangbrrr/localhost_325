package handler

import (
	"context"
	"github.com/tangbo/twatt/mond/service/biz.demo/app"
	"github.com/tangbo/twatt/mond/service/biz.demo/domain/demo"
	"github.com/tangbo/twatt/mond/wind"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

var (
	resourceLock sync.Mutex
	resourceInit bool
)

func (m *BizdemoService) ResourceInit(ctx context.Context) error {
	resourceLock.Lock()
	defer resourceLock.Unlock()
	if resourceInit {
		panic("ResourceInit已经执行过了")
	}
	rdb, err := wind.GetRedisDbManager().GetClient("master")
	if err != nil {
		return err
	}

	dbm := wind.GetMongoDbManager()
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)

	demoColl, err := dbm.GetCollection("master.meta.demo")
	if err != nil {
		return err
	}
	index := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "id", Value: 1},
			},
			Options: options.Index().SetBackground(true).SetUnique(true),
		},
	}
	demoColl.Indexes().CreateMany(ctx, index, opts)

	demoSvc := demo.NewService(demoColl, rdb)

	m.app = app.NewApp(rdb, demoSvc)

	resourceInit = true
	return nil
}
