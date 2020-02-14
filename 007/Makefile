install:
	go get github.com/stianeikeland/go-rpio
build:
	env GOOS=linux GOARCH=arm GOARM=5 go build gopio.go
transfer:
	scp gopio pi@192.168.88.205:/home/pi/
run:
	ssh -t pi@192.168.88.205 "./gopio"
