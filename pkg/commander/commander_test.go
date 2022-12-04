package commander

import (
	"errors"
	"fmt"
	"os/exec"
	"reflect"
	"testing"
)

type FakeCommander struct {
	CombinedOutputFunc func() ([]byte, error)
}

func (sc FakeCommander) CombinedOutput() ([]byte, error) {
	return sc.CombinedOutputFunc()
}

var mockResponse = []byte("mock response")

func TestNewCommander(t *testing.T) {
	type args struct {
		command  string
		args     []string
		filename string
	}
	tests := []struct {
		name string
		args args
		want Runner
	}{
		{
			name: "test",
			args: args{
				command:  "test",
				args:     []string{"test"},
				filename: "test",
			},
			want: &Commander{
				Cmd: exec.Command("test", "test", "test"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCommander(tt.args.command, tt.args.args, tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommander() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommand(t *testing.T) {
	fakeExecCommander := ExecCommander
	defer func() { ExecCommander = fakeExecCommander }()
	ExecCommander = func(command string, args []string, filename interface{}) Runner {
		fmt.Printf("exec.Command() for %v called with %v, %v, and %v\n", t.Name(), command, args, filename)
		return FakeCommander{
			CombinedOutputFunc: func() ([]byte, error) {
				return mockResponse, errors.New("mock error")
			},
		}
	}
	type args struct {
		command  string
		args     []string
		filename interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				command:  "ls",
				args:     []string{"-la"},
				filename: ".",
			},
			want:    mockResponse,
			wantErr: true,
		},
		{
			name: "test",
			args: args{
				command:  "ls",
				args:     []string{"-la"},
				filename: []string{".", ".."},
			},
			want:    mockResponse,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Command(tt.args.command, tt.args.args, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Command() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Command() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestGitDiff(t *testing.T) {
	fakeExecCommander := ExecCommander
	defer func() { ExecCommander = fakeExecCommander }()
	ExecCommander = func(command string, args []string, filename interface{}) Runner {
		fmt.Printf("exec.Command() for %v called with %v, %v, and %v\n", t.Name(), command, args, filename)
		return FakeCommander{
			CombinedOutputFunc: func() ([]byte, error) {
				return mockResponse, errors.New("mock error")
			},
		}
	}

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				filename: "test",
			},
			want:    mockResponse,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GitDiff(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("GitDiff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GitDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitRestore(t *testing.T) {
	fakeExecCommander := ExecCommander
	defer func() { ExecCommander = fakeExecCommander }()
	ExecCommander = func(command string, args []string, filename interface{}) Runner {
		fmt.Printf("exec.Command() for %v called with %v, %v, and %v\n", t.Name(), command, args, filename)
		return FakeCommander{
			CombinedOutputFunc: func() ([]byte, error) {
				// golang return error
				return mockResponse, errors.New("mock error")
			},
		}
	}
	type args struct {
		files []string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				files: []string{"test"},
			},
			want:    mockResponse,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GitRestore(tt.args.files)
			if (err != nil) != tt.wantErr {
				t.Errorf("GitRestore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GitRestore() = %v, want %v", got, tt.want)
			}
		})
	}
}
