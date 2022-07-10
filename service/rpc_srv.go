package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/et-zone/eids/sonyflake"
	pb "github.com/et-zone/proto_api/ids"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Srv struct {
	pb.UnimplementedIDsInterfaceServer
}

func (s *Srv)GetID(ctx context.Context, req *pb.Request) (res *pb.ResponseID,err error){
	res=&pb.ResponseID{
		Code: pb.Code_SUCC,
	}
	res.ID,err= sonyflake.NextID()
	if err!=nil{
		res.Code=pb.Code_FIELD
		return
	}
	return
}

func (s *Srv)GetServID(ctx context.Context, req *pb.Request) (res *pb.Response ,err error){
	mid,err	:= sonyflake.EIDCliet.MachineID()

	res=&pb.Response{ServID: mid,}
	if err!=nil{
		return res,errors.New("ger serv id err ")
	}
	return res,nil
}

func run(port int)error{
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return err
	}
	pb.RegisterIDsInterfaceServer(s, &Srv{})
	log.Printf("eids server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
	return err
}

func RunServ (machineID,port int)error{
	if machineID<0{
		log.Fatalf("RunServ err , machineID can not < 0 ")
		return errors.New("RunServ err , machineID can not < 0 ")
	}
	err:=sonyflake.EIDCliet.InitSonyFlake(int32(machineID))
	if err!=nil{
		log.Fatalf("RunServ err , failed to InitSonyFlake : %v", err)
		return err
	}
	err= run(port)
	if err!=nil{
		return err
	}
	return err
}

func SrvInitSize18 (){
	sonyflake.EIDCliet.SetByteSzie(sonyflake.B_e18)
}
