# Tools

## BLE Scanner

Used to scan for Bluetooth devices.

```
go run ./blescanner/
```

## AP Connect

Used to connect an Arduino with WiFi to an access point.

```
tinygo flash -target nano-rp2040 -ldflags="-X main.ssid=thessid -X main.pass=thepassword" ./apconnect/
```
