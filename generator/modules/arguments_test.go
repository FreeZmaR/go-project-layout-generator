package modules

import "testing"

func TestMakeLineArgumentsTemplate(t *testing.T) {
	type testCase struct {
		name string
		args []argument
		want string
	}

	tt := []testCase{
		{
			name: "Empty",
			want: "",
		},
		{
			name: "One argument",
			args: []argument{{"ctx", "context.Context"}},
			want: "ctx context.Context",
		},
		{
			name: "Two arguments",
			args: []argument{
				{"ctx", "context.Context"},
				{"some", "someType"},
			},
			want: "ctx context.Context, some someType",
		},
		{
			name: "Three arguments with one type",
			args: []argument{
				{"one", "someType"},
				{"some", "someType"},
				{"another", "someType"},
			},
			want: "one, some, another someType",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := makeLineArgumentsTemplate(tc.args...)
			if got != tc.want {
				t.Errorf("Wrong result. Got: %s | Want: %s", got, tc.want)
			}
		})
	}
}

func TestMakeWithNewLinesArguments(t *testing.T) {
	type testCase struct {
		name string
		args []argument
		want string
	}

	tt := []testCase{
		{
			name: "Empty",
			want: "",
		},
		{
			name: "One argument",
			args: []argument{{"ctx", "context.Context"}},
			want: "\n\tctx context.Context,\n",
		},
		{
			name: "Two arguments",
			args: []argument{
				{"ctx", "context.Context"},
				{"some", "someType"},
			},
			want: "\n\tctx context.Context,\n\tsome someType,\n",
		},
		{
			name: "Three arguments with one type",
			args: []argument{
				{"one", "someType"},
				{"some", "someType"},
				{"another", "someType"},
			},
			want: "\n\tone, some, another someType,\n",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := makeWithNewLinesArguments(tc.args...)
			if got != tc.want {
				t.Errorf("Wrong result. Got: %s | Want: %s", got, tc.want)
			}
		})
	}
}

func TestHasContextArgument(t *testing.T) {
	type testCase struct {
		name string
		args []argument
		want bool
	}

	tt := []testCase{
		{
			name: "Empty",
			want: false,
		},
		{
			name: "One argument",
			args: []argument{{"ctx", "context.Context"}},
			want: true,
		},
		{
			name: "Two arguments",
			args: []argument{
				{"ctx", "context.Context"},
				{"some", "someType"},
			},
			want: true,
		},
		{
			name: "Three arguments with one type",
			args: []argument{
				{"one", "someType"},
				{"some", "someType"},
				{"another", "someType"},
			},
			want: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := hasContextArgument(tc.args)
			if got != tc.want {
				t.Errorf("Wrong result. Got: %t | Want: %t", got, tc.want)
			}
		})
	}
}

func TestGroupArgumentsByType(t *testing.T) {
	type testCase struct {
		name string
		args []argument
		want map[string][]argument
	}

	tt := []testCase{
		{
			name: "Empty",
			want: map[string][]argument{},
		},
		{
			name: "One argument",
			args: []argument{{"ctx", "context.Context"}},
			want: map[string][]argument{"context.Context": {{"ctx", "context.Context"}}},
		},
		{
			name: "Two arguments",
			args: []argument{
				{"ctx", "context.Context"},
				{"some", "someType"},
			},
			want: map[string][]argument{
				"context.Context": {{"ctx", "context.Context"}},
				"someType":        {{"some", "someType"}},
			},
		},
		{
			name: "Three arguments with one type",
			args: []argument{
				{"one", "someType"},
				{"some", "someType"},
				{"another", "someType"},
			},
			want: map[string][]argument{
				"someType": {{"one", "someType"}, {"some", "someType"}, {"another", "someType"}},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := groupArgumentsByType(tc.args)
			if len(got) != len(tc.want) {
				t.Errorf("Wrong result. Got: %v | Want: %v", got, tc.want)
			}

			for k, v := range got {
				if len(v) != len(tc.want[k]) {
					t.Errorf("Wrong result. Got: %v | Want: %v", got, tc.want)
				}
			}
		})
	}
}
