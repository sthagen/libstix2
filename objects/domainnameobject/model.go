// Copyright 2015-2019 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package domainnameobject

import (
	"github.com/freetaxii/libstix2/objects"
	"github.com/freetaxii/libstix2/objects/properties"
)

/*
DomainNameObject - This type implements the STIX 2.1 DomainNameObject SCO.
The Domain Name object represents the properties of a network domain name.
*/
type DomainNameObject struct {
	objects.CommonObjectProperties
	properties.ResolvesToRefsProperty
	properties.ValueProperty
}
