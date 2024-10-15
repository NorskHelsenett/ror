import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CreateRoutingModule } from './create.routing';
import { SharedModule } from '../shared/shared.module';
import { InputDropdownComponent } from '../shared/components/input-dropdown/input-dropdown.component';
import { ProviderComponent } from '../shared/components/provider/provider.component';

@NgModule({
  declarations: [],
  imports: [CommonModule, CreateRoutingModule, SharedModule, InputDropdownComponent, ProviderComponent],
})
export class CreateModule {}
