# Bitcoin Wallet

This project implements a Bitcoin wallet in Go.

## Description

1. **Type Alias**:

   ```go
   type Bitcoin float64 //  Represents the amount of Bitcoin in the wallet, aliased to `float64`.
   ```

2. **Struct**:

   ```go
       // Represents a Bitcoin wallet with fields for balance and mutex for concurrency safety.
       type Wallet struct {
           balance Bitcoin
           mutex   sync.Mutex
       }
   ```

3. **Struct Methods**:
   - `New(amount Bitcoin) *Wallet`: Creates a new wallet with the specified amount of Bitcoin.
   - `(b Bitcoin) String() string`: Converts a Bitcoin amount to its string representation.
   - `(b *Bitcoin) Deposit(amount Bitcoin) error`: Adds the specified amount of Bitcoin to the wallet.
   - `(b *Bitcoin) Withdraw(amount Bitcoin) error`: Deducts the specified amount of Bitcoin from the wallet.
   - `(b *Bitcoin) Balance() Bitcoin`: Returns the current balance of the wallet.

## Usage

```go
package main

import (
    "fmt"
    "sync"

    w "github.com/salawhaaat/bcwallet/wallet"
)

func main() {
    wallet := w.New(w.Bitcoin(100))
    var wg sync.WaitGroup
    wg.Add(10)
    for i := 0; i < 5; i++ {
        go func() {
            defer wg.Done()
            err := wallet.Withdraw(w.Bitcoin(5))
            if err != nil {
                panic(err)
            }
        }()
        go func() {
            defer wg.Done()
            err := wallet.Deposit(w.Bitcoin(10))
            if err != nil {
                panic(err)
            }
        }()
    }
    wg.Wait()
    fmt.Println("Current balance:", wallet.Balance())
}
```

## Testing

To run the tests, execute the following command from the root directory of the project:

```sh
go test ./wallet -v
```
