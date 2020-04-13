# udpoclient
Simple UDP connection in Go

## How to run

### Cloning
You can "git clone" my repo with :

```
git clone https://github.com/TRedzepagic/tcpclient.git
```
Then run with :

```
go run main.go "Address:Port" 
```
or 
```
go run main.go ":Port" (localhost)
```

Doing this in tandem with my udp listen server will open a netcat-like environment for sending messages which the server will display (on its end) and log.

Upon successfully sending a message, the server will respond to the client.