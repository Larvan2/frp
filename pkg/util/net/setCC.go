package net

import (
	"github.com/metacubex/quic-go"
	c "github.com/metacubex/quic-go/congestion"

	"github.com/fatedier/frp/pkg/util/net/bbr"
	"github.com/fatedier/frp/pkg/util/net/cubic"
)

const (
	DefaultStreamReceiveWindow     = 15728640 // 15 MB/s
	DefaultConnectionReceiveWindow = 67108864 // 64 MB/s
)

func SetCongestionController(quicConn quic.Connection, cc string, cwnd int) {
	if cwnd == 0 {
		cwnd = 32
	}
	switch cc {
	case "cubic":
		quicConn.SetCongestionControl(
			cubic.NewCubicSender(
				cubic.DefaultClock{},
				cubic.GetInitialPacketSize(quicConn),
				false,
			),
		)
	case "new_reno":
		quicConn.SetCongestionControl(
			cubic.NewCubicSender(
				cubic.DefaultClock{},
				cubic.GetInitialPacketSize(quicConn),
				true,
			),
		)
	case "bbr":
		quicConn.SetCongestionControl(
			bbr.NewBbrSender(
				bbr.DefaultClock{},
				bbr.GetInitialPacketSize(quicConn),
				c.ByteCount(cwnd),
			),
		)
	}
}
