import { Component, Input } from '@angular/core';
import { FormGroup, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { TranslateModule } from '@ngx-translate/core';
import { DropdownModule } from 'primeng/dropdown';
import { CheckboxModule } from 'primeng/checkbox';

@Component({
  selector: 'app-input-dropdown',
  standalone: true,
  imports: [TranslateModule, DropdownModule, ReactiveFormsModule, CheckboxModule, FormsModule],
  templateUrl: './input-dropdown.component.html',
  styleUrl: './input-dropdown.component.scss',
})
export class InputDropdownComponent {
  @Input() formGroup: FormGroup | undefined;
  @Input() controlName: string | undefined;

  @Input() couldOverride: boolean;
  @Input() groupSupport: boolean;
  @Input() showClear: boolean;
  @Input() clearTextOnOverride: boolean;

  @Input() options: any[] | undefined;
  @Input() optionLabel: string | undefined;
  @Input() optionValue: string | undefined;

  @Input() filterBy: string | undefined;

  @Input() placeholder: string | undefined;
  @Input() emptyListMessage: string | undefined;

  override: boolean;

  constructor() {}

  onOverride() {
    if (this.clearTextOnOverride) {
      this.formGroup?.get(this.controlName)?.setValue('');
      this.formGroup?.get(this.controlName)?.markAsPristine();
    }
  }
}
