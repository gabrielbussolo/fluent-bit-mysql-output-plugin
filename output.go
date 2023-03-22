package main

import "C"
import (
	"fmt"
	"github.com/fluent/fluent-bit-go/output"
	"log"
	"unsafe"
)

var (
	pluginName = "mysql"
	version    = "dev-version"
)

//export FLBPluginRegister
func FLBPluginRegister(def unsafe.Pointer) int {
	return output.FLBPluginRegister(def, "mysql", fmt.Sprintf("%s output plugin %s", pluginName, version))
}

//export FLBPluginInit
func FLBPluginInit(plugin unsafe.Pointer) int {
	standardFields := []string{"Host", "Database", "Table", "User", "Password", "Port", "MinPoolSize", "MaxPoolSize", "Async"}
	for _, field := range standardFields {
		key := output.FLBPluginConfigKey(plugin, field)
		log.Printf("[info] [mysql] Key: %s, Value: %s\n", field, key)
	}
	return output.FLB_OK
}

//export FLBPluginFlush
func FLBPluginFlush(data unsafe.Pointer, length C.int, tag *C.char) int {
	return output.FLB_OK
}

//export FLBPluginExit
func FLBPluginExit() int {
	return output.FLB_OK
}

func main() {
}
