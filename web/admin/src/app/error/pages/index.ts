import { ForbiddenComponent } from './forbidden/forbidden.component';
import { NotFoundComponent } from './not-found/not-found.component';
import { ServerErrorComponent } from './server-error/server-error.component';
import { UnauthorizedComponent } from './unauthorized/unauthorized.component';

export * from './forbidden/forbidden.component';
export * from './not-found/not-found.component';
export * from './server-error/server-error.component';
export * from './unauthorized/unauthorized.component';

export const errorPages = [ForbiddenComponent, NotFoundComponent, ServerErrorComponent, UnauthorizedComponent];
