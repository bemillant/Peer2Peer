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

// func (p *peer) oldPingToAll() {
// 	request := &ping.PassToken{SkrrrtNumber: p.s}
// 	for id, client := range p.clients {
// 		reply, err := client.ping(p.ctx, request)
// 		if err != nil {
// 			fmt.Println("something went wrong")
// 		}
// 		fmt.Printf("Got reply from id %v: %v\n", id, reply.Amount)
// 	}
// }

// func (p *peer) ping(ctx context.Context, req *ping.PassToken) (*ping.Reply, error) {
// 	// id := req.Id
// 	// p.amountOfPings[id] += 1

// 	// rep := &ping.Reply{Amount: p.amountOfPings[id]}
// 	return rep, nil
// }

func (p *peer) passToken(ctx context.Context) (*ping.Reply, error) {

	if p.wantToEnterCS == true {
		p.skrrrtNumber++
	}

	// message := &ping.PassToken{
	// 	SkrrrtNumber: p.skrrrtNumber,
	// }

	p.hasToken = false

	reply := &ping.Reply{
		Message: "Token has been passed succesfully",
	}
	log.Print(reply.GetMessage)
	return reply, nil

}

func (p *peer) waitForToken() {
	for !p.hasToken {
		// do nothing
	}

	p.passToken(p.ctx)
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
