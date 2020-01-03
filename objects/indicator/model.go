// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package indicator

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

/* Indicator - This type implements the STIX 2 Indicator SDO and defines all of
the properties and methods needed to create and work with this object. All of
the methods not defined local to this type are inherited from the individual
properties. */
type Indicator struct {
	objects.CommonObjectProperties
	properties.NameProperty
	properties.DescriptionProperty
	IndicatorTypes []string `json:"indicator_types,omitempty"`
	Pattern        string   `json:"pattern,omitempty"`
	PatternType    string   `json:"pattern_type,omitempty"`
	PatternVersion string   `json:"pattern_version,omitempty"`
	ValidFrom      string   `json:"valid_from,omitempty"`
	ValidUntil     string   `json:"valid_until,omitempty"`
	properties.KillChainPhasesProperty
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

/* New - This function will create a new STIX Indicator object and return it as
a pointer. It will also initialize the object by setting all of the basic
properties. */
func New() *Indicator {
	var obj Indicator
	obj.InitSDO("indicator")
	return &obj
}