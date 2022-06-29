package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	c, _ := rocketmq.NewPushConsumer(
		//consumer.WithGroupName("testGroup"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"192.168.1.9:9876"})),
	)
	err := c.Subscribe("test", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Println("收到消息：")
			//fmt.Printf("subscribe callback: %v \n", msgs[i])
			fmt.Print(msgs[i].MsgId, "  ", string(msgs[i].Body), "  ")
			fmt.Println(msgs[i].GetProperty("name"))
		}

		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	//开始消费
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	//不能马上退出，要等到收到消息
	time.Sleep(time.Hour)
	err = c.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error: %s", err.Error())
	}
}
