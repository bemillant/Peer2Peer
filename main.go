package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	ping "github.com/NaddiNadja/peer-to-peer/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type peer struct {
	ping.UnimplementedPingServer
	id            int32
	amountOfPings map[int32]int32
	clients       map[int32]ping.PingClient
	ctx           context.Context
	skrrrtNumber  int32
	wantToEnterCS bool
	neighbour     ping.PingClient
	hasToken      bool
}

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 5001

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := &peer{
		id:            ownPort,
		clients:       make(map[int32]ping.PingClient),
		ctx:           ctx,
		neighbour:     nil,
		wantToEnterCS: false,
		hasToken:      false,
		skrrrtNumber:  0,
	}

	if ownPort == 5001 {
		p.hasToken = true

		// This is a place in the code that is guaranteed to only run once, at startup
		// The following method call ensures the critical section is empty in the beginning.
		p.wipeCriticalSection()
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
		port := int32(5001) + int32(i)

		if port == ownPort {
			continue
		}

		var conn *grpc.ClientConn
		fmt.Printf("Trying to dial: %v\n", port)
		insecure := insecure.NewCredentials()
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithTransportCredentials(insecure), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}

		log.Printf("--- Succesfully dialed to %v\n", port)

		defer conn.Close()
		c := ping.NewPingClient(conn)
		p.clients[port] = c
	}

	p.setNeighbour()

	//scanner := bufio.NewScanner(os.Stdin)
	for {

		var message string
		fmt.Scan(&message)

		switch message {
		case "pass":
			p.PassTokenToNeighbour()
			continue
		case "requestCS":
			p.requestCriticalSection()
			continue
		case "accessCS":
			p.handleCriticalSection()
			continue
		default:
			continue
		}
	}
}

// func incTime(){
// 	for
// }

func (p *peer) Ping(ctx context.Context, req *ping.Request) (*ping.Reply, error) {
	id := req.Id
	p.amountOfPings[id] += 1

	rep := &ping.Reply{Amount: p.amountOfPings[id]}
	return rep, nil
}

func (p *peer) sendPingToAll() {
	request := &ping.Request{Id: p.id}

	for id, client := range p.clients {
		reply, err := client.Ping(p.ctx, request)
		if err != nil {
			fmt.Println("something went wrong")
		}
		fmt.Printf("Got reply from id %v: %v\n", id, reply.Amount)
	}
}

func (p *peer) Token(ctx context.Context, pass *ping.Pass) (*ping.Acknowledgement, error) {
	Ack := &ping.Acknowledgement{
		Message: "Token has succesfully been passed",
	}
	p.hasToken = true

	log.Printf("token has been received from %v", pass.Id)
	return Ack, nil
}

func (p *peer) PassTokenToNeighbour() {
	if p.hasToken {
		token := &ping.Pass{
			Message: "Passing on token",
			Id:      p.id,
		}

		ack, err := p.neighbour.Token(p.ctx, token)
		if err != nil {
			fmt.Println("something went wrong when trying to pass the token")
		}

		p.hasToken = false
		log.Printf("Token succesfully passed to client at port %v with message: %v", ack.Message, p.getNeighbourID())

	} else {
		log.Print("does not possess token, so cannot pass")
	}
}

// // method to randomise request for critical sections
// func (p *peer) randomUpdateWantToEnterCS() {
// 	rand.Seed(time.Now().UnixNano())
// 	n := rand.Intn(10) // n will be between 0 and 10
// 	time.Sleep(time.Duration(n) * time.Second)
// 	p.requestCriticalSection()
// }

func (p *peer) setNeighbour() {

	if p.id == 5003 {
		p.neighbour = p.clients[5001]
	} else {
		p.neighbour = p.clients[p.id+1]
	}
}

func (p *peer) getNeighbourID() int32 {

	if p.id == 5003 {
		return 5001
	} else {
		return p.id + 1
	}
}

func (p *peer) writeToFile(message string) {

	f, err := os.OpenFile(
		"critical_section.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	if _, err := f.WriteString(message + "\n"); err != nil {
		log.Println(err)
	}
}

func (p *peer) wipeCriticalSection() {
	if err := os.Truncate("critical_section.log", 0); err != nil {
		log.Print("Failed to truncate: %v", err)
	}
}

func (p *peer) requestCriticalSection() {
	p.wantToEnterCS = true
	fmt.Printf("peer with Id: %v now request to enter the Critical section \n", p.id)
}

func (p *peer) handleCriticalSection() {

	if p.wantToEnterCS && p.hasToken {
		p.writeToFile(p.generateCSMessage())
		p.wantToEnterCS = false
	} else if p.hasToken {
		fmt.Println("no request made, so cannot access critical section")
	} else {
		fmt.Println("does not have token, so cannot access critical section")
	}
}

func (p *peer) generateCSMessage() string {
	var message string
	fmt.Println("Input text to write to critical section: ")
	fmt.Scanln(&message)
	return message
}
