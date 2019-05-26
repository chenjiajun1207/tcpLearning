/*echo 的作用： 熟悉服务端被动接受新连接 、 收发数据 、 被动处理连接断开 。 每个
连接是独立服务的，连接之间 没有关联 。 在消息内容方面 echo 有一些变种：比如做
成一 问一答的方式，收到的请求和发送响应的内容不一样，这时候要考虑打包与拆包
格式的设计，进一步还可以写简单的 HTTP 服务 。
*/
package main

import (
	"net"
	"time"

	log15 "github.com/chenjiajun1207/tcpLearning/Log"
)

func main() {
	accept, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		log15.Error(err.Error())
		return
	}
	for {
		conn, err := accept.Accept()
		if err != nil {
			log15.Error(err.Error())
			return
		}
		go func() {
			defer conn.Close()
			for {
				err := conn.SetReadDeadline(time.Now().Add(3 * time.Second))
				if err != nil {
					log15.Error(err.Error())
					return
				}
				buf := make([]byte, 10, 10)
				n, err := conn.Read(buf)
				if err != nil {
					log15.Error("", err.Error(), n)
					return
				}
				n, err = conn.Write(buf)
				if err != nil {
					log15.Error("", err.Error())
					return
				}
			}
		}()
	}
	return
}
