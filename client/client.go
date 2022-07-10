package client

import (
	"context"
	"fmt"
	"time"

	_ "github.com/et-zone/eids/sonyflake"
	pb "github.com/et-zone/proto_api/ids"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)
var cli *client

type client struct {
	Conn *grpc.ClientConn
	Cancel context.CancelFunc
	Ctx context.Context
	Pbclient pb.IDsInterfaceClient
	ServID int32
	timeout int
}

func RunClient(ip string,port ,timeout int)error{
	conn, err := grpc.Dial(ip+":"+fmt.Sprintf("%v",port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("did not connect: %v", err)
		return err
	}
	// Contact the server and print out its response.
	//ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	cli=&client{Conn: conn, Cancel: cancel, Ctx: ctx, Pbclient: pb.NewIDsInterfaceClient(conn),timeout:timeout}
	machineID,err:=getServID()
	if err!=nil{
		log.Println("init ids err , not connect server")
		return err
	}

	//err=sonyflake.EIDCliet.InitSonyFlake(machineID)
	//if err!=nil{
	//	log.Fatalf("RunClient err , failed to InitSonyFlake : %v", err)
	//	return err
	//}

	cli.ServID=machineID
	return err
}

func (c *client)Close(){
	c.Conn.Close()
}

func NextID()(uint64,error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cli.timeout)*time.Second)
	defer cancel()
	r,err:=cli.Pbclient.GetID(ctx,&pb.Request{})
	if r==nil{
		return 0,err
	}
	return r.ID,err
}
func getServID()(int32,error){
	r,err:=cli.Pbclient.GetServID(cli.Ctx,&pb.Request{})
	return r.ServID,err
}