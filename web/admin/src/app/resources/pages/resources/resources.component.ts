import { ChangeDetectionStrategy, Component } from '@angular/core';
import { ResourceTableComponent } from '../../components/resource-table/resource-table.component';
import { TranslateModule } from '@ngx-translate/core';

@Component({
  selector: 'app-resources',
  standalone: true,
  imports: [TranslateModule, ResourceTableComponent],
  providers: [],
  templateUrl: './resources.component.html',
  styleUrl: './resources.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ResourcesComponent {}
