package io

import (
	"errors"
	"testing"
)

func TestImageExtensionValidation(t *testing.T) {
	testCases := []struct {
		name string
		args string
		want error
	}{
		{
			name: "BMP file return unsupported file format",
			args: "test.bmp",
			want: errors.New("unsupported file format"),
		},
		{
			name: "PNG file return unsupported file format",
			args: "test.png",
			want: errors.New("unsupported file format"),
		},
		{
			name: "TIFF file return unsupported file format",
			args: "test.bmp",
			want: errors.New("unsupported file format"),
		},
		{
			name: "JPEG file return no error (nil)",
			args: "test.jpeg",
			want: nil,
		},
		{
			name: "JPG file return no error (nil)",
			args: "test.jpg",
			want: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := imageExtensionValidation(tc.args)
			if got != nil && got.Error() != tc.want.Error() {
				t.Errorf("expected error %v, but got %v", tc.want, got)
			} else if got == nil && tc.want != nil {
				t.Errorf("expected error %v, but got %v", tc.want, got)
			} else if got != nil && tc.want == nil {
				t.Errorf("expected error %v, but got %v", tc.want, got)
			}
		})
	}
}
