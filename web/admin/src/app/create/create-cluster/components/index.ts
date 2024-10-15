import { LocationStepComponent } from './location-step/location-step.component';
import { MetadataStepComponent } from './metadata-step/metadata-step.component';
import { ResourcesStepComponent } from './resources-step/resources-step.component';
import { SummaryStepComponent } from './summary-step/summary-step.component';
import { SummaryComponent } from './summary/summary.component';

export * from './location-step/location-step.component';
export * from './metadata-step/metadata-step.component';
export * from './resources-step/resources-step.component';
export * from './summary-step/summary-step.component';
export * from './summary/summary.component';

export const createClusterComponents = [LocationStepComponent, MetadataStepComponent, ResourcesStepComponent, SummaryComponent, SummaryStepComponent];
