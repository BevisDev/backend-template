package com.bevis.masterservice.utils;

import java.text.DateFormat;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

public class DateTimeUtils {
    public static final int MILLIS_PER_SECOND = 1000;
    public static final int MILLIS_PER_MINUTE = 60 * 1000; // 1 minute
    public static final int MILLIS_PER_HOUR = 60 * 60 * 1000; // 1 hour
    public static final int MILLIS_PER_DAY = 24 * 60 * 60 * 1000; // 1 day

    public static Date convertStringToDate(String str, String pattern) throws ParseException {
        DateFormat formatter = new SimpleDateFormat(pattern);
        return formatter.parse(str);
    }

    public static String convertDateToString(Date date, String pattern) {
        DateFormat dateFormat = new SimpleDateFormat(pattern);
        return dateFormat.format(date);
    }
}
