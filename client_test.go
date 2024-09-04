package client

import (
	"fmt"
	"testing"
	"time"
)

const (
	host = "localhost"
)

func TestClient_GetStateHardware(t *testing.T) {
	client := NewClient(host)
	err := client.Connect()
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer client.Disconnect()

	request, err := NewRequest(MessageType_GetStateHardware, nil)
	if err != nil {
		t.Fatalf("Failed to NewRequest: %v", err)
	}
	response, err := client.SendItem(request)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Received response: %v\n", string(response.BytesOrPanic()))
}

func TestClient_SetCommand(t *testing.T) {
	client := NewClient(host)
	err := client.Connect()
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer client.Disconnect()

	request, err := NewRequest(MessageType_SetCommand, nil)
	if err != nil {
		t.Fatalf("Failed to NewRequest: %v", err)
	}
	response, err := client.SendItem(request)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Received response: %v\n", string(response.BytesOrPanic()))
}

func TestClient_GetSetup(t *testing.T) {
	client := NewClient(host)
	err := client.Connect()
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer client.Disconnect()

	request, err := NewRequest(MessageType_GetSetup, nil)
	if err != nil {
		t.Fatalf("Failed to NewRequest: %v", err)
	}
	response, err := client.SendItem(request)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Received response: %v\n", string(response.BytesOrPanic()))
}

func TestClient_SetSetup(t *testing.T) {
	client := NewClient(host)
	err := client.Connect()
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer client.Disconnect()
	request, err := NewRequest(MessageType_GetSetup, nil)
	if err != nil {
		t.Fatalf("Failed to NewRequest: %v", err)
	}
	// var setupSubsystem SetupSubsystem

	response, err := client.SendItem(request)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
		return
	}
	setupSubsystem, err := ParseResponseAndCast[SetupSubsystem](response)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
		return
	}
	requestSetSetup, err := NewRequest(MessageType_SetSetup, &setupSubsystem.Setup)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	setSetupResponse, err := client.SendItem(requestSetSetup)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Received response: %s\n", setSetupResponse.BytesOrPanic())

	// После этого вызова возможно потребуется переподключение,
	// добавляем задержку, чтобы дать время на перезагрузку
	time.Sleep(2 * time.Second)
}

func TestClient_GetStatistics(t *testing.T) {
	client := NewClient(host)
	err := client.Connect()
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer client.Disconnect()

	request, err := NewRequest(MessageType_GetStatistics, nil)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	response, err := client.SendItem(request)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Received response: %v\n", response.BytesOrPanic())
}
