package utils;

import java.text.SimpleDateFormat;
import java.util.Date;

public class Utils {

    public static String getTimecode() {
        var sdf = new SimpleDateFormat("dd.MM.yyyy HH:mm:ss.SSS");
        return sdf.format(new Date());
    }

}