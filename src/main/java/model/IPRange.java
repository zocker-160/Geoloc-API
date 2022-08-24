package model;

public class IPRange {

    private final long startIP;
    private final long endIP;

    public IPRange(long startIP, long endIP) {
        this.startIP = startIP;
        this.endIP = endIP;
    }

    public boolean isInRange(long ip) {
        return startIP <= ip && ip <= endIP;
    }
}
