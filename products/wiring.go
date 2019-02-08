//+build wireinject

package products

import (
	"github.com/google/wire"
	"github.com/mvcatsifma/golang-wire-modules/db"
)

var dbClientSet = wire.NewSet(db.NewClient, wire.Bind(new(IDbClient), new(db.DbClient)))

func InitializeService() *Service {
	wire.Build(dbClientSet, NewService)
	return &Service{}
}
