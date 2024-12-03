package p2p

const (
	IncomingMessage = 0x1
	IncomingStream = 0x2
)


// RPC holds any arbitrary data that is being send over each transport  between two nodes in the network


type RPC struct{
	From string
	Payload []byte
	Stream bool //To know when we are streaming, then we know we need to lock and the unlock so we don't receive any messages while rhe read loop
}