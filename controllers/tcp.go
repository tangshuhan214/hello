package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/axgle/mahonia"
	"github.com/hashicorp/consul/api"
	"hello/common"
	"net"
	"os"
	"strconv"
)

type TcpControllers struct {
	beego.Controller
}

var Urlcache cache.Cache

func (tcp *TcpControllers) PrintLine() {
	fmt.Print("123123123")
	tcp.Ctx.WriteString("123123123")
}

func (tcp *TcpControllers) One() {
	Urlcache, _ = cache.NewCache("memory", `{"interval":10}`)
	Urlcache.Put("test", "tttttttt", 0)
	data := tcp.Ctx.Input.RequestBody
	params := map[string]interface{}{}
	_ = json.Unmarshal(data, &params)
	result := Urlcache.Get("test")
	tcp.Data["json"] = result
	tcp.ServeJSON()
}

func (tcp *TcpControllers) Two() {
	ot := make([]string, 0)
	for i := 0; i < 100; i++ {
		ot = append(ot, strconv.Itoa(i)+"->")
	}

	c, e := common.NewScatterSlice(ot, func(v interface{}) interface{} {
		a := v.(string)
		return a + "========="
	})
	if e != nil {
		logs.Error(e)
	}

	fmt.Println(len(c))

	tcp.Data["json"] = c
	tcp.ServeJSON()
}

func (tcp *TcpControllers) Three() {
	data := tcp.Ctx.Input.RequestBody
	params := map[string]interface{}{}
	_ = json.Unmarshal(data, &params)

	/*conn, _ := grpc.Dial(
		"",
		grpc.WithInsecure(),
		// 负载均衡，使用 consul 作服务发现
		grpc.WithBalancer(grpc.RoundRobin(grpclb.NewConsulResolver(
			"127.0.0.1:8500", "grpc.health.v1.add",
		))),
	)*/

	client, _ := api.NewClient(api.DefaultConfig()) //非默认情况下需要设置实际的参数
	services, _ := client.Agent().Services()
	fmt.Println(services)



	result := common.PostJson("http://"+services["spring-cloud-consul-producer-8511"].Address+":"+strconv.Itoa(services["spring-cloud-consul-producer-8511"].Port)+"/hello", params)
	tcp.Data["json"] = result
	tcp.ServeJSON()
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
	//data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader(buffer[:msg]), simplifiedchinese.GBK.NewEncoder()))
	result := mahonia.NewEncoder("gb2312").ConvertString(string(buffer[:msg]))
	fmt.Println(conn.RemoteAddr().String(), "服务器反馈：", result, msgBack, "；实际发送了", len(words))
	_, _ = conn.Write([]byte("ok")) //在告诉服务器，它的反馈收到了。
	return result
}
