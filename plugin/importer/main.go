package main

import (
	"os"

	sdk "github.com/PlakarKorp/go-kloset-sdk"
	"github.com/PlakarKorp/integration-tar/importer"
)

func main() {
	sdk.EntrypointImporter(os.Args, importer.NewTarImporter)
}
