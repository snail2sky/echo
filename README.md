!!!

## Current command functionality has been merged into https://github.com/snail2sky/bbx
- usage
```bash
bbx echo --help
```

# echo
TCP or UDP echo server

## build
```bash
go build
```

## run
```bash
# run 
./echo
# help
./echo --help
```

## usage
- ./echo [-buf-size uint] [-host string] [-port int] [-protocol string]
    - -buf-size ECHO server receive buffer size. (default 1024)
    - -host ECHO server listen on this address. (default "0.0.0.0")
    - -port ECHO server listen on this port. (default 7890)
    - -protocol ECHO server use tcp/udp/all. (default "tcp")
