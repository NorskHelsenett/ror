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

import { ProviderComponent } from '../shared/components/provider/provider.component';
import { ButtonModule } from 'primeng/button';
import { HighlightModule } from 'ngx-highlightjs';
import { HighlightLineNumbers } from 'ngx-highlightjs/line-numbers';
import { clusterServices } from './services';
import { TypesService } from '../resources/services/types.service';
import { Resourcesv2FilterComponent } from '../resourcesv2/components/resourcesv2-filter/resourcesv2-filter.component';
import { ResourcesV2ListComponent } from '../resourcesv2/components/resources-v2-list/resources-v2-list.component';
import { SidebarModule } from 'primeng/sidebar';
import { ResourceV2DetailsComponent } from '../resourcesv2/components/resource-v2-details/resource-v2-details.component';

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
    Resourcesv2FilterComponent,
    ResourcesV2ListComponent,
    SidebarModule,
    ResourceV2DetailsComponent,
  ],
  exports: [ClustersComponent, ...clustersPages, ...clusterComponents],
  providers: [ConfirmationService, ...clusterServices, TypesService],
})
export class ClustersModule {}
