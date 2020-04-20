# udpclient
Simple UDP connection in Go

## How to run

### Cloning
You can "git clone" my repo with :

```
git clone https://github.com/TRedzepagic/udpclient.git
```
Then run with :

```
go run main.go "Address:Port" 
```

Doing this in tandem with my udp listen server will open a netcat-like environment for sending messages which the server will display (on its end) and log.

Upon successfully sending a message, the server will respond to the client. Also, every N seconds (can be configured in udplistener) the server will send a timer tick to the client, which will then be logged into a file.