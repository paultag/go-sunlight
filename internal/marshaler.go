package internal

import (
	"fmt"
	"time"
)

type Time struct {
	time.Time
}

func (t Time) UnmarshalJSON(in []byte) error {
	var strin = string(in[1 : len(in)-1])
	out, err := time.Parse("2006-01-02 15:04:05", strin)

	if err != nil {
		return err
	}

	fmt.Printf("%s", out)
	return nil
}
