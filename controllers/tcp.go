package controllers

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net"
	"os"
)

type TcpControllers struct {
	beego.Controller
}

func (tcp *TcpControllers) PrintLine() {
	fmt.Print("123123123")
	tcp.Ctx.WriteString("123123123")
}

func (tcp *TcpControllers) TcpOnline() {
	server := "171.221.203.106:5014"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)

	if err != nil {
		fmt.Println("Fatal error: ", os.Stderr, err)
		os.Exit(1)
	}

	//建立服务器连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		os.Exit(1)
	}

	fmt.Println("connection success")
	resp := sender(conn)
	fmt.Println("send over")
	tcp.Ctx.WriteString(resp)
}

func sender(conn *net.TCPConn) (resp string) {
	words := "000144{\"channelno\":\"1\",\"transtype\":\"1\",\"countno\":\"042012\",\"terminal_serialno\":\"11620170426052342501\",\"amount\":\"1.00\",\"auth_code\":\"130227498519276010\"}"
	gbk := mahonia.NewEncoder("gbk").ConvertString(words)
	msgBack, err := conn.Write([]byte(gbk)) //给服务器发信息

	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), "服务器反馈")
		os.Exit(1)
	}
	buffer := make([]byte, 1024)
	msg, err := conn.Read(buffer) //接受服务器信息
	data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader(buffer[:msg]), simplifiedchinese.GBK.NewEncoder()))
	result := mahonia.NewEncoder("GBK").ConvertString(string(data))
	fmt.Println(conn.RemoteAddr().String(), "服务器反馈：", result, msgBack, "；实际发送了", len(words))
	_, _ = conn.Write([]byte("ok")) //在告诉服务器，它的反馈收到了。
	return result
}
