package p2p

// Message holds any or arbitrary data that is being send over each transport  between two nodes in the network
type Message struct{
	Payload []byte

}