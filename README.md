# simple-proxy

The lantern only serve the computer that installs it, simple-proxy adapts to the lantern, allowing the LAN device to share the lantern.

# Install

You can install from source (assume you have go installed):

```
# on lantern server
go get github.com/lincolnlee/simple-proxy/cmd/simple-proxy-for-lantern
```

It's recommended to disable cgo when compiling simple-proxy. This will prevent the go runtime from creating too many threads for dns lookup.

# Usage

simple-proxy-for-lantern program will look for `config.json` in the current directory. You can use `-c` option to specify another configuration file.

Configuration file is in json format and you can download the sample [`config.json`](https://github.com/lincolnlee/simple-proxy/blob/master/sp/config.go), change the following values:

```
server          your server ip
server_port     server port
client      only this ip is allowed access simple-proxy，not limited by default
password        a password used to identify whether access is allowed，not limited by default
Proxy         lantern ip
ProxyPort          lantern port
```

Run `simple-proxy-for-lantern` on your server. To run it in the background, run `simple-proxy-for-lantern > log &`.

On client, change proxy settings of your browser to

```
http/https  server:server_port
```

## Command line options

Command line options can override settings from configuration files. Use `-h` option to see all available options.

```
simple-proxy-for-lantern -s server -p server_port -b client
    -k password -l proxy
    -x proxy_port -c config.json
```

