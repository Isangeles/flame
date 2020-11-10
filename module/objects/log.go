/*
 * log.go
 *
 * Copyright 2020 Dariusz Sikora <dev@isangeles.pl>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
 * MA 02110-1301, USA.
 *
 *
 */

package objects

import (
	"time"

	"github.com/isangeles/flame/data/res"
)

// Struct for character log.
type Log struct {
	channel  chan string
	messages []Message
}

// Struct for character log message.
type Message struct {
	time time.Time
	text string
}

// NewLog creates new log.
func NewLog() *Log {
	l := Log{channel: make(chan string, 3)}
	return &l
}

// Add adds new message to log.
func (l *Log) Add(m string) {
	msg := Message{time.Now(), m}
	select {
	case l.channel <- m:
	default:
	}
	l.messages = append(l.messages, msg)
}

// Channel returns log channel.
func (l *Log) Channel() chan string {
	return l.channel
}

// Messages returns all messages from log.
func (l *Log) Messages() []Message {
	return l.messages
}

// Clear clears log messages.
func (l *Log) Clear() {
	l.messages = make([]Message, 0)
}

// Apply applies specified data on the object log.
func (l *Log) Apply(data res.ObjectLogData) {
	l.Clear()
	for _, md := range data.Messages {
		m := Message{md.Time, md.Text}
		l.messages = append(l.messages, m)
	}
}

// Data creates data resource for object log.
func (l *Log) Data() (data res.ObjectLogData) {
	for _, m := range l.Messages() {
		md := res.ObjectLogMessageData{m.Time(), m.String()}
		data.Messages = append(data.Messages, md)
	}
	return
}

// String returns message text.
func (m Message) String() string {
	return m.text
}

// Time returns message time.
func (m Message) Time() time.Time {
	return m.time
}
