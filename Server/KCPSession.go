package Server

import (
	"net"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/xtaci/kcp-go"
)

var ActiveTimeout int64 = 30

type Sender func(session ISession, bytes []byte, length int)

type KCPSession struct {
	sid            uint32
	userid         uint32
	listener       ISessionListener
	sender         Sender
	remoteEndPoint *net.UDPAddr
	lastActiveTime int64
	nextUpdateTime uint32
	active         bool
	recvData1      *chan []byte
	recvData2      *chan []byte

	recvData3           chan []byte
	closerecv           chan int
	// receiveInMainActive bool

	needKCPUpdate bool
	sessionPing   uint32
	logger        *log.Entry

	Kcp *kcp.KCP
}

func NewKCPSession(_sid uint32, _sender Sender, _listener ISessionListener, kcpconv uint32, logger *log.Entry) *KCPSession {
	kcpSession := new(KCPSession)
	kcpSession.logger = logger
	kcpSession.logger.Info("New a KCPSession")
	kcpSession.sid = _sid
	//TODO:
	kcpSession.userid = _sid
	kcpSession.sender = _sender
	kcpSession.listener = _listener

	kcpSession.lastActiveTime = 0
	kcpSession.nextUpdateTime = 0
	kcpSession.sessionPing = 0
	kcpSession.active = false
	c1 := make(chan []byte, 128)
	kcpSession.recvData1 = &c1
	c2 := make(chan []byte, 128)
	kcpSession.recvData2 = &c2
	//--------
	kcpSession.recvData3 = make(chan []byte, 128)
	kcpSession.closerecv = make(chan int, 5)
	// kcpSession.receiveInMainActive = false
	//--------

	kcpSession.needKCPUpdate = false

	kcpSession.Kcp = kcp.NewKCP(kcpconv, kcpSession.HandKcpSend)
	kcpSession.Kcp.NoDelay(1, 10, 2, 1)
	kcpSession.Kcp.WndSize(128, 128)
	// kcpSession.Initialize()

	kcpSession.logger.Info("KCPSession Created")
	return kcpSession
}

func (kcpSession *KCPSession) Initialize() {
	go kcpSession.DoReceiveInMain()
}

func (kcpSession *KCPSession) SetUserId(uid uint32) {
	kcpSession.userid = uid
}

func (kcpSession *KCPSession) HandKcpSend(buf []byte, size int) {
	// kcpSession.logger.Debug("After len: ", len(buf), size)
	kcpSession.sender(kcpSession, buf, size)
}

func (kcpSession *KCPSession) Active(addr *net.UDPAddr) {
	kcpSession.lastActiveTime = time.Now().Unix()
	kcpSession.active = true

	// kcpSession.logger.Warn("Remote addr: ", addr, "session id: ",kcpSession.sid)
	kcpSession.remoteEndPoint = addr
}

func (kcpSession *KCPSession) GetUid() uint32 {

	return kcpSession.userid
}

func (kcpSession *KCPSession) GetId() uint32 {

	return kcpSession.sid
}

func (kcpSession *KCPSession) Ping() uint32 {

	return kcpSession.sessionPing
}

func (kcpSession *KCPSession) SetPing(ping uint32) {
	kcpSession.sessionPing = ping
	return
}

func (kcpSession *KCPSession) IsActive() bool {
	if !kcpSession.active {
		return false
	} else {
		dt := time.Now().Unix() - kcpSession.lastActiveTime
		if dt > ActiveTimeout {
			kcpSession.active = false
		}
		return kcpSession.active
	}
}

func (kcpSession *KCPSession) IsAuth() bool {

	return kcpSession.userid > 0
}

func (kcpSession *KCPSession) SetAuth(userId uint32) {
	kcpSession.userid = userId
	return
}

func (kcpSession *KCPSession) Send(cnt []byte, length int) bool {
	if !kcpSession.IsActive() {
		//log.Println("Client close")
		kcpSession.logger.Info("Client Closed, Send failed")
		return false
	}
	// kcpSession.logger.Debug("Before len : ", len(cnt))
	i := kcpSession.Kcp.Send(cnt)
	//log.Println("KCPSession *kcpSession.Send() i: ",i)
	// kcpSession.logger.Debug("Send successfully, data len: ", i)
	return i >= 0
}

func (kcpSession *KCPSession) GetRemoteEndPoint() *net.UDPAddr {
	return kcpSession.remoteEndPoint
}

func (kcpSession *KCPSession) DoReceiveInGateWay(buf []byte, size int) {
	// kcpSession.logger.Debug("DoReceiveinGateway")
	*kcpSession.recvData1 <- buf[:size]

	//-----------
	kcpSession.recvData3 <- buf[:size]
	//-----------
}

// StopReceive : signal of close goroutine : DoReceiveInMain
func (kcpSession *KCPSession) StopReceive() {

}

func (kcpSession *KCPSession) DoReceiveInMain() {
	// kcpSession.logger.Debug("DoReceiveinMain")
	// kcpSession.receiveInMainActive = true

	// tmp := kcpSession.recvData
	// kcpSession.recvData = kcpSession.recvData2
	// kcpSession.recvData2 = tmp
	kcpSession.recvData1, kcpSession.recvData2 = kcpSession.recvData2, kcpSession.recvData1

	for true {
		select {
		// ----------
		case data := <- *kcpSession.recvData2:
			// ----------
			// kcpSession.logger.Debug("DoReceiveInMain of KCPSession received data len: ", len(data))
			ret := kcpSession.Kcp.Input(data, true, true)
			if ret < 0 {
				//log.Println("not a correct package ",ret)
				log.Warn("DeReceiveInMain of KCPSession, data is not a correct package, input ret: ", ret)
				return
			}
			kcpSession.needKCPUpdate = true
			
			for size := kcpSession.Kcp.PeekSize(); size > 0; size = kcpSession.Kcp.PeekSize() {
				buf := make([]byte, size)
				if kcpSession.Kcp.Recv(buf) > 0 {
					kcpSession.listener.OnReceive(kcpSession, buf, size)
				}
			}
		default:
			return
		}
	}
}

func (kcpSession *KCPSession) Tick(currentTime uint32) {
	kcpSession.DoReceiveInMain()
	current := currentTime
	if kcpSession.needKCPUpdate || current >= kcpSession.nextUpdateTime {
		kcpSession.Kcp.Update()
		kcpSession.nextUpdateTime = kcpSession.Kcp.Check()
		kcpSession.needKCPUpdate = false
	}
}
