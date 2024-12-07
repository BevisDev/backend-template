package com.bevis.masterservice.utils;

import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.regex.Pattern;

public class ValidateUtils {

    /**
     * Validate check null or empty
     *
     * @param object input
     * @return true if null or empty
     */
    public static <T> boolean isNullOrEmpty(T object) {
        if (null == object) {
            return true;
        }

        if (object instanceof String) {
            return ((String) object).isEmpty();
        }

        if (object instanceof List<?>) {
            return ((List<?>) object).isEmpty();
        }

        if (object instanceof Set<?>) {
            return ((Set<?>) object).isEmpty();
        }

        if (object instanceof Map<?, ?>) {
            return ((Map<?, ?>) object).isEmpty();
        }

        if (object instanceof String[]) {
            return ((String[]) object).length == 0;
        }

        return false;
    }

    public static <T> boolean isNotNullOrEmpty(T object) {
        return !isNullOrEmpty(object);
    }

    /**
     * validate length of string
     *
     * @param str    from request
     * @param length length of string
     * @return result validation
     */
    public static boolean isValidLengthStr(String str, int length) {
        return isNotNullOrEmpty(str) && str.length() <= length;
    }

    /**
     * validate pattern of string
     *
     * @param str        from request
     * @param patternStr pattern validation
     * @return result validation
     */
    public static boolean isValidPattern(String str, String patternStr) {
        Pattern pattern = Pattern.compile(patternStr);
        return pattern.matcher(str).matches();
    }
}
