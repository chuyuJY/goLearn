package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	common "goLearn/chatRoom/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [1024 * 4]byte
}

func (tf *Transfer) ReadMessage() (msg common.Message, err error) {
	fmt.Println("等待客户端发送数据......")
	_, err = tf.Conn.Read(tf.Buf[:4])
	if err != nil {
		err = errors.New("read message length error")
		return
	}

	// 转换成可读数据
	messageLength := binary.BigEndian.Uint32(tf.Buf[:4])
	// 其实是原buf[5:messageLength]，只不过前四个已被读出
	n, err := tf.Conn.Read(tf.Buf[:messageLength])
	if n != int(messageLength) || err != nil {
		//fmt.Println("conn.Read failed, err =", err)
		err = errors.New("read message body error")
		return
	}

	// 反序列化消息
	err = json.Unmarshal(tf.Buf[:messageLength], &msg)
	if err != nil {
		//fmt.Println("json.Unmarshal failed, err =", err)
		err = errors.New("message unmarshal error")
		return
	}

	return
}

func (tf *Transfer) WriteMessage(data []byte) (err error) {
	// 发送消息的长度
	// 1. 先获取到data的长度，转成[]byte{}
	messageLen := uint32(len(data))
	binary.BigEndian.PutUint32(tf.Buf[:4], messageLen)
	_, err = tf.Conn.Write(tf.Buf[:4])
	if err != nil {
		fmt.Println("conn.Write err =", err)
		return err
	}
	// 2. 发送data本身
	n, err := tf.Conn.Write(data)
	if n != int(messageLen) || err != nil {
		fmt.Println("conn.Write err =", err)
		return
	}
	return
}
