package searchclient

import (
	pbsearch "github.com/dfuse-io/pbgo/dfuse/search/v1"
)

type MatchOrError struct {
	match *pbsearch.SearchMatch
	err   error
}
