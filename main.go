package main /* import "s32x.com/kathisto" */

import (
	"fmt"
	"os"

	"s32x.com/kathisto/api"
)

var (
	version   = "0.5"
	userAgent = fmt.Sprintf("kathisto/%s", version)
	pubDir    = getenv("PUBLIC_ADDR", "/dist")
	host      = getenv("HOST", "")
	port      = getenv("PORT", "80")
)

func main() { api.Start(version, userAgent, pubDir, host, port) }

// getenv retrieves a variable from the environment and falls back to a passed
// default value if the key doesn't exist
func getenv(key, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}
