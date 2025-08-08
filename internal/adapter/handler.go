package adapter

import "github.com/jasonkwh/coffeed/proto"

type DaemonHandler struct {
	proto.UnimplementedCoffeedServiceServer
}
