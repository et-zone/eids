package client

import (
	"context"
	"fmt"
	"github.com/et-zone/eids/conf"
	"github.com/et-zone/eids/public"
	pb "github.com/et-zone/proto_api/ids"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)
var cli *client

type client struct {
	Conn *grpc.ClientConn
	Cancel context.CancelFunc
	Ctx context.Context
	Pbclient pb.IDsInterfaceClient
	ServID int32
}

func InitClient(ip string,port ,timeout int){
	conn, err := grpc.Dial(ip+":"+fmt.Sprintf("%v",port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	cli=&client{Conn: conn, Cancel: cancel, Ctx: ctx, Pbclient: pb.NewIDsInterfaceClient(conn)}
	mid,err:=getServID()
	if err!=nil{
		panic("init ids err , not get machine_id ")
	}
	conf.ConMsg=&conf.Conf{
		ServID: int(mid),
		Port: port,
	}
	public.InitIDS()
	cli.ServID=mid
}

func (c *client)Close(){
	c.Conn.Close()
}

func NextID()(int64,error){
	r,err:=cli.Pbclient.GetID(cli.Ctx,&pb.Request{ServID: &cli.ServID})
	return int64(r.ID),err
}
func getServID()(int32,error){
	r,err:=cli.Pbclient.GetServID(cli.Ctx,&pb.Request{})
	return r.ServID,err
}