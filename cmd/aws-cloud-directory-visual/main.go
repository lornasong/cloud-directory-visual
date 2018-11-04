package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/clouddirectory"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/lornasong/aws-cloud-directory-visual/src/directory"
	"github.com/lornasong/aws-cloud-directory-visual/src/handlers"
	"github.com/lornasong/aws-cloud-directory-visual/src/visual"
	"github.com/pkg/errors"
)

func main() {

	e := echo.New()

	c, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	sess := session.Must(session.NewSession())
	client := clouddirectory.New(sess)
	dir := directory.New(client, c.CloudDirectoryArn, c.CloudDirectorySchemaArn)
	v := visual.New(dir)

	e.Static("/static", "node_modules")
	e.File("/", "public/index.html")
	e.GET("/find", handlers.FindRoot(v))
	e.GET("/find/:id", handlers.Find(v))
	e.Start(":8000")
}

type config struct {
	CloudDirectoryArn       string `split_words:"true" required:"true"`
	CloudDirectorySchemaArn string `split_words:"true" required:"true"`
}

func loadConfig() (*config, error) {
	var c config
	err := envconfig.Process("AWS", &c)
	if err != nil {
		return nil, errors.Wrap(err, "error loading config")
	}
	return &c, nil
}
