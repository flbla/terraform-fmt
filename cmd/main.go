package main

import (
	"flag"
	"martijnvdp/terraform-fmt.git/handlers"
)

func main() {
	var dir string
	flag.StringVar(&dir, "path", ".", "Path to the root Terraform module.")
	flag.Parse()
	handlers.FormatCode(dir)
}
