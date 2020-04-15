# High-level client for dfuse Search

Augmented, high-performance search client, that combines data from `kvdb` storage when
archives provide only pointers.

## Scope

The dfuse Search service's search results contain only pointers to
matching transaction data (except in live, which it brings along to
avoid races).

The actual data object is stored in a database through
https://github.com/dfuse-io/kvdb and its interface.

Different services querying search directly can now benefit from the
optimizations of this library to speed up access to search results'
full data, by resolving pointers through `kvdb`.


## Installation & Usage

See the different protocol-specific `dfuse` binaries at https://github.com/dfuse-io/dfuse#protocols

Current `search` implementations:

* [**dfuse for EOSIO**](https://github.com/dfuse-io/dfuse-eosio)
* **dfuse for Ethereum**, soon to be open sourced

## Contributing

**Issues and PR in this repo related strictly to the search client.**

Report any protocol-specific issues in their
[respective repositories](https://github.com/dfuse-io/dfuse#protocols)

**Please first refer to the general
[dfuse contribution guide](https://github.com/dfuse-io/dfuse#contributing)**,
if you wish to contribute to this code base.


## License

[Apache 2.0](LICENSE)
