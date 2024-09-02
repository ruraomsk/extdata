package back

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

const port = 8888

type IClient interface {
	Connect() error
	Disconnect() error
	SendMessage(v interface{}, result interface{}) error
	Send(v []byte) ([]byte, error)
}

func NewClient(host string) IClient {
	return &Client{
		host: host,
		port: port,
	}
}

type Client struct {
	conn   net.Conn
	host   string
	port   int
	mutex  sync.Mutex
	isOpen bool
}

func (c *Client) Connect() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.isOpen {
		return nil
	}

	var err error
	c.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", c.host, c.port))
	if err != nil {
		return err
	}

	c.isOpen = true
	return nil
}

func (c *Client) Disconnect() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if !c.isOpen {
		return nil
	}

	err := c.conn.Close()
	if err != nil {
		return err
	}

	c.isOpen = false
	return nil
}

func (c *Client) Send(v []byte) ([]byte, error) {
	//TODO FIX
	return []byte{}, nil
}

func (c *Client) SendMessage(v interface{}, result interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if !c.isOpen {
		return fmt.Errorf("connection is not established")
	}

	// Сериализуем входное сообщение в JSON
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	// Добавляем символ конца строки и отправляем сообщение
	_, err = c.conn.Write(append(data, Endline))
	if err != nil {
		return err
	}
	// Ждем и читаем ответ
	reader := bufio.NewReader(c.conn)
	response, err := reader.ReadBytes(Endline)
	if err != nil {
		return err
	}

	// Удаляем символ конца строки и десериализуем ответ
	response = response[:len(response)-1]
	err = json.Unmarshal(response, &result)
	if err != nil {
		fmt.Printf("Unmarshel res %v error: %v", string(response), err)
		return err
	}

	return nil
}

func (c *Client) Reconnect() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Закрываем текущее соединение, если оно открыто
	if c.isOpen {
		err := c.conn.Close()
		if err != nil {
			return fmt.Errorf("failed to close current connection: %v", err)
		}
		c.isOpen = false
	}

	// Пытаемся установить новое соединение
	var err error
	c.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", c.host, c.port))
	if err != nil {
		return fmt.Errorf("failed to reconnect: %v", err)
	}

	c.isOpen = true
	return nil
}
