package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	ping "github.com/bemillant/Peer2Peer/grpc"
	"google.golang.org/grpc"
)

type peer struct {
	ping.UnimplementedPingServer
	id            int32
	clients       map[int32]ping.PingClient
	ctx           context.Context
	skrrrtNumber  int32
	wantToEnterCS bool
	neighbour     ping.PingClient
	hasToken      bool
	// clientConnectionStrings map[int32]ping.PingServer
}

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 5000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := &peer{
		id:            ownPort,
		clients:       make(map[int32]ping.PingClient),
		ctx:           ctx,
		wantToEnterCS: false,
		neighbour:     nil,
		hasToken:      false,
		skrrrtNumber:  0,
	}

	// Create listener tcp on port ownPort
	list, err := net.Listen("tcp", fmt.Sprintf(":%v", ownPort))
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	grpcServer := grpc.NewServer()
	ping.RegisterPingServer(grpcServer, p)

	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to server %v", err)
		}
	}()

	for i := 0; i < 3; i++ {
		port := int32(5000) + int32(i)

		if port == ownPort {
			continue
		}

		var conn *grpc.ClientConn
		fmt.Printf("Trying to dial: %v\n", port)
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		defer conn.Close()
		c := ping.NewPingClient(conn)
		p.clients[port] = c
	}

	p.setNeighbour()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// p.sendPingToAll()
	}

	go p.waitForToken()
}

func (p *peer) receiveToken() {
	p.hasToken = true
}

// passToken method should pass the ability to enter the CS and update the skrrrtNumber to its neighbour
func (p *peer) passToken(ctx context.Context, msg *ping.SkrrrtNumber) (*ping.Reply, error) {

	if p.wantToEnterCS == true {
		p.skrrrtNumber++
	}

	reply, err := p.neighbour.ping(p.ctx)
	if err != nil {
		fmt.Println("something went wrong")
	}

	p.hasToken = false

	log.Print(reply.GetMessage)
	return reply, nil

}

func (p *peer) waitForToken() {
	for !p.hasToken {
		// do nothing
	}

	p.passToken(p.ctx, &ping.SkrrrtNumber{
		SkrrrtNumber: p.skrrrtNumber,
	})
}

func (p *peer) randomUpdateWantToEnterCS() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10) // n will be between 0 and 10
	time.Sleep(time.Duration(n) * time.Second)
	p.wantToEnterCS = true
}

func (p *peer) setNeighbour() {

	if p.id == 5002 {
		p.neighbour = p.clients[5000]
	} else {
		p.neighbour = p.clients[p.id+1]
	}
}

func (p *peer) oldPingToAll() {
	request := &ping.PassToken{SkrrrtNumber: p.s}
	for id, client := range p.clients {
		reply, err := client.ping(p.ctx, request)
		if err != nil {
			fmt.Println("something went wrong")
		}
		fmt.Printf("Got reply from id %v: %v\n", id, reply.Amount)
	}
}

func (p *peer) ping(ctx context.Context) (*ping.Reply, error) {
	rep := &ping.Reply{Message: "send token"}
	return rep, nil
}

// func sendMessage(client *Client, serverConnection gRPC.TimeAskServiceClient) {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	for scanner.Scan() {
// 		input := scanner.Text()

// 		if input == "exit" {
// 			client.stream.Send(&gRPC.Message{
// 				Message:    "exit",
// 				Clientname: client.name,
// 			})

// 			client.connection.Close()

// 		} else {
// 			log.Printf("(Message sent from this client: '%s')", input)
// 			client.stream.Send(&gRPC.Message{
// 				Clientname:       client.name,
// 				Message:          input,
// 				LamportTimestamp: client.lamportTime,
// 			})
// 		}
// 	}
// }
