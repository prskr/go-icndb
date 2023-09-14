package logging

import (
	"flag"
	"log/slog"
)

var _ flag.Value = (*LevelVar)(nil)

type LevelVar struct {
	Value *slog.LevelVar
}

func (l *LevelVar) String() string {
	return l.Value.String()
}

func (l *LevelVar) Set(s string) error {
	return l.Value.UnmarshalText([]byte(s))
}
