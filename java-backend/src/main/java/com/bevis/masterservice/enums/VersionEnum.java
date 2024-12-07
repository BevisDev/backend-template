package com.bevis.masterservice.enums;

import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.experimental.FieldDefaults;

@Getter
@AllArgsConstructor
@NoArgsConstructor
@FieldDefaults(level = AccessLevel.PRIVATE)
public enum VersionEnum {
    VERSION_1(1),
    VERSION_2(2),
    VERSION_3(3),
    VERSION_4(4),
    VERSION_5(5),
    ;

    int version;
}
