import * as adminPages from './pages/index';
import * as adminDatacentersPages from './pages/datacenters/pages/index';
import * as adminPricesPages from './pages/prices/pages/index';
import * as adminProjectsPages from './pages/projects/pages/index';
import { Routes, RouterModule } from '@angular/router';
import { AdminReadGuard } from '../core/guards/admin-read.guard';

import * as datacentersPages from '../datacenters/pages';
import * as configPages from '../configuration/pages';
import * as apikeysPages from '../apikey/pages';
import * as aclPages from '../acl/pages';
import { AdminCreateGuard as AdminOwnerGuard } from '../core/guards/admin-create.guard';

const routes: Routes = [
  {
    path: 'acl',
    canActivate: [AdminOwnerGuard],
    component: aclPages.AclComponent,
  },
  {
    path: 'acl/create',
    canActivate: [AdminOwnerGuard],
    component: aclPages.AclCreateUpdateComponent,
  },
  {
    path: 'acl/:id/edit',
    canActivate: [AdminOwnerGuard],
    component: aclPages.AclCreateUpdateComponent,
  },
  {
    path: 'apikeys',
    canActivate: [AdminOwnerGuard],
    component: apikeysPages.ApikeysComponent,
  },
  {
    path: 'apikeys',
    canActivate: [AdminOwnerGuard],
    component: apikeysPages.ApikeysComponent,
  },
  {
    path: 'datacenter',
    component: adminPages.AdminDatacentersComponent,
  },
  {
    path: 'datacenter/create',
    canActivate: [AdminReadGuard],
    component: adminDatacentersPages.AdminDatacenterCreateComponent,
  },
  {
    path: 'datacenter/:datacenterName/edit',
    canActivate: [AdminReadGuard],
    component: adminDatacentersPages.AdminDatacenterCreateComponent,
  },
  {
    path: 'prices',
    canActivate: [AdminReadGuard],
    component: adminPages.AdminPricesComponent,
  },
  {
    path: 'prices/create',
    canActivate: [AdminReadGuard],
    component: adminPricesPages.AdminPricesCreateComponent,
  },
  {
    path: 'prices/:id/edit',
    canActivate: [AdminReadGuard],
    component: adminPricesPages.AdminPricesCreateComponent,
  },
  {
    path: 'auditlogs',
    canActivate: [AdminOwnerGuard],
    component: adminPages.AdminAuditlogsComponent,
  },
  {
    path: 'projects',
    component: adminPages.ProjectsComponent,
  },
  {
    path: 'projects/create',
    canActivate: [AdminReadGuard],
    component: adminProjectsPages.ProjectsCreateComponent,
  },
  {
    path: 'projects/:id/edit',
    canActivate: [AdminReadGuard],
    component: adminProjectsPages.ProjectsCreateComponent,
  },
  {
    path: 'projects/:id',
    component: adminProjectsPages.ProjectDetailsComponent,
  },
  {
    path: 'datacenter/:datacenterName',
    component: datacentersPages.DatacenterDetailComponent,
  },
  {
    path: 'configuration',
    canActivate: [AdminOwnerGuard],
    component: configPages.ConfigurationComponent,
  },
  {
    path: 'configuration/operatorconfig/create',
    canActivate: [AdminOwnerGuard],
    component: configPages.ConfigOperatorconfigCreateUpdateComponent,
  },
  {
    path: 'configuration/operatorconfig/:id/edit',
    canActivate: [AdminOwnerGuard],
    component: configPages.ConfigOperatorconfigCreateUpdateComponent,
  },
  {
    path: 'configuration/task/create',
    canActivate: [AdminOwnerGuard],
    component: configPages.ConfigTaskCreateUpdateComponent,
  },
  {
    path: 'configuration/task/:id/edit',
    canActivate: [AdminOwnerGuard],
    component: configPages.ConfigTaskCreateUpdateComponent,
  },
  {
    path: 'configuration/desiredversion/create',
    canActivate: [AdminOwnerGuard],
    component: configPages.ConfigDesiredversionCreateUpdateComponent,
  },
  {
    path: 'configuration/desiredversion/edit',
    canActivate: [AdminOwnerGuard],
    component: configPages.ConfigDesiredversionCreateUpdateComponent,
  },
  {
    path: 'policyreports',
    canActivate: [AdminReadGuard],
    component: adminPages.PolicyReportsComponent,
  },
  {
    path: 'vulnerabilityreports',
    canActivate: [AdminReadGuard],
    component: adminPages.VulnerabilityReportsComponent,
  },
  {
    path: 'compliancereports',
    canActivate: [AdminReadGuard],
    component: adminPages.ComplianceReportsComponent,
  },
];

export const AdminRoutingModule = RouterModule.forChild(routes);
