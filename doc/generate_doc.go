package main

import (
	"bytes"
	"io/ioutil"
	"text/template"

	"github.com/sirupsen/logrus"

	"github.com/rancher/norman/types"
	clusterSchema "github.com/rancher/types/apis/cluster.cattle.io/v3/schema"
	managementSchema "github.com/rancher/types/apis/management.cattle.io/v3/schema"
	projectSchema "github.com/rancher/types/apis/project.cattle.io/v3/schema"
)

type templateContent struct {
	Schemas []*types.Schema
}

const (
	documentFileName = "api-doc.md"
)

func main() {
	logrus.Info("generating api document for rancher types")
	schemas := types.NewSchemas().
		AddSchemas(managementSchema.Schemas).
		AddSchemas(clusterSchema.Schemas).
		AddSchemas(projectSchema.Schemas)
	schemaMap := templateContent{
		Schemas: schemas.Schemas(),
	}
	tpl, err := template.ParseFiles("./template.tpl")
	if err != nil {
		logrus.Fatal(err)
	}
	buffer := bytes.NewBuffer([]byte{})
	if err := tpl.Execute(buffer, schemaMap); err != nil {
		logrus.Fatal(err)
	}
	if err := ioutil.WriteFile(documentFileName, buffer.Bytes(), 0666); err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("generate api document %s success", documentFileName)
}

func (t templateContent) GetPath(path string) string {
	switch path {
	case "/v3/cluster":
		return "/v3/clusters/<cluster_id>/"
	case "/v3/project":
		return "/v3/projects/<project_id>/"
	case "/v3":
		return "/v3/"
	case "/v3-public":
		return "/v3-public/"
	default:
		panic("invalid path " + path)
	}
}
