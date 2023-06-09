package pubsubservice

import "net/netip"

const TopicUrl = "mem://updates"

type updateMessagePayload struct {
	Username string     `json:username`
	Ip       netip.Addr `json:ip`
}
