# Peer2Peer

------- EXPLANATION OF IMPLEMENTATION ---------

    OBS:
    We have made a rudimentary implementaion of the automatic passing of the token. We decided to keep the below text since it still explains our approach. The only thing that has changed is the fact that the clients no longer has to manually pass the token. 

    OBS 2: 
    The SimulationLog.txt file contains the logs from a simulation of a runthrough of the application, that showcase the different state and actions a client can be in and take.

    In our implementation we followed the "Token Ring" approach, which ensures one-at-a-time-access to the critical section, by only making it possible to access when a client is in possession of the token.
    This token is passed from peer to peer. Ideally, and to meet the requirement of liveliness, the passing of the token should happen automatically. This might be accomplished by implementing some sort of timer, that forces the client to pass the token, has it not already done so, within a short timeframe.

    We didn't do this, since it wasn't obvious to us how to do this properly, and since we were short on time.

    Instead, each client has to manually pass the token through terminal commands, which is not ideal, but it serves to slowly demonstrate the different states a client can be in, and the different reactions to the different actions a client might take.

    These are the possible commands a client can issue through the terminal:

        pass --attempts to pass the token. This is only succesful, if the client actually has the token.

        requestCS --this changes the clients own wantsToAccessCS boolean to true. If this boolean is false the       client cannot get access to the critical section, since a client is supposed to request access first.

        accessCS --this let's the client access the critical section, if it both holds the token, and has requested access. In our implementation, the critical section is simply a permission to write a line of text to a file called "critical_section.log"
    
    We believe this implementation, though not perfect, demonstrates how a Token Ring might be implemented to ensure safety and liveliness in a peer to peer system, that utilizes no central server, and communicates only through messages using gRPC.

------- INSTRUCTIONS ------------

To run the program, open 3 separate terminals at the project directory.

If you have MAKE installed (otherwise, see below)

    Run the following commands, one in each terminal:

        make client0

        make client1

        make client2

If you don't have MAKE:

    Run the following commands, one in each terminal:

        go run main.go 0

        go run main.go 1

        go run main.go 2
