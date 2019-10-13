/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package location

import (
	"agent/log"
	"bufio"
	"errors"
	"os"
	"strings"
)

const (
	machineInfoPath = "/etc/machine-info"
)

// Configurer represents a system that can configure the node location
type Configurer struct {
	debug bool
}

// NewConfigurer creates a new instance of the Configurer
func NewConfigurer() *Configurer {
	configurer := new(Configurer)
	return configurer
}

// Name returns the name of the configurer
func (configurer *Configurer) Name() string {
	return "location"
}

// SetDebug sets the debugging output flag
func (configurer *Configurer) SetDebug(debug bool) {
	configurer.debug = debug
}

// Configure triggers the configurer to configure the node with the given configuration
func (configurer *Configurer) Configure(value interface{}) error {
	if value == nil {
		if configurer.debug {
			log.Debugln("Removing location from machine-info")
		}
		return removeLocationFromMachineInfo()
	}
	if location, ok := value.(string); ok {
		if configurer.debug {
			log.Debugln("Setting location in machine-info to", location)
		}
		return setLocationInMachineInfo(location)
	}
	return errors.New("Location was not a string")
}

func setLocationInMachineInfo(location string) error {
	lines := readLinesExceptLocationFromMachineInfo()
	lines = append(lines, "LOCATION="+location)
	return writeLinesToMachineInfo(lines)
}

func removeLocationFromMachineInfo() error {
	lines := readLinesExceptLocationFromMachineInfo()
	return writeLinesToMachineInfo(lines)
}

func readLinesExceptLocationFromMachineInfo() []string {
	file, err := os.Open(machineInfoPath)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "LOCATION=") {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Errorf("Error reading lines from file %s: %s", machineInfoPath, err)
	}
	return lines
}

func writeLinesToMachineInfo(lines []string) error {
	file, err := os.OpenFile(machineInfoPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
