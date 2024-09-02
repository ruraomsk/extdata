package back

import (
	"encoding/json"
	"fmt"
)

type MessageItem struct {
	Type  MessageType     `json:"type"`
	Data  json.RawMessage `json:"data"`
	Error string          `json:"error"`
}

type MessageType string

type ResponseMessage struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

const (
	MessageType_GetStateHardware MessageType = "GetStateHardware"
	MessageType_SetCommand       MessageType = "SetCommand"
	MessageType_GetSetup         MessageType = "GetSetup"
	MessageType_SetSetup         MessageType = "SetSetup"
	MessageType_GetStatistics    MessageType = "GetStatistics"
)

func NewMessage(messageType MessageType, data interface{}) (*MessageItem, error) {
	var rawData []byte
	var err error
	if data != nil {
		// Проверка на соответствие типа data типу messageType
		switch messageType {
		case MessageType_GetStateHardware, MessageType_GetSetup, MessageType_GetStatistics:
			if _, ok := data.(*Message); !ok {
				return nil, fmt.Errorf("invalid data type for %s, expected *Message", messageType)
			}
		case MessageType_SetCommand:
			if _, ok := data.(*CommandForDevice); !ok {
				return nil, fmt.Errorf("invalid data type for %s, expected *SetCommand", messageType)
			}
		case MessageType_SetSetup:
			if _, ok := data.(*SetupSubsystem); !ok {
				return nil, fmt.Errorf("invalid data type for %s, expected *SetupSubsystem", messageType)
			}
		default:
			return nil, fmt.Errorf("unknown message type: %s", messageType)
		}

		// Сериализация данных в JSON
		rawData, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}

	return &MessageItem{
		Type: messageType,
		Data: rawData,
	}, nil
}

func (m *MessageItem) ParseRequest() (interface{}, error) {
	if len(m.Data) == 0 {
		return nil, nil
	}
	var result interface{}
	switch m.Type {
	case MessageType_GetStateHardware:
		result = new(Message)
	case MessageType_SetCommand:
		result = new(CommandForDevice)
	case MessageType_GetSetup:
		result = new(Message)
	case MessageType_SetSetup:
		result = new(SetupSubsystem)
	case MessageType_GetStatistics:
		result = new(GetStatistics)
	default:
		return nil, fmt.Errorf("unknown message type: %s", m.Type)
	}

	err := json.Unmarshal(m.Data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m *MessageItem) ParseResponse() (interface{}, error) {
	if len(m.Data) == 0 {
		return nil, nil
	}
	var result interface{}

	switch m.Type {
	case MessageType_GetStateHardware:
		result = new(StateHardware)
	case MessageType_SetCommand:
		result = new(ResponseMessage)
	case MessageType_GetSetup:
		result = new(SetupSubsystem)
	case MessageType_SetSetup:
		result = new(ResponseMessage)
	case MessageType_GetStatistics:
		result = new(RepStatistics)
	default:
		return nil, fmt.Errorf("unknown message type: %s", m.Type)
	}

	err := json.Unmarshal(m.Data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m *MessageItem) Bytes() ([]byte, error) {
	// Сериализация MessageItem в JSON и добавление конца строки
	rawData, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return append(rawData, '\n'), nil
}

func ParseMessage(data []byte) (MessageItem, error) {
	var item MessageItem
	err := json.Unmarshal(data, &item)
	return item, err
}
