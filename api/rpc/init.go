package rpc

// InitRPC Init all  RPC services before the main server runs
func InitRPC() {
	initCoreRPC()
	initInteractRPC()
}
