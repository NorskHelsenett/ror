package resourceupdatev2

import (
	"reflect"
	"testing"
	"time"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestResourceCacheWorkqueue_ItemCount(t *testing.T) {

	workquevalues := []ResourceCacheWorkqueueObject{
		{
			SubmittedTime: time.Now(),
			RetryCount:    0,
			ResourceUpdate: &apiresourcecontracts.ResourceUpdateModel{
				Uid: "3c99c410-3cdd-11ee-be56-0242ac120002",
			},
		},
		{
			SubmittedTime: time.Now().Add(time.Hour),
			RetryCount:    0,
			ResourceUpdate: &apiresourcecontracts.ResourceUpdateModel{
				Uid: "3c99c410-3cdd-11ee-be56-0242ac120012",
			},
		},
	}
	tests := []struct {
		name string
		wq   ResourceCacheWorkqueue
		want int
	}{
		{
			name: "Test ItemCount",
			wq:   workquevalues,
			want: 2,
		}, {
			name: "Test ItemCount empty que",
			wq:   []ResourceCacheWorkqueueObject{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.wq.ItemCount(); got != tt.want {
				t.Errorf("ResourceCacheWorkqueue.ItemCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceCacheWorkqueue_GetByUid(t *testing.T) {
	type args struct {
		uid string
	}

	time1 := time.Now().Add(time.Hour)
	time2 := time.Now().Add(time.Hour * 2)
	workquevalues := []ResourceCacheWorkqueueObject{
		{
			SubmittedTime: time1,
			RetryCount:    0,
			ResourceUpdate: &apiresourcecontracts.ResourceUpdateModel{
				Uid: "3c99c410-3cdd-11ee-be56-0242ac120002",
			},
		},
		{
			SubmittedTime: time2,
			RetryCount:    0,
			ResourceUpdate: &apiresourcecontracts.ResourceUpdateModel{
				Uid: "3c99c410-3cdd-11ee-be56-0242ac120012",
			},
		},
	}

	tests := []struct {
		name  string
		m     ResourceCacheWorkqueue
		args  args
		want  ResourceCacheWorkqueueObject
		want1 int
	}{
		{
			name: "Test GetByUid empty request",
			m:    workquevalues,
			args: args{
				uid: "",
			},
			want:  ResourceCacheWorkqueueObject{},
			want1: 0,
		}, {
			name: "Test GetByUid",
			m:    workquevalues,
			args: args{
				uid: "3c99c410-3cdd-11ee-be56-0242ac120012",
			},
			want: ResourceCacheWorkqueueObject{
				SubmittedTime: time2,
				RetryCount:    0,
				ResourceUpdate: &apiresourcecontracts.ResourceUpdateModel{
					Uid: "3c99c410-3cdd-11ee-be56-0242ac120012",
				},
			},
			want1: 1,
		}, {
			name: "Test GetByUid Nonexistent uid",
			m:    workquevalues,
			args: args{
				uid: "3c99c410-3cdd-11ee-be56-0242ac120022",
			},
			want:  ResourceCacheWorkqueueObject{},
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.m.GetByUid(tt.args.uid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceCacheWorkqueue.GetByUid() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ResourceCacheWorkqueue.GetByUid() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestResourceCacheWorkqueue_Add(t *testing.T) {
	type args struct {
		resourceUpdate *apiresourcecontracts.ResourceUpdateModel
	}

	testresourceUpdate := &apiresourcecontracts.ResourceUpdateModel{
		Uid: "3c99c410-3cdd-11ee-be56-0242ac120012",
	}
	time1 := time.Now().Add(time.Hour)
	time2 := time.Now().Add(time.Hour * 2)

	var testworkque ResourceCacheWorkqueue = []ResourceCacheWorkqueueObject{
		{
			SubmittedTime: time1,
			RetryCount:    0,
			ResourceUpdate: &apiresourcecontracts.ResourceUpdateModel{
				Uid: "3c99c410-3cdd-11ee-be56-0242ac120002",
			},
		},
		{
			SubmittedTime: time2,
			RetryCount:    0,
			ResourceUpdate: &apiresourcecontracts.ResourceUpdateModel{
				Uid: "3c99c410-3cdd-11ee-be56-0242ac120012",
			},
		},
	}

	var testworkqueadded ResourceCacheWorkqueue = []ResourceCacheWorkqueueObject{
		{
			SubmittedTime: time1,
			RetryCount:    0,
			ResourceUpdate: &apiresourcecontracts.ResourceUpdateModel{
				Uid: "3c99c410-3cdd-11ee-be56-0242ac120002",
			},
		},
		{
			SubmittedTime: time2,
			RetryCount:    1,
			ResourceUpdate: &apiresourcecontracts.ResourceUpdateModel{
				Uid: "3c99c410-3cdd-11ee-be56-0242ac120012",
			},
		},
	}
	var testworkqueadded2 ResourceCacheWorkqueue = []ResourceCacheWorkqueueObject{
		{
			SubmittedTime: time1,
			RetryCount:    0,
			ResourceUpdate: &apiresourcecontracts.ResourceUpdateModel{
				Uid: "3c99c410-3cdd-11ee-be56-0242ac120002",
			},
		},
		{
			SubmittedTime: time2,
			RetryCount:    1,
			ResourceUpdate: &apiresourcecontracts.ResourceUpdateModel{
				Uid: "3c99c410-3cdd-11ee-be56-0242ac120012",
			},
		}, {
			SubmittedTime: time2,
			RetryCount:    0,
			ResourceUpdate: &apiresourcecontracts.ResourceUpdateModel{
				Uid: "3c99c410-3cdd-11ee-be56-0242ac120022",
			},
		},
	}
	tests := []struct {
		name string
		m    *ResourceCacheWorkqueue
		args args
		want *ResourceCacheWorkqueue
	}{
		{
			name: "Test Add existing uid",
			m:    &testworkque,
			args: args{
				resourceUpdate: testresourceUpdate,
			},
			want: &testworkqueadded,
		}, {
			name: "Test Add new uid",
			m:    &testworkque,
			args: args{
				resourceUpdate: &apiresourcecontracts.ResourceUpdateModel{
					Uid: "3c99c410-3cdd-11ee-be56-0242ac120022",
				},
			},
			want: &testworkqueadded2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Add(tt.args.resourceUpdate)
			opts := cmpopts.IgnoreTypes(time.Now())
			if !cmp.Equal(tt.m, tt.want, opts) {
				t.Errorf("%s failed", tt.name)
			}
		})
	}
}

func TestResourceCacheWorkqueue_DeleteByUid(t *testing.T) {
	type args struct {
		uid string
	}
	var testworkqueadded ResourceCacheWorkqueue = []ResourceCacheWorkqueueObject{
		{
			SubmittedTime: time.Now(),
			RetryCount:    1,
			ResourceUpdate: &apiresourcecontracts.ResourceUpdateModel{
				Uid: "3c99c410-3cdd-11ee-be56-0242ac120012",
			},
		},
	}
	var testworkqueadded2 ResourceCacheWorkqueue = []ResourceCacheWorkqueueObject{
		{
			SubmittedTime: time.Now(),
			RetryCount:    1,
			ResourceUpdate: &apiresourcecontracts.ResourceUpdateModel{
				Uid: "3c99c410-3cdd-11ee-be56-0242ac120012",
			},
		}, {
			SubmittedTime: time.Now(),
			RetryCount:    0,
			ResourceUpdate: &apiresourcecontracts.ResourceUpdateModel{
				Uid: "3c99c410-3cdd-11ee-be56-0242ac120022",
			},
		},
	}

	var testworkqueempty ResourceCacheWorkqueue

	tests := []struct {
		name string
		m    *ResourceCacheWorkqueue
		args args
		want *ResourceCacheWorkqueue
	}{
		{
			name: "Test Remove uid",
			m:    &testworkqueadded2,
			args: args{
				uid: "3c99c410-3cdd-11ee-be56-0242ac120022",
			},
			want: &testworkqueadded,
		}, {
			name: "Test Remove uid - empty uid",
			m:    &testworkqueadded,
			args: args{
				uid: "",
			},
			want: &testworkqueadded,
		}, {
			name: "Test Remove uid - remove last que item",
			m:    &testworkqueadded,
			args: args{
				uid: "3c99c410-3cdd-11ee-be56-0242ac120012",
			},
			want: &testworkqueempty,
		}, {
			name: "Test Remove uid - remove from empty que",
			m:    &testworkqueempty,
			args: args{
				uid: "3c99c410-3cdd-11ee-be56-0242ac120012",
			},
			want: &testworkqueempty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.DeleteByUid(tt.args.uid)
			if !cmp.Equal(tt.m, tt.want, cmpopts.IgnoreTypes(time.Now()), cmpopts.EquateEmpty()) {
				t.Errorf("%s failed", tt.name)
			}
		})
	}
}
