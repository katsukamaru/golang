package thumnails

import "testing"

func Test_openFile(t *testing.T) {
	type args struct {
		filename string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Open file", args{"./test.jpg"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := openFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("openFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
