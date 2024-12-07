package com.bevis.masterservice.enums;

import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.experimental.FieldDefaults;

@Getter
@AllArgsConstructor
@FieldDefaults(level = AccessLevel.PRIVATE, makeFinal = true)
public enum ErrorCode {
    // =========================================================== //
    // ------------------ API Gateway Reject Code 4xx
    // =========================================================== //
    REJECT_METHOD_NOT_ALLOW("RJ1000", "Invalid Method."),
    REJECT_INVALID_MEDIA_TYPE("RJ1001", "Invalid Media Type."),
    REJECT_NOT_FOUND("RJ1002", "Not Found."),
    REJECT_INVALID_CREDENTIALS("RJ1003", "Security credentials is incorrect."),
    REJECT_NOT_AUTHORIZED_ACCESS("RJ1004", "User is not authorized to access this API."),
    REJECT_INVALID_API_KEY("RJ1005", "Invalid API Key."),

    // API Code 4xx -- Client error
    INVALID_REQUEST("C1000", "Invalid Request."),
    INVALID_EMAIL("C1001", "Invalid Email."),

    // API Code 5xx - Server error
    SERVER_ERROR("A1000", "Internal Server Error."),
    SERVER_DOWN("A1001", "Server is down or under maintenance."),
    SERVER_TIMEOUT("A1002", "Server gateway is timed out"),
    ;

    String code;
    String message;
}
