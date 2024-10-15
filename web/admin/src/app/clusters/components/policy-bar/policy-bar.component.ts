import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-policy-bar',
  templateUrl: './policy-bar.component.html',
  styleUrls: ['./policy-bar.component.scss'],
})
export class PolicyBarComponent {
  @Input() failed: number;
  @Input() passed: number;
  @Input() error: number;
  @Input() warning: number;
  @Input() skipped: number;
  @Input() total: number;
}
