package back

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

	request := Message{Message: "GetStateHardware"}
	var response StateHardware

	err = client.SendMessage(request, &response)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Received response: %+v\n", response)
}

func TestClient_SetCommand(t *testing.T) {
	client := NewClient(host)
	err := client.Connect()
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer client.Disconnect()

	request := Message{Message: "SetCommand"}
	var response string

	err = client.SendMessage(request, &response)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Received response: %s\n", response)
}

func TestClient_GetSetup(t *testing.T) {
	client := NewClient(host)
	err := client.Connect()
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer client.Disconnect()

	request := Message{Message: "GetSetup"}
	var response SetupSubsystem

	err = client.SendMessage(request, &response)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Received response: %+v\n", response)
}

func TestClient_SetSetup(t *testing.T) {
	client := NewClient(host)
	err := client.Connect()
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer client.Disconnect()
	request := Message{Message: "GetSetup"}
	var response SetupSubsystem

	err = client.SendMessage(request, &response)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
		return
	}
	requestSetSetup := SetupSubsystem{
		Message: "SetSetup",
		Setup:   response.Setup,
	}

	var responseItem string

	err = client.SendMessage(requestSetSetup, &responseItem)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Received response: %s\n", responseItem)

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

	request := Message{Message: "GetStatistics"}
	var response RepStatistics

	err = client.SendMessage(request, &response)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Received response: %+v\n", response)
}
