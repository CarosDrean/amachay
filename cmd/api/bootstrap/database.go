package bootstrap

import (
	"github.com/CarosDrean/amachay/kit/db"
)

func newDB(config Configuration) db.DB {
	return config.Database.Name
}
