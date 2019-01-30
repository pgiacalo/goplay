
/*
 * Note that the Dump1090 program sends ADS-B messages in SBS1 (BaseStation) format
 * which is a fixed position, comma-separated string, as follows.
 *
 * Position     Field	            Description
 * --------     -----               -----------
 *  0           message_type	    See MessageType.
 *  1           transmission_type	See TransmissionType.
 *  2           session_id	        String. Database session record number.
 *  3           aircraft_id	        String. Database aircraft record number.
 *  4           hex_ident	        String. 24-bit ICACO ID, in hex. (e.g., 3C64EC)
 *  5           flight_id	        String. Database flight record number.
 *  6           generated_date	    String. Date the message was generated.
 *  7           generated_time	    String. Time the message was generated.
 *  8           logged_date	        String. Date the message was logged.
 *  9           logged_time	        String. Time the message was logged.
 * 10           callsign	        String. Eight character flight ID or callsign.
 * 11           altitude	        Integer. Mode C Altitude relative to 1013 mb (29.92" Hg).
 * 12           ground_speed	    Integer. Speed over ground.
 * 13           track	            Integer. Ground track angle (relative to geographic north)
 * 14           lat	                Float. Latitude in decimal degrees.
 * 15           lon	                Float. Longitude in decimal degrees.
 * 16           vertical_rate	    Integer. Climb rate.
 * 17           squawk	            String. Assigned Mode A squawk code.
 * 18           alert	            Boolean. Flag to indicate that squawk has changed.
 * 19           emergency	        Boolean. Flag to indicate emergency code has been set.
 * 20           spi	                Boolean. Flag to indicate Special Position Indicator has been set.
 * 21           is_on_ground	    Boolean. Flag to indicate ground squat switch is active.
 *
 * There are 8 Transmission "MSG" sub types:
 *  1 ES identification and category (callsign)
 *      MSG,1,145,256,7404F2,11267,2008/11/28,23:48:18.611,2008/11/28,23:53:19.161,RJA1118,,,,,,,,,,,
 *
 *  2 ES surface position message (ground speed, track, lat, lon, vertical rate, is_on_ground)[23 fields]
 *      MSG,2,496,603,400CB6,13168,2008/10/13,12:24:32.414,2008/10/13,12:28:52.074,,,0,76.4,258.3,54.05735,-4.38826,,,,,,0
 *
 *  3 ES airborne position message (altitude, lat, lon, alert, emergency, spi, is_on_ground)[22 fields]
 *      MSG,3,496,211,4CA2D6,10057,2008/11/28,14:53:50.594,2008/11/28,14:58:51.153,,37000,,,51.45735,-1.02826,,,0,0,0,0
 *
 *  4 ES airborne velocity message (ground speed, track, vertical rate)
 *      MSG,4,496,469,4CA767,27854,2010/02/19,17:58:13.039,2010/02/19,17:58:13.368,,,288.6,103.2,,,-832,,,,,
 *
 *  5 Surveillance alt message (altitude, alert, spi, is_on_ground)
 *      MSG,5,496,329,394A65,27868,2010/02/19,17:58:12.644,2010/02/19,17:58:13.368,,10000,,,,,,,0,,0,0
 *
 *  6 Surveillance ID message (altitude, squawk, alert, emergency, spi, is_on_ground)
 *      MSG,6,496,237,4CA215,27864,2010/02/19,17:58:12.846,2010/02/19,17:58:13.368,,33325,,,,,,0271,0,0,0,0
 *
 *  7 Air-to-air message (altitude, is_on_ground)
 *      MSG,7,496,742,51106E,27929,2011/03/06,07:57:36.523,2011/03/06,07:57:37.054,,3775,,,,,,,,,,0
 *
 *  8 All call reply (is_on_ground)
 *      MSG,8,496,194,405F4E,27884,2010/02/19,17:58:13.244,2010/02/19,17:58:13.368,,,,,,,,,,,,0
 *
 *
 *  Only MSG types 2 and 3 have position (latitude and longitude) information.
 *
 */
