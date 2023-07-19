package utils

import (
	"os"
	"runtime"
)

// CurrentTrace returns the current trace of the function
func CurrentTrace() string {
	counter, _, line, success := runtime.Caller(1)

	if !success {
		println("functionName: runtime.Caller: failed")
		os.Exit(1)
	}

	// 'userservice/pkg/services.Login\n/Users/kagan.meric/PROJECT_ZERO/userservice/pkg/services/login.go\n\u0085'
	str := runtime.FuncForPC(counter).Name() + " | " + string(rune(line))

	return str
}
