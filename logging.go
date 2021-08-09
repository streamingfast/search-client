package searchclient

import (
	"os"

	"github.com/dfuse-io/logging"
	"go.uber.org/zap"
)

var traceEnabled = false
var zlog = zap.NewNop()

func init() {
	logging.Register("github.com/streamingfast/search-client", &zlog)

	if os.Getenv("TRACE") == "true" {
		traceEnabled = true
	}
}
