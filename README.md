# kata-tdd-go
Simple Golang+TDD kata from https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking

# Description

Problem to solve: Create a code that prompts in stdout a
countdown starting at 3 like:
```
3
2
1
Go!
```

Each version commit defines a milestone reached:
 * [V1: Working version, but test lasts 4 seconds](https://github.com/jaumebecks/kata-tdd-go/commit/09c79e73fd5cae3117297b56c9eb4670782f8eea)
 * [V2: Working version, mocked tests lasts less than 100ms, but not covering all cases](https://github.com/jaumebecks/kata-tdd-go/commit/a25c005afa50c64e926d4cfae7a13d26aea7f5f0)
 * [V2-Breaker: Broken version, but test still pass, so it's still not fully tested](https://github.com/jaumebecks/kata-tdd-go/commit/cad0ebc2acf8f2c698ba728b93227d67b53abf6c)
 * [V3: Working version, mocked tests lasts less than 100ms, and fixes V2-Breaker](https://github.com/jaumebecks/kata-tdd-go/commit/42c35843dbcc80fd6f5c446d89ab4711c8415467)
 * [V3-Refactored: Removing unused spy](https://github.com/jaumebecks/kata-tdd-go/commit/166b8b94b12ef37ed3397efe4acb0edd8e5cc2ce)
 * [V4: Added new feature&test: Configurable Sleeper](https://github.com/jaumebecks/kata-tdd-go/commit/76275a90603f6ebd544a7084a2114a6dd68d1965)

# Commands

In project root:
```sh
go test # Run tests
go run countdown.go # Run code functionality
```
