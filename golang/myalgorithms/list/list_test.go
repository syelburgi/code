package lists

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		want  *List
		want1 bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := New()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("New() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestList_Search(t *testing.T) {
	type fields struct {
		head *node
		size int
	}
	type args struct {
		d interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				head: tt.fields.head,
				size: tt.fields.size,
			}
			if got := l.Search(tt.args.d); got != tt.want {
				t.Errorf("List.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Insert(t *testing.T) {
	type fields struct {
		head *node
		size int
	}
	type args struct {
		d interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				head: tt.fields.head,
				size: tt.fields.size,
			}
			if got := l.Insert(tt.args.d); got != tt.want {
				t.Errorf("List.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Delete(t *testing.T) {
	type fields struct {
		head *node
		size int
	}
	type args struct {
		d interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				head: tt.fields.head,
				size: tt.fields.size,
			}
			if got := l.Delete(tt.args.d); got != tt.want {
				t.Errorf("List.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
