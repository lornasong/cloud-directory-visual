package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/clouddirectory"
	"github.com/kelseyhightower/envconfig"
	"github.com/lornasong/aws-cloud-directory-visual/src/directory"
	"github.com/lornasong/aws-cloud-directory-visual/src/visual"
	"github.com/pkg/errors"
)

func main() {
	c, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	sess := session.Must(session.NewSession())
	client := clouddirectory.New(sess)
	dir := directory.New(client, c.CloudDirectoryArn, c.CloudDirectorySchemaArn)

	v := visual.New(dir)

	id := ""
	out, err := v.DescribeObject(id)
	if err != nil {
		log.Fatalf("Failed: %s", err)
	}
	fmt.Printf("%s", out)
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
