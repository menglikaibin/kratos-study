package data

import (
	"blog/internal/conf"
	"blog/internal/data/ent"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewArticleRepo)

// Data .
type Data struct {
	// TODO warpped database client
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, error) {
	//log := log.NewHelper("data", logger)
	client, err := ent.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	if err != nil {
		//log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		//log.Errorf("failed creating schema resources: %v", err)
		return nil, err
	}

	//return &Data{}, nil
	return &Data{
		db: client,
	}, nil
}
