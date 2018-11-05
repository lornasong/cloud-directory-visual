package main

import (
	"flag"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/clouddirectory"
	"github.com/labstack/echo"
	"github.com/lornasong/aws-cloud-directory-visual/src/directory"
	"github.com/lornasong/aws-cloud-directory-visual/src/handlers"
	"github.com/lornasong/aws-cloud-directory-visual/src/visual"
)

func main() {
	cdArn := flag.String("arn", "", "aws cloud directory arn")
	cdSchemaArn := flag.String("schemaArn", "", "aws cloud directory schema arn")
	flag.Parse()

	if len(*cdArn) == 0 {
		log.Fatal("Error: missing required value for command-line flag '-arn=$(AWS_CLOUD_DIRECTORY_ARN)'")
	}

	if len(*cdSchemaArn) == 0 {
		log.Fatal("Error: missing required value for command-line flag '-schemaArn=$(AWS_CLOUD_DIRECTORY_SCHEMA_ARN)'")
	}

	sess := session.Must(session.NewSession())
	client := clouddirectory.New(sess)
	dir := directory.New(client, *cdArn, *cdSchemaArn)
	v := visual.New(dir)

	e := echo.New()
	e.Static("/static", "node_modules")
	e.File("/", "public/index.html")
	e.GET("/find", handlers.FindRoot(v))
	e.GET("/find/:id", handlers.Find(v))
	e.Start(":8000")
}
