## Ruine -- Let's ruin your Golang project

> A Go implementation of Clojure's [lein-ruin](https://github.com/canweriotnow/lein-ruin)

This project as it's name implies destroys Go projects in a certain directory. It does this by replacing all occurrences of `error`, `err` and `.`  to an empty string. Happy coding!!!

### Why Ruine

Programming in Go is fun. This project gives you time to write even more Go.

```bash

$ go get github.com/adelowo/ruine
$ ruine -dir="someruinnabledirectory"

```

> `-dir` would default to the current directory. See `ruine -h`
