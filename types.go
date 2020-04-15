package searchclient

import (
	pbsearch "github.com/dfuse-io/pbgo/dfuse/search/v1"
)

type MatchOrError struct {
	Match *pbsearch.SearchMatch
	Err   error
}
