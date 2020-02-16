package main

import (
	"fmt"
	"os"
	"time"
	"log"
	"os/exec"
	"os/signal"
	"syscall"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var (
	c MQTT.Client
)
func init() {
}

var onMessage MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	run()
}

func init_mqtt() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://www.ruichen.top:1883")
	// opts.SetClientID("go-simple")
	opts.SetUsername("admin")
	opts.SetPassword("123123123")

  //create and start a client using the above ClientOptions
	c = MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

  //subscribe to the topic /go-mqtt/sample and request messages to be delivered
  //at a maximum qos of zero, wait for the receipt to confirm the subscription
  if token := c.Subscribe("pi/ip", 0, onMessage); token.Wait() && token.Error() != nil {
    fmt.Println(token.Error())
    os.Exit(1)
  }
}

func publish(topic, message string) {
    token := c.Publish(topic, 0, false, message)
    token.Wait()
}

// 执行脚本的函数，返回结果和错误
func RunScriptFile(script string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c",  script)

	if output, err := cmd.Output(); err != nil {
		return "", err
	} else {
		return string(output), nil
	}
}

func run () {
	pwd, _ := os.Getwd()
	if output, err := RunScriptFile(pwd + "/getip.sh"); err == nil {
		publish("pi/connect", output)
	}else{
		log.Fatal(err)
	}
}

func main() {
	init_mqtt()
	time.Sleep(time.Second * 1)
	//合建chan
    c := make(chan os.Signal)
    //监听指定信号 ctrl+c kill
    signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
    //阻塞直到有信号传入
    go run()
    //阻塞直至有信号传入
    s := <-c
    fmt.Println("退出信号", s)
	
}