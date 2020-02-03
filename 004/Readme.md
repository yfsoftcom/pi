# Camera

- [ ] [https://tutorials-raspberrypi.com/raspberry-pi-security-camera-livestream-setup/](https://tutorials-raspberrypi.com/raspberry-pi-security-camera-livestream-setup/)

- [x] [https://randomnerdtutorials.com/video-streaming-with-raspberry-pi-camera/](https://randomnerdtutorials.com/video-streaming-with-raspberry-pi-camera/)

- [x] [https://www.electronicsforu.com/electronics-projects/make-video-streaming-camera-with-raspberry-pi](https://www.electronicsforu.com/electronics-projects/make-video-streaming-camera-with-raspberry-pi)

## Enable camera

```bash
sudo raspi-config
```

## Camera Detail

```bash
$ v4l2-ctl -V
```

```
pi@raspberrypi:~ $ v4l2-ctl -V
Format Video Capture:
	Width/Height  : 1024/768
	Pixel Format  : 'JPEG'
	Field         : None
	Bytes per Line: 0
	Size Image    : 786432
	Colorspace    : JPEG (JFIF/ITU601)
	Flags         : 
```

## Open the live stream

```bash
$ sudo apt-get install vlc
$ raspivid -o - -t 0 -n | cvlc -vvv stream:///dev/stdin --sout '#rtp{sdp=rtsp://:8554/}' :demux=h264
```

## The live stream

`rtsp://pi:8554/`