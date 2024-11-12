package resourceupdatev2

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

func Test_hashList_markActive(t *testing.T) {

	type fields struct {
		Items []hashItem
	}
	type args struct {
		uid string
	}

	testfields := fields{
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
	}
	testfieldsupdated := fields{
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
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{
			name:   "Test markActive valid input",
			fields: testfields,
			args: args{
				uid: "3c99c410-3cdd-11ee-be56-0242ac120002",
			},
			want: testfieldsupdated,
		}, {
			name:   "Test markActive on unknown uid",
			fields: testfields,
			args: args{
				uid: "3c99c410-3cdd-11ee-be56-0242ac120012",
			},
			want: testfields,
		}, {
			name: "Test markActive on empty hashlist",
			fields: fields{
				Items: []hashItem{},
			},
			args: args{
				uid: "3c99c410-3cdd-11ee-be56-0242ac120012",
			},
			want: fields{
				Items: []hashItem{},
			},
		}, {
			name:   "Test markActive on empty uid",
			fields: testfields,
			args: args{
				uid: "",
			},
			want: testfields,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hl := &hashList{
				Items: tt.fields.Items,
			}
			hl.markActive(tt.args.uid)
			assert.Equal(t, hl.Items, tt.want.Items)
		})
	}
}

func Test_hashList_getHashByUid(t *testing.T) {
	type fields struct {
		Items []hashItem
	}
	type args struct {
		uid string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   hashItem
		want1  int
	}{
		{
			name: "Test getHashByUid, empty hashlist",
			fields: fields{
				Items: []hashItem{},
			},
			args: args{
				uid: "3c99c410-3cdd-11ee-be56-0242ac120012",
			},
			want:  hashItem{},
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hl := hashList{
				Items: tt.fields.Items,
			}
			got, got1 := hl.getHashByUid(tt.args.uid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hashList.getHashByUid() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("hashList.getHashByUid() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_hashList_getInactiveUid(t *testing.T) {
	type fields struct {
		Items []hashItem
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "getInactive",
			fields: fields{
				Items: []hashItem{
					{
						Uid:    "3c99c410-3cdd-11ee-be56-0242ac120002",
						Hash:   "dabfadfd",
						Active: false,
					},
					{
						Uid:    "3c99c410-3cdd-11ee-be56-0242ac120012",
						Hash:   "dabfadf3",
						Active: false,
					},
					{
						Uid:    "3c99c410-3cdd-11ee-be56-0242ac12022",
						Hash:   "dab6fadf",
						Active: true,
					},
				},
			},
			want: []string{
				"3c99c410-3cdd-11ee-be56-0242ac120002",
				"3c99c410-3cdd-11ee-be56-0242ac120012",
			},
		}, {
			name: "getInactive - empty result",
			fields: fields{
				Items: []hashItem{
					{
						Uid:    "3c99c410-3cdd-11ee-be56-0242ac120002",
						Hash:   "dabfadfd",
						Active: true,
					},
					{
						Uid:    "3c99c410-3cdd-11ee-be56-0242ac120012",
						Hash:   "dabfadf3",
						Active: true,
					},
					{
						Uid:    "3c99c410-3cdd-11ee-be56-0242ac12022",
						Hash:   "dab6fadf",
						Active: true,
					},
				},
			},
			want: []string{},
		}, {
			name: "getInactive - empty list",
			fields: fields{
				Items: []hashItem{},
			},
			want: []string{},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hl := hashList{
				Items: tt.fields.Items,
			}
			if !cmp.Equal(hl.getInactiveUid(), tt.want, cmpopts.EquateEmpty()) {
				t.Errorf("%s failed", tt.name)
			}
		})
	}
}

func Test_hashList_checkUpdateNeeded(t *testing.T) {
	type fields struct {
		Items []hashItem
	}
	type args struct {
		uid  string
		hash string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "checkUpdateNeeded - no need to update",
			fields: fields{
				Items: []hashItem{
					{
						Uid:    "3c99c410-3cdd-11ee-be56-0242ac120002",
						Hash:   "dabfadfd",
						Active: true,
					},
					{
						Uid:    "3c99c410-3cdd-11ee-be56-0242ac120012",
						Hash:   "dabfadf3",
						Active: true,
					},
					{
						Uid:    "3c99c410-3cdd-11ee-be56-0242ac12022",
						Hash:   "dab6fadf",
						Active: true,
					},
				},
			},
			args: args{
				uid:  "3c99c410-3cdd-11ee-be56-0242ac12022",
				hash: "dab6fadf",
			},
			want: false,
		}, {
			name: "checkUpdateNeeded - need to update",
			fields: fields{
				Items: []hashItem{
					{
						Uid:    "3c99c410-3cdd-11ee-be56-0242ac120002",
						Hash:   "dabfadfd",
						Active: true,
					},
					{
						Uid:    "3c99c410-3cdd-11ee-be56-0242ac120012",
						Hash:   "dabfadf3",
						Active: true,
					},
					{
						Uid:    "3c99c410-3cdd-11ee-be56-0242ac12022",
						Hash:   "dab6fadf",
						Active: true,
					},
				},
			},
			args: args{
				uid:  "3c99c410-3cdd-11ee-be56-0242ac12022",
				hash: "dab6faff",
			},
			want: true,
		}, {
			name: "checkUpdateNeeded - new uid",
			fields: fields{
				Items: []hashItem{
					{
						Uid:    "3c99c410-3cdd-11ee-be56-0242ac120002",
						Hash:   "dabfadfd",
						Active: true,
					},
					{
						Uid:    "3c99c410-3cdd-11ee-be56-0242ac120012",
						Hash:   "dabfadf3",
						Active: true,
					},
					{
						Uid:    "3c99c410-3cdd-11ee-be56-0242ac12022",
						Hash:   "dab6fadf",
						Active: true,
					},
				},
			},
			args: args{
				uid:  "3c99c410-3cdd-11ee-be56-0242ac12032",
				hash: "dab6faff",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := hashList{
				Items: tt.fields.Items,
			}
			if got := rc.checkUpdateNeeded(tt.args.uid, tt.args.hash); got != tt.want {
				t.Errorf("hashList.checkUpdateNeeded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashList_updateHash(t *testing.T) {

	type fields struct {
		Items []hashItem
	}
	type args struct {
		uid  string
		hash string
	}

	testfields := fields{
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
		},
	}
	testfieldsupdated := fields{
		Items: []hashItem{
			{
				Uid:    "3c99c410-3cdd-11ee-be56-0242ac120002",
				Hash:   "1234",
				Active: false,
			},
			{
				Uid:    "3c99c410-3cdd-11ee-be56-0242ac120022",
				Hash:   "123467ff",
				Active: false,
			},
		},
	}
	testfieldsupdated2 := fields{
		Items: []hashItem{
			{
				Uid:    "3c99c410-3cdd-11ee-be56-0242ac120002",
				Hash:   "1234",
				Active: false,
			},
			{
				Uid:    "3c99c410-3cdd-11ee-be56-0242ac120022",
				Hash:   "123467ff",
				Active: false,
			},
			{
				Uid:    "3c99c410-3cdd-11ee-be56-0242ac120032",
				Hash:   "123467",
				Active: true,
			},
		},
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   fields
	}{
		{
			name:   "Test updateHash valid input",
			fields: testfields,
			args: args{
				uid:  "3c99c410-3cdd-11ee-be56-0242ac120022",
				hash: "123467ff",
			},
			want: testfieldsupdated,
		}, {
			name:   "Test updateHash add new hash",
			fields: testfieldsupdated,
			args: args{
				uid:  "3c99c410-3cdd-11ee-be56-0242ac120032",
				hash: "123467",
			},
			want: testfieldsupdated2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hl := &hashList{
				Items: tt.fields.Items,
			}
			hl.updateHash(tt.args.uid, tt.args.hash)
			assert.Equal(t, hl.Items, tt.want.Items)
		})
	}
}
