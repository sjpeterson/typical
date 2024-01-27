# typical

Within half an hour or so of starting any greenfield project in Go, I tend to
re-implement one or more of the generic types in this package. While there are
idiomatic ways to handle each one in Go, a simple generic wrapper makes the
interface much nicer, more readable, and reduces the risk of error.

The package provides the following types, each in their own subpackage:

- Maybe (a.k.a optional, nullable)
- Set
- Stack
- Queue

A common design pattern for all of them is that the internals are kept private
by exposing an interface and 

## Maybe

Maybe is a value that may or may not exist. In various languages, it is also
known as nullable, Option, Optional, etc. In Go, the idiomatic way to express
them is to use two separate variables - one for the value itself, and a bool
signifying whether the value is valid or not, e.g. in the returns of (some)
fallible built-ins:

    if val, ok := myMap["key"]; ok {
        fmt.Printf("The value of %v is %v", key val)
    }

Maybe co-locates the two variables and hides them in order to protect against
accidental misuse. It also helps with readability. Developers are often tempted
to take the shortcut and use the zero value as a stand-in for no value. This is
fine in some cases where the zero value has no meaning, but obviously not fine
in cases where it does. It is especially dangerous in situations where it has
a meaning but is not expected to happen, as more often than not, developers
will forget to check for it.

## Set

The idiomatic way to implement a set of strings in Go is `map[string]struct{}`.
It doesn't look too bad at first glance, but the more you start to use it, the
more it chafes. Adding an element is

    mySet[element] = struct{}{}

which is readable enough, if a bit circumlocutory. Checking if an element is a
member of the set is still simple, but it's getting less readable:

    if _, ok := mySet[element]; ok {
        ...
    }

When we get into things you typically want to do with sets however, that's
where it really starts going downhill. Let's compute the intersection of two
sets into a new set:

    intersection := make(map[string]struct{})
    for element := range setA {
        if _, isAMember := setB[element] {
            intersection[element] = struct{}{}
        }
    }

With typical's `set.Set`, the above actions are

    mySet.Add(element)

,

    if mySet.Contains(element) {
        ...
    }

, and

    intersection := set.Intersection(setA, setB)

. Nice, tidy, and readable. As it should be.

## Stack and Queue

In Go, stacks and queues are just slices that you append to and take values
from according to specific rules. It's straightforward enough, but rather
wordy. And although it's easy to understand, readability isn't great. Compared
to

    n := len(myStack) - 1
    if n >= 0 {
        item := myStack[n]
        myStack = myStack[:n]
        doSomethingWith(item)
    }

I would rather

    if item, ok := myStack.Pop(); ok {
        doSomethingWith(item)
    }

any day of the week.
