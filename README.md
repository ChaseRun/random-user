# random-user

A simple Go package providing embedded random user portrait images. Contains 200 high-quality portraits (100 men, 100 women) provided by [randomuser.me](https://randomuser.me) that can be accessed without external API calls.


## Features

- **200 embedded portraits** - 100 men and 100 women (200x200px each)
- **No external dependencies** - Uses only Go standard library
- **No API calls** - All images embedded in the binary
- **Simple API** - Just 3 functions
- **Zero configuration** - Just import and use

## Installation

```bash
go get github.com/ChaseRun/random-user
```

## Usage

```go
package main

import (
    "io/ioutil"
    "log"

    randomuser "github.com/ChaseRun/random-user"
)

func main() {
    // Get a random portrait (man or woman)
    portrait := randomuser.RandomImage()

    // Get a random man's portrait
    manPortrait := randomuser.RandomManImage()

    // Get a random woman's portrait
    womanPortrait := randomuser.RandomWomanImage()

    // Save to file
    err := ioutil.WriteFile("portrait.jpg", portrait, 0644)
    if err != nil {
        log.Fatal(err)
    }
}
```

## API

### `RandomImage() []byte`
Returns a random portrait image (either man or woman).

### `RandomManImage() []byte`
Returns a random man's portrait image.

### `RandomWomanImage() []byte`
Returns a random woman's portrait image.

All functions return JPEG image data as `[]byte`.

## Example

```go
package main

import (
    "fmt"
    "io/ioutil"

    randomuser "github.com/ChaseRun/random-user"
)

func main() {
    // Get and save a random portrait
    portrait := randomuser.RandomImage()
    ioutil.WriteFile("random.jpg", portrait, 0644)
    fmt.Printf("Saved random portrait (%d bytes)\n", len(portrait))

    // Get and save a man's portrait
    man := randomuser.RandomManImage()
    ioutil.WriteFile("man.jpg", man, 0644)
    fmt.Printf("Saved man portrait (%d bytes)\n", len(man))

    // Get and save a woman's portrait
    woman := randomuser.RandomWomanImage()
    ioutil.WriteFile("woman.jpg", woman, 0644)
    fmt.Printf("Saved woman portrait (%d bytes)\n", len(woman))
}
```

## Use Cases

- **User avatars** - Generate random avatars for new users
- **Placeholder images** - Use during development and testing
- **Demo data** - Populate applications with realistic portraits
- **Profile pictures** - Default images for user profiles

## Testing

```bash
# Run tests
go test

# Run benchmarks
go test -bench=.
```

## Image Details

- **Size**: 200x200 pixels
- **Format**: JPEG
- **Total images**: 200 (100 men, 100 women)
- **Binary size**: ~10-15MB with embedded images

## License

MIT License - See [LICENSE](LICENSE) file.