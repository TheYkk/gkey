/*
Copyright 2020 Kaan Karakaya

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gkey

import "testing"

func TestGenPass(t *testing.T) {
	type args struct {
		password string
		realm    string
		length   int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test normal password",
			args: args{
				password: "ykk",
				realm:    "github.com",
				length:   16,
			},
			want:    "vIvpK(xm0)|#?::l",
			wantErr: false,
		}, {
			name: "Test password without realm",
			args: args{
				password: "ykk",
				realm:    "",
				length:   16,
			},
			want:    "",
			wantErr: true,
		}, {
			name: "Test password without password",
			args: args{
				password: "",
				realm:    "github.com",
				length:   16,
			},
			want:    "",
			wantErr: true,
		}, {
			name: "Test password without length",
			args: args{
				password: "ykk",
				realm:    "github.com",
				length:   0,
			},
			want:    "",
			wantErr: true,
		}, {
			name: "Test password integers",
			args: args{
				password: "123",
				realm:    "123",
				length:   16,
			},
			want:    "By8Q>d,RHL2-{rki",
			wantErr: false,
		}, {
			name: "Test password with space",
			args: args{
				password: "yk k",
				realm:    "git hub.com",
				length:   16,
			},
			want:    "dHan&Ghe*l8HmQK~",
			wantErr: false,
		}, {
			name: "Test password with special chars",
			args: args{
				password: "!!@@##",
				realm:    "%%^^&&",
				length:   16,
			},
			want:    "K%0IHV7r)z!E>183",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenPass(tt.args.password, tt.args.realm, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenPass() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenPass() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Hexto94(t *testing.T) {
	type args struct {
		base byte
		max  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
		{
			name: "Test 194",
			args: args{
				base: 194,
				max:  94,
			},
			want: 6,
		}, {
			name: "Test 211",
			args: args{
				base: 211,
				max:  94,
			},
			want: 23,
		}, {
			name: "Test 94",
			args: args{
				base: 94,
				max:  94,
			},
			want: 0,
		}, {
			name: "Test 255",
			args: args{
				base: 255,
				max:  94,
			},
			want: 67,
		}, {
			name: "Test 0",
			args: args{
				base: 0,
				max:  94,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hexto94(tt.args.base, tt.args.max); got != tt.want {
				t.Errorf("hexto94() = %v, want %v", got, tt.want)
			}
		})
	}
}
