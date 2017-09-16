package main

import "net"

func main() {
	logger, err := emitio.New()
	if err != nil {
		panic(err)
	}
	logger.Info(
		"hello there",
		emitio.String("test", "value"),
		emitio.IP("from", net.ParseIP("2001:db8::68")),
		emitio.Double("rate", 0.863),
	)
	defer eio.Close()
}
