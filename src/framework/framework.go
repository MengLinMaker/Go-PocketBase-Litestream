package framework

import (
	"log"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

type Framework struct {
	Pb *pocketbase.PocketBase
}

func New() Framework {
	return Framework{
		Pb: pocketbase.New(),
	}
}

type ServeEventHandler func(e *core.ServeEvent)

func (f Framework) AddRoutes(handler ServeEventHandler) Framework {
	f.Pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		handler(e)
		return nil
	})
	return f
}

func (f Framework) Start() {
	// Automigrate data
	migratecmd.MustRegister(f.Pb, f.Pb.RootCmd, migratecmd.Config{
		Dir:         "./migrations",
		Automigrate: true,
	})
	if err := f.Pb.Start(); err != nil {
		log.Fatal(err)
	}
}

func (f Framework) Dao() *daos.Dao {
	return f.Pb.Dao()
}

func (f Framework) DB() dbx.Builder {
	return f.Pb.Dao().DB()
}
