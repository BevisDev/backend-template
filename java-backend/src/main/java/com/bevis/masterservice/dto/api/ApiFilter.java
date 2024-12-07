package com.bevis.masterservice.dto.api;

import java.io.Serial;
import java.io.Serializable;

import com.bevis.masterservice.consts.Const;
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
public class ApiFilter extends ApiRequest implements Serializable {
    @Serial
    private static final long serialVersionUID = 3122949941812274044L;

    int page;
    int size;
    String sort;
    String search;

    public ApiFilter() {
        this.sort = Const.DESC_SORT;
    }
}
