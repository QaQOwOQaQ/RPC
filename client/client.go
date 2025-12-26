package client

import "geerpc/codec"

// 一次 RPC 调用所需要的信息
type Call struct {
	Seq           uint64      //
	ServiceMethod string      // <service>.<method>
	Args          interface{} // arguments to the function
	Reply         interface{} // reply from the function
	Error         error       // if error occurs, it will be set
	Done          chan *Call  // notifies when the call is complete
}

func (call *Call) done() {
	call.Done <- call
}

// Client represents an RPC client.
// There may be multiple outstanding Calls associated
// with a single Client, and a Client may be used by
// multiple goroutines simulteously.
type Client struct {
	cc codec.Codec
}
