# Dot

Simple get/set with dot-notation for Go that can operate on maps, structs and interfaces.

## Usage

### Get

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
fallbackText, err := Get(sample, "x", "y", "b")
if err != nil {
    // handle the error
}

// fallbackText will be equal to 8
```

### Set

```go
obj := make(map[string]interface{})
err := Set(obj, "X", "test34")
if err != nil {
	// handle err
}

// obj.X will be "test34"
```

### Additional Getters (TODO: Enhance Details)

- GetString
- GetInt64
- GetFloat64

Notable details:

- Property access is case-insensitive
- Failure to find a value at the provided property with Get will result in an error (you can still choose to ignore 
the error and count on a nil value, if you wish) 