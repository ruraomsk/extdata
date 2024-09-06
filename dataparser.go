package client

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
	Message string `json:"message"`
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

func NewRequest(messageType MessageType, data interface{}) (*MessageItem, error) {
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

func NewResponse(messageType MessageType, data []byte) (*MessageItem, error) {
	var rawData []byte

	// Проверка на соответствие типа data типу messageType
	if data != nil {
		switch messageType {
		case MessageType_GetStateHardware:
			var mess MessageItem
			err := json.Unmarshal([]byte(data), &mess)
			if err != nil {
				return nil, fmt.Errorf("invalid data type for %s, expected *StateHardware", messageType)
			}
			rawData, err = json.Marshal(mess)
			if err != nil {
				return nil, err
			}
		case MessageType_GetSetup:
			var mess SetupSubsystem
			err := json.Unmarshal([]byte(data), &mess)
			if err != nil {
				return nil, fmt.Errorf("invalid data type for %s, expected *SetupSubsystem", messageType)
			}
			rawData, err = json.Marshal(mess)
			if err != nil {
				return nil, err
			}
		case MessageType_GetStatistics:
			var mess RepStatistics
			err := json.Unmarshal([]byte(data), &mess)
			if err != nil {
				return nil, fmt.Errorf("invalid data type for %s, expected *RepStatistics", messageType)
			}

			rawData, err = json.Marshal(mess)
			if err != nil {
				return nil, err
			}
		case MessageType_SetCommand:
			var mess ResponseMessage
			err := json.Unmarshal([]byte(data), &mess)
			if err != nil {
				return nil, fmt.Errorf("invalid data type for %s, expected *ResponseMessage", messageType)
			}
			rawData, err = json.Marshal(mess)
			if err != nil {
				return nil, err
			}
		case MessageType_SetSetup:
			var mess ResponseMessage
			err := json.Unmarshal([]byte(data), &mess)
			if err != nil {
				return nil, fmt.Errorf("invalid data type for %s, expected *ResponseMessage", messageType)
			}
			rawData, err = json.Marshal(mess)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unknown message type: %s", messageType)
		}

		// Сериализация данных в JSON
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
	bs, err := m.BytesRaw()
	if err != nil {
		return nil, err
	}
	return append(bs, '\n'), nil
}

func (m *MessageItem) BytesOrPanic() []byte {
	bs, err := m.BytesRaw()
	if err != nil {
		panic(err)
	}
	return append(bs, '\n')
}

func (m *MessageItem) BytesRaw() ([]byte, error) {
	// Сериализация MessageItem в JSON и добавление конца строки
	rawData, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return rawData, nil
}

func ParseMessage(data []byte) (*MessageItem, error) {
	var item MessageItem
	err := json.Unmarshal(data, &item)
	return &item, err
}
