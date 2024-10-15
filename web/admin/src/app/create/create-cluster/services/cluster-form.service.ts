import { Injectable } from '@angular/core';
import { FormGroup } from '@angular/forms';

@Injectable({
  providedIn: 'root',
})
export class ClusterFormService {
  clusterForm: FormGroup | undefined;

  selectedProvider: any | undefined;

  getNodePoolSum(): number {
    let sum = 0;
    let capasity = this.clusterForm.get('capasity').value;
    if (!capasity || capasity.length === 0) {
      return sum;
    }

    capasity.forEach((x: any) => {
      sum = sum + x.price * x.count;
    });

    return sum;
  }
}
