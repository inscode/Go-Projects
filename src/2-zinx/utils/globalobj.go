package utils

import (
	"2-zinx/ziface"
	"io/ioutil"
	"encoding/json"
)

//存储一切有关zinx框架的全局参数，共其他模块使用
//一些参数使用json文件配置

type GlobalObj struct {
	//当前zinx全局的server对象
	TCPServer ziface.IServer

	//当前服务器主机监听的IP
	Host string

	//当前服务器主机监听的端口
	TCPPort int

	//当前服务器的名称
	Name string

	//当前服务器的版本号
	Version string

	//当前服务器主机运行的最大连接数
	MaxConn int

	//当前Zinx框架数据包的最大值
	MaxPackageSize uint32

	WorkerPoolSize   uint32 //业务工作Worker池的数量
	MaxWorkerTaskLen uint32 //业务工作Worker对应负责的任务队列最大任务存储数量
}

var GlobalObject *GlobalObj

//从zinx.json加载自定义的参数
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	//将json数据解析到struct中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}

}

//默认初始化当前的GlobalObject
func init() {
	//如果配置文件没有加载，默认的值
	GlobalObject = &GlobalObj{
		Name:             "ZinxServerAPP",
		Version:          "V1.0",
		TCPPort:          8999,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,   //worker工作池的队列的个数
		MaxWorkerTaskLen: 1024, //每个worker对应的消息队列的任务的数量最大值
	}
	//尝试从conf/zinx.json中去加载一些用户自定义的参数
	GlobalObject.Reload()
}
