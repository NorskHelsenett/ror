import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { createClusterPages } from './pages';
import { createClusterComponents } from './components';
import { CreateClusterRoutingModule } from './create-cluster.routing';

import { TranslateModule } from '@ngx-translate/core';

import { AutoFocusModule } from 'primeng/autofocus';
import { StepsModule } from 'primeng/steps';
import { DropdownModule } from 'primeng/dropdown';
import { ChipsModule } from 'primeng/chips';
import { InputNumberModule } from 'primeng/inputnumber';
import { RadioButtonModule } from 'primeng/radiobutton';
import { TagModule } from 'primeng/tag';
import { SharedModule } from 'primeng/api';
import { InputDropdownComponent } from '../../shared/components/input-dropdown/input-dropdown.component';
import { ProviderComponent } from '../../shared/components/provider/provider.component';

@NgModule({
  declarations: [...createClusterPages, ...createClusterComponents],
  imports: [
    CommonModule,
    CreateClusterRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    StepsModule,
    TranslateModule,
    DropdownModule,
    ChipsModule,
    InputNumberModule,
    RadioButtonModule,
    SharedModule,
    TagModule,
    AutoFocusModule,
    ProviderComponent,
    InputDropdownComponent,
  ],
})
export class CreateClusterModule {}
