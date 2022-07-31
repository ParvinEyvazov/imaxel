package main

import (
	"crypto"
	"encoding/hex"
	"syscall/js"
)

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("imageDiff", js.FuncOf(hash))
	<-done
}

func hash(this js.Value, args []js.Value) interface{} {
	h := crypto.SHA512.New()
	h.Write([]byte(args[0].String()))

	return hex.EncodeToString(h.Sum(nil))
}
