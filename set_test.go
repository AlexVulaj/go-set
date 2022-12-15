package go_set

import (
	"reflect"
	"testing"
)

func TestSet_Add(t *testing.T) {
	type args struct {
		item int
	}
	tests := []struct {
		name       string
		s          Set[int]
		args       args
		want       bool
		desiredSet Set[int]
	}{
		{
			name:       "Add to empty set",
			s:          NewSet[int](),
			args:       args{1},
			want:       true,
			desiredSet: NewSet(1),
		},
		{
			name:       "Add to non-empty set",
			s:          NewSet(1),
			args:       args{2},
			want:       true,
			desiredSet: NewSet(1, 2),
		},
		{
			name:       "Add already existing item",
			s:          NewSet(1),
			args:       args{1},
			want:       false,
			desiredSet: NewSet(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Add(tt.args.item); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.s, tt.desiredSet) {
				t.Errorf("Add() resulted in %v, want %v", tt.s.ToString(), tt.desiredSet.ToString())
			}
		})
	}
}

func TestSet_AddAll(t *testing.T) {
	type args struct {
		items []int
	}
	tests := []struct {
		name       string
		s          Set[int]
		args       args
		want       bool
		desiredSet Set[int]
	}{
		{
			name:       "Add single to empty set",
			s:          NewSet[int](),
			args:       args{[]int{1}},
			want:       true,
			desiredSet: NewSet(1),
		},
		{
			name:       "Add single to non-empty set",
			s:          NewSet(1),
			args:       args{[]int{2}},
			want:       true,
			desiredSet: NewSet(1, 2),
		},
		{
			name:       "Add single already existing item",
			s:          NewSet(1),
			args:       args{[]int{1}},
			want:       false,
			desiredSet: NewSet(1),
		},
		{
			name:       "Add multiple to empty set",
			s:          NewSet[int](),
			args:       args{[]int{1, 2}},
			want:       true,
			desiredSet: NewSet(1, 2),
		},
		{
			name:       "Add multiple to non-empty set",
			s:          NewSet(1),
			args:       args{[]int{2, 3}},
			want:       true,
			desiredSet: NewSet(1, 2, 3),
		},
		{
			name:       "Add multiple already existing items",
			s:          NewSet(1, 2),
			args:       args{[]int{1, 2}},
			want:       false,
			desiredSet: NewSet(1, 2),
		},
		{
			name:       "Add multiple partially existing items",
			s:          NewSet(1, 2),
			args:       args{[]int{2, 3}},
			want:       true,
			desiredSet: NewSet(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.AddAll(tt.args.items...); got != tt.want {
				t.Errorf("AddAll() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(tt.s, tt.desiredSet) {
				t.Errorf("Add() resulted in %v, want %v", tt.s.ToString(), tt.desiredSet.ToString())
			}
		})
	}
}

func TestSet_Clear(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
	}{
		{
			name: "Clearing an already empty set",
			s:    NewSet[int](),
		},
		{
			name: "Clearing a non-empty set",
			s:    NewSet(1, 2, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Clear()

			if tt.s.Clear(); len(tt.s) != 0 {
				t.Errorf("Clear() resulted in %v, want empty set", tt.s.ToString())
			}
		})
	}
}

