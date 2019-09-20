# Dot

[![Build Status](https://travis-ci.org/markdicksonjr/dot.svg?branch=master)](https://travis-ci.org/markdicksonjr/dot)
[![Coverage Status](https://coveralls.io/repos/github/markdicksonjr/dot/badge.svg?branch=master)](https://coveralls.io/github/markdicksonjr/dot?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/markdicksonjr/dot)](https://goreportcard.com/report/github.com/markdicksonjr/dot)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Simple get/set with dot-notation for Go that can operate on maps, structs and interfaces.

## Usage

### Get

Get will retrieve the value at the specified dot path.  It will return an error if the property is not found.

```go
sample := SampleStruct{
    A: SampleStructA{
        Text: "foo",
    },
    B: 8,
    C: "something",
}

// nested get, make "a" lowercase to demonstrate case-insensitivity
aText, err := Get(sample, "a.text")
if err != nil {
    // handle the error
}

// aText will be equal to "foo"

sampleMap := make(map[string]interface{})
sampleMap["D"] = "onmap"

dText, err := Get(sampleMap, "D")
if err != nil {
    // handle the error
}

// dText will be "onmap"

```

A fallback mechanism is also available, in which you define a list of properties and the first property with a value
will be used for the Get function.

```go
sample := SampleStruct{
    A: 0.65,
    B: 8,
    C: "something",
}

// nested get, make "a" lowercase to demonstrate case-insensitivity
fallbackText, err := dot.Get(sample, "x", "y", "b")
if err != nil {
    // handle the error
}

// fallbackText will be equal to 8
```

### Set

Sets the value at the dot-property provided.  Will create map[string]interface{} for any missing levels along the way.
If that poses an issue (struct setting, perhaps), it's recommended to allocate structures accordingly.

```go
obj := make(map[string]interface{})
err := dot.Set(obj, "X", "test34")
if err != nil {
	// handle err
}

// obj.X will be "test34"
```

### Keys

Gets the list of keys for an arbitrary structure (non-recursively).  In the result below, the result will be ["A", "B"], 
though it's best to not assume the elements are ordered:

```go
testStruct := TestStruct{
    A: false,
    B: map[string]interface{}{
        "A": 1,
    }
}

keysFromStruct := dot.Keys(testStruct)
```

### Extend

Writes any non-nil, non-default value from the right object to the left object.

TODO: example needed

### KeysRecursive

Just like Keys, only recursive

### KeysRecursiveLeaves

Just like KeysRecursive, except it returns only items with no "children"

### Additional Getters (TODO: Enhance Details)

- GetString
- GetInt64
- GetFloat64

Notable details:

- Property access is case-insensitive
- Failure to find a value at the provided property with Get will result in an error (you can still choose to ignore 
the error and count on a nil value, if you wish) 

## Common Errors

`object must be a pointer to a struct`

This error means that you're passing a struct to dot.Set, when you should be passing a pointer to the struct.  Add a &
before the struct you're passing.
