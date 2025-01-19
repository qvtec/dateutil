# dateutil

A lightweight utility package for date manipulations in Go.

## Features

- `AddDays` - Add or subtract days from a given date.
- `AddMonths` - Add or subtract months, with end-of-month adjustment.
- `AddYears` - Add or subtract years, considering leap years.

- `FirstOfMonth` - Get the first day of the month for a given date.
- `EndOfMonth` - Get the last day of the month for a given date.


## Installation

```bash
go get github.com/qvtec/dateutil
```

## Usage

``` go
package main

import (
	"fmt"
	"time"

	"github.com/qvtec/dateutil"
)

func main() {
	now := time.Now()
	fmt.Println("Today:", now)
	fmt.Println("10 days later:", dateutil.AddDays(now, 10))
	fmt.Println("2 months later:", dateutil.AddMonths(now, 2))
	fmt.Println("1 year ago:", dateutil.AddYears(now, -1))
	fmt.Println("First of this month:", dateutil.FirstOfMonth(now))
	fmt.Println("End of this month:", dateutil.EndOfMonth(now))
}
```
