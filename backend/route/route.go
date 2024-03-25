package route

import (
	"net/http"

	"github.com/ayush/ide/handler"
)

func SetupRoutes() {
	http.HandleFunc("/api/run", handler.HandleRunCode)
}
