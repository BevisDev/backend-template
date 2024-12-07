package com.bevis.masterservice.service.healthcheck;

import org.springframework.stereotype.Component;

import com.bevis.masterservice.config.healthcheck.HealthCheckProps;
import com.bevis.masterservice.consts.Const;
import com.bevis.masterservice.dto.api.ApiResponse;
import com.bevis.masterservice.dto.healthcheck.HealthCheckResp;
import com.bevis.masterservice.enums.StatusEnum;

import lombok.AccessLevel;
import lombok.RequiredArgsConstructor;
import lombok.experimental.FieldDefaults;

@Component
@RequiredArgsConstructor
@FieldDefaults(level = AccessLevel.PRIVATE, makeFinal = true)
public class HealthCheckServiceImpl implements HealthCheckService {

    HealthCheckProps healthCheckProps;

    @Override
    public ApiResponse<HealthCheckResp> getStatusSystem() {
        HealthCheckResp healthCheckResp = HealthCheckResp.builder()
                .appName(healthCheckProps.getName())
                .appVersion(healthCheckProps.getVersion())
                .buildTime(healthCheckProps.getBuildTime())
                .status(StatusEnum.SUCCESS.getCapitalValue())
                .service(Const.HEALTH_CHECK_SYSTEM)
                .build();
        return ApiResponse.<HealthCheckResp>builder().data(healthCheckResp).build();
    }
}
