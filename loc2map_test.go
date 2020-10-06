package loc2map

import "testing"

func TestConvert(t *testing.T) {
	type args struct {
		lat      float64
		lng      float64
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "fail", args:args{}, wantErr: true},
		{name: "paris", args:args{
			lat:      48.845776,
			lng:      2.314182,
			filePath: "paris.png",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Convert(tt.args.lat, tt.args.lng, tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}