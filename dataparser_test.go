package client

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNewRequest_GetStateHardware_WithMessage(t *testing.T) {
	msg := &Message{Message: "get"}
	item, err := NewRequest(MessageType_GetStateHardware, msg)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	if item.Type != MessageType_GetStateHardware {
		t.Errorf("Type = %q, want %q", item.Type, MessageType_GetStateHardware)
	}
	if len(item.Data) == 0 {
		t.Error("Data should not be empty")
	}
	var decoded Message
	if err := json.Unmarshal(item.Data, &decoded); err != nil {
		t.Fatalf("unmarshal Data: %v", err)
	}
	if decoded.Message != "get" {
		t.Errorf("decoded.Message = %q, want %q", decoded.Message, "get")
	}
}

func TestNewRequest_GetStateHardware_WrongDataType(t *testing.T) {
	_, err := NewRequest(MessageType_GetStateHardware, &CommandForDevice{})
	if err == nil {
		t.Fatal("expected error for wrong data type")
	}
}

func TestNewRequest_SetCommand_WithCommandForDevice(t *testing.T) {
	cmd := &CommandForDevice{Plan: 1, Phase: 2}
	item, err := NewRequest(MessageType_SetCommand, cmd)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	if item.Type != MessageType_SetCommand {
		t.Errorf("Type = %q, want %q", item.Type, MessageType_SetCommand)
	}
	var decoded CommandForDevice
	if err := json.Unmarshal(item.Data, &decoded); err != nil {
		t.Fatalf("unmarshal Data: %v", err)
	}
	if decoded.Plan != 1 || decoded.Phase != 2 {
		t.Errorf("decoded = %+v", decoded)
	}
}

func TestNewRequest_SetCommand_WrongDataType(t *testing.T) {
	_, err := NewRequest(MessageType_SetCommand, &Message{})
	if err == nil {
		t.Fatal("expected error for wrong data type")
	}
}

func TestNewRequest_SetSetup_WithSetupSubsystem(t *testing.T) {
	setup := &SetupSubsystem{}
	item, err := NewRequest(MessageType_SetSetup, setup)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	if item.Type != MessageType_SetSetup {
		t.Errorf("Type = %q, want %q", item.Type, MessageType_SetSetup)
	}
}

func TestNewRequest_SetSetup_WrongDataType(t *testing.T) {
	_, err := NewRequest(MessageType_SetSetup, &Message{})
	if err == nil {
		t.Fatal("expected error for wrong data type")
	}
}

func TestNewRequest_GetSetup_WithMessage(t *testing.T) {
	msg := &Message{Message: "get"}
	item, err := NewRequest(MessageType_GetSetup, msg)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	if item.Type != MessageType_GetSetup {
		t.Errorf("Type = %q, want %q", item.Type, MessageType_GetSetup)
	}
}

func TestNewRequest_GetStatistics_WithMessage(t *testing.T) {
	msg := &Message{Message: "all"}
	item, err := NewRequest(MessageType_GetStatistics, msg)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	if item.Type != MessageType_GetStatistics {
		t.Errorf("Type = %q, want %q", item.Type, MessageType_GetStatistics)
	}
}

func TestNewRequest_UnknownMessageType(t *testing.T) {
	_, err := NewRequest(MessageType_GetBlinds, &Message{})
	if err == nil {
		t.Fatal("expected error for unknown message type")
	}
}

func TestNewRequest_NilData(t *testing.T) {
	item, err := NewRequest(MessageType_GetStateHardware, nil)
	if err != nil {
		t.Fatalf("NewRequest with nil data: %v", err)
	}
	if item.Type != MessageType_GetStateHardware {
		t.Errorf("Type = %q", item.Type)
	}
	if item.Data != nil {
		t.Errorf("Data should be nil, got len=%d", len(item.Data))
	}
}

func TestNewResponse_GetStateHardware_ValidJSON(t *testing.T) {
	// NewResponse для GetStateHardware анмаршалит data в MessageItem
	data := []byte(`{"type":"GetStateHardware","data":null,"error":""}`)
	item, err := NewResponse(MessageType_GetStateHardware, data)
	if err != nil {
		t.Fatalf("NewResponse: %v", err)
	}
	if item.Type != MessageType_GetStateHardware {
		t.Errorf("Type = %q, want %q", item.Type, MessageType_GetStateHardware)
	}
}

func TestNewResponse_GetStateHardware_InvalidJSON(t *testing.T) {
	_, err := NewResponse(MessageType_GetStateHardware, []byte("not json"))
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestNewResponse_GetSetup_ValidJSON(t *testing.T) {
	data := []byte(`{"Setup":{}}`)
	item, err := NewResponse(MessageType_GetSetup, data)
	if err != nil {
		t.Fatalf("NewResponse: %v", err)
	}
	if item.Type != MessageType_GetSetup {
		t.Errorf("Type = %q", item.Type)
	}
}

func TestNewResponse_GetStatistics_ValidJSON(t *testing.T) {
	data := []byte(`{"counts":[],"ocupaes":[]}`)
	item, err := NewResponse(MessageType_GetStatistics, data)
	if err != nil {
		t.Fatalf("NewResponse: %v", err)
	}
	if item.Type != MessageType_GetStatistics {
		t.Errorf("Type = %q", item.Type)
	}
}

func TestNewResponse_SetCommand_ValidJSON(t *testing.T) {
	data := []byte(`{"message":"OK","success":true,"error":""}`)
	item, err := NewResponse(MessageType_SetCommand, data)
	if err != nil {
		t.Fatalf("NewResponse: %v", err)
	}
	if item.Type != MessageType_SetCommand {
		t.Errorf("Type = %q", item.Type)
	}
}

func TestNewResponse_SetSetup_ValidJSON(t *testing.T) {
	data := []byte(`{"message":"OK","success":true}`)
	item, err := NewResponse(MessageType_SetSetup, data)
	if err != nil {
		t.Fatalf("NewResponse: %v", err)
	}
	if item.Type != MessageType_SetSetup {
		t.Errorf("Type = %q", item.Type)
	}
}

func TestNewResponse_GetJournal_ValidJSON(t *testing.T) {
	data := []byte(`{"Setup":{}}`)
	item, err := NewResponse(MessageType_GetJournal, data)
	if err != nil {
		t.Fatalf("NewResponse: %v", err)
	}
	if item.Type != MessageType_GetJournal {
		t.Errorf("Type = %q", item.Type)
	}
}

func TestNewResponse_GetBlinds_ValidJSON(t *testing.T) {
	data := []byte(`{"message":"OK","success":true}`)
	item, err := NewResponse(MessageType_GetBlinds, data)
	if err != nil {
		t.Fatalf("NewResponse: %v", err)
	}
	if item.Type != MessageType_GetBlinds {
		t.Errorf("Type = %q", item.Type)
	}
}

func TestNewResponse_UnknownMessageType(t *testing.T) {
	_, err := NewResponse(MessageType_GetDiagrams, []byte("{}"))
	if err == nil {
		t.Fatal("expected error for unknown message type")
	}
}

func TestNewResponse_NilData(t *testing.T) {
	item, err := NewResponse(MessageType_GetStateHardware, nil)
	if err != nil {
		t.Fatalf("NewResponse with nil: %v", err)
	}
	if item.Data != nil {
		t.Errorf("Data should be nil")
	}
}

func TestMessageItem_ParseRequest_EmptyData(t *testing.T) {
	m := &MessageItem{Type: MessageType_GetStateHardware, Data: nil}
	v, err := m.ParseRequest()
	if err != nil {
		t.Fatalf("ParseRequest: %v", err)
	}
	if v != nil {
		t.Errorf("expected nil for empty data, got %v", v)
	}
}

func TestMessageItem_ParseRequest_GetStateHardware(t *testing.T) {
	data := []byte(`{"message":"test"}`)
	m := &MessageItem{Type: MessageType_GetStateHardware, Data: data}
	v, err := m.ParseRequest()
	if err != nil {
		t.Fatalf("ParseRequest: %v", err)
	}
	msg, ok := v.(*Message)
	if !ok {
		t.Fatalf("expected *Message, got %T", v)
	}
	if msg.Message != "test" {
		t.Errorf("Message = %q, want test", msg.Message)
	}
}

func TestMessageItem_ParseRequest_SetCommand(t *testing.T) {
	data := []byte(`{"plan":5,"phase":3}`)
	m := &MessageItem{Type: MessageType_SetCommand, Data: data}
	v, err := m.ParseRequest()
	if err != nil {
		t.Fatalf("ParseRequest: %v", err)
	}
	cmd, ok := v.(*CommandForDevice)
	if !ok {
		t.Fatalf("expected *CommandForDevice, got %T", v)
	}
	if cmd.Plan != 5 || cmd.Phase != 3 {
		t.Errorf("got %+v", cmd)
	}
}

func TestMessageItem_ParseRequest_GetSetup(t *testing.T) {
	data := []byte(`{"message":"get"}`)
	m := &MessageItem{Type: MessageType_GetSetup, Data: data}
	v, err := m.ParseRequest()
	if err != nil {
		t.Fatalf("ParseRequest: %v", err)
	}
	msg, ok := v.(*Message)
	if !ok {
		t.Fatalf("expected *Message, got %T", v)
	}
	if msg.Message != "get" {
		t.Errorf("Message = %q", msg.Message)
	}
}

func TestMessageItem_ParseRequest_SetSetup(t *testing.T) {
	data := []byte(`{"Setup":{}}`)
	m := &MessageItem{Type: MessageType_SetSetup, Data: data}
	v, err := m.ParseRequest()
	if err != nil {
		t.Fatalf("ParseRequest: %v", err)
	}
	_, ok := v.(*SetupSubsystem)
	if !ok {
		t.Fatalf("expected *SetupSubsystem, got %T", v)
	}
}

func TestMessageItem_ParseRequest_GetStatistics(t *testing.T) {
	data := []byte(`{"type":"all","start":"2025-01-01T00:00:00Z","end":"2025-01-02T00:00:00Z"}`)
	m := &MessageItem{Type: MessageType_GetStatistics, Data: data}
	v, err := m.ParseRequest()
	if err != nil {
		t.Fatalf("ParseRequest: %v", err)
	}
	_, ok := v.(*GetStatistics)
	if !ok {
		t.Fatalf("expected *GetStatistics, got %T", v)
	}
}

func TestMessageItem_ParseRequest_GetJournal(t *testing.T) {
	data := []byte(`{"journal":["a","b"]}`)
	m := &MessageItem{Type: MessageType_GetJournal, Data: data}
	v, err := m.ParseRequest()
	if err != nil {
		t.Fatalf("ParseRequest: %v", err)
	}
	rep, ok := v.(*RepJournal)
	if !ok {
		t.Fatalf("expected *RepJournal, got %T", v)
	}
	if len(rep.Journal) != 2 {
		t.Errorf("len(Journal) = %d, want 2", len(rep.Journal))
	}
}

func TestMessageItem_ParseRequest_UnknownType(t *testing.T) {
	m := &MessageItem{Type: MessageType_GetBlinds, Data: []byte("{}")}
	_, err := m.ParseRequest()
	if err == nil {
		t.Fatal("expected error for unknown message type")
	}
}

func TestMessageItem_ParseRequest_InvalidJSON(t *testing.T) {
	m := &MessageItem{Type: MessageType_GetStateHardware, Data: []byte("not json")}
	_, err := m.ParseRequest()
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestMessageItem_ParseResponse_EmptyData(t *testing.T) {
	m := &MessageItem{Type: MessageType_GetStateHardware, Data: nil}
	v, err := m.ParseResponse()
	if err != nil {
		t.Fatalf("ParseResponse: %v", err)
	}
	if v != nil {
		t.Errorf("expected nil for empty data, got %v", v)
	}
}

func TestMessageItem_ParseResponse_GetStateHardware(t *testing.T) {
	data := []byte(`{"message":"StateHardware"}`)
	m := &MessageItem{Type: MessageType_GetStateHardware, Data: data}
	v, err := m.ParseResponse()
	if err != nil {
		t.Fatalf("ParseResponse: %v", err)
	}
	sh, ok := v.(*StateHardware)
	if !ok {
		t.Fatalf("expected *StateHardware, got %T", v)
	}
	if sh.Message != "StateHardware" {
		t.Errorf("Message = %q", sh.Message)
	}
}

func TestMessageItem_ParseResponse_SetCommand(t *testing.T) {
	data := []byte(`{"message":"OK","success":true,"error":""}`)
	m := &MessageItem{Type: MessageType_SetCommand, Data: data}
	v, err := m.ParseResponse()
	if err != nil {
		t.Fatalf("ParseResponse: %v", err)
	}
	_, ok := v.(*ResponseMessage)
	if !ok {
		t.Fatalf("expected *ResponseMessage, got %T", v)
	}
}

func TestMessageItem_ParseResponse_GetSetup(t *testing.T) {
	data := []byte(`{"Setup":{}}`)
	m := &MessageItem{Type: MessageType_GetSetup, Data: data}
	v, err := m.ParseResponse()
	if err != nil {
		t.Fatalf("ParseResponse: %v", err)
	}
	_, ok := v.(*SetupSubsystem)
	if !ok {
		t.Fatalf("expected *SetupSubsystem, got %T", v)
	}
}

func TestMessageItem_ParseResponse_GetStatistics(t *testing.T) {
	data := []byte(`{"counts":[],"ocupaes":[]}`)
	m := &MessageItem{Type: MessageType_GetStatistics, Data: data}
	v, err := m.ParseResponse()
	if err != nil {
		t.Fatalf("ParseResponse: %v", err)
	}
	_, ok := v.(*RepStatistics)
	if !ok {
		t.Fatalf("expected *RepStatistics, got %T", v)
	}
}

func TestMessageItem_ParseResponse_GetJournal(t *testing.T) {
	data := []byte(`{"journal":[]}`)
	m := &MessageItem{Type: MessageType_GetJournal, Data: data}
	v, err := m.ParseResponse()
	if err != nil {
		t.Fatalf("ParseResponse: %v", err)
	}
	_, ok := v.(*RepJournal)
	if !ok {
		t.Fatalf("expected *RepJournal, got %T", v)
	}
}

func TestMessageItem_ParseResponse_UnknownType(t *testing.T) {
	m := &MessageItem{Type: MessageType_GetBlinds, Data: []byte("{}")}
	_, err := m.ParseResponse()
	if err == nil {
		t.Fatal("expected error for unknown message type")
	}
}

func TestMessageItem_BytesRaw(t *testing.T) {
	m := &MessageItem{Type: MessageType_GetStateHardware, Data: []byte(`{"message":"x"}`)}
	raw, err := m.BytesRaw()
	if err != nil {
		t.Fatalf("BytesRaw: %v", err)
	}
	if raw[len(raw)-1] == '\n' {
		t.Error("BytesRaw should not end with newline")
	}
	var decoded MessageItem
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if decoded.Type != MessageType_GetStateHardware {
		t.Errorf("Type = %q", decoded.Type)
	}
}

func TestMessageItem_Bytes(t *testing.T) {
	m := &MessageItem{Type: MessageType_GetStateHardware, Data: []byte(`{"message":"x"}`)}
	bs, err := m.Bytes()
	if err != nil {
		t.Fatalf("Bytes: %v", err)
	}
	if len(bs) == 0 || bs[len(bs)-1] != '\n' {
		t.Errorf("Bytes should end with newline, got len=%d", len(bs))
	}
}

func TestMessageItem_BytesOrPanic(t *testing.T) {
	m := &MessageItem{Type: MessageType_GetStateHardware, Data: []byte(`{"message":"x"}`)}
	bs := m.BytesOrPanic()
	if len(bs) == 0 || bs[len(bs)-1] != '\n' {
		t.Errorf("BytesOrPanic should end with newline")
	}
}

func TestParseMessage_Valid(t *testing.T) {
	payload := []byte(`{"type":"GetStateHardware","data":{"message":"hello"},"error":""}`)
	item, err := ParseMessage(payload)
	if err != nil {
		t.Fatalf("ParseMessage: %v", err)
	}
	if item.Type != MessageType_GetStateHardware {
		t.Errorf("Type = %q", item.Type)
	}
	if len(item.Data) == 0 {
		t.Error("Data should not be empty")
	}
}

func TestParseMessage_InvalidJSON(t *testing.T) {
	_, err := ParseMessage([]byte("invalid"))
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestParseMessage_EmptyObject(t *testing.T) {
	item, err := ParseMessage([]byte("{}"))
	if err != nil {
		t.Fatalf("ParseMessage: %v", err)
	}
	if item.Type != "" {
		t.Errorf("Type = %q for empty object", item.Type)
	}
}

func TestRoundtrip_NewRequest_ParseRequest(t *testing.T) {
	msg := &Message{Message: "roundtrip"}
	item, err := NewRequest(MessageType_GetStateHardware, msg)
	if err != nil {
		t.Fatalf("NewRequest: %v", err)
	}
	v, err := item.ParseRequest()
	if err != nil {
		t.Fatalf("ParseRequest: %v", err)
	}
	parsed, ok := v.(*Message)
	if !ok {
		t.Fatalf("expected *Message, got %T", v)
	}
	if parsed.Message != "roundtrip" {
		t.Errorf("Message = %q, want roundtrip", parsed.Message)
	}
}

func TestRoundtrip_ParseMessage_Bytes(t *testing.T) {
	original := &MessageItem{
		Type: MessageType_GetStatistics,
		Data: mustMarshal(RepStatistics{
			Counts:  []Counts{{Time: time.Now(), Values: []int{1, 2}}},
			Ocupaes: nil,
		}),
	}
	raw, err := original.BytesRaw()
	if err != nil {
		t.Fatalf("BytesRaw: %v", err)
	}
	parsed, err := ParseMessage(raw)
	if err != nil {
		t.Fatalf("ParseMessage: %v", err)
	}
	if parsed.Type != original.Type {
		t.Errorf("Type = %q, want %q", parsed.Type, original.Type)
	}
}

func mustMarshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
