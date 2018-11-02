package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/clouddirectory"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/lornasong/aws-cloud-directory-visual/src/directory"
	"github.com/lornasong/aws-cloud-directory-visual/src/visual"
	"github.com/pkg/errors"
)

// Visualize TODO:
func Visualize(ctx echo.Context) error {
	c, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	sess := session.Must(session.NewSession())
	client := clouddirectory.New(sess)
	dir := directory.New(client, c.CloudDirectoryArn, c.CloudDirectorySchemaArn)
	v := visual.New(dir)

	out, err := v.GenerateProfile("/")
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, out)
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
