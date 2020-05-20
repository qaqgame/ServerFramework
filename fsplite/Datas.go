package fsplite

import (
	"math/rand"
	"time"
)

const (
	// ServerFrameInterval : ms of frames interval
	ServerFrameInterval = 66
	// ServerTimeout : ms of client timeout
	ServerTimeout = 15000
	// ClientFrameRateMultiple :
	ClientFrameRateMultiple = 2
	// MaxFrameID :
	MaxFrameID = -1
	// AUTH :
	AUTH = 1008
	// EnableSpeedUp :
	EnableSpeedUp = true
	// DefaultSpeedUp :
	DefaultSpeedUp = 1
	// JitterBufferSize :
	JitterBufferSize = 0
	// EnableAutoBuffer :
	EnableAutoBuffer = true
	// SessionActiveTimeout :
	SessionActiveTimeout = 30
)

// GameStates
const (
	// None : 初始状态
	None = iota + 1000
	// Create : 游戏创建状态
	Create
	// GameBegin : 游戏开始状态
	GameBegin
	// RoundBegin : 回合开始
	RoundBegin
	// ControlStart : 可以开始操作
	ControlStart
	// RoundEnd : 回合结束
	RoundEnd
	// GameEnd : 游戏结束
	GameEnd
	// GameExit :
	GameExit
)

const (
	// NormalExit :
	NormalExit = iota + 100
)

// Empty : judge if fspframe is empty
func (fspframe *FSPFrame) Empty() bool {
	return len(fspframe.Msgs) == 0 || fspframe.Msgs == nil
}

var reftime time.Time
var sid uint32 = 0
var randomSeed int32
func init() {
	reftime = time.Now()
	rand.Seed(time.Now().UnixNano())
	randomSeed = rand.Int31()
}

// NewDefaultFspParam : Create a default FspPrame
func NewDefaultFspParam(host string, port int) *FSPParam {
	fspparam := new(FSPParam)
	// default param
	fspparam.Host = host
	fspparam.Port = int32(port)
	fspparam.Sid = 0
	fspparam.ServerFrameInterval = ServerFrameInterval
	fspparam.ServerTimeout = ServerTimeout
	fspparam.ClientFrameRateMultiple = ClientFrameRateMultiple
	fspparam.UseLocal = false
	fspparam.AuthID = AUTH
	fspparam.MaxFrameID = MaxFrameID
	fspparam.EnableSpeedUp = EnableSpeedUp
	fspparam.EnableAutoBuffer = EnableAutoBuffer
	fspparam.DefaultSpeed = DefaultSpeedUp
	fspparam.JitterBufferSize = JitterBufferSize
	fspparam.RandomSeed = randomSeed

	return fspparam
}

// TODO: use interface
type GameProcess interface {

}

type FSPGameI interface {
	OnStateGameCreate()
	OnStateGameBegin()
	OnStateRoundBegin()
	OnStateControlStart()
	OnStateRoundEnd()
	OnStateGameEnd()
	IsGameEnd() bool
	SetGameState(int, int, int)
	EnterFrame()
	AddCmdToCurrFrame(int32, []byte)
	AddMsgToCurrFrame(uint32, *FSPMessage)
	Release()
	AddPlayer(uint32, *FSPSession, uint32, uint32, uint32) *FSPPlayer
	GetGameID() uint32
	// session中收到msg时调用
	OnGameBeginCallBack(*FSPPlayer, *FSPMessage)
	// 将消息添加到frame后调用
	OnGameBeginMsgAddCallBack()
	// 生成GameBegin消息
	CreateGameBeginMsg() []byte

	OnRoundBeginCallBack(*FSPPlayer, *FSPMessage)
	OnRoundBeginMsgAddCallBack()
	CreateRoundMsg() []byte

	OnControlStartCallBack(*FSPPlayer, *FSPMessage)
	OnControlStartMsgAddCallBack()
	CreateControlStartMsg() []byte

	OnRoundEndCallBack(*FSPPlayer, *FSPMessage)
	OnRoundEndMsgAddCallBack()
	CreateRoundEndMsg() []byte

	OnGameEndCallBack(*FSPPlayer, *FSPMessage)
	OnGameEndMsgAddCallBack()
	CreateGameEndMsg() []byte
}