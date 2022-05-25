package main

import (
	"fmt"
	"os"
	"syscall/js"

	"github.com/GoogleContainerTools/kpt-functions-catalog/functions/go/set-namespace/transformer"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
)

func main() {
	doc := js.Global().Get("document")
	jsInput := doc.Call("getElementById", "krmfninput").Get("value").String()

	output, err := fn.Run(fn.ResourceListProcessorFunc(transformer.SetNamespace), []byte(jsInput))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	doc.Call("getElementById", "krmfnoutput").Set("value", string(output))
}
