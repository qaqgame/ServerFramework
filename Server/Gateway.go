package Server

import (
	"encoding/binary"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

type Gateway struct {
	mapSession      map[uint32]ISession
	conn            *net.UDPConn

	isRunning       bool
	recvBuf         []byte
	listener        ISessionListener
	port            int
	closeSignal     chan int
	lastClearSessionTime int64

	recvRunning     bool
	rwMutex         sync.RWMutex
}

func NewGateway(_port int, _listener ISessionListener) *Gateway {
	log.Println("New Gateway")
	gateway := new(Gateway)
	gateway.port = _port
	gateway.listener = _listener

	gateway.mapSession = make(map[uint32]ISession)
	gateway.recvRunning = false
	gateway.isRunning = false
	gateway.conn = nil
	gateway.recvBuf = make([]byte,4096)
	gateway.closeSignal = make(chan int, 2)
	gateway.rwMutex = sync.RWMutex{}

	gateway.Initialize()

	return gateway
}

func (gateway *Gateway) Initialize() {
	gateway.start()
}

func (gateway *Gateway) start() {
	log.Println("start Gateway")
	gateway.isRunning = true

	udpAddr,err := net.ResolveUDPAddr("udp4","0.0.0.0:"+strconv.Itoa(gateway.port))
	if err != nil {
		log.Fatal("Start gateway err(resolveUDPAddr): ",err)
	}

	gateway.conn,err  = net.ListenUDP("udp4",udpAddr)
	if err != nil {
		log.Fatal("Start gateway err(listenUdp): ", err)
	}

	log.Println("Listen udp:",gateway.port,udpAddr)

	go gateway.Recv()
}

func (gateway *Gateway) Clean() {
	gateway.mapSession = make(map[uint32]ISession)
	gateway.Close()
}

func (gateway *Gateway) Close() {
	gateway.isRunning = false

	if gateway.recvRunning {
		gateway.closeSignal <- -1
	}

	_ = gateway.conn.Close()
	gateway.conn = nil
}

func (gateway *Gateway) GetSession(sid uint32) ISession {
	return gateway.mapSession[sid]
}

// 接受数据的协程
func (gateway *Gateway) Recv() {
	log.Println("start recv goroutine")
	gateway.recvRunning = true

	for gateway.isRunning {
		select {
		case v := <-gateway.closeSignal:
			if v == -1 {
				gateway.recvRunning = false
				return
			}
		default:
			// todo
			gateway.DoReceiveInGoroutine()
		}
	}

	gateway.recvRunning = false
}

func (gateway *Gateway) DoReceiveInGoroutine() {
	log.Println(time.Now().Unix(),"Start receive in goroutine")
	sidBuf := make([]byte,4)

	// lis,err := kcp.DialWithOptions(":"+strconv.Itoa(gateway.port),nil,0,0)


	n, addr, err := gateway.conn.ReadFromUDP(gateway.recvBuf)
	if err != nil {
		log.Println("error DoReceiveInThread err: ",err)
	}
	log.Println(time.Now().Unix(),"Received data: ", n)
	if n > 0 {
		//buferReader := bytes.NewBuffer(gateway.recvBuf)
		//_,_ = buferReader.Read(sidBuf)

		sidBuf = gateway.recvBuf[24:28]

		var kcpsession ISession = nil
		uid := binary.BigEndian.Uint32(sidBuf)
		log.Println("read ",uid)

		if uid == 0 {
			sid := SId.NewId()
			kcpsession = NewKCPSession(sid, gateway.HandSessionSender, gateway.listener,1)
			log.Println("sid = ", sid)
			gateway.rwMutex.Lock()
			gateway.mapSession[sid]=kcpsession
			gateway.rwMutex.Unlock()
		} else {
			log.Println(gateway.recvBuf[:n])
			log.Println("uid != 0, uid = ",uid)
			kcpsession = gateway.mapSession[uid]
		}

		if kcpsession != nil {
			kcpsession.Active(addr)
			kcpsession.DoReceiveInGateWay(gateway.recvBuf,n)
		} else {
			log.Println("useless package in DoReceiveInGoroutine")
		}
	}
}

func (gateway *Gateway) HandSessionSender(session ISession,buf []byte, size int) {
	log.Println(time.Now().Unix(),"sid: ",session.GetId())
	log.Println("Gateway *Gateway.HandSessionSender() size",size,len(buf))
	if gateway.conn != nil {
		n, err := gateway.conn.WriteToUDP(buf[:size], session.GetRemoteEndPoint())
		log.Println("写了",n,"字节")
		if err != nil {
			log.Println("HandSessionSender error: ",err)
		}
	} else {
		log.Println("HandSessionSender: conn has been closed")
	}

	log.Println(time.Now().Unix(),"end send")
}

func (gateway *Gateway) Tick() {
	if gateway.isRunning {
		discrepancy := uint32(time.Now().Sub(refTime) / time.Millisecond)
		currentTime := time.Now().Unix()
		if currentTime - gateway.lastClearSessionTime > ActiveTimeout {
			gateway.lastClearSessionTime = currentTime
			gateway.ClearNoActionSession()
		}

		for _,v := range gateway.mapSession {
			v.Tick(discrepancy)
		}
	}
}

func (gateway *Gateway) ClearNoActionSession() {
	for k,v := range gateway.mapSession {
		if !v.IsActive() {
			gateway.rwMutex.Lock()
			delete(gateway.mapSession,k)
			gateway.rwMutex.Unlock()
		}
	}
}