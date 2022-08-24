package model;

import java.text.ParseException;

public class IPEntry {

    private final IPRange range;

    private final String countryCode;
    private final String country;
    private final String state;
    private final String city;

    private final Geolocation geolocation;

    public static IPEntry parseLine(String line) throws ParseException {
        var entries = line.split(",");

        if (entries.length < 8)
            throw new ParseException("Invalid data input: "+line, 0);

        long ipStart = Long.parseLong(entries[0]);
        long ipEnd = Long.parseLong(entries[1]);
        var ipRange = new IPRange(ipStart, ipEnd);

        String countryCode = entries[2];
        String country = entries[3];
        String state = entries[4];
        String city = entries[entries.length - 3];

        float lattit = Float.parseFloat(entries[entries.length - 2]);
        float longit = Float.parseFloat(entries[entries.length - 1]);
        var geoloc = new Geolocation(lattit, longit);

        return new IPEntry(
                ipRange,
                countryCode, country, state, city,
                geoloc);
    }

    public IPEntry(
            IPRange range,
            String countryCode, String country, String state, String city,
            Geolocation geolocation) {
        this.range = range;
        this.countryCode = countryCode;
        this.country = country;
        this.state = state;
        this.city = city;
        this.geolocation = geolocation;
    }

    public IPRange getRange() {
        return range;
    }

    public String getCountryCode() {
        return countryCode;
    }

    public String getCountry() {
        return country;
    }

    public String getState() {
        return state;
    }

    public String getCity() {
        return city;
    }

    public Geolocation getGeolocation() {
        return geolocation;
    }
}
