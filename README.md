# readable-random

[![GoDoc](https://godoc.org/github.com/sharadbhat/readable-random?status.svg)](https://godoc.org/github.com/sharadbhat/readable-random)
[![Go Report Card](https://goreportcard.com/badge/github.com/sharadbhat/readable-random)](https://goreportcard.com/report/github.com/sharadbhat/readable-random)

A Go package to generate readable random phrase.

[GitHub Repository](https://github.com/sharadbhat/readable-random)

## Get started

To install,

```sh
go get github.com/sharadbhat/readable-random
```

## Usage Instructions

To use the package, first import it.

```go
import readable "github.com/sharadbhat/readable-random"
```

To generate a random phrase,

```go
phrase := readable.Generate() // Eg: ForgetfulEgg
```

Has a method to allow for custom word count, custom separator and optional title casing.

```go
phrase := readable.GenerateSpecial(false, 4, "_") // Eg: a_purple_psychotic_animal
```

Also has methods to return single adjective or noun.

```go
adjective := readable.Adjective() // Eg: absurd

noun := readable.Noun() // Eg: opinion
```

## Uses

- This can be used to name and fetch user uploaded resources.

  Example: `https://example.com/photos/ForgetfulEgg`

- Container names.

For best results, use a word count value of 2, 3, 4, or 5.
