/* Do not change, this code is generated from Golang structs */

export class ResourceQueryFilter {
  field?: string;
  value?: string;
  type?: string;
  operator?: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.field = source['field'];
    this.value = source['value'];
    this.type = source['type'];
    this.operator = source['operator'];
  }
}
export class ResourceQueryOrder {
  field?: string;
  descending?: boolean;
  index?: number;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.field = source['field'];
    this.descending = source['descending'];
    this.index = source['index'];
  }
}
export class RorResourceOwnerReference {
  scope: string;
  subject: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.scope = source['scope'];
    this.subject = source['subject'];
  }
}
export class GroupVersionKind {
  Group: string;
  Version: string;
  Kind: string;

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.Group = source['Group'];
    this.Version = source['Version'];
    this.Kind = source['Kind'];
  }
}
export class ResourceQuery {
  versionkind?: GroupVersionKind;
  uids?: string[];
  ownerrefs?: RorResourceOwnerReference[];
  fields?: string[];
  order?: ResourceQueryOrder[];
  filters?: ResourceQueryFilter[];
  offset?: number;
  limit?: number;
  additionalresources?: GroupVersionKind[];

  constructor(source: any = {}) {
    if ('string' === typeof source) source = JSON.parse(source);
    this.versionkind = this.convertValues(source['versionkind'], GroupVersionKind);
    this.uids = source['uids'];
    this.ownerrefs = this.convertValues(source['ownerrefs'], RorResourceOwnerReference);
    this.fields = source['fields'];
    this.order = this.convertValues(source['order'], ResourceQueryOrder);
    this.filters = this.convertValues(source['filters'], ResourceQueryFilter);
    this.offset = source['offset'];
    this.limit = source['limit'];
    this.additionalresources = this.convertValues(source['additionalresources'], GroupVersionKind);
  }

  convertValues(a: any, classs: any, asMap: boolean = false): any {
    if (!a) {
      return a;
    }
    if (a.slice) {
      return (a as any[]).map((elem) => this.convertValues(elem, classs));
    } else if ('object' === typeof a) {
      if (asMap) {
        for (const key of Object.keys(a)) {
          a[key] = new classs(a[key]);
        }
        return a;
      }
      return new classs(a);
    }
    return a;
  }
}
