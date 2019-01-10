CREATE TABLE ADSB_MESSAGE (
	MessageType          VARCHAR(10),  --Field 1: (MSG STA ID AIR SEL or CLK)
	TransmissionType     SMALLINT,     --Field 2: MSG sub types 1 to 8. Not used by other message types.
	SessionID            INT,          --Field 3: Database Session record number
	AircraftID           VARCHAR(10),  --Field 4: Database Aircraft record number
	HexIdent             VARCHAR(10),  --Field 5: Aircraft Mode S hexadecimal code
	FlightID             VARCHAR(10),  --Field 6: Database Flight record number
	DateMessageGenerated VARCHAR(10),  --Field 7:
	TimeMessageGenerated VARCHAR(10),  --Field 8:
	DateMessageLogged    VARCHAR(10),  --Field 9:
	TimeMessageLogged    VARCHAR(10),  --Field 10:
	Callsign             VARCHAR(10),  --Field 11: An eight digit flight ID - can be flight number or registration (or even nothing).
	Altitude             INT,          --Field 12: Mode C altitude. Height relative to 1013.2mb (Flight Level). Not height AMSL..
	GroundSpeed          SMALLINT,     --Field 13: Speed over ground (not indicated airspeed)
	Track                SMALLINT,     --Field 14: Track of aircraft (not heading). Derived from the velocity E/W and velocity N/S
	Latitude             NUMERIC(5,3),  --Field 15: North and East positive. South and West negative. (e.g., 36.373)
	Longitude            NUMERIC(6,3),  --Field 16: North and East positive. South and West negative. (e.g., -120.397)
	VerticalRate         SMALLINT,           --Field 17: 64ft resolution
	Squawk               VARCHAR(10),   --Field 18: Assigned Mode Person squawk code.
	Alert                VARCHAR(10),   --Field 19: (Squawk change)	 Flag to indicate squawk has changed.
	Emergency            VARCHAR(10),   --Field 20: Flag to indicate emergency code has been set
	SPI                  VARCHAR(10),   --Field 21: (Ident Flag to indicate transponder Ident has been activated.
	IsOnGround           VARCHAR(10),   --Field 22: Flag to indicate ground squat switch is active
    TS_Local             TIMESTAMP with time zone DEFAULT now(),                 --Additional field. Local timestamp when the message was received.
    TS_UTC               TIMESTAMP with time zone DEFAULT timezone('UTC', now()) --Additional field. UTC timestamp when the message was received.
);
