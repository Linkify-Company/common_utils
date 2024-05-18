package logger

import (
	"github.com/Linkify-Company/common_utils/errify"
	"log/slog"
)

type dev struct {
	logger *slog.Logger
}

func (d *dev) Info(err errify.IError) {
	d.logger.Info(format(err))
}

func (d *dev) Infof(format string, args ...interface{}) {
	d.logger.Info(format, args...)
}

func (d *dev) Error(err errify.IError) {
	d.logger.Error(format(err))
}

func (d *dev) Errorf(format string, args ...interface{}) {
	d.logger.Error(format, args...)
}

func (d *dev) Warn(err errify.IError) {
	d.logger.Warn(format(err))
}

func (d *dev) Warnf(format string, args ...interface{}) {
	d.logger.Warn(format, args...)
}

func (d *dev) Debug(err errify.IError) {
	d.logger.Debug(format(err))
}

func (d *dev) Debugf(format string, args ...interface{}) {
	d.logger.Debug(format, args...)
}
