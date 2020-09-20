package main

import (
	"context"
	"fmt"
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
		log.Println("연결자체에러자나나", err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	resp, err := cli.Put(ctx, "sample_key", "sample_value")
	cancel()
	if err != nil {
		// handle error!
		log.Println("명령어 전송단계 에러, ", err)
	}
	// use the response

	log.Println(resp)

	watchChan := cli.Watch(context.Background(), "sample_key")
	fmt.Println("set WATCH on " + "sample_key")

	//go func() {
	//	fmt.Println("started goroutine for PUT...")
	//	for {
	//		cli.Put(context.Background(), "sample_key", time.Now().String())
	//		fmt.Println("populated " + "sample_key" + " with a value..")
	//		time.Sleep(2 * time.Second)
	//	}
	//
	//}()

	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			fmt.Printf("Event received! %s executed on %q with value %q\n", event.Type, event.Kv.Key, event.Kv.Value)

		}

	}


}
