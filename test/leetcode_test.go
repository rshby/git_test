package test

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestSegitiga(t *testing.T) {
	segitiga := func(n int) {
		for i := 1; i <= n; i++ {
			for j := 1; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}

	segitiga(5)
}

func TestIO(t *testing.T) {
	var buf bytes.Buffer

	x := http.Client{}
	req := http.Request{}

	response, err := x.Do(&req)

	if response != nil {
		all, err := io.ReadAll(response.Body)
	}

	csv.NewReader()

}
