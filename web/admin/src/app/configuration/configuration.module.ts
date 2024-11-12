import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { SharedModule } from '../shared/shared.module';
import { TranslateModule } from '@ngx-translate/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { TableModule } from 'primeng/table';
import { ButtonModule } from 'primeng/button';
import { configPages } from './pages';
import { configComponents } from './components';
import { RouterModule } from '@angular/router';
import { InputSwitchModule } from 'primeng/inputswitch';
import { DropdownModule } from 'primeng/dropdown';
import { ConfirmDialogModule } from 'primeng/confirmdialog';
import { ToastModule } from 'primeng/toast';
import { HighlightModule } from 'ngx-highlightjs';
import { HighlightLineNumbers } from 'ngx-highlightjs/line-numbers';
import { RippleModule } from 'primeng/ripple';
import { InputTextModule } from 'primeng/inputtext';

@NgModule({
  declarations: [...configPages, ...configComponents],
  imports: [
    CommonModule,
    RouterModule,
    SharedModule,
    TranslateModule,
    TableModule,
    ButtonModule,
    InputSwitchModule,
    DropdownModule,
    FormsModule,
    ReactiveFormsModule,
    ConfirmDialogModule,
    ToastModule,
    HighlightModule,
    HighlightLineNumbers,
    RippleModule,
    InputTextModule,
  ],
})
export class ConfigurationModule {}
