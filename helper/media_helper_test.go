package helper

import "testing"

func TestImageUploadHelper(t *testing.T) {
	type args struct {
		input interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ImageUploadHelper(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImageUploadHelper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ImageUploadHelper() got = %v, want %v", got, tt.want)
			}
		})
	}
}
