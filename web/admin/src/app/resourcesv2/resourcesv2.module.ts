import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ResourcesV2ModulesRoutingModule } from './resourcesv2.routing';
import { ResourcesV2QueryService } from './services/resources-v2-query.service';
import { TypesService } from '../resources/services/types.service';

@NgModule({
  declarations: [],
  imports: [CommonModule, ResourcesV2ModulesRoutingModule],
  providers: [ResourcesV2QueryService, TypesService],
})
export class ResourcesV2Module {}
