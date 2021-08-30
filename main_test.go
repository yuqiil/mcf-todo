package main

import (
	"log"
	"os"
	"testing"
)

func TestFilePathWalkDir(t *testing.T) {
	// Get path of current directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	type args struct {
		root string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Case1", // Running the script from the somemodule sub-directory
			args: args{
				root: dir + "/somedir/somemodule",
			},
			want:    dir + "/somedir/somemodule/somefile.js\n" + dir + "/somedir/somemodule/someotherfile.js",
			wantErr: false,
		},
		{
			name: "Case2", // Running from the somedir3 sub-directory
			args: args{
				root: dir + "/somedir3",
			},
			want:    dir + "/somedir3/another_file.js",
			wantErr: false,
		},
		{
			name: "Case3", // Running from the somedir2 sub-directory
			args: args{
				root: dir + "/somedir2",
			},
			want:    dir + "/somedir2/anotherdir/index.js\n" + dir + "/somedir2/anotherdir/yet_another_dir/index.js\n" + dir + "/somedir2/index.js",
			wantErr: false,
		},
		{
			name: "Case4", // Running from the yet_another_dir sub-directory
			args: args{
				root: dir + "/somedir2/anotherdir/yet_another_dir",
			},
			want:    dir + "/somedir2/anotherdir/yet_another_dir/index.js",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FilePathWalkDir(tt.args.root)
			if (err != nil) != tt.wantErr {
				t.Errorf("FilePathWalkDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FilePathWalkDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
