package com.bevis.masterservice.enums;

import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.experimental.FieldDefaults;

@AllArgsConstructor
@NoArgsConstructor
@FieldDefaults(level = AccessLevel.PRIVATE)
@Getter
public enum DateTimeEnum {

    // =========================================================== //
    // ------------------ Stored in the ISO 8601 format
    // =========================================================== //
    YYYY_MM_DD("yyyy-MM-dd"),
    MM_DD_YYYY("MM/dd/yyyy"),
    YYYY_MM_DD_HH_MM_SS("yyyy-MM-dd HH:mm:ss"),
    YYYY_MM_DD_T_HH_MM_SSZ("yyyy-MM-dd'T'HH:mm:ssZ"),
    YYYY_MM_DD_T_HH_MM_SS("yyyy-MM-dd'T'HH:mm:ss"),
    YYYY_MM_DD_T_HH_MM_SS_SSSZ("yyyy-MM-dd'T'HH:mm:ss:SSSZ"),

    // =========================================================== //
    // ------------------ Date Time format
    // =========================================================== //
    DD_MMM_YYYY_HH_MM_SS("dd-MMM-yyyy HH:mm:ss"),
    DD_MMM_YYYY_DASH("dd-MMM-yyyy"),
    DD_MM_YY_SLASHED("dd/MM/yy"),
    MM_YYYY_SLASHED("MM/yyyy"),
    MMM_YYYY_SLASHED("MMM/yyyy"),
    YYYYMMDD("yyyyMMdd"),
    HH_MM("HH:mm"),
    ;

    String pattern;
}
