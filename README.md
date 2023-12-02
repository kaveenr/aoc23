# Templates

## File Template

```go
package main

func Day2_Part1(input string) int {
	return 0
}

func Day2_Part2(input string) int {
	return 0
}

```

## Test Template

```go

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day2_Part1(t *testing.T) {
	result := Day1_Part1(``)
	assert.Equal(t, 0, result)
}

func Test_Day2_Part2(t *testing.T) {
	result := Day1_Part2(``)
	assert.Equal(t, 0, result)
}

```