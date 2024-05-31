package ashcmd

import (
	"reflect"
	"testing"
)

func TestAvast_parseOutput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name  string
		avast *Avast
		args  args
		want  *ScanResult
	}{
		{
			name:  "Basic",
			avast: NewAvast(""),
			args: args{
				input: `
C:\Users\Popo\Downloads\eicar.com.txt   EICAR Test-NOT virus!!!
# ----------------------------------------------------------------
# Nombre de fichier scannÚs : 1
# Nombre de dossiers scannÚs : 0
# Nombre de fichiers infectÚs : 1
# Taille totale des fichiers scannÚs : 68
# Base de donnÚes virale : 240531-2, 31/5/24
# DurÚe totale du scan : 0:0:0`,
			},
			want: &ScanResult{
				IsDetected:    true,
				Malware:       "EICAR Test-NOT virus!!!",
				EngineVersion: "240531-2, 31/5/24",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.avast.parseOutput(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Avast.parseOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}
