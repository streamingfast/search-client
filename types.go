package searchclient

import (
	pbsearch "github.com/streamingfast/pbgo/dfuse/search/v1"
)

type MatchOrError struct {
	Match *pbsearch.SearchMatch
	Err   error
}
