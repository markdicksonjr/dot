# Dot

Simple get/set with dot-notation for Go that can operate on maps, structs and interfaces.

```go
type SampleStructA struct {
    Text string
}

type SampleStruct struct {
    A SampleStructA
    B int
    C string
}

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

```

A fallback mechanism is also available, in which you define a list of properties and the first property with a value
will be used for the Get function.

```go
type SampleStruct struct {
    A float64
    B int
    C string
}

sample := SampleStruct{
    A: 0.65,
    B: 8,
    C: "something",
}

// nested get, make "a" lowercase to demonstrate case-insensitivity
fallbackText, err := Get(sample, "x", "y", "b")
if err != nil {
    // handle the error
}

// fallbackText will be equal to 8
```

Notable details:

- Properties are case-insensitive
- Failure to find a value at the provided property with Get will result in an error (you can still choose to ignore 
the error and count on a nil value, if you wish) 