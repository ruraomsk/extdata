package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/tumarsal/scaner"
)

func TestJsons(t *testing.T) {
	err := scaner.JsonsWithContent("./jsons", func(content string) error {
		data, err := ParseMessage([]byte(content))
		if err != nil {
			return err
		}
		fmt.Println(data.Type)
		fmt.Println(data.Data)
		fmt.Println(data.Error)
		fmt.Println("--------------------------------")
		time.Sleep(1 * time.Second)
		return nil
	})
	if err != nil {
		t.Fatalf("JsonsWithContent: %v", err)
	}
}
