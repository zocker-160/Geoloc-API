package model;

public class Geolocation {

    private final float latitude;
    private final float longitude;

    public Geolocation(float latitude, float longitude) {
        this.latitude = latitude;
        this.longitude = longitude;
    }

    public String toStringTuple() {
        return latitude+","+longitude;
    }
}
