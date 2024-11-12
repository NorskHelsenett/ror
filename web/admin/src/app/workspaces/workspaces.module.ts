import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { TranslateModule } from '@ngx-translate/core';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { WorkspacesComponent } from './workspaces.component';
import { WorkspacesRoutingModule } from './workspaces.routing';
import { workspacesPages } from './pages';
import { workspaceComponents } from './components';
import { SharedModule } from '../shared/shared.module';
import { ButtonModule } from 'primeng/button';
import { DropdownModule } from 'primeng/dropdown';
import { InputTextModule } from 'primeng/inputtext';
import { TableModule } from 'primeng/table';

@NgModule({
  declarations: [WorkspacesComponent, ...workspacesPages, workspaceComponents],
  imports: [
    CommonModule,
    WorkspacesRoutingModule,
    TranslateModule,
    SharedModule,
    FormsModule,
    ReactiveFormsModule,
    ButtonModule,
    DropdownModule,
    InputTextModule,
    TableModule,
  ],
})
export class WorkspacesModule {}
