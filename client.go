package client

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
	// just send bytes
	SendBytes(v []byte) ([]byte, error)
	// parse bytes to message item and send
	SendItemBytes(v []byte) ([]byte, error)
	// send meesage items
	SendItem(v *MessageItem) (*MessageItem, error)
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
	reader *bufio.Reader
	writer *bufio.Writer
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
	c.reader = bufio.NewReader(c.conn)
	c.writer = bufio.NewWriter(c.conn)
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

func (c *Client) SendItemBytes(v []byte) ([]byte, error) {
	m, err := ParseMessage(v)
	if err != nil {
		return nil, err
	}
	response, err := c.SendItem(m)
	if err != nil {
		return nil, err
	}
	return response.BytesRaw()
}

func (c *Client) SendItem(v *MessageItem) (*MessageItem, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	resBytes, err := c.SendBytes(bytes)
	if err != nil {
		return nil, err
	}
	resMesage, err := NewResponse(v.Type, resBytes)
	if err != nil {
		return nil, err
	}
	return resMesage, nil
}

func (c *Client) SendBytes(v []byte) ([]byte, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if !c.isOpen {
		return nil, fmt.Errorf("connection is not established")
	}

	// Добавляем символ конца строки и отправляем сообщение
	_, err := c.writer.Write(v)
	if err != nil {
		return nil, err
	}
	c.writer.WriteByte(Endline)
	c.writer.Flush()
	// Ждем и читаем ответ
	response, err := c.reader.ReadBytes(Endline)
	if err != nil {
		return nil, err
	}
	// Удаляем символ конца строки из ответа

	return response, nil
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
	_, err = c.writer.Write(data)
	if err != nil {
		return err
	}
	c.writer.WriteByte(Endline)
	c.writer.Flush()
	// Ждем и читаем ответ
	response, err := c.reader.ReadBytes(Endline)
	if err != nil {
		return err
	}
	// Удаляем символ конца строки и десериализуем ответ
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
	return c.Connect()
}
