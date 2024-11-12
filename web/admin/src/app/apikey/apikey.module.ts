import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

import { ButtonModule } from 'primeng/button';
import { TableModule } from 'primeng/table';
import { ConfirmDialogModule } from 'primeng/confirmdialog';

import { apikeysPages } from './pages';
import { apikeysComponents } from './components';
import { SharedModule } from '../shared/shared.module';
import { TranslateModule } from '@ngx-translate/core';
import { ConfirmationService } from 'primeng/api';
import { ToastModule } from 'primeng/toast';
import { DropdownModule } from 'primeng/dropdown';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CalendarModule } from 'primeng/calendar';
import { InputTextModule } from 'primeng/inputtext';

@NgModule({
  declarations: [...apikeysPages, ...apikeysComponents],
  exports: [...apikeysPages, ...apikeysComponents],
  imports: [
    CommonModule,
    RouterModule,
    FormsModule,
    ReactiveFormsModule,
    SharedModule,
    TranslateModule,
    TableModule,
    ButtonModule,
    ConfirmDialogModule,
    ToastModule,
    DropdownModule,
    FormsModule,
    CalendarModule,
    InputTextModule,
  ],
  providers: [ConfirmationService],
})
export class ApikeyModule {}
