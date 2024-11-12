import { ClipboardModule } from 'ngx-clipboard';
import { TranslateModule } from '@ngx-translate/core';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { UserprofileComponent } from './userprofile.component';
import { RouterModule, Routes } from '@angular/router';
import { SharedModule } from '../shared/shared.module';
import { ApikeyModule } from '../apikey/apikey.module';
import { userprofilePages } from './pages';
import { userprofileComponents } from './components';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { TabViewModule } from 'primeng/tabview';
import { ToggleButtonModule } from 'primeng/togglebutton';
import { CalendarModule } from 'primeng/calendar';
import { ConfirmDialogModule } from 'primeng/confirmdialog';
import { TableModule } from 'primeng/table';
import { InputTextModule } from 'primeng/inputtext';

const routes: Routes = [
  {
    path: '',
    component: UserprofileComponent,
  },
];

@NgModule({
  declarations: [UserprofileComponent, ...userprofileComponents, ...userprofilePages],
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    RouterModule.forChild(routes),
    TranslateModule,
    ClipboardModule,
    SharedModule,
    ApikeyModule,
    TabViewModule,
    ToggleButtonModule,
    CalendarModule,
    ConfirmDialogModule,
    TableModule,
    InputTextModule,
  ],
})
export class UserprofileModule {}
