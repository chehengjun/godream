package main

import (
	"log"
	"net/rpc"
	"time"
)

func main() {
	log.Println("This is rpc test.")
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:5021")
	log.Println("clinet", client)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var args = "hello rp****c"
	var reply string

	for {
		time.Sleep(5e9)
		log.Println("**********conn spool", client)
		err = client.Call("RDB.Add", args, &reply)
		if err != nil {
			client.Close()
		}

		log.Println("**********conn spool close", client)
		log.Println("rpc call error:", err)

	}

}
