package main

import (
    "fmt"
	"os"
)

func WriteToKmlFile(filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
				
	writeKmlOpenning(f);

	Aircraft[] aircraftArray = tracker.getAircrafts();
	if (aircraftArray.length != 0){				
		writeAircraftKmlPlacemark(aircraftArray, f);
	}

	// GeoFenceUtil util = GeoFenceUtil.getInstance();
	// String geoFenceKml = util.getGeoFenceKml();
	// if (geoFenceKml != null){
	// 	writeGeoFence(geoFenceKml, f);
	// }

	writeKmlClosing(f);
	f.Sync()
		
}

// must be synchronized
func writeAircraftKmlPlacemark(Aircraft[] acs, Writer out) error {
		
	for(int i = 0; i < acs.length; i++) {
		if(acs[i].hasPositions()) {
			writePlacemark(acs[i], out);
			//System.out.println("Server.writeKML(): writing placemark. " + acs[i].getHexId());
		} else {
			System.out.println("Server.writeKML(): this aicraft does not have position data. " + acs[i].getHexId());
		}
	}
	
}

	func writeGeoFence(fenceKml string, out *os.File) error {
		_, err := f.WriteString(fenceKml);
		if err != nil {
			return err
		}
		return nil
	}
	
	func writePlacemark(Aircraft ac, *os.File out) err {
		String callSign = ac.hasCallsign() ? ac.getCallsign() : ac.flightId;
		Aircraft.Position curPos = ac.getCurrentPosition();
		double groundTrack = ac.getGroundTrack();
		
		String iconColor = "Afffffff";
		if (callSign.equals(FMC.getInstance().getDroneId())){
			iconColor = "Afffff00";
		}
		
		out.WriteString("<Placemark>");
		  out.WriteString( "<description>Generated Aircraft Placemark</description>"); 
		  out.WriteString("<name>");out.WriteString(callSign);out.WriteString("</name>");
		  out.WriteString("<Style>");
		    out.WriteString("<IconStyle>");
		    out.WriteString("<scale>0.8</scale>");
		    //Not sure why we need to add 90 degrees to the ground track below.
		    //But it works to get the airplane icon pointed in the track direction.
		    out.WriteString("<heading>" + ((int)groundTrack) + "</heading>");
		      out.WriteString("<Icon>");
		        out.WriteString("<href>http://maps.google.com/mapfiles/kml/pal2/icon56.png</href>");
		      out.WriteString("</Icon>");
		    out.WriteString("</IconStyle>");
		    out.WriteString("<LineStyle>"); 
				out.WriteString("<color>" + iconColor + "</color>");
				out.WriteString("<width>3</width>");
			out.WriteString("</LineStyle>");
		  out.WriteString("</Style>");
		  out.WriteString("<MultiGeometry>");
			  out.WriteString("<Point>");
			    out.WriteString("<extrude>0</extrude>");
			    out.WriteString("<altitudeMode>absolute</altitudeMode>");
			    out.WriteString("<coordinates>");
			    	out.WriteString(String.valueOf(curPos.getLongitude()));
			    	out.WriteString(",");out.WriteString(String.valueOf(curPos.getLatitude()));
			    	out.WriteString(",");out.WriteString(String.valueOf(curPos.getAltitude()));
			    out.WriteString("</coordinates>");
			  out.WriteString("</Point>");
			  out.WriteString("<LineString>");
			    out.WriteString("<extrude>0</extrude>");
			    out.WriteString("<altitudeMode>absolute</altitudeMode>");
			    out.WriteString("<coordinates>");
			    	writeCoordinateListOfPositions(ac, 1, out);
			    out.WriteString("</coordinates>");
			  out.WriteString("</LineString>");
		  out.WriteString("</MultiGeometry>");
		out.WriteString("</Placemark>");
	}
	
	func writeCoordinateListOfPositions(Aircraft ac, int startPosition, *os.File out) err {
		if (ac != null){
			Aircraft.Position[] positions = ac.getPositions();
			for(int i = startPosition; i < positions.length; i++) {
				Aircraft.Position pos = positions[i];
				out.WriteString(String.valueOf(pos.getLongitude()));
				out.WriteString(",");
				out.WriteString(String.valueOf(pos.getLatitude()));
				out.WriteString(",");
				out.WriteString(String.valueOf(pos.getAltitude()));
				if(i + 1 < positions.length) {
					out.WriteString(" ");
				}
			}
		}
	}

	private synchronized void writeKmlOpenning(Writer out) throws IOException {
		out.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\r\n");
		out.WriteString("<kml xmlns=\"http://www.opengis.net/kml/2.2\">\r\n");
//		out.WriteString("<kml xmlns=\"http://earth.google.com/kml/2.0\">\r\n");
		out.WriteString("<Document>");
		out.WriteString("<name>Aircraft</name>\r\n");		
	}

	private synchronized void writeKmlClosing(Writer out) throws IOException {
		out.WriteString("</Document>\r\n");
		out.WriteString("</kml>");
	}

