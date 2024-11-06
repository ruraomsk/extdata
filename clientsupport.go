package client

import (
	"fmt"
	"time"
)

var subclient IClient
var err error

func Start(host string) {
	time.Sleep(10 * time.Second)
	for {
		for {
			fmt.Println("Начинаем эмулировать клиента back")
			subclient = NewClient(host)
			err = subclient.Connect()
			if err != nil {
				fmt.Printf("Failed to connect: %v", err)
				time.Sleep(10 * time.Second)
				continue
			} else {
				break
			}
		}
		for {
			time.Sleep(time.Second)
			if err = exchange(MessageType_GetBlinds); err != nil {
				break
			}
			// time.Sleep(time.Second)
			// if err = exchange(MessageType_GetStateHardware); err != nil {
			// 	break
			// }
			// time.Sleep(time.Second)
			// if err = exchange(MessageType_GetStatistics); err != nil {
			// 	break
			// }
			// time.Sleep(time.Second)
			// if err = exchange(MessageType_SetCommand); err != nil {
			// 	break
			// }
			// time.Sleep(time.Second)
			// if err = exchange(MessageType_GetSetup); err != nil {
			// 	break
			// }
			// time.Sleep(time.Second)
			// if err = exchange(MessageType_SetSetup); err != nil {
			// 	break
			// }
		}
		fmt.Printf("error %v", err)
		time.Sleep(5 * time.Second)
		subclient.Disconnect()
	}
}
func exchange(mt MessageType) error {
	request, err := NewRequest(mt, nil)
	if err != nil {
		return fmt.Errorf("failed to nNewRequest %v: %v", mt, err)
	}
	response, err := subclient.SendItem(request)
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}
	fmt.Printf("Received response: %v\n", string(response.BytesOrPanic()))
	return nil
}
