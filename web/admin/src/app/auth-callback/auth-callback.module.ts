import { TranslateModule } from '@ngx-translate/core';
import { NgModule } from '@angular/core';
import { CommonModule, NgOptimizedImage } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';
import { AuthCallbackComponent } from './auth-callback.component';
import { SharedModule } from '../shared/shared.module';

const routes: Routes = [
  {
    path: '',
    component: AuthCallbackComponent,
  },
];

@NgModule({
  declarations: [AuthCallbackComponent],
  imports: [CommonModule, RouterModule.forChild(routes), SharedModule, TranslateModule, NgOptimizedImage],
})
export class AuthCallbackModule {}
