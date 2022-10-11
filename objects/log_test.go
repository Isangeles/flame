/*
 * log_test.go
 *
 * Copyright 2022 Dariusz Sikora <ds@isangeles.dev>
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
	"testing"
	"time"

	"github.com/isangeles/flame/data/res"
)

// TestLogAdd tests adding message to log.
func TestLogAdd(t *testing.T) {
	log := NewLog()
	msg := NewMessage("Message", true)
	log.Add(msg)
	logMsg := log.Messages()[0]
	if logMsg.Text != msg.Text {
		t.Fatalf("Message not added to the log")
	}
	if !logMsg.Translated {
		t.Fatalf("Translated flag not set")
	}
}

// TestLogApply tests appying data resource for log.
func TestLogApply(t *testing.T) {
	var logData res.ObjectLogData
	msg := res.ObjectLogMessageData{true, "Message 1", time.Now()}
	logData.Messages = append(logData.Messages, msg)
	msg = res.ObjectLogMessageData{true, "Message 2", time.Now()}
	logData.Messages = append(logData.Messages, msg)
	msg = res.ObjectLogMessageData{false, "Message 3", time.Now()}
	logData.Messages = append(logData.Messages, msg)
	log := NewLog()
	log.Apply(logData)
	if len(log.Messages()) != 3 {
		t.Fatalf("Wrong number of message after apply: %d != 3",
			len(log.Messages()))
	}
}

// TestLogData tests creating data resource for log.
func TestLogData(t *testing.T) {
	log := NewLog()
	log.Add(NewMessage("Message 1", true))
	log.Add(NewMessage("Message 2", true))
	log.Add(NewMessage("Message 3", false))
	logData := log.Data()
	if len(logData.Messages) != 3 {
		t.Fatalf("Wrong number of messages in data resource: %d != 3",
			len(logData.Messages))
	}
}
