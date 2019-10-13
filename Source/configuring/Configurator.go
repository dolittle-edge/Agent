/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Dolittle. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package configuring

import (
	"agent/log"
	"agent/provisioning"
	"bytes"
	"encoding/json"
	"errors"
	"strings"
)

// Configurator represents a system that coordinates node configurers
type Configurator struct {
	configuration map[string]interface{}
	configurers   map[string]ICanConfigureNode
	debug         bool
}

// NewConfigurator creates a new instance of the Configurer
func NewConfigurator(provisioner *provisioning.Provider, configurers []ICanConfigureNode) (*Configurator, error) {
	configurator := new(Configurator)
	configurator.configuration = make(map[string]interface{})
	configurator.configurers = make(map[string]ICanConfigureNode)

	for _, configurer := range configurers {
		if _, exists := configurator.configurers[configurer.Name()]; exists {
			return nil, errors.New("Multiple configurers for type " + configurer.Name())
		}
		configurator.configurers[configurer.Name()] = configurer
	}

	configurator.startConfigurationListener(provisioner)

	return configurator, nil
}

// SetDebug sets the debugging output flag
func (configurator *Configurator) SetDebug(debug bool) {
	configurator.debug = debug

	for _, configurer := range configurator.configurers {
		configurer.SetDebug(debug)
	}
}

func (configurator *Configurator) startConfigurationListener(provisioner *provisioning.Provider) {
	listener := make(chan provisioning.Node)
	go func() {
		firstConfiguration := true
		for {
			node := <-listener
			newConfiguration := make(map[string]interface{})
			for name, configuration := range node.Configuration {
				newConfiguration[strings.ToLower(name)] = configuration
			}
			configurator.runConfigurers(newConfiguration, firstConfiguration)
			configurator.configuration = newConfiguration
			firstConfiguration = false
		}
	}()
	provisioner.Listen(listener)
}

func (configurator *Configurator) runConfigurers(newConfiguration map[string]interface{}, force bool) {
	for name := range configurator.configuration {
		if _, exists := newConfiguration[name]; !exists {
			log.Informationf("Configuration for type %s was removed\n", name)
			configurator.runConfigurer(name, nil, false)
		}
	}

	for name, configuration := range newConfiguration {
		oldConfig, err := json.Marshal(configurator.configuration[name])
		if err != nil {
			log.Errorf("Could not marshal existing configuration for type %s: %v\n", name, configurator.configuration[name])
		}
		newConfig, err := json.Marshal(configuration)
		if err != nil {
			log.Errorf("Could not marshal new configuration for type %s: %v\n", name, configuration)
		}
		if force || !bytes.Equal(oldConfig, newConfig) {
			log.Informationf("Configuration for type %s changed\n", name)
			configurator.runConfigurer(name, configuration, true)
		}
	}
}

func (configurator *Configurator) runConfigurer(name string, value interface{}, warnIfNoConfigurer bool) {
	if configurer, exists := configurator.configurers[name]; exists {
		err := configurer.Configure(value)
		if err != nil {
			log.Errorf("Configurer for type %s failed: %s\n", name, err)
		}
		return
	}
	if warnIfNoConfigurer {
		log.Warningf("No configurer set for type %s\n", name)
	}
}
