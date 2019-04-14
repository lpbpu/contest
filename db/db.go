package db

import (
	"fmt"
	"strconv"
)

type ipAssignment struct {
	IPRangeStart string `json:"ipRangeStart"`
	IPRangeEnd   string `json:"ipRangeEnd"`
}

type route struct {
	Target string `json:"target"`
	Via    string `json:"via"`
}

type v6assignmode struct {
	Zt      bool `json:"zt"`
	Rfc4193 bool `json:"rfc4193"`
	V6plane bool `json:"6plane"`
}

type v4assignmode struct {
	Zt bool `json:"zt"`
}

type ncapabilities struct {
	ID      uint64 `json:"id"`
	Default bool   `json:"default"`
}

type ntag struct {
	ID      uint32 `json:"id"`
	Default bool   `json:"default"`
}

type Networkconfig struct {
	Authtokens        interface{}     `json:"authtokens"`
	Capabilities      []ncapabilities `json:"capabilities"`
	Creationtime      *uint64         `json:"creationTime"`
	Enablebroadcast   *bool           `json:"enableBroadcast"`
	ID                *string         `json:"id"`
	IpassignmentPools []ipAssignment  `json:"ipAssignmentPools"`
	Mtu               *uint16         `json:"mtu"`
	Multicastlimit    *uint16         `json:"multicastLimit"`
	Name              *string         `json:"Name"`
	Nwid              *string         `json:"nwid"`
	Objtype           *string         `json:"objtype"`
	Private           *bool           `json:"private"`
	Remotetracelevel  *int            `json:"remoteTraceLevel"`
	Remotetracetarget *string         `json:"remoteTraceTarget"`
	Revision          *int            `json:"revision"`
	Routes            []route         `json:"routes"`
	Rules             interface{}     `json:"rules"`
	Tags              []ntag          `json:"tags"`
	V4assignmode      *v4assignmode   `json:"v4assignmode"`
	V6assignmode      *v6assignmode   `json:"v6assignmode"`
}

type Memberconfig struct {
	Activebridge                 *bool       `json:"activeBridge"`
	Noautoassignips              *bool       `json:"noAutoAssignIps"`
	Remotetracetarget            *string     `json:"remoteTraceTarget"`
	Remotetracelevel             *int        `json:"remoteTraceLevel"`
	Authorized                   *bool       `json:"authorized"`
	Lastauthorizedcredentialtype *string     `json:"lastAuthorizedCredentialType"`
	Lastauthorizedcredential     *string     `json:"lastAuthorizedCredential"`
	Lastauthorizedtime           *uint64     `json:"lastAuthorizedTime	"`
	Lastdeauthorizedtime         *uint64     `json:"lastDeauthorizedTime"`
	Ipassignments                []string    `json:"ipAssignments"`
	Tags                         interface{} `json:"tags"`
	Capabilities                 interface{} `json:"capabilities"`
	ID                           *string     `json:"id"`
	Address                      *string     `json:"address"`
	Nwid                         *string     `json:"nwid"`
	Creationtime                 *uint64     `json:"creationTime"`
	Identity                     *string     `json:"identity"`
	Objtype                      *string     `json:"objtype"`
	Revision                     *int        `json:"revision"`
	Vmajor                       *int        `json:"vMajor"`
	Vminor                       *int        `json:"vMinor"`
	Vproto                       *int        `json:"vProto"`
	Vrev                         *int        `json:"vRev"`
}

type Network struct {
	config            Networkconfig
	members           map[uint64]Memberconfig
	activemembers     []uint64
	authorizedmembers []uint64
}

type DB struct {
	path     string
	networks map[uint64]Network
}

func NewDB(path string) *DB {
	var db *DB

	networks := make(map[uint64]Network)

	db = &DB{
		path:     path,
		networks: networks,
	}

	return db
}

func (db *DB) Onnetworkchanged(old *Networkconfig, new *Networkconfig, push bool) {

	if new != nil {
		var network Network

		id, err := strconv.ParseUint(*new.ID, 16, 64)
		if err != nil {
			return
		}

		fmt.Printf("id=%d\n", id)

		network, ok := db.networks[id]
		if !ok {
			db.networks[id] = Network{}
		}

		network.config = *new

		if push {
			fmt.Printf("push to controller network changed\n")
		}
	} else if old != nil {
		id, err := strconv.ParseUint(*old.ID, 16, 64)
		if err != nil {
			return
		}

		delete(db.networks, id)

	}

}

func (db *DB) Onmemberchanged(old *Memberconfig, new *Memberconfig, push bool) {
	if old != nil {
		/*
			memberid, err1 := strconv.ParseUint(*old.ID, 16, 64)
			if err1 != nil {
				return
			}

			networkid, err2 := strconv.ParseUint(*old.Nwid, 16, 64)
			if err2 != nil {
				return
			}*/

	}
}
