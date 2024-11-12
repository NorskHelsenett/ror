import { AbstractControl, AsyncValidatorFn, ValidationErrors } from '@angular/forms';
import { Observable, of, catchError } from 'rxjs';
import { map } from 'rxjs/operators';
import { ClustersService } from '../../core/services/clusters.service';

export default class ClusterNameValidator {
  static validName(clusterService: ClustersService): AsyncValidatorFn {
    return (control: AbstractControl): Observable<ValidationErrors | null> => {
      const id = control?.value;
      if (!id) {
        return of({
          error: 'error',
        });
      }
      return clusterService.exists(id).pipe(
        map((clusterExisting: boolean) => {
          if (clusterExisting) {
            return {
              nonUniqueUsername: true,
            };
          } else {
            return null;
          }
        }),
        catchError(() => null),
      );
    };
  }
}
