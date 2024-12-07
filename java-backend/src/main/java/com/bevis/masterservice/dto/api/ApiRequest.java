package com.bevis.masterservice.dto.api;

import java.io.Serial;
import java.io.Serializable;
import java.util.Date;

import org.apache.commons.lang3.builder.ToStringBuilder;
import org.apache.commons.lang3.builder.ToStringStyle;

import com.bevis.masterservice.utils.CommonUtils;
import com.bevis.masterservice.utils.ValidateUtils;
import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonInclude;

import lombok.AccessLevel;
import lombok.Getter;
import lombok.Setter;
import lombok.experimental.FieldDefaults;

@Getter
@Setter
@JsonInclude(JsonInclude.Include.NON_NULL)
@JsonIgnoreProperties(ignoreUnknown = true)
@FieldDefaults(level = AccessLevel.PROTECTED)
public class ApiRequest implements Serializable {
    @Serial
    private static final long serialVersionUID = 2905646205909847233L;

    String requestId;

    Date requestAt;

    public ApiRequest() {
        if (ValidateUtils.isNullOrEmpty(requestId)) {
            this.requestId = CommonUtils.randomUUID();
        }
        if (ValidateUtils.isNullOrEmpty(requestAt)) {
            this.requestAt = new Date();
        }
    }

    @Override
    public String toString() {
        return ToStringBuilder.reflectionToString(this, ToStringStyle.JSON_STYLE);
    }
}
