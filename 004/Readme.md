# Camera

- [https://tutorials-raspberrypi.com/raspberry-pi-security-camera-livestream-setup/](https://tutorials-raspberrypi.com/raspberry-pi-security-camera-livestream-setup/)


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

```