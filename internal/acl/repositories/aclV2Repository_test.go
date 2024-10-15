package aclrepository

import (
	"github.com/NorskHelsenett/ror/internal/mocks/identitymocks"
	"reflect"
	"testing"

	aclmodels "github.com/NorskHelsenett/ror/pkg/models/acl"

	identitymodels "github.com/NorskHelsenett/ror/pkg/models/identity"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_createACLV2FilterByScope(t *testing.T) {
	type args struct {
		identity identitymodels.Identity
		scope    aclmodels.Acl2Scope
	}
	tests := []struct {
		name string
		args args
		want []bson.M
	}{
		{
			name: "EmptyGroup",
			args: args{
				identity: identitymocks.ValiduserWithGroups([]string{""}),
				scope:    aclmodels.Acl2ScopeCluster,
			},
			want: []bson.M{
				{
					"$match": bson.M{
						"group": bson.M{
							"$in": bson.A{
								"Unknown-Unauthorized",
							},
						},
					},
				},
				{
					"$match": bson.M{
						"$or": bson.A{
							bson.M{
								"scope": aclmodels.Acl2ScopeCluster,
							},
							bson.M{
								"scope": aclmodels.Acl2ScopeRor,
								"subject": bson.M{
									"$in": []string{
										string(aclmodels.Acl2ScopeCluster),
										string(aclmodels.Acl2RorSubjectGlobal),
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "NoGroup",
			args: args{
				identity: identitymocks.ValiduserWithGroups([]string{}),
				scope:    aclmodels.Acl2ScopeCluster,
			},
			want: []bson.M{
				{
					"$match": bson.M{
						"group": bson.M{
							"$in": bson.A{
								"Unknown-Unauthorized",
							},
						},
					},
				},
				{
					"$match": bson.M{
						"$or": bson.A{
							bson.M{
								"scope": aclmodels.Acl2ScopeCluster,
							},
							bson.M{
								"scope": aclmodels.Acl2ScopeRor,
								"subject": bson.M{
									"$in": []string{
										string(aclmodels.Acl2ScopeCluster),
										string(aclmodels.Acl2RorSubjectGlobal),
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "SingleGroup",
			args: args{
				identity: identitymocks.ValiduserWithGroups([]string{"T1-A-TEST-Admin@test.nhn.no"}),
				scope:    aclmodels.Acl2ScopeCluster,
			},
			want: []bson.M{
				{
					"$match": bson.M{
						"group": bson.M{
							"$in": bson.A{
								"T1-A-TEST-Admin@test.nhn.no",
							},
						},
					},
				},
				{
					"$match": bson.M{
						"$or": bson.A{
							bson.M{
								"scope": aclmodels.Acl2ScopeCluster,
							},
							bson.M{
								"scope": aclmodels.Acl2ScopeRor,
								"subject": bson.M{
									"$in": []string{
										string(aclmodels.Acl2ScopeCluster),
										string(aclmodels.Acl2RorSubjectGlobal),
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "MultiGroup",
			args: args{
				identity: identitymocks.ValiduserWithGroups([]string{"T1-A-TEST-Admin@test.nhn.no", "T1-A-TEST-User01@test.nhn.no", "T1-A-TEST-User02@test.nhn.no", "T1-A-TEST-User03@test.nhn.no", "T1-A-TEST-User04@test.nhn.no", "T1-A-TEST-User05@test.nhn.no", "T1-A-TEST-User06@test.nhn.no"}),
				scope:    aclmodels.Acl2ScopeCluster,
			},
			want: []bson.M{
				{
					"$match": bson.M{
						"group": bson.M{
							"$in": bson.A{
								"T1-A-TEST-Admin@test.nhn.no",
								"T1-A-TEST-User01@test.nhn.no",
								"T1-A-TEST-User02@test.nhn.no",
								"T1-A-TEST-User03@test.nhn.no",
								"T1-A-TEST-User04@test.nhn.no",
								"T1-A-TEST-User05@test.nhn.no",
								"T1-A-TEST-User06@test.nhn.no",
							},
						},
					},
				},
				{
					"$match": bson.M{
						"$or": bson.A{
							bson.M{
								"scope": aclmodels.Acl2ScopeCluster,
							},
							bson.M{
								"scope": aclmodels.Acl2ScopeRor,
								"subject": bson.M{
									"$in": []string{
										string(aclmodels.Acl2ScopeCluster),
										string(aclmodels.Acl2RorSubjectGlobal),
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createACLV2FilterByScope(tt.args.identity, tt.args.scope); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createACLV2FilterByScope() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createACLV2FilterByScopeSubject(t *testing.T) {
	type args struct {
		identity identitymodels.Identity
		scope    aclmodels.Acl2Scope
		subject  aclmodels.Acl2Subject
	}
	tests := []struct {
		name string
		args args
		want []bson.M
	}{
		{
			name: "Empty Group",
			args: args{
				identity: identitymocks.ValiduserWithGroups([]string{""}),
				scope:    aclmodels.Acl2ScopeCluster,
				subject:  aclmodels.Acl2Subject("t-test-001"),
			},
			want: []bson.M{
				{
					"$match": bson.M{
						"group": bson.M{
							"$in": bson.A{
								"Unknown-Unauthorized",
							},
						},
					},
				},
				{
					"$match": bson.M{
						"$or": bson.A{
							bson.M{
								"scope":   aclmodels.Acl2ScopeCluster,
								"subject": aclmodels.Acl2Subject("t-test-001"),
							},
							bson.M{
								"scope": aclmodels.Acl2ScopeRor,
								"subject": bson.M{
									"$in": []string{
										string(aclmodels.Acl2ScopeCluster),
										string(aclmodels.Acl2RorSubjectGlobal),
									},
								},
							},
						},
					},
				},
			},
		}, {
			name: "No Group",
			args: args{
				identity: identitymocks.ValiduserWithGroups([]string{}),
				scope:    aclmodels.Acl2ScopeCluster,
				subject:  aclmodels.Acl2Subject("t-test-001"),
			},
			want: []bson.M{
				{
					"$match": bson.M{
						"group": bson.M{
							"$in": bson.A{
								"Unknown-Unauthorized",
							},
						},
					},
				},
				{
					"$match": bson.M{
						"$or": bson.A{
							bson.M{
								"scope":   aclmodels.Acl2ScopeCluster,
								"subject": aclmodels.Acl2Subject("t-test-001"),
							},
							bson.M{
								"scope": aclmodels.Acl2ScopeRor,
								"subject": bson.M{
									"$in": []string{
										string(aclmodels.Acl2ScopeCluster),
										string(aclmodels.Acl2RorSubjectGlobal),
									},
								},
							},
						},
					},
				},
			},
		}, {
			name: "SingleGroup",
			args: args{
				identity: identitymocks.ValiduserWithGroups([]string{"T1-A-TEST-Admin@test.nhn.no"}),
				scope:    aclmodels.Acl2ScopeCluster,
				subject:  aclmodels.Acl2Subject("t-test-001"),
			},
			want: []bson.M{
				{
					"$match": bson.M{
						"group": bson.M{
							"$in": bson.A{
								"T1-A-TEST-Admin@test.nhn.no",
							},
						},
					},
				},
				{
					"$match": bson.M{
						"$or": bson.A{
							bson.M{
								"scope":   aclmodels.Acl2ScopeCluster,
								"subject": aclmodels.Acl2Subject("t-test-001"),
							},
							bson.M{
								"scope": aclmodels.Acl2ScopeRor,
								"subject": bson.M{
									"$in": []string{
										string(aclmodels.Acl2ScopeCluster),
										string(aclmodels.Acl2RorSubjectGlobal),
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "MultiGroup",
			args: args{
				identity: identitymocks.ValiduserWithGroups([]string{"T1-A-TEST-Admin@test.nhn.no", "T1-A-TEST-User01@test.nhn.no", "T1-A-TEST-User02@test.nhn.no", "T1-A-TEST-User03@test.nhn.no", "T1-A-TEST-User04@test.nhn.no", "T1-A-TEST-User05@test.nhn.no", "T1-A-TEST-User06@test.nhn.no"}),
				scope:    aclmodels.Acl2ScopeCluster,
				subject:  aclmodels.Acl2Subject("t-test-001"),
			},
			want: []bson.M{
				{
					"$match": bson.M{
						"group": bson.M{
							"$in": bson.A{
								"T1-A-TEST-Admin@test.nhn.no",
								"T1-A-TEST-User01@test.nhn.no",
								"T1-A-TEST-User02@test.nhn.no",
								"T1-A-TEST-User03@test.nhn.no",
								"T1-A-TEST-User04@test.nhn.no",
								"T1-A-TEST-User05@test.nhn.no",
								"T1-A-TEST-User06@test.nhn.no",
							},
						},
					},
				},
				{
					"$match": bson.M{
						"$or": bson.A{
							bson.M{
								"scope":   aclmodels.Acl2ScopeCluster,
								"subject": aclmodels.Acl2Subject("t-test-001"),
							},
							bson.M{
								"scope": aclmodels.Acl2ScopeRor,
								"subject": bson.M{
									"$in": []string{
										string(aclmodels.Acl2ScopeCluster),
										string(aclmodels.Acl2RorSubjectGlobal),
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createACLV2FilterByScopeSubject(tt.args.identity, tt.args.scope, tt.args.subject)

			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("createACLV2FilterByScopeSubject() = %v, want %v", got, tt.want)
			// }
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("MakeGatewayInfo() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_compileAccessSum(t *testing.T) {
	type args struct {
		existing aclmodels.AclV2ListItemAccess
		new      aclmodels.AclV2ListItemAccess
	}
	tests := []struct {
		name string
		args args
		want aclmodels.AclV2ListItemAccess
	}{
		{
			name: "AddRead",
			args: args{
				existing: denyallACL,
				new: aclmodels.AclV2ListItemAccess{
					Read: true,
				},
			},
			want: aclmodels.AclV2ListItemAccess{
				Read: true,
			},
		}, {
			name: "AddCreate",
			args: args{
				existing: denyallACL,
				new: aclmodels.AclV2ListItemAccess{
					Create: true,
				},
			},
			want: aclmodels.AclV2ListItemAccess{
				Create: true,
			},
		}, {
			name: "AddUpdate",
			args: args{
				existing: denyallACL,
				new: aclmodels.AclV2ListItemAccess{
					Update: true,
				},
			},
			want: aclmodels.AclV2ListItemAccess{
				Update: true,
			},
		}, {
			name: "AddDelete",
			args: args{
				existing: denyallACL,
				new: aclmodels.AclV2ListItemAccess{
					Delete: true,
				},
			},
			want: aclmodels.AclV2ListItemAccess{
				Delete: true,
			},
		}, {
			name: "AddOwner",
			args: args{
				existing: denyallACL,
				new: aclmodels.AclV2ListItemAccess{
					Owner: true,
				},
			},
			want: aclmodels.AclV2ListItemAccess{
				Owner: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compileAccessSum(tt.args.existing, tt.args.new); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compileAccessSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
