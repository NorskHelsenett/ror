import { TranslateModule } from '@ngx-translate/core';
import { NgModule } from '@angular/core';
import { CommonModule, NgOptimizedImage } from '@angular/common';
import { ClustersComponent } from './clusters.component';
import { clustersPages } from './pages';
import { clusterComponents } from './components';
import { SharedModule } from '../shared/shared.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MultiSelectModule } from 'primeng/multiselect';
import { TableModule } from 'primeng/table';
import { TooltipModule } from 'primeng/tooltip';
import { AutoCompleteModule } from 'primeng/autocomplete';
import { SelectButtonModule } from 'primeng/selectbutton';
import { TabViewModule } from 'primeng/tabview';
import { DropdownModule } from 'primeng/dropdown';
import { ChipsModule } from 'primeng/chips';
import { InputNumberModule } from 'primeng/inputnumber';
import { CardModule } from 'primeng/card';
import { BadgeModule } from 'primeng/badge';
import { TabMenuModule } from 'primeng/tabmenu';
import { ToastModule } from 'primeng/toast';
import { ConfirmDialogModule } from 'primeng/confirmdialog';
import { ConfirmationService } from 'primeng/api';
import { TagModule } from 'primeng/tag';
import { OrganizationChartModule } from 'primeng/organizationchart';
import { InputTextareaModule } from 'primeng/inputtextarea';
import { RippleModule } from 'primeng/ripple';
import { DialogModule } from 'primeng/dialog';
import { ClusterVulnerabilityComponent } from './components/cluster-vulnerability/cluster-vulnerability.component';

import { ProviderComponent } from '../shared/components/provider/provider.component';
import { ButtonModule } from 'primeng/button';
import { HighlightModule } from 'ngx-highlightjs';
import { HighlightLineNumbers } from 'ngx-highlightjs/line-numbers';
import { clusterServices } from './services';
import { SidebarModule } from 'primeng/sidebar';
import { ResourceV2DetailsComponent } from '../resourcesv2/components/resource-v2-details/resource-v2-details.component';
import { IngressComponent } from './pages/ingress/ingress.component';
import { ClusterIngressTableComponent } from './components/cluster-ingress-table/cluster-ingress-table.component';
import { ClusterResourceTableComponent } from './components/cluster-resource-table/cluster-resource-table.component';

@NgModule({
  declarations: [ClustersComponent, ...clustersPages, ...clusterComponents],
  imports: [
    CommonModule,
    TranslateModule,
    SharedModule,
    FormsModule,
    ReactiveFormsModule,
    AutoCompleteModule,
    TooltipModule,
    SelectButtonModule,
    ButtonModule,
    DropdownModule,
    ChipsModule,
    InputNumberModule,
    TableModule,
    MultiSelectModule,
    CardModule,
    TabViewModule,
    TabMenuModule,
    BadgeModule,
    ToastModule,
    ConfirmDialogModule,
    TagModule,
    OrganizationChartModule,
    InputTextareaModule,
    RippleModule,
    NgOptimizedImage,
    ProviderComponent,
    HighlightModule,
    HighlightLineNumbers,
    DialogModule,
    ClusterVulnerabilityComponent,

    SidebarModule,
    ResourceV2DetailsComponent,
    IngressComponent,
    ClusterIngressTableComponent,
    ClusterResourceTableComponent,
  ],
  exports: [ClustersComponent, ...clustersPages, ...clusterComponents],
  providers: [ConfirmationService, ...clusterServices],
})
export class ClustersModule {}
