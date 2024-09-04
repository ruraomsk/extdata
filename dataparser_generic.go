//go:build go1.18

package client

import (
	"encoding/json"
)

// ParseResponseAndCast - generic функция для разбора ответа и приведения его к указанному типу.
// Эта функция компилируется только если версия Go 1.18 или выше.
func ParseResponseAndCast[T any](m *MessageItem) (*T, error) {
	if len(m.Data) == 0 {
		return nil, nil
	}

	var result T
	err := json.Unmarshal(m.Data, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
