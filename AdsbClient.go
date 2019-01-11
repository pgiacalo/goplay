package main

import (
	"bufio"
	//"flag"
	"fmt"
	"net"
	"os"
	//	"strconv"
	"strings"
	//	"github.com/davecgh/go-spew/"
)

/*
	see: http://woodair.net/sbs/article/barebones42_socket_data.htm

	This struct holds the data received from aircraft via ADS-B messages (ADSB).

	It is received in a format called "SBS1 (BaseStation, beast) format".
	These ADSB messages are provided by the dump1090 program and are available on port 30003 (by default).
	Messages (MSG) received from aircraft may be one of eight types.

	reference: http://woodair.net/sbs/article/barebones42_socket_data.htm

	Examples:
	MSG,4,5,211,4CA2D6,10057,2008/11/28,14:53:49.986,2008/11/28,14:58:51.153,,,408.3,146.4,,,64,,,,,
	MSG,8,5,211,4CA2D6,10057,2008/11/28,14:53:50.391,2008/11/28,14:58:51.153,,,,,,,,,,,,0
	MSG,4,5,211,4CA2D6,10057,2008/11/28,14:53:50.391,2008/11/28,14:58:51.153,,,408.3,146.4,,,64,,,,,
	MSG,3,5,211,4CA2D6,10057,2008/11/28,14:53:50.594,2008/11/28,14:58:51.153,,37000,,,51.45735,-1.02826,,,0,0,0,0
	MSG,8,5,812,ABBEE3,10095,2008/11/28,14:53:50.594,2008/11/28,14:58:51.153,,,,,,,,,,,,0

	SEL,,496,2286,4CA4E5,27215,2010/02/19,18:06:07.710,2010/02/19,18:06:07.710,RYR1427
	ID,,496,7162,405637,27928,2010/02/19,18:06:07.115,2010/02/19,18:06:07.115,EZY691A
	AIR,,496,5906,400F01,27931,2010/02/19,18:06:07.128,2010/02/19,18:06:07.128
	STA,,5,179,400AE7,10103,2008/11/28,14:58:51.153,2008/11/28,14:58:51.153,RM
	CLK,,496,-1,,-1,2010/02/19,18:18:19.036,2010/02/19,18:18:19.036
	MSG,1,145,256,7404F2,11267,2008/11/28,23:48:18.611,2008/11/28,23:53:19.161,RJA1118,,,,,,,,,,,
	MSG,2,496,603,400CB6,13168,2008/10/13,12:24:32.414,2008/10/13,12:28:52.074,,,0,76.4,258.3,54.05735,-4.38826,,,,,,0
	MSG,3,496,211,4CA2D6,10057,2008/11/28,14:53:50.594,2008/11/28,14:58:51.153,,37000,,,51.45735,-1.02826,,,0,0,0,0
	MSG,4,496,469,4CA767,27854,2010/02/19,17:58:13.039,2010/02/19,17:58:13.368,,,288.6,103.2,,,-832,,,,,
	MSG,5,496,329,394A65,27868,2010/02/19,17:58:12.644,2010/02/19,17:58:13.368,,10000,,,,,,,0,,0,0
	MSG,6,496,237,4CA215,27864,2010/02/19,17:58:12.846,2010/02/19,17:58:13.368,,33325,,,,,,0271,0,0,0,0
	MSG,7,496,742,51106E,27929,2011/03/06,07:57:36.523,2011/03/06,07:57:37.054,,3775,,,,,,,,,,0
	MSG,8,496,194,405F4E,27884,2010/02/19,17:58:13.244,2010/02/19,17:58:13.368,,,,,,,,,,,,0
*/

/*
	Type adsbMsg has fields for all of the values for the 22 fields of ADSB MSG messages.
	Note that there are 8 different types of MSG messages that are identified by their Transmission Type identifier.
	Depending of the Transmission Type, some of the fields will be empty.
*/
type adsbMsg struct {
	MessageType          string  //Field 1: (MSG, STA, ID, AIR, SEL, or CLK)
	TransmissionType     int     //Field 2: MSG sub types 1 to 8. Not used by other message types.
	SessionID            int     //Field 3: Database Session record number
	AircraftID           string  //Field 4: Database Aircraft record number
	HexIdent             string  //Field 5: Aircraft Mode S hexadecimal code
	FlightID             string  //Field 6: Database Flight record number
	DateMessageGenerated string  //Field 7:
	TimeMessageGenerated string  //Field 8:
	DateMessageLogged    string  //Field 9:
	TimeMessageLogged    string  //Field 10:
	Callsign             string  //Field 11: An eight digit flight ID - can be flight number or registration (or even nothing).
	Altitude             int     //Field 12: Mode C altitude. Height relative to 1013.2mb (Flight Level). Not height AMSL..
	GroundSpeed          int     //Field 13: Speed over ground (not indicated airspeed)
	Track                int     //Field 14: Track of aircraft (not heading). Derived from the velocity E/W and velocity N/S
	Latitude             float32 //Field 15: North and East positive. South and West negative.
	Longitude            float32 //Field 16: North and East positive. South and West negative.
	VerticalRate         int     //Field 17: 64ft resolution
	Squawk               string  //Field 18: Assigned Mode Person squawk code.
	Alert                string  //Field 19: (Squawk change)	 Flag to indicate squawk has changed.
	Emergency            string  //Field 20: Flag to indicate emergency code has been set
	SPI                  string  //Field 21: (Ident Flag to indicate transponder Ident has been activated.
	IsOnGround           string  //Field 22: Flag to indicate ground squat switch is active
}

type Key struct {
	AircraftID, MessageType string
}

var aircraftMsgTypeCount = make(map[Key]int)

func main() {

	const dump1090SocketAddress = "127.0.0.1:30003"

	// connect to the socket
	conn, err := net.Dial("tcp", dump1090SocketAddress)
	if err != nil {
		fmt.Printf("Error: unable to connect to socket. Is rtl-sdr connected with dump1090 running? "+
			"Note: launch dump1090 with alias \"adsb\" or \"adsbfile\")? \n%s", err)
		os.Exit(1)
	}

	// loop forever to read messages arriving from the socket
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Print("Error while reading from socket:\n", err)
			os.Exit(1)
		}

		fmt.Print("Message received from dump1090 socket: " + message)
		var aMessage []string = strings.Split(message, ",")
		var msg adsbMsg = adsbMsg{
			MessageType: aMessage[0],
			//TransmissionType:     strconv.Atoi(aMessage[1]),
			//SessionID : aMessage[2],
			AircraftID:           aMessage[3],
			HexIdent:             aMessage[4],
			FlightID:             aMessage[5],
			DateMessageGenerated: aMessage[6],
			TimeMessageGenerated: aMessage[7],
			DateMessageLogged:    aMessage[8],
			TimeMessageLogged:    aMessage[9],
			Callsign:             aMessage[10],
			//Altitude:             aMessage[11],
			//GroundSpeed:          aMessage[12],
			//Track:                aMessage[13],
			//Latitude:             aMessage[14],
			//Longitude:            aMessage[15],
			//VerticalRate:         aMessage[16],
			Squawk:     aMessage[17],
			Alert:      aMessage[18],
			Emergency:  aMessage[19],
			SPI:        aMessage[20],
			IsOnGround: aMessage[21],
		}
		aircraftMsgTypeCount[Key{aMessage[3], aMessage[0]}]++

		fmt.Printf("msg:%v\n", msg)

	}
}
