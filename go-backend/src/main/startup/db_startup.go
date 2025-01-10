package startup

import "github.com/BevisDev/backend-template/src/main/helper/db"

func startDB(state string) {
	db.NewDb(state)
}
