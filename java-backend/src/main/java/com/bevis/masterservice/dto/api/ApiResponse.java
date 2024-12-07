package com.bevis.masterservice.dto.api;

import java.io.Serial;
import java.io.Serializable;
import java.util.Date;
import java.util.List;

import org.apache.commons.lang3.builder.ToStringBuilder;
import org.apache.commons.lang3.builder.ToStringStyle;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;

import lombok.*;

@Getter
@Setter
@Builder
@JsonInclude(JsonInclude.Include.NON_NULL)
@JsonIgnoreProperties(ignoreUnknown = true)
public class ApiResponse<T> implements Serializable {
    @Serial
    private static final long serialVersionUID = -5952913922941280904L;

    @Builder.Default
    boolean success = true;

    T data;

    String requestId;

    @Builder.Default
    Date responseAt = new Date();

    List<Error> errors;

    public static class Error {
        String errorCode;
        String message;
    }

    @Override
    public String toString() {
        return ToStringBuilder.reflectionToString(this, ToStringStyle.JSON_STYLE);
    }
}
