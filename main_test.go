package dvcardlib

import (
	"bufio"
	"reflect"
	"testing"
)

func TestLoadAllCardLibsFromFolder(t *testing.T) {
	type args struct {
		folder string
	}
	tests := []struct {
		name    string
		args    args
		want    []CardLib
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadAllCardLibsFromFolder(tt.args.folder)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadAllCardLibsFromFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadAllCardLibsFromFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCardLibFile(t *testing.T) {
	type args struct {
		reader *bufio.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsCardLibFile(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsCardLibFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsCardLibFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadCardLibFromFile(t *testing.T) {
	type args struct {
		reader *bufio.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    CardLib
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadCardLibFromFile(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadCardLibFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadCardLibFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
