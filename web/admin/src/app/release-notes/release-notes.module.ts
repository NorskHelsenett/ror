import { NgModule } from '@angular/core';
import { CommonModule, NgOptimizedImage } from '@angular/common';
import { ReleaseNotesRoutingModule } from './release-notes.routing';
import { ReleaseNotesComponent } from './release-notes.component';
import { TranslateModule } from '@ngx-translate/core';

@NgModule({
  declarations: [ReleaseNotesComponent],
  imports: [CommonModule, ReleaseNotesRoutingModule, TranslateModule, NgOptimizedImage],
})
export class ReleaseNotesModule {}
