package server

import "embed"

//go:embed *
var SchemaDir embed.FS

const SchemaFilePath = "openapi.yaml"
