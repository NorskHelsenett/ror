package resourceupdatev2

import (
	"testing"

	"github.com/go-co-op/gocron"
	"github.com/stretchr/testify/assert"
)

func Test_resourcecache_CleanupRunning(t *testing.T) {
	type fields struct {
		HashList       hashList
		Workqueue      ResourceCacheWorkqueue
		cleanupRunning bool
		scheduler      *gocron.Scheduler
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Cleanup is running",
			fields: fields{
				cleanupRunning: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := resourcecache{
				HashList:       tt.fields.HashList,
				Workqueue:      tt.fields.Workqueue,
				cleanupRunning: tt.fields.cleanupRunning,
				scheduler:      tt.fields.scheduler,
			}
			if got := rc.CleanupRunning(); got != tt.want {
				t.Errorf("resourcecache.CleanupRunning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resourcecache_MarkActive(t *testing.T) {
	type fields struct {
		HashList       hashList
		Workqueue      ResourceCacheWorkqueue
		cleanupRunning bool
		scheduler      *gocron.Scheduler
	}

	testfields := fields{
		HashList: hashList{
			Items: []hashItem{
				{
					Uid:    "3c99c410-3cdd-11ee-be56-0242ac120002",
					Hash:   "1234",
					Active: false,
				},
				{
					Uid:    "3c99c410-3cdd-11ee-be56-0242ac120022",
					Hash:   "12345",
					Active: false,
				},
				{
					Uid:    "3c99c410-3cdd-11ee-be56-0242ac130022",
					Hash:   "12346",
					Active: false,
				},
			},
		},
		Workqueue:      ResourceCacheWorkqueue{},
		cleanupRunning: false,
	}

	testfieldsupdated := fields{
		HashList: hashList{
			Items: []hashItem{
				{
					Uid:    "3c99c410-3cdd-11ee-be56-0242ac120002",
					Hash:   "1234",
					Active: true,
				},
				{
					Uid:    "3c99c410-3cdd-11ee-be56-0242ac120022",
					Hash:   "12345",
					Active: false,
				},
				{
					Uid:    "3c99c410-3cdd-11ee-be56-0242ac130022",
					Hash:   "12346",
					Active: false,
				},
			},
		},
		Workqueue:      ResourceCacheWorkqueue{},
		cleanupRunning: false,
	}
	type args struct {
		uid string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{
			name:   "Test MarkActive valid input",
			fields: testfields,
			args: args{
				uid: "3c99c410-3cdd-11ee-be56-0242ac120002",
			},
			want: testfieldsupdated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := resourcecache{
				HashList:       tt.fields.HashList,
				Workqueue:      tt.fields.Workqueue,
				cleanupRunning: tt.fields.cleanupRunning,
				scheduler:      tt.fields.scheduler,
			}
			rc.MarkActive(tt.args.uid)
			assert.Equal(t, rc.HashList, tt.want.HashList)
		})
	}
}

func Test_resourcecache_runWorkqueScheduler(t *testing.T) {
	type fields struct {
		HashList       hashList
		Workqueue      ResourceCacheWorkqueue
		cleanupRunning bool
		scheduler      *gocron.Scheduler
	}
	testfields := fields{
		HashList: hashList{
			Items: []hashItem{
				{
					Uid:    "3c99c410-3cdd-11ee-be56-0242ac120002",
					Hash:   "1234",
					Active: false,
				},
				{
					Uid:    "3c99c410-3cdd-11ee-be56-0242ac120022",
					Hash:   "12345",
					Active: false,
				},
				{
					Uid:    "3c99c410-3cdd-11ee-be56-0242ac130022",
					Hash:   "12346",
					Active: false,
				},
			},
		},
		Workqueue:      ResourceCacheWorkqueue{},
		cleanupRunning: false,
	}

	tests := []struct {
		name   string
		fields fields
		want   fields
	}{
		{
			name:   "Test runWorkqueScheduler empty workque",
			fields: testfields,
			want:   testfields,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := resourcecache{
				HashList:       tt.fields.HashList,
				Workqueue:      tt.fields.Workqueue,
				cleanupRunning: tt.fields.cleanupRunning,
				scheduler:      tt.fields.scheduler,
			}
			rc.runWorkqueScheduler()
			assert.Equal(t, rc.Workqueue, tt.want.Workqueue)
		})
	}
}
