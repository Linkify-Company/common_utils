package logger

import (
	"github.com/Linkify-Company/common_utils/errify"
	"log/slog"
)

type prod struct {
	logger *slog.Logger
}

func (d *prod) Info(err errify.IError) {
	d.logger.Info(format(err))
}

func (d *prod) Infof(format string, args ...interface{}) {
	d.logger.Info(format, args...)
}

func (d *prod) Error(err errify.IError) {
	d.logger.Error(format(err))
}

func (d *prod) Errorf(format string, args ...interface{}) {
	d.logger.Error(format, args...)
}

func (d *prod) Warn(err errify.IError) {
	d.logger.Warn(format(err))
}

func (d *prod) Warnf(format string, args ...interface{}) {
	d.logger.Warn(format, args...)
}

func (d *prod) Debug(err errify.IError) {
	d.logger.Debug(format(err))
}

func (d *prod) Debugf(format string, args ...interface{}) {
	d.logger.Debug(format, args...)
}
