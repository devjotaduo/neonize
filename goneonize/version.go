package main

import "C"

//export GetVersion
func GetVersion() *C.char {
	// jotaduo: precisa bater com __GONEONIZE_VERSION__ do lado Python
	// (0.4.3.post0) — o load_goneonize compara e, se diferir, BAIXA o
	// binario oficial do upstream por cima deste. O valor commitado na tag
	// era "0.3.13" (o CI upstream regenera no build; build manual nao).
	version := "0.4.3.post0"
	return C.CString(version)
}
