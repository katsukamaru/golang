package thumnails

import "testing"

func TestImageFile(t *testing.T) {
	type args struct {
		infile string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"file open test", args{"./test.jpg"}, "./test.jpg", true},
		{"file nonexist open test", args{"./test_nonexitst.jpg"}, "no such file or directory", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ImageFile(tt.args.infile)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImageFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ImageFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
