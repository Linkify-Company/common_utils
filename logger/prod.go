package logger

import (
	"fmt"
	"github.com/Linkify-Company/common_utils/errify"
	"log/slog"
)

type prod struct {
	logger *slog.Logger
}

func (p *prod) Info(err errify.IError) {
	p.logger.Info(format(err))
}

func (p *prod) Infof(format string, args ...interface{}) {
	p.logger.Info(format, args...)
}

func (p *prod) Error(err errify.IError) {
	p.logger.Error(format(err))
}

func (p *prod) Errorf(format string, args ...interface{}) {
	p.logger.Error(format, args...)
}

func (p *prod) Warn(err errify.IError) {
	p.logger.Warn(format(err))
}

func (p *prod) Warnf(format string, args ...interface{}) {
	p.logger.Warn(format, args...)
}

func (p *prod) Debug(err errify.IError) {
	p.logger.Debug(format(err))
}

func (p *prod) Debugf(format string, args ...interface{}) {
	p.logger.Debug(format, args...)
}

func (p *prod) Panic(err errify.IError) {
	panic(format(err))
}

func (p *prod) Panicf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}
