package model;

import java.nio.file.NoSuchFileException;
import java.util.List;

public class IPDatabase {

    private static IPDatabase instance;
    private List<IPEntry> ipData;

    private static long IPtoDecimal(String ip) {
        //https://stackoverflow.com/questions/11548273/convert-an-ip-address-to-its-decimal-equivalent-in-java
        var addrArray = ip.split("\\.");

        long num = 0;
        for (int i = 0; i < addrArray.length; i++) {
            int power = 3 - i;
            num += ((Integer.parseInt(addrArray[i]) % 256 * Math.pow(256, power)));
        }

        System.out.println(num);

        return num;
    }

    public static void initialize(List<IPEntry> data) {
        IPDatabase.instance = new IPDatabase(data);
    }

    public static IPDatabase getInstance() {
        if (IPDatabase.instance == null)
            throw new RuntimeException("Database was not initialized!");

        return IPDatabase.instance;
    }

    private IPDatabase(List<IPEntry> data) {
        this.ipData = data;
    }

    public String getCountry(String ip) throws NumberFormatException, NoSuchFieldException {
        if (ip.startsWith("127.0.0"))
            throw new NumberFormatException("you have got to be kidding me");

        long ipDec = IPDatabase.IPtoDecimal(ip);
        long startTime = System.currentTimeMillis();

        for (var entry : ipData) {
            if (entry.getRange().isInRange(ipDec)) {
                System.out.println("Search time: " + (System.currentTimeMillis() - startTime) + "ms");
                return entry.getCountry();
            }
        }

        throw new NoSuchFieldException(ip+" not found in database");
    }
}
