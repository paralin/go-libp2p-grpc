package p2pgrpc

import (
	inet "github.com/libp2p/go-libp2p-net"
	manet "github.com/multiformats/go-multiaddr-net"
	"net"
)

// streamConn represents a net.Conn wrapped to be compatible with net.conn
type streamConn struct {
	inet.Stream
}

// LocalAddr returns the local address.
func (c *streamConn) LocalAddr() net.Addr {
	addr, err := manet.ToNetAddr(c.Stream.Conn().LocalMultiaddr())
	if err != nil {
		return fakeLocalAddr()
	}
	return addr
}

// RemoteAddr returns the remote address.
func (c *streamConn) RemoteAddr() net.Addr {
	addr, err := manet.ToNetAddr(c.Stream.Conn().RemoteMultiaddr())
	if err != nil {
		return fakeRemoteAddr()
	}
	return addr
}

var _ net.Conn = &streamConn{}
