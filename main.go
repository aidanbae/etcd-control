package main

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		log.Println(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	resp, err := cli.Put(ctx, "sample_key", "sample_value")
	cancel()
	if err != nil {
		// handle error!
		log.Println(err)
	}
	// use the response

	log.Println(resp)
}