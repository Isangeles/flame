/*
 * serial_test.go
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

package serial

import (
	"fmt"
	"testing"
)

type testObject struct {
	id, serial string
}

// Tests function for assiging unique
// serial values
func TestRegister(t *testing.T) {
	ob1 := new(testObject)
	ob2 := new(testObject)
	ob1.id, ob2.id = "test", "test"
	Register(ob1)
	Register(ob2)
	if ob1.Serial() == ob2.Serial() {
		t.Errorf("Not unique serial values: %s == %s",
			ob1.Serial(), ob2.Serial())
	}
}

// Tests function for retrieving
// registered objects.
func TestObject(t *testing.T) {
	ob1 := new(testObject)
	ob2 := new(testObject)
	ob1.id, ob2.id = "test", "test"
	Register(ob1)
	Register(ob2)
	serialer := Object(ob1.ID(), ob1.Serial())
	if serialer == nil {
		t.Errorf("Object object not found after assignment: %s %s",
			ob1.ID(), ob1.Serial())
	}
}

// Tests reset function.
func TestReset(t *testing.T) {
	ob1 := new(testObject)
	ob1.id = "test"
	Register(ob1)
	Reset()
	ob2 := new(testObject)
	ob2.id = ob1.id
	Register(ob2)
	if ob2.Serial() != "0" {
		t.Errorf("Not first value assigned after reset: %s != 0",
			ob2.Serial())
	}
}

// Tests concurrent calls on Register and Object functions.
func TestConcurrentAccess(t *testing.T) {
	add := func (){
		for i := 0; i < 1000; i ++ {
			ob := testObject{"test", ""}
			Register(&ob)
		}
	}
	retrieve := func (){
		for i := 1000; i > 0; i -- {
			Object("test", fmt.Sprintf("%d", i))
		}
	}
	go add()
	go retrieve()
	go add()
	go retrieve()
	retrieve()
}

// Returns id.
func (to *testObject) ID() string {
	return to.id
}

// Returns serial.
func (to *testObject) Serial() string {
	return to.serial
}

// Sets serial value.
func (to *testObject) SetSerial(s string) {
	to.serial = s
}
