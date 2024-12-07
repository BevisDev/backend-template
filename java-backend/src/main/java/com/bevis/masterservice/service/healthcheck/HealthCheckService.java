package com.bevis.masterservice.service.healthcheck;

import com.bevis.masterservice.dto.api.ApiResponse;
import com.bevis.masterservice.dto.healthcheck.HealthCheckResp;

public interface HealthCheckService {
    /**
     * Get status system internal
     *
     * @return health check system
     */
    ApiResponse<HealthCheckResp> getStatusSystem();
}
