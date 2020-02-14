all: build transfer run

install:
	go get github.com/stianeikeland/go-rpio
	sudo apt install mosquitto mosquitto-clients
build:
	env GOOS=linux GOARCH=arm GOARM=5 go build gopio.go
transfer:
	scp gopio pi@192.168.88.205:/home/pi/
run:
	ssh -t pi@192.168.88.205 "./gopio"

pub:
	mosquitto_pub -h www.ruichen.top -t go-mqtt/sample/toggle -m "5" -u "admin" -P "123123123"
