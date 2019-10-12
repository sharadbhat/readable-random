# readable-random

A Go package to generate readable random phrase.

[GitHub repository](https://github.com/sharadbhat/readable-random)

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
phrase := readable.Generate() // Eg: ForgetfulHarshEgg
```

Also has a method to allow for custom word count, custom separator and optional title casing.

```go
phrase := readable.GenerateSpecial(false, 4, "_") // Eg: a_purple_psychotic_animal
```

## Uses

This can be used to add to the end of a URL.

Example: https://example.com/photos/ForgetfulHarshEgg

For best results, use an wordCount value of 3, 4, or 5.
