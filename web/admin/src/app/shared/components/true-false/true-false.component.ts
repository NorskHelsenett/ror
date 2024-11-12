import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-true-false',
  templateUrl: './true-false.component.html',
  styleUrls: ['./true-false.component.scss'],
})
export class TrueFalseComponent {
  @Input()
  trueOrFalse: boolean;

  @Input()
  labelTrue: string;

  @Input()
  labelFalse: string;
}
