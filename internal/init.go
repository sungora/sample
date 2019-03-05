package internal

import (
	"time"

	"github.com/sungora/app/servhttp"
	"github.com/sungora/app/servhttp/middleware"
)

// Init базовая инициализация приложения
func Init() (code int) {

	servhttp.RootMiddleware(middleware.Main(time.Second * 3))
	servhttp.NotFound(middleware.NotFound)

	return
}
