package client

import (
	"LunaGO/server/conn"
	"bytes"
	"encoding/binary"
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
			if code == -1 {
				log.Println("time to leave.")
				quit <- true
			}
		}
	}()
	return client
}

func (c *Client) SendLogin() {
	buf := bytes.NewBuffer([]byte{})
	writeMessageCode(buf, 0)
	err := binary.Write(buf, binary.LittleEndian, int32(c.ID))
	if err != nil {
		log.Println("send login failed:", err.Error())
	}
	c.connection.SendBytes(buf.Bytes())
}

func (c *Client) SendClose() {
	buf := bytes.NewBuffer([]byte{})
	writeMessageCode(buf, -1)
	c.connection.SendBytes(buf.Bytes())
}

func writeMessageCode(buf *bytes.Buffer, code int32) {
	binary.Write(buf, binary.LittleEndian, code)
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
