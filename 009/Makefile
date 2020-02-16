all: build transfer run

install:
	go get github.com/stianeikeland/go-rpio
	sudo apt install mosquitto mosquitto-clients
build:
	env GOOS=linux GOARCH=arm GOARM=5 go build startup.go
transfer:
	scp startup pi@192.168.88.205:/home/pi/
	scp getip.sh pi@192.168.88.205:/home/pi/
run:
	ssh -t pi@192.168.88.205 "./startup"

pub:
	mosquitto_pub -h www.ruichen.top -t pi/ip -m "1" -u "admin" -P "123123123"
sub:
	mosquitto_sub -h www.ruichen.top -t pi/connect -u "admin" -P "123123123"
