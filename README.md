# High-level client for StreamingFast Search

[![reference](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://pkg.go.dev/github.com/streamingfast/search-client)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Augmented, high-performance search client, that combines data from `kvdb` storage when
archives provide only pointers.


## Scope

The StreamingFast Search service's search results contain only pointers to
matching transaction data (except in live, which it brings along to
avoid races).

The actual data object is stored in a database through
https://github.com/streamingfast/kvdb and its interface.

Different services querying search directly can now benefit from the
optimizations of this library to speed up access to search results'
full data, by resolving pointers through `kvdb`.


## Installation & Usage

See the different protocol-specific `StreamingFast` binaries at https://github.com/streamingfast/streamingfast#protocols

Current `search` implementations:

* [**EOSIO on StreamingFast**](https://github.com/streamingfast/sf-eosio)
* [**Ethereum on StreamingFast**](https://github.com/streamingfast/sf-ethereum)

## Contributing

**Issues and PR in this repo related strictly to the search client.**

Report any protocol-specific issues in their
[respective repositories](https://github.com/streamingfast/streamingfast#protocols)

**Please first refer to the general
[StreamingFast contribution guide](https://github.com/streamingfast/streamingfast/blob/master/CONTRIBUTING.md)**,
if you wish to contribute to this code base.


## License

[Apache 2.0](LICENSE)
