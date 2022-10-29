package functions

import "strings"

type AppFunctions struct{}

func (f *AppFunctions) ExtractIP(addr string) string {
	return strings.Split(addr, ":")[0]
}
