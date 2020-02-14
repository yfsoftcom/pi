package main

import (
	"fmt"
	"os"
	"time"
	"strconv"
	"os/signal"
	"syscall"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/stianeikeland/go-rpio"
)

var (
	is_close = true
	c MQTT.Client
)
func init() {
	open()
}

func open() {
	if !is_close {
		return
	}
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}
	is_close = false
}
func cleanup () {
	is_close = true
	rpio.Close()	
}

func toggle(pinCode uint8)  {
	open()
	pin := rpio.Pin(pinCode)
	pin.Output()
	pin.Toggle()
}

func high(pinCode uint8) {
	open()
	pin := rpio.Pin(pinCode)
	pin.Output()
	pin.High()
}
func low(pinCode uint8) {
	open()
	pin := rpio.Pin(pinCode)
	pin.Output()
	pin.Low()
}

func read(pinCode uint8) (interface{}){
	open()
	pin := rpio.Pin(pinCode)
	pin.Input()
	res := pin.Read()
	return res
}

func convert(src string) uint8 {
	u, _ := strconv.ParseUint(src, 10, 8)
	return uint8(u)
}

var onMessage MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	pin := convert(string(msg.Payload()))
	toggle(pin)
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
  if token := c.Subscribe("go-mqtt/sample/toggle", 0, onMessage); token.Wait() && token.Error() != nil {
    fmt.Println(token.Error())
    os.Exit(1)
  }
}

func publish(topic, message string) {
    token := c.Publish(topic, 0, false, message)
    token.Wait()
}

func run () {

	for i :=0 ; i< 10 ; i++ {
		publish("go-mqtt/sample/toggle", `5`)
		time.Sleep(time.Second )
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