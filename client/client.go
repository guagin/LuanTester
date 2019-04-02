package client

import (
	"LunaGO/server/conn"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

type Client struct {
	connection *conn.Connection
	ID         int32
	packets    chan []byte
}

func New(quit chan<- bool, ID int32) *Client {
	log.Println("begin Dial")
	c, err := net.Dial("tcp", ":55555")
	if err != nil {
		log.Println("dial error:", err)
	}
	log.Println("Dial ok")
	cc := conn.NewConnection(c)
	client := &Client{
		connection: cc,
		packets:    make(chan []byte, 256),
		ID:         ID,
	}
	go cc.StartReceiving(client.packets)
	go func() {
		for {
			packet := <-client.packets
			code, err := GetMessageCode(packet)
			if err != nil {
				return
			}
			log.Println("client receive:", code)
			if code == 1 {
				log.Println("time to leave.")
				close(client.packets)
				quit <- true
				return
			}
		}
	}()
	return client
}

func (c *Client) SendLogin() {
	log.Println("send login")
	buf := bytes.NewBuffer([]byte{})
	writeMessageCode(buf, 0)
	err := binary.Write(buf, binary.LittleEndian, int32(c.ID))

	playerID := fmt.Sprintf("player_%d", c.ID)
	if err != nil {
		log.Println("send login failed:", err.Error())
		return
	}
	err = binary.Write(buf, binary.LittleEndian, int32(len(playerID)))
	if err != nil {
		log.Println("send login failed:", err.Error())
		return
	}
	_, err = buf.Write([]byte(playerID))
	if err != nil {
		log.Println("send login failed:", err.Error())
		return
	}
	c.connection.SendBytes(buf.Bytes())
}

func (c *Client) SendClose() {
	log.Println("send close")
	buf := bytes.NewBuffer([]byte{})
	writeMessageCode(buf, 1)
	c.connection.SendBytes(buf.Bytes())
}

func writeMessageCode(buf *bytes.Buffer, code int32) {
	err := binary.Write(buf, binary.LittleEndian, code)
	if err != nil {
		log.Println("writeMessageCode failed:", err.Error())
	}
}

func GetMessageCode(packet []byte) (int32, error) {
	var code int32
	buf := bytes.NewReader(packet[0:4])
	err := binary.Read(buf, binary.LittleEndian, &code)
	if err != nil {
		log.Println("failed to get message code:", err.Error())
		return 0, err
	}
	return code, nil
}

func GetData(packet []byte) []byte {
	return packet[4:]
}
