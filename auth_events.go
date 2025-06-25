package steam

import (
	"github.com/anovokreschenov/go-steam/protocol/protobuf"
	. "github.com/anovokreschenov/go-steam/protocol/steamlang"
	"github.com/anovokreschenov/go-steam/steamid"
)

type LoggedOnEvent struct {
	Result         EResult
	ExtendedResult EResult
	AccountFlags   EAccountFlags
	ClientSteamId  steamid.SteamId `json:",string"`
	Body           *protobuf.CMsgClientLogonResponse
}

type LogOnFailedEvent struct {
	Result EResult
}

type LoginKeyEvent struct {
	UniqueId uint32
	LoginKey string
}

type LoggedOffEvent struct {
	Result EResult
}

type MachineAuthUpdateEvent struct {
	Hash []byte
}

type AccountInfoEvent struct {
	PersonaName          string
	Country              string
	CountAuthedComputers int32
	AccountFlags         EAccountFlags
	FacebookId           uint64 `json:",string"`
	FacebookName         string
}

// Returned when Steam is down for some reason.
// A disconnect will follow, probably.
type SteamFailureEvent struct {
	Result EResult
}
