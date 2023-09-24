package secretkeeper

import (
	"errors"
	"fmt"
	"math/rand"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/everesthack-incubator/secret-keeper/pkg/commander"
	"github.com/everesthack-incubator/secret-keeper/pkg/config"
	log "github.com/sirupsen/logrus"
)

func getValues(c <-chan string) []string {
	var r []string

	for i := range c {
		r = append(r, i)
	}

	return r
}

type FakeCommander struct {
	CombinedOutputFunc func() ([]byte, error)
}

func (sc FakeCommander) CombinedOutput() ([]byte, error) {
	return sc.CombinedOutputFunc()
}

func TestNewVaultDiffer(t *testing.T) {
	tests := []struct {
		name string
		want *SecretKeeper
	}{
		{
			name: "TestNewVaultDiffer",
			want: &SecretKeeper{
				filePatterns: nil,
				logLevel:     0x0,
				vaultTool:    "",
				encryptArgs:  nil,
				decryptArgs:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSecretKeeper(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVaultDiffer() = %+#v, want %+#v", got, tt.want)
			}
		})
	}
}

func TestVaultDiffer_GetEncryptArgs(t *testing.T) {
	type fields struct {
		secrets     []string
		logLevel    log.Level
		vaultTool   string
		encryptArgs []string
		decryptArgs []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "TestGetEncryptArgs",
			fields: fields{
				secrets:     nil,
				logLevel:    0x0,
				vaultTool:   "",
				encryptArgs: []string{"encrypt", "-field", "value", "-format", "json"},
				decryptArgs: nil,
			},
			want: []string{"encrypt", "-field", "value", "-format", "json"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &SecretKeeper{
				filePatterns: tt.fields.secrets,
				logLevel:     tt.fields.logLevel,
				vaultTool:    tt.fields.vaultTool,
				encryptArgs:  tt.fields.encryptArgs,
				decryptArgs:  tt.fields.decryptArgs,
			}
			if got := a.GetEncryptArgs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VaultDiffer.GetEncryptArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVaultDiffer_GetDecryptArgs(t *testing.T) {
	type fields struct {
		secrets     []string
		logLevel    log.Level
		vaultTool   string
		encryptArgs []string
		decryptArgs []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "TestGetEncryptArgs",
			fields: fields{
				secrets:     nil,
				logLevel:    0x0,
				vaultTool:   "",
				encryptArgs: nil,
				decryptArgs: []string{"encrypt", "-field", "value", "-format", "json"},
			},
			want: []string{"encrypt", "-field", "value", "-format", "json"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &SecretKeeper{
				filePatterns: tt.fields.secrets,
				logLevel:     tt.fields.logLevel,
				vaultTool:    tt.fields.vaultTool,
				encryptArgs:  tt.fields.encryptArgs,
				decryptArgs:  tt.fields.decryptArgs,
			}
			if got := a.GetDecryptArgs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VaultDiffer.GetDecryptArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVaultDiffer_GetVaultCommand(t *testing.T) {
	type fields struct {
		secrets     []string
		logLevel    log.Level
		vaultTool   string
		encryptArgs []string
		decryptArgs []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestGetEncryptArgs",
			fields: fields{
				secrets:     nil,
				logLevel:    0x0,
				vaultTool:   "ansible-vault",
				encryptArgs: nil,
				decryptArgs: []string{"encrypt", "-field", "value", "-format", "json"},
			},
			want: "ansible-vault",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &SecretKeeper{
				filePatterns: tt.fields.secrets,
				logLevel:     tt.fields.logLevel,
				vaultTool:    tt.fields.vaultTool,
				encryptArgs:  tt.fields.encryptArgs,
				decryptArgs:  tt.fields.decryptArgs,
			}
			if got := a.GetVaultCommand(); got != tt.want {
				t.Errorf("VaultDiffer.GetVaultCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVaultDiffer_InitConfig(t *testing.T) {
	type fields struct {
		secrets     []string
		logLevel    log.Level
		vaultTool   string
		encryptArgs []string
		decryptArgs []string
	}
	type args struct {
		config config.Config
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SecretKeeper
	}{
		{
			name: "TestInitConfig",
			fields: fields{
				secrets:     nil,
				logLevel:    0x0,
				vaultTool:   "",
				encryptArgs: nil,
				decryptArgs: nil,
			},
			args: args{
				config: config.Config{
					VaultTool:    "ansible-vault",
					Debug:        true,
					FilePatterns: []string{"*.vault.yml"},
					EncryptArgs:  []string{"encrypt", "-field", "value", "-format", "json"},
					DecryptArgs:  []string{"decrypt", "-field", "value", "-format", "json"},
				},
			},
			want: &SecretKeeper{
				filePatterns: []string{"*.vault.yml"},
				logLevel:     log.DebugLevel,
				vaultTool:    "ansible-vault",
				encryptArgs:  []string{"encrypt", "-field", "value", "-format", "json"},
				decryptArgs:  []string{"decrypt", "-field", "value", "-format", "json"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &SecretKeeper{}
			a.InitConfig(tt.args.config)
			if !reflect.DeepEqual(a, tt.want) {
				t.Errorf("VaultDiffer.InitConfig() = %v, want %v", a, tt.want)
			}
		})
	}
}

func TestVaultDiffer_MatchFiles(t *testing.T) {
	type fields struct {
		secrets     []string
		logLevel    log.Level
		vaultTool   string
		encryptArgs []string
		decryptArgs []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "TestMatchFiles",
			fields: fields{
				secrets:     []string{"*.go", "["},
				logLevel:    0x0,
				vaultTool:   "",
				encryptArgs: nil,
				decryptArgs: nil,
			},
			want: []string{"secret_keeper.go", "secret_keeper_test.go"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &SecretKeeper{
				filePatterns: tt.fields.secrets,
				logLevel:     tt.fields.logLevel,
				vaultTool:    tt.fields.vaultTool,
				encryptArgs:  tt.fields.encryptArgs,
				decryptArgs:  tt.fields.decryptArgs,
			}
			ch := a.MatchFiles()
			got := getValues(ch)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VaultDiffer.MatchFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVaultDiffer_Clean(t *testing.T) {
	fakeExecCommander := commander.ExecCommander
	passedFileChannel := make(chan []string)
	defer func() { commander.ExecCommander = fakeExecCommander }()
	commander.ExecCommander = func(command string, args []string, filename interface{}) commander.Runner {
		fmt.Printf("exec.Command() for %v called with %v, %v, and %v\n", t.Name(), command, args, filename)

		// convert filename to string only

		return FakeCommander{
			CombinedOutputFunc: func() ([]byte, error) {
				// return error at random
				if rand.Intn(10)%2 == 0 {
					return []byte{}, nil
				}
				go func() {
					// check type of filename
					switch f := filename.(type) {
					case string:
						passedFileChannel <- []string{f}
					default:
						passedFileChannel <- filename.([]string)
					}
				}()
				return []byte{}, errors.New("error")
			},
		}
	}
	type fields struct {
		secrets     []string
		logLevel    log.Level
		vaultTool   string
		encryptArgs []string
		decryptArgs []string
	}
	type args struct {
		files <-chan string
	}

	channel := make(chan string)

	glob, _ := filepath.Glob("*.go")

	go func(files []string) {
		for _, file := range files {
			channel <- file
		}
		close(channel)
	}(glob)

	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "TestClean",
			fields: fields{
				secrets:     []string{"*.go"},
				logLevel:    0x0,
				vaultTool:   "",
				encryptArgs: nil,
				decryptArgs: nil,
			},
			args: args{
				files: channel,
			},
			want: glob,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &SecretKeeper{
				filePatterns: tt.fields.secrets,
				logLevel:     tt.fields.logLevel,
				vaultTool:    tt.fields.vaultTool,
				encryptArgs:  tt.fields.encryptArgs,
				decryptArgs:  tt.fields.decryptArgs,
			}
			ch := a.Clean(tt.args.files)
			got := getValues(ch)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VaultDiffer.Clean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVaultDiffer_Differ(t *testing.T) {
	fakeExecCommander := commander.ExecCommander
	passedFileChannel := make(chan string)
	defer func() { commander.ExecCommander = fakeExecCommander }()
	commander.ExecCommander = func(command string, args []string, filename interface{}) commander.Runner {
		fmt.Printf("exec.Command() for %v called with %v, %v, and %v\n", t.Name(), command, args, filename)
		return FakeCommander{
			CombinedOutputFunc: func() ([]byte, error) {
				// return error at random
				if rand.Intn(10)%2 == 0 {
					return []byte{}, errors.New("error")
				}
				go func() {
					passedFileChannel <- filename.(string)
				}()
				return []byte{}, nil
			},
		}
	}

	channel := make(chan string)

	glob, _ := filepath.Glob("../*/*.go")

	go func(files []string) {
		for _, file := range files {
			channel <- file
		}
		close(channel)
	}(glob)

	type fields struct {
		secrets     []string
		logLevel    log.Level
		vaultTool   string
		encryptArgs []string
		decryptArgs []string
	}
	type args struct {
		files <-chan string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "TestDiffer",
			fields: fields{
				secrets:     []string{"*.go"},
				logLevel:    0x0,
				vaultTool:   "",
				encryptArgs: nil,
				decryptArgs: nil,
			},
			args: args{
				files: channel,
			},
			want: glob,
		},
	}
	for i, tt := range tests {
		// if we don't close the channel, the test will hang. So we close it on the last test
		go func(i int) {
			if i == len(tests)-1 {
				time.Sleep(10000 * time.Millisecond)
				close(passedFileChannel)
			}
		}(i)

		t.Run(tt.name, func(t *testing.T) {
			a := &SecretKeeper{
				filePatterns: tt.fields.secrets,
				logLevel:     tt.fields.logLevel,
				vaultTool:    tt.fields.vaultTool,
				encryptArgs:  tt.fields.encryptArgs,
				decryptArgs:  tt.fields.decryptArgs,
			}
			ch := a.Differ(tt.args.files)
			got := getValues(ch)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VaultDiffer.Differ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVaultDiffer_Encrypt(t *testing.T) {
	fakeExecCommander := commander.ExecCommander
	passedFileChannel := make(chan string)
	defer func() { commander.ExecCommander = fakeExecCommander }()
	commander.ExecCommander = func(command string, args []string, filename interface{}) commander.Runner {
		fmt.Printf("exec.Command() for %v called with %v, %v, and %v\n", t.Name(), command, args, filename)
		return FakeCommander{
			CombinedOutputFunc: func() ([]byte, error) {
				// return error at random
				if rand.Intn(10)%2 == 0 {
					return []byte{}, errors.New("error")
				}
				go func() {
					passedFileChannel <- filename.(string)
				}()
				return []byte{}, nil
			},
		}
	}

	channel := make(chan string)

	glob, _ := filepath.Glob("../*/*.go")

	go func(files []string) {
		for _, file := range files {
			channel <- file
		}
		close(channel)
	}(glob)

	type fields struct {
		secrets     []string
		logLevel    log.Level
		vaultTool   string
		encryptArgs []string
		decryptArgs []string
	}
	type args struct {
		files <-chan string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "TestEncrypt",
			fields: fields{
				secrets:     []string{"*.go"},
				logLevel:    0x0,
				vaultTool:   "",
				encryptArgs: nil,
				decryptArgs: nil,
			},
			args: args{
				files: channel,
			},
		},
	}
	for i, tt := range tests {

		// if we don't close the channel, the test will hang. So we close it on the last test
		go func(i int) {
			if i == len(tests)-1 {
				time.Sleep(1000 * time.Millisecond)
				close(passedFileChannel)
			}
		}(i)

		t.Run(tt.name, func(t *testing.T) {
			a := &SecretKeeper{
				filePatterns: tt.fields.secrets,
				logLevel:     tt.fields.logLevel,
				vaultTool:    tt.fields.vaultTool,
				encryptArgs:  tt.fields.encryptArgs,
				decryptArgs:  tt.fields.decryptArgs,
			}
			ch := a.Encrypt(tt.args.files)
			got := getValues(ch)
			sort.Strings(got)
			want := getValues(passedFileChannel)
			sort.Strings(want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("VaultDiffer.Encrypt() = %v, want %v", got, want)
			}
		})
	}
}

func TestVaultDiffer_Decrypt(t *testing.T) {
	fakeExecCommander := commander.ExecCommander
	passedFileChannel := make(chan string)
	defer func() { commander.ExecCommander = fakeExecCommander }()
	commander.ExecCommander = func(command string, args []string, filename interface{}) commander.Runner {
		fmt.Printf("exec.Command() for %v called with %v, %v, and %v\n", t.Name(), command, args, filename)
		return FakeCommander{
			CombinedOutputFunc: func() ([]byte, error) {
				// return error at random
				if rand.Intn(10)%2 == 0 {
					return []byte{}, errors.New("error")
				}
				go func() {
					passedFileChannel <- filename.(string)
				}()
				return []byte{}, nil
			},
		}
	}

	channel := make(chan string)

	glob, _ := filepath.Glob("../*/*.go")

	go func(files []string) {
		for _, file := range files {
			channel <- file
		}
		close(channel)
	}(glob)

	type fields struct {
		secrets     []string
		logLevel    log.Level
		vaultTool   string
		encryptArgs []string
		decryptArgs []string
	}
	type args struct {
		files <-chan string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "TestDecrypt",
			fields: fields{
				secrets:     []string{"*.go"},
				logLevel:    0x0,
				vaultTool:   "",
				encryptArgs: nil,
				decryptArgs: nil,
			},
			args: args{
				files: channel,
			},
			want: glob,
		},
	}
	for i, tt := range tests {
		// if we don't close the channel, the test will hang. So we close it on the last test
		go func(i int) {
			if i == len(tests)-1 {
				time.Sleep(1000 * time.Millisecond)
				close(passedFileChannel)
			}
		}(i)
		t.Run(tt.name, func(t *testing.T) {
			a := &SecretKeeper{
				filePatterns: tt.fields.secrets,
				logLevel:     tt.fields.logLevel,
				vaultTool:    tt.fields.vaultTool,
				encryptArgs:  tt.fields.encryptArgs,
				decryptArgs:  tt.fields.decryptArgs,
			}
			ch := a.Decrypt(tt.args.files)
			got := getValues(ch)
			sort.Strings(got)
			want := getValues(passedFileChannel)
			sort.Strings(want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("VaultDiffer.Decrypt() = %v, want %v", got, want)
			}
		})
	}
}
