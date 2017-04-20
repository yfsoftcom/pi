# yfsoft/PI
运行在 pi 3 上的项目

# 001
websocket连接服务端

# Depdents

- fpm
- websocket lib (`pip install websocket-client`)


# 树莓派设置开机启动python脚本
#### 设置python脚本开机启动

编辑 `/etc/rc.local`

`sudo vim /etc/rc.local`

在 `exit 0` 上一行输入：

```bash
echo "Starting PI LOGIN"
python /home/pi/workspace/pi/001/libs/ws_client.py &
```

最后设置开机启动就好了