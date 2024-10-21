package main

import (
	_ "github.com/rostamzad1904/Warp/hiddify_extension"

	"github.com/hiddify/hiddify-core/extension/server"
)

func main() {
	server.StartTestExtensionServer()
}
