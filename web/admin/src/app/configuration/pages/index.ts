import { ConfigOperatorconfigCreateUpdateComponent } from './config-operatorconfig-create-update/config-operatorconfig-create-update.component';
import { ConfigTaskCreateUpdateComponent } from './config-task-create-update/config-task-create-update.component';
import { ConfigDesiredversionCreateUpdateComponent } from './config-desiredversion-create-update/config-desiredversion-create-update.component';
import { ConfigurationComponent } from './configuration/configuration.component';

export * from './config-operatorconfig-create-update/config-operatorconfig-create-update.component';
export * from './config-task-create-update/config-task-create-update.component';
export * from './config-desiredversion-create-update/config-desiredversion-create-update.component';
export * from './configuration/configuration.component';

export const configPages = [
  ConfigOperatorconfigCreateUpdateComponent,
  ConfigTaskCreateUpdateComponent,
  ConfigDesiredversionCreateUpdateComponent,
  ConfigurationComponent,
];
