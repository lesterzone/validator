# Validator

Simple validator based on the idea of validate as simple as possible.
There's just one goal, validate and return a map with errors

## Is this package for you ?

Keep reading.

### Goal

Since I'm working with mgo(mongo db driver for golang) I need a way to validate
my `models` ( structs ).

### Usage

1. Your `model`

```go
type Event struct {
    Title       string `bson:"title" json:"title"`
    Description string `bson:"title" json:"title"`
}
```

2. Your Validations ( based on validator )

```go
func (e *Event) Validate() (bool, map[string]string){
    v : validator.Validator()
    v.Present("title", e.Title) 

    return v.Valid, v.Errors
}
```

3. Your object validated

```go
// e defined
valid, errors := e.Validate()
```

## Install

Just copy validator.go and validator_test.go to your project, change it as
wish and don't expect until this package is updated from this repo.


