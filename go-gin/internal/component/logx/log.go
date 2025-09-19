package logx

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type default_log struct {
	ctx context.Context
}

// 【重要】TODO: 这是一个非常经典且巧妙的 Go 语言技巧，用来【在编译期强制检查一个类型是否实现了某个接口】
// 【在编译期强制检查 *default_log 是否实现了 Logger 接口，如果没实现，编译失败】
var _ Logger = (*default_log)(nil)

func WithContext(ctx context.Context) Logger {
	return &default_log{ctx}
}

func (l *default_log) Debug(keyword string, data any) {
	l.print(zerolog.DebugLevel, keyword, data)
}

func (l *default_log) Debugf(keyword string, format string, data ...any) {
	l.Debug(keyword, fmt.Sprintf(format, data...))
}

func (l *default_log) Info(keyword string, data any) {
	l.print(zerolog.InfoLevel, keyword, data)
}

func (l *default_log) Infof(keyword string, format string, data ...any) {
	l.Info(keyword, fmt.Sprintf(format, data...))
}

func (l *default_log) Warn(keyword string, data any) {
	l.print(zerolog.WarnLevel, keyword, data)
}

func (l *default_log) Warnf(keyword string, format string, data ...any) {
	l.Warn(keyword, fmt.Sprintf(format, data...))
}

func (l *default_log) Error(keyword string, data any) {
	l.print(zerolog.ErrorLevel, keyword, data)
}

func (l *default_log) Errorf(keyword string, format string, data ...any) {
	l.Error(keyword, fmt.Sprintf(format, data...))
}

func (l *default_log) print(level zerolog.Level, keyword string, data any) {
	log.WithLevel(level).Ctx(l.ctx).Str("keyword", keyword).Any("data", data).Send()
}
