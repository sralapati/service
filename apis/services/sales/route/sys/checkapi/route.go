package checkapi

import "github.com/sralapati/service/foundation/web"

func Routes(app *web.App) {
	app.HandleFunc("GET /liveness", liveness)
	app.HandleFunc("GET /readiness", readiness)
}
