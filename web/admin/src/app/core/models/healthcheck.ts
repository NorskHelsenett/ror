import { HealthCheckService } from './healthcheckservice';

export interface HealthCheck {
  services: Array<HealthCheckService>;
}
