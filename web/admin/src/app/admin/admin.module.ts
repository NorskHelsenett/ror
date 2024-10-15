import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { TranslateModule } from '@ngx-translate/core';
import { SharedModule } from '../shared/shared.module';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { adminPages } from './pages';
import { AdminRoutingModule } from './admin.routing';
import { adminDatacentersPages } from './pages/datacenters/pages';
import { adminPricesPages } from './pages/prices/pages';

import { InputTextModule } from 'primeng/inputtext';
import { InputTextareaModule } from 'primeng/inputtextarea';
import { InputNumberModule } from 'primeng/inputnumber';
import { InputSwitchModule } from 'primeng/inputswitch';
import { CalendarModule } from 'primeng/calendar';
import { ButtonModule } from 'primeng/button';
import { TableModule } from 'primeng/table';
import { FieldsetModule } from 'primeng/fieldset';
import { CardModule } from 'primeng/card';
import { MultiSelectModule } from 'primeng/multiselect';
import { ToggleButtonModule } from 'primeng/togglebutton';
import { ConfirmDialogModule } from 'primeng/confirmdialog';
import { ConfirmationService } from 'primeng/api';
import { DropdownModule } from 'primeng/dropdown';
import { ChipsModule } from 'primeng/chips';
import { TagModule } from 'primeng/tag';
import { TabViewModule } from 'primeng/tabview';

import { adminProjectsPages } from './pages/projects/pages';
import { ToastModule } from 'primeng/toast';
import { DatacentersModule } from '../datacenters/datacenters.module';
import { ConfigurationModule } from '../configuration/configuration.module';
import { ApikeyModule } from '../apikey/apikey.module';
import { AclModule } from '../acl/acl.module';
import { RippleModule } from 'primeng/ripple';
import { TooltipModule } from 'primeng/tooltip';
import { IconFieldModule } from 'primeng/iconfield';
import { InputIconModule } from 'primeng/inputicon';

@NgModule({
  declarations: [...adminPages, ...adminDatacentersPages, ...adminPricesPages, ...adminProjectsPages],
  imports: [
    AdminRoutingModule,
    CommonModule,
    SharedModule,
    TranslateModule,
    FormsModule,
    ReactiveFormsModule,
    InputTextModule,
    InputTextareaModule,
    InputNumberModule,
    InputSwitchModule,
    CalendarModule,
    ButtonModule,
    TableModule,
    FieldsetModule,
    CardModule,
    MultiSelectModule,
    ToggleButtonModule,
    ConfirmDialogModule,
    ToastModule,
    DropdownModule,
    ChipsModule,
    TagModule,
    DatacentersModule,
    ConfigurationModule,
    ApikeyModule,
    AclModule,
    TabViewModule,
    RippleModule,
    TooltipModule,
    IconFieldModule,
    InputIconModule,
  ],
  providers: [ConfirmationService],
})
export class AdminModule {}
