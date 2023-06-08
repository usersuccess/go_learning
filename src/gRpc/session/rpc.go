package session

import (
	"encoding/binary"
	"io"
	"net"
)

type Session struct {
	conn net.Conn
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		conn: conn,
	}
}

func (s *Session) Write(data []byte) error {
	// 4字节头部+可变体长度
	buf := make([]byte, 4+len(data))
	//写入头部,记录数据长度
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	//把数据放在4后面
	copy(buf[4:], data)
	_, err := s.conn.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

//从连接读取数据
func (s *Session) Read() ([]byte, error) {
	//读取头部记录的长度
	header := make([]byte, 4)
	_, err := io.ReadFull(s.conn, header)
	if err != nil {
		return nil, err
	}
	//读取数据
	dataLen := binary.BigEndian.Uint32(header)
	data := make([]byte, dataLen)
	_, err = io.ReadFull(s.conn, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
