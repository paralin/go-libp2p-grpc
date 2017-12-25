package p2pgrpc

import "net"

// fakeLocalAddr returns a dummy local address.
func fakeLocalAddr() net.Addr {
	localIp := net.ParseIP("127.0.0.1")
	return &net.TCPAddr{IP: localIp, Port: 0}
}

// fakeRemoteAddr returns a dummy remote address.
func fakeRemoteAddr() net.Addr {
	remoteIp := net.ParseIP("127.1.0.1")
	return &net.TCPAddr{IP: remoteIp, Port: 0}
}
