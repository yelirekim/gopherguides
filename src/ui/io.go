package ui

import "io"

type IO interface {
	In() io.Reader
	Out() io.Writer
	Err() io.Writer
}

type stdio struct{}
