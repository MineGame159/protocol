// SPDX-FileCopyrightText: Copyright 2020 The Go Language Server Authors
// SPDX-License-Identifier: BSD-3-Clause

package protocol

import (
	"testing"
)

func TestNewVersion(t *testing.T) {
	tests := []struct {
		name string
		i    int32
	}{
		{
			name: "Valid",
			i:    5000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := NewVersion(tt.i)
			if got := NewVersion(tt.i); *got != *want {
				t.Errorf("NewVersion(%v) = %v, want %v", tt.i, *got, *want)
			}
		})
	}
}
