/*
 * Copyright 2021 Teo Mrnjavac <teo.mrnjavac@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package logger

import (
	"github.com/sirupsen/logrus"
)

type Log struct {
	logrus.Entry
}

func (logger *Log) WithPrefix(prefix string) *logrus.Entry {
	return logger.WithField("prefix", prefix)
}

func New(baseLogger *logrus.Logger, defaultPrefix string) *Log {
	logger := new(Log)
	logger.Logger = baseLogger
	logger.Data = make(logrus.Fields, 5)
	logger.Data["prefix"] = defaultPrefix
	return logger
}
