package p2p

// Message holds any or arbitrary data that is being send over each transport  between two nodes in the network
type RPC struct{
	From string
	Payload []byte

}