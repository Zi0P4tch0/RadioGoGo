// Copyright (c) 2023 Matteo Pacini
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package common

// StationQuery represents the type of query that can be performed on a radio station.
type StationQuery string

// The following constants represent the different types of queries that can be performed on a radio station.
const (
	StationQueryAll                StationQuery = ""                   // Returns all radio stations.
	StationQueryByUuid             StationQuery = "byuuid"             // Returns radio stations by UUID.
	StationQueryByName             StationQuery = "byname"             // Returns radio stations by name.
	StationQueryByNameExact        StationQuery = "bynameexact"        // Returns radio stations by exact name.
	StationQueryByCodec            StationQuery = "bycodec"            // Returns radio stations by codec.
	StationQueryByCodecExact       StationQuery = "bycodecexact"       // Returns radio stations by exact codec.
	StationQueryByCountry          StationQuery = "bycountry"          // Returns radio stations by country.
	StationQueryByCountryExact     StationQuery = "bycountryexact"     // Returns radio stations by exact country.
	StationQueryByCountryCodeExact StationQuery = "bycountrycodeexact" // Returns radio stations by exact country code.
	StationQueryByState            StationQuery = "bystate"            // Returns radio stations by state.
	StationQueryByStateExact       StationQuery = "bystateexact"       // Returns radio stations by exact state.
	StationQueryByLanguage         StationQuery = "bylanguage"         // Returns radio stations by language.
	StationQueryByLanguageExact    StationQuery = "bylanguageexact"    // Returns radio stations by exact language.
	StationQueryByTag              StationQuery = "bytag"              // Returns radio stations by tag.
	StationQueryByTagExact         StationQuery = "bytagexact"         // Returns radio stations by exact tag.
)
