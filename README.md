# Mooncake Gambling

Mooncake Gambling is a golang library for implement the Mooncake Gambling function.

## Usage

```go
package main

import (
	"fmt"
	"strings"

	mooncake_gambling "github.com/gitsang/mooncake-gambling"
)

func main() {
	_, results := mooncake_gambling.Gamble()
	fmt.Println(strings.Join(results[0:6], " "))
	fmt.Println(strings.Join(results[6:], " "))
}
```

Result

```text
⚁ ⚀ ⚅ ⚃ ⚅ ⚂
一秀
```
