package IPCWork

import (
	"net"
	"net/http"
	"net/rpc"
	log "github.com/sirupsen/logrus"
	"strconv"
	"fmt"
)
type IPCManager struct {
	myId       int
	myPort     int
	isRunning  bool
	logger     *log.Entry
	stopSignal chan int
}

func NewIPCManager(id int) *IPCManager {
	ipc := new(IPCManager)
	ipc.myId = id
	ipc.myPort = GetIPCInfo(ipc.myId).Port
	fmt.Println("MYPORT",ipc.myPort)
	ipc.logger = log.WithFields(log.Fields{"Server":"IPCManager of"+strconv.Itoa(ipc.myId)})
	ipc.isRunning = false
	ipc.stopSignal = make(chan int, 2)

	return ipc
}

func (ipc *IPCManager) RegisterRPC(m interface{}) {
	rpc.Register(m)
	rpc.HandleHTTP()
	l,e := net.Listen("tcp",":"+strconv.Itoa(ipc.myPort))
	ipc.logger.Info("port is",ipc.myPort)
	if e != nil {
		ipc.logger.Warn("Listen tcp error",e)
	}
	go http.Serve(l,nil)
}

func (ipc *IPCManager) Clean() {
	ipc.Stop()
}

func (ipc *IPCManager) Start() {

}

func (ipc *IPCManager) Stop() {
	ipc.stopSignal <- 1
}

func (ipc *IPCManager) CallRpc(args, reply interface{},port int, rpcname string) bool {
	c, err := rpc.DialHTTP("tcp","127.0.0.1:"+strconv.Itoa(port))
	defer c.Close()
	if err != nil {
		ipc.logger.Warn("Dial tcp error:",err)
	}

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}
	if err.Error() != "unexpected EOF" {
		ipc.logger.Warn("rpc call error:",err)
	}
	return false
}