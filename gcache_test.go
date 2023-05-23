package gcache

import (
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	type args struct {
		ttl time.Duration
	}
	tests := []struct {
		name string
		args args
		want *Cache
	}{
		{name: "testnew_", args: args{ttl: time.Duration(time.Second)}, want: New(time.Duration(time.Second))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.ttl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestNew2(t *testing.T) {
// 	type args struct {
// 		ttl time.Duration
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want any
// 	}{
// 		{name: "testnew_", args: args{ttl: time.Duration(time.Second)}, want: New2(time.Duration(time.Second), time.Duration(time.Second))},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := New(tt.args.ttl); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("New() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestCache_Set(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "TestCache_Set_1",
			args: args{
				key:   "test1",
				value: 123,
			},
			want: 123,
		},
		{
			name: "TestCache_Set_2",
			args: args{
				key:   "test2",
				value: `{"name":"jeffotoni"}`,
			},
			want: `{"name":"jeffotoni"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(200 * time.Millisecond)
			c.Set(tt.args.key, tt.args.value)
			got, exist := c.Get(tt.args.key)
			if exist {
				switch got.(type) {
				case int:
					vgot := got.(int)
					if vgot != tt.want {
						t.Errorf("Cache.Get() = %v, want %v", vgot, tt.want)
					}
				case string:
					vgot := got.(string)
					if vgot != tt.want {
						t.Errorf("Cache.Get() = %v, want %v", vgot, tt.want)
					}
				}
			}
			time.Sleep(300 * time.Millisecond)
			_, exist = c.Get(tt.args.key)
			if exist {
				t.Errorf("Cache item should have been expired and not exist")
			}
		})
	}
}

func TestCache_Get(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "TestCache_Set_1",
			args: args{
				key:   "test1",
				value: 123,
			},
			want: 123,
		},
		{
			name: "TestCache_Set_2",
			args: args{
				key:   "test2",
				value: `{"name":"jeffotoni"}`,
			},
			want: `{"name":"jeffotoni"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(200 * time.Millisecond)
			c.Set(tt.args.key, tt.args.value)
			got, exist := c.Get(tt.args.key)
			if exist {
				switch got.(type) {
				case int:
					vgot := got.(int)
					if vgot != tt.want {
						t.Errorf("Cache.Get() = %v, want %v", vgot, tt.want)
					}
				case string:
					vgot := got.(string)
					if vgot != tt.want {
						t.Errorf("Cache.Get() = %v, want %v", vgot, tt.want)
					}
				}
			}
			time.Sleep(300 * time.Millisecond)
			_, exist = c.Get(tt.args.key)
			if exist {
				t.Errorf("Cache item should have been expired and not exist")
			}
		})
	}
}

// func TestCache_Delete(t *testing.T) {
// 	type fields struct {
// 		mu    sync.Mutex
// 		ttl   time.Duration
// 		items map[string]*Item
// 		heap  expirationHeap
// 	}
// 	type args struct {
// 		key string
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := &Cache{
// 				mu:    tt.fields.mu,
// 				ttl:   tt.fields.ttl,
// 				items: tt.fields.items,
// 				heap:  tt.fields.heap,
// 			}
// 			c.Delete(tt.args.key)
// 		})
// 	}
// }

// func TestCache_cleanEvic(t *testing.T) {
// 	type fields struct {
// 		mu    sync.Mutex
// 		ttl   time.Duration
// 		items map[string]*Item
// 		heap  expirationHeap
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := &Cache{
// 				mu:    tt.fields.mu,
// 				ttl:   tt.fields.ttl,
// 				items: tt.fields.items,
// 				heap:  tt.fields.heap,
// 			}
// 			c.cleanEvic()
// 		})
// 	}
// }

// func TestCache_evict(t *testing.T) {
// 	type fields struct {
// 		mu    sync.Mutex
// 		ttl   time.Duration
// 		items map[string]*Item
// 		heap  expirationHeap
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := &Cache{
// 				mu:    tt.fields.mu,
// 				ttl:   tt.fields.ttl,
// 				items: tt.fields.items,
// 				heap:  tt.fields.heap,
// 			}
// 			c.evict()
// 		})
// 	}
// }

// func Test_expirationHeap_Len(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		h    expirationHeap
// 		want int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.h.Len(); got != tt.want {
// 				t.Errorf("expirationHeap.Len() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_expirationHeap_Less(t *testing.T) {
// 	type args struct {
// 		i int
// 		j int
// 	}
// 	tests := []struct {
// 		name string
// 		h    expirationHeap
// 		args args
// 		want bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
// 				t.Errorf("expirationHeap.Less() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_expirationHeap_Swap(t *testing.T) {
// 	type args struct {
// 		i int
// 		j int
// 	}
// 	tests := []struct {
// 		name string
// 		h    expirationHeap
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.h.Swap(tt.args.i, tt.args.j)
// 		})
// 	}
// }

// func Test_expirationHeap_Push(t *testing.T) {
// 	type args struct {
// 		x interface{}
// 	}
// 	tests := []struct {
// 		name string
// 		h    *expirationHeap
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.h.Push(tt.args.x)
// 		})
// 	}
// }

// func Test_expirationHeap_Pop(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		h    *expirationHeap
// 		want interface{}
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := tt.h.Pop(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("expirationHeap.Pop() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
