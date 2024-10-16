package util

import (
	"bytes"
	"car_controller/config"
	"errors"
	"net"
)

var Connect net.Conn

func InitMessageConnect(carConfig config.CarConfig) error {
	if Connect == nil {
		var buffer bytes.Buffer
		buffer.WriteString(carConfig.Ip)
		buffer.WriteString(":")
		buffer.WriteString(carConfig.Port)
		// 创建 UDP 连接
		conn, err := net.Dial("udp", buffer.String())
		Connect = conn
		if err != nil {
			return err
		}
	}
	return nil
}

func UnMessageConnect() error {
	if Connect != nil {
		err := Connect.Close()
		if err != nil {
			return err
		} else {
			Connect = nil
			return nil
		}
	}
	return errors.New("无保持的连接")
}

func SendInstructions(instructions string) (err error) {
	if Connect != nil {
		_, err = Connect.Write([]byte(instructions))
		return err
	}
	return errors.New("连接已断开")
}

func TestInstructions() (err error) {
	if Connect != nil {
		_, err = Connect.Write([]byte("test"))
		return err
	}
	return errors.New("未连接")
}
