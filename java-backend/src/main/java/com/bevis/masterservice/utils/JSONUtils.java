package com.bevis.masterservice.utils;

import org.apache.commons.lang3.StringUtils;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;

public class    JSONUtils {

    /**
     * JSON to Class
     *
     * @param message JSON String
     * @param clazz   class
     * @param <T>     Object
     * @return return Object
     */
    public static <T> T toObject(String message, Class<T> clazz) throws JsonProcessingException {
        ObjectMapper mapper = new ObjectMapper();
        mapper.configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false);
        return mapper.readValue(message, clazz);
    }

    /**
     * JSON to class reference
     *
     * @param message  JSON string
     * @param clazzRef type refer
     * @param <T>      Object eg List, Map
     * @return Object
     */
    public static <T> T toObject(String message, TypeReference<T> clazzRef) throws JsonProcessingException {
        ObjectMapper mapper = new ObjectMapper();
        mapper.configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false);
        return mapper.readValue(message, clazzRef);
    }

    /**
     * write Object to JSON String
     *
     * @param obj object
     * @return new string
     */
    public static <T> String toJSON(T obj) throws JsonProcessingException {
        ObjectMapper mapper = new ObjectMapper();
        return mapper.writeValueAsString(obj);
    }

    /**
     * compare JSON
     *
     * @param obj1 Object 1
     * @param obj2 Object 2
     * @return result of compare
     */
    public static <T, E> boolean compareJSON(T obj1, E obj2) throws JsonProcessingException {
        String var1 = toJSON(obj1);
        String var2 = toJSON(obj2);
        return StringUtils.equals(var1, var2);
    }
}
