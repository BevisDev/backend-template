package com.bevis.masterservice.utils;

import java.util.UUID;

public class CommonUtils {

    /**
     * random UUID
     * eg 7df55507-9995-4c26-8ea7-0360c066b61e
     *
     * @return random UUID
     */
    public static String randomUUID() {
        return UUID.randomUUID().toString();
    }
}
