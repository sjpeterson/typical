# typical

Within half an hour or so of starting any greenfield project in Go, I tend to
re-implement one or more of the generic types in this package. While there are
idiomatic ways to handle each one in Go, a simple generic wrapper makes the
interface much nicer, more readable, and reduces the risk of error.

The package provides the following types, each in their own subpackage:

- Set
- Maybe (a.k.a optional, nullable)
- Stack
- Queue
