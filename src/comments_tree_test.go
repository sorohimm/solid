package src

import (
	"reflect"
	"testing"
)

func TestCommentsTree_SetBodies(t1 *testing.T) {
	type fields struct {
		Id      int
		Body    string
		Replies []CommentsTree
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "ok", fields: fields{Id: 1, Body: "", Replies: []CommentsTree{
			{Id: 2, Body: "", Replies: []CommentsTree{}},
			{Id: 3, Body: "", Replies: []CommentsTree{}},
		},
		},
			wantErr: false,
		},
		{name: "bad ids", fields: fields{Id: 1, Body: "", Replies: []CommentsTree{
			{Id: 342, Body: "", Replies: []CommentsTree{}},
			{Id: 101, Body: "", Replies: []CommentsTree{}},
		},
		},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &CommentsTree{
				Id:      tt.fields.Id,
				Body:    tt.fields.Body,
				Replies: tt.fields.Replies,
			}
			if err := t.SetBodies(); (err != nil) != tt.wantErr {
				t1.Errorf("SetBodies() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommentsTree_getIdsList(t1 *testing.T) {
	type fields struct {
		Id      int
		Body    string
		Replies []CommentsTree
	}
	type args struct {
		ids *[]int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		expect  []int
		wantErr bool
	}{
		{name: "ok", fields: fields{Id: 1, Body: "", Replies: []CommentsTree{
			{Id: 2, Body: "", Replies: []CommentsTree{}},
			{Id: 3, Body: "", Replies: []CommentsTree{}},
		},
		},
			args:    args{ids: new([]int)},
			expect:  []int{1, 2, 3},
			wantErr: false,
		},

		{name: "bad read", fields: fields{Id: 1, Body: "", Replies: []CommentsTree{
			{Id: 2, Body: "", Replies: []CommentsTree{}},
			{Id: 3, Body: "", Replies: []CommentsTree{}},
		},
		},
			args:    args{ids: new([]int)},
			expect:  []int{1, 2, 6},
			wantErr: true,
		},

		{name: "different_len_01", fields: fields{Id: 1, Body: "", Replies: []CommentsTree{
			{Id: 2, Body: "", Replies: []CommentsTree{}},
		},
		},
			args:    args{ids: new([]int)},
			expect:  []int{1, 2, 6},
			wantErr: true,
		},

		{name: "different_len_02", fields: fields{Id: 1, Body: "", Replies: []CommentsTree{
			{Id: 2, Body: "", Replies: []CommentsTree{}},
			{Id: 3, Body: "", Replies: []CommentsTree{}},
		},
		},
			args:    args{ids: new([]int)},
			expect:  []int{1, 2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &CommentsTree{
				Id:      tt.fields.Id,
				Body:    tt.fields.Body,
				Replies: tt.fields.Replies,
			}
			t.getIdsList(tt.args.ids)

			if !reflect.DeepEqual(*tt.args.ids, tt.expect) && !tt.wantErr {
				t1.Errorf("getBodies() got = %v, want %v", tt.args.ids, tt.expect)
			}
		})
	}
}

func TestCommentsTree_getBodies(t1 *testing.T) {
	type fields struct {
		Id      int
		Body    string
		Replies []CommentsTree
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[int]string
		wantErr bool
	}{
		{name: "ok", fields: fields{Id: 1, Body: "", Replies: []CommentsTree{
			{Id: 2, Body: "", Replies: []CommentsTree{}},
			{Id: 3, Body: "", Replies: []CommentsTree{}},
		},
		},
			want: map[int]string{1: "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
				2: "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla",
				3: "et iusto sed quo iure\nvoluptatem occaecati omnis eligendi aut ad\nvoluptatem doloribus vel accusantium quis pariatur\nmolestiae porro eius odio et labore et velit aut",
			},

			wantErr: false,
		},

		{name: "wrong ids", fields: fields{Id: 1, Body: "", Replies: []CommentsTree{
			{Id: 200, Body: "", Replies: []CommentsTree{}},
			{Id: 3, Body: "", Replies: []CommentsTree{}},
		},
		},
			want: map[int]string{1: "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
				2: "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla",
				3: "et iusto sed quo iure\nvoluptatem occaecati omnis eligendi aut ad\nvoluptatem doloribus vel accusantium quis pariatur\nmolestiae porro eius odio et labore et velit aut",
			},

			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &CommentsTree{
				Id:      tt.fields.Id,
				Body:    tt.fields.Body,
				Replies: tt.fields.Replies,
			}
			got, err := t.getBodies()
			if (err != nil) != tt.wantErr {
				t1.Errorf("getBodies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("getBodies() got = %v, want %v", got, tt.want)
			}
		})
	}
}
