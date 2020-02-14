package main

import (
	"fmt"
	"time"
	"log"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/stianeikeland/go-rpio"
)

var (
	is_close = true
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

func read(pinCode uint8, isPullUp bool) (interface{}){
	open()
	pin := rpio.Pin(pinCode)
	pin.Input()
	if isPullUp {
		pin.PullUp()
	}else{
		pin.PullDown()
	}
	res := pin.Read()
	return res
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello there")
}

func CleanHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	cleanup()
    fmt.Fprintf(w, "1")
}

func convert(src string) uint8 {
	u, _ := strconv.ParseUint(src, 10, 8)
	return uint8(u)
}
func ToggleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := vars["pin"]
	toggle(convert(pin))
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "1")
}

func TurnOnHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := vars["pin"]
	high(convert(pin))
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "1")
}

func TurnOffHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := vars["pin"]
	low(convert(pin))
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "1")
}

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pin := vars["pin"]
	isPulldown := vars["pulldown"]
	state := read(convert(pin), isPulldown == "up")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "%d", state)
}

func main() {
	
	r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
    r.HandleFunc("/clean", CleanHandler)
	r.HandleFunc("/toggle/{pin}", ToggleHandler)
	r.HandleFunc("/on/{pin}", TurnOnHandler)
	r.HandleFunc("/off/{pin}", TurnOffHandler)
	r.HandleFunc("/read/{pin}/{pulldown}", ReadHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8007",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}