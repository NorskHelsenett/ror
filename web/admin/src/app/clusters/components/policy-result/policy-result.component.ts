import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-policy-result',
  templateUrl: './policy-result.component.html',
  styleUrls: ['./policy-result.component.scss'],
})
export class PolicyResultComponent {
  @Input() result: string;
}
