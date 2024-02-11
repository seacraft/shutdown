// Copyright 2024 The seacraft Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http:www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package posixsignal provides a listener for a posix signal. By default
it listens for SIGINT and SIGTERM, but others can be chosen in NewPosixSignalManager.
When ShutdownFinish is called it exits with os.Exit(0)
*/
package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

// Name defines shutdown ShutdownManager name.
const Name = "PosixSignalManager"

// PosixSignalManager implements ShutdownManager interface that is added
// to GracefulShutdown. Initialize with NewPosixSignalManager.
type PosixSignalManager struct {
	signals []os.Signal
}

// NewPosixSignalManager initializes the PosixSignalManager.
// As arguments you can provide os.Signal-s to listen to, if none are given,
// it will default to SIGINT and SIGTERM.
func NewPosixSignalManager(sig ...os.Signal) *PosixSignalManager {
	if len(sig) == 0 {
		sig = make([]os.Signal, 2)
		sig[0] = os.Interrupt
		sig[1] = syscall.SIGTERM
	}

	return &PosixSignalManager{
		signals: sig,
	}
}

// GetName returns name of this ShutdownManager.
func (posixSignalManager *PosixSignalManager) GetName() string {
	return Name
}

// Start starts listening for posix signals.
func (posixSignalManager *PosixSignalManager) Start(gs GSInterface) error {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, posixSignalManager.signals...)

		// Block until a signal is received.
		<-c

		gs.StartShutdown(posixSignalManager)
	}()

	return nil
}

// ShutdownStart does nothing.
func (posixSignalManager *PosixSignalManager) ShutdownStart() error {
	return nil
}

// ShutdownFinish exits the app with os.Exit(0).
func (posixSignalManager *PosixSignalManager) ShutdownFinish() error {
	os.Exit(0)

	return nil
}
