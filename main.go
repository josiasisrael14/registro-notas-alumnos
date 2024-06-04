package main

import (
	_ "embed"

	"notas/bootstrap"
)

//go:embed boot.yaml
var boot []byte

func main() {
	bootstrap.Run2(boot)
}
