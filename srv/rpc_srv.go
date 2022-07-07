package srv

import (
	"context"
	"errors"
	"github.com/et-zone/eids/public"
	pb "github.com/et-zone/proto_api/ids"
)

type Srv struct {
	//pb.UnimplementedIDsInterfaceServer
}

func (s *Srv)GetID(ctx context.Context, req *pb.Request) (res *pb.ResponseID,err error){
	res=&pb.ResponseID{
		Code: pb.Code_SUCC,
	}
	res.ID,err=public.GetCli().NextID()
	if err!=nil{
		res.Code=pb.Code_FIELD
		return
	}
	return
}

func (s *Srv)GetServID(ctx context.Context, req *pb.Request) (res *pb.Response ,err error){
	mid:=public.GetCli().GetMachineID()
	res=&pb.Response{
		ServID: int32(mid),
	}
	if mid<0{
		return res,errors.New("ger serv id err ")
	}
	return res,nil
}