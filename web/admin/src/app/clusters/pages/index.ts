import { ClusterDetailsComponent } from './cluster-details/cluster-details.component';
import { ClusterDetailsEditComponent } from './cluster-details-edit/cluster-details-edit.component';
import { ClusterMetadataPageComponent } from './cluster-metadata-page/cluster-metadata-page.component';
import { ClusterNewComponent } from './cluster-new/cluster-new.component';
import { IngressDetailsComponent } from './ingress-details/ingress-details.component';

export * from './cluster-details/cluster-details.component';
export * from './cluster-details-edit/cluster-details-edit.component';
export * from './cluster-new/cluster-new.component';
export * from './cluster-metadata-page/cluster-metadata-page.component';
export * from './ingress-details/ingress-details.component';

export const clustersPages = [
  ClusterDetailsComponent,
  ClusterDetailsEditComponent,
  ClusterMetadataPageComponent,
  ClusterNewComponent,
  IngressDetailsComponent,
];
