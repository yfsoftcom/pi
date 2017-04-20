# yfsoft/PI
运行在 pi 3 上的项目

# 001
websocket连接服务端

# Depdents

- fpm
- websocket lib 
  - Install package ( `pip install -U socketIO-client` )

# 树莓派设置开机启动python脚本
#### 设置python脚本开机启动

编辑 `/etc/rc.local`

`sudo vim /etc/rc.local`

在 `exit 0` 上一行输入：

```bash
echo "Starting PI LOGIN"
python /home/pi/workspace/pi/001/libs/ws_client.py &
```

# 通过文件配置 wifi 

Mac Address: B8-27-EB-9D-85-D7

用编辑器打开interfaces文件

`$ sudo vim /etc/network/interfaces`

在 `auto wlan0` 下面找到默认的 wifi 热点信息的配置文件路径

```bash
allow-hotplug wlan0
iface wlan0 inet dhcp
    wpa-conf /etc/wpa_supplicant/wpa_supplicant.conf
```

编辑文件 `/etc/wpa_supplicant/wpa_supplicant.conf`

`sudo vim /etc/wpa_supplicant/wpa_supplicant.conf`

```bash
ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
#update_config=1

# home
network={
        ssid="YfsoftCOM"
        psk="741235896"
        key_mgmt=WPA-PSK
}
# office
network={
        ssid="205"
        psk="123456789"
        key_mgmt=WPA-PSK
}
```

输入命令启用无线网卡

`sudo ifup wlan0`

重启网卡
`sudo /etc/init.d/networking restart`
