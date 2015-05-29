# Goupil assertion

As Go doesn't have assertion system, here is a simple AssertEqual function used for testing generating an error in case of unexpected given value.

## Usage

```go
import (
  "github.com/GoGoupil/assert"
  "testing"
)

func TestSomething(t *testing.T) {
  value, expected := true, false
  assert.AssertEqual(t, value, expected) // Generate an error
}
```
