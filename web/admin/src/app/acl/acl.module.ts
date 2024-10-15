import { TagModule } from 'primeng/tag';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { aclPages } from './pages';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { SharedModule } from '../shared/shared.module';
import { TranslateModule } from '@ngx-translate/core';
import { TableModule } from 'primeng/table';
import { ToastModule } from 'primeng/toast';
import { ConfirmDialogModule } from 'primeng/confirmdialog';
import { RouterModule } from '@angular/router';
import { DropdownModule } from 'primeng/dropdown';
import { MultiSelectModule } from 'primeng/multiselect';
import { aclComponents } from './component';
import { InputSwitchModule } from 'primeng/inputswitch';

import { InputDropdownComponent } from '../shared/components/input-dropdown/input-dropdown.component';

@NgModule({
  declarations: [...aclPages, ...aclComponents],
  imports: [
    CommonModule,
    RouterModule,
    FormsModule,
    ReactiveFormsModule,
    SharedModule,
    TranslateModule,
    TableModule,
    ToastModule,
    ConfirmDialogModule,
    TagModule,
    DropdownModule,
    MultiSelectModule,
    InputSwitchModule,
    InputDropdownComponent,
  ],
  exports: [...aclPages],
})
export class AclModule {}