func TestSet_Contains(t *testing.T) {
	type args struct {
		item int
	}
	tests := []struct {
		name string
		s    Set[int]
		args args
		want bool
	}{
		{
			name: "Empty set contains",
			s:    NewSet[int](),
			args: args{1},
			want: false,
		},
		{
			name: "Non-empty set contains value",
			s:    NewSet(1, 2, 3),
			args: args{1},
			want: true,
		},
		{
			name: "Non-empty set does not contain value",
			s:    NewSet(1, 2, 3),
			args: args{4},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.item); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_ContainsAll(t *testing.T) {
	type args struct {
		items []int
	}
	tests := []struct {
		name string
		s    Set[int]
		args args
		want bool
	}{
		{
			name: "Empty set contains",
			s:    NewSet[int](),
			args: args{[]int{1}},
			want: false,
		},
		{
			name: "Non-empty set contains single",
			s:    NewSet(1),
			args: args{[]int{1}},
			want: true,
		},
		{
			name: "Non-empty does not contain single",
			s:    NewSet(1),
			args: args{[]int{2}},
			want: false,
		},
		{
			name: "Non-empty set contains all",
			s:    NewSet(1, 2, 3),
			args: args{[]int{1, 2, 3}},
			want: true,
		},
		{
			name: "Non-empty set contains some",
			s:    NewSet(1, 2),
			args: args{[]int{1, 2, 3}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ContainsAll(tt.args.items...); got != tt.want {
				t.Errorf("ContainsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		want bool
	}{
		{
			name: "Empty set",
			s:    NewSet[int](),
			want: true,
		},
		{
			name: "Not empty",
			s:    NewSet(1),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Remove(t *testing.T) {
	type args struct {
		item int
	}
	tests := []struct {
		name       string
		s          Set[int]
		args       args
		want       bool
		desiredSet Set[int]
	}{
		{
			name:       "Remove from empty set",
			s:          NewSet[int](),
			args:       args{1},
			want:       false,
			desiredSet: NewSet[int](),
		},
		{
			name:       "Remove from non-empty set",
			s:          NewSet(1),
			args:       args{1},
			want:       true,
			desiredSet: NewSet[int](),
		},
		{
			name:       "Remove item that doesn't exist from non-empty set",
			s:          NewSet(1),
			args:       args{2},
			want:       false,
			desiredSet: NewSet(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Remove(tt.args.item); got != tt.want {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.s, tt.desiredSet) {
				t.Errorf("Add() resulted in %v, want %v", tt.s.ToString(), tt.desiredSet.ToString())
			}
		})
	}
}

func TestSet_RemoveAll(t *testing.T) {
	type args struct {
		items []int
	}
	tests := []struct {
		name       string
		s          Set[int]
		args       args
		want       bool
		desiredSet Set[int]
	}{
		{
			name:       "Remove single from empty set",
			s:          NewSet[int](),
			args:       args{[]int{1}},
			want:       false,
			desiredSet: NewSet[int](),
		},
		{
			name:       "Remove single from non-empty set",
			s:          NewSet(1),
			args:       args{[]int{1}},
			want:       true,
			desiredSet: NewSet[int](),
		},
		{
			name:       "Remove single that doesn't exist in set",
			s:          NewSet(1),
			args:       args{[]int{2}},
			want:       false,
			desiredSet: NewSet(1),
		},
		{
			name:       "Remove multiple from empty set",
			s:          NewSet[int](),
			args:       args{[]int{1, 2}},
			want:       false,
			desiredSet: NewSet[int](),
		},
		{
			name:       "Remove multiple from non-empty set",
			s:          NewSet(1, 2),
			args:       args{[]int{1, 2}},
			want:       true,
			desiredSet: NewSet[int](),
		},
		{
			name:       "Remove multiple non-existing items",
			s:          NewSet(1, 2),
			args:       args{[]int{3, 4}},
			want:       false,
			desiredSet: NewSet(1, 2),
		},
		{
			name:       "Remove multiple partially existing items",
			s:          NewSet(1, 2),
			args:       args{[]int{2, 3}},
			want:       true,
			desiredSet: NewSet(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RemoveAll(tt.args.items...); got != tt.want {
				t.Errorf("RemoveAll() = %v, want %v", got, tt.want)
			}
		})
		if !reflect.DeepEqual(tt.s, tt.desiredSet) {
			t.Errorf("Add() resulted in %v, want %v", tt.s.ToString(), tt.desiredSet.ToString())
		}
	}
}

func TestSet_Size(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		want int
	}{
		{
			name: "Empty set",
			s:    NewSet[int](),
			want: 0,
		},
		{
			name: "Single item",
			s:    NewSet(1),
			want: 1,
		},
		{
			name: "Multiple items",
			s:    NewSet(1, 2, 3),
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_ToSlice(t *testing.T) {
	tests := []struct {
		name string
		s    Set[int]
		want []int
	}{
		{
			name: "Empty set",
			s:    NewSet[int](),
			want: []int{},
		},
		{
			name: "Single item",
			s:    NewSet(1),
			want: []int{1},
		},
		{
			name: "Multiple items",
			s:    NewSet(1, 2, 3),
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.ToSlice()
			if len(got) != len(tt.want) {
				t.Errorf("Original set has length %v but ToSlice() created a Slice of length %v", len(got), len(tt.want))
			}
			for _, item := range tt.want {
				if _, ok := tt.s[item]; !ok {
					t.Errorf("ToSlice() produced a slice with items %v, but expected items %v", got, tt.s.ToString())
				}
			}
			if got := tt.s.ToSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_RetainAll(t *testing.T) {
	type args struct {
		items []int
	}
	tests := []struct {
		name       string
		s          Set[int]
		args       args
		want       bool
		desiredSet Set[int]
	}{
		{
			name:       "Retain single from empty set",
			s:          NewSet[int](),
			args:       args{[]int{1}},
			want:       false,
			desiredSet: NewSet[int](),
		},
		{
			name:       "Retain single from non-empty set",
			s:          NewSet(1),
			args:       args{[]int{1}},
			want:       false,
			desiredSet: NewSet[int](1),
		},
		{
			name:       "Retain single that doesn't exist in set",
			s:          NewSet(1),
			args:       args{[]int{2}},
			want:       true,
			desiredSet: NewSet[int](),
		},
		{
			name:       "Retain multiple from empty set",
			s:          NewSet[int](),
			args:       args{[]int{1, 2}},
			want:       false,
			desiredSet: NewSet[int](),
		},
		{
			name:       "Retain multiple from non-empty set",
			s:          NewSet(1, 2),
			args:       args{[]int{1, 2}},
			want:       false,
			desiredSet: NewSet[int](1, 2),
		},
		{
			name:       "Retain multiple non-existing items",
			s:          NewSet(1, 2),
			args:       args{[]int{3, 4}},
			want:       true,
			desiredSet: NewSet[int](),
		},
		{
			name:       "Retain multiple partially existing items",
			s:          NewSet(1, 2),
			args:       args{[]int{2, 3}},
			want:       true,
			desiredSet: NewSet(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.RetainAll(tt.args.items...); got != tt.want {
				t.Errorf("retainAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
