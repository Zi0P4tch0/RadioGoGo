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

package api

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"radiogogo/common"
	"radiogogo/data"
	"radiogogo/mocks"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBrowserImplEscapesIPV6(t *testing.T) {

	mockDNSLookupService := mocks.MockDNSLookupService{
		LookupIPFunc: func(host string) ([]string, error) {
			return []string{"2001:db8::1"}, nil
		},
	}

	mockHttpClient := mocks.MockHttpClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return nil, io.EOF
		},
	}

	browser, err := NewRadioBrowserWithDependencies(&mockDNSLookupService, &mockHttpClient)
	assert.NoError(t, err)

	assert.Equal(t, "http://[2001:db8::1]/json", browser.(*RadioBrowserImpl).baseUrl.String())

	assert.NoError(t, err)

}

func TestBrowserImplNewRadioBrowserWithDependencies(t *testing.T) {

	t.Run("returns an error if the DNS lookup fails", func(t *testing.T) {

		mockDNSLookupService := mocks.MockDNSLookupService{
			LookupIPFunc: func(host string) ([]string, error) {
				return nil, errors.New("dns")
			},
		}

		mockHttpClient := mocks.MockHttpClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				return nil, errors.New("http")
			},
		}

		_, err := NewRadioBrowserWithDependencies(&mockDNSLookupService, &mockHttpClient)

		assert.Error(t, err)
		assert.Equal(t, "dns", err.Error())

	})

	t.Run("returns an error if the URL parsing fails", func(t *testing.T) {

		mockDNSLookupService := mocks.MockDNSLookupService{
			LookupIPFunc: func(host string) ([]string, error) {
				return []string{"&!@#*)!@)@)@"}, nil
			},
		}

		mockHttpClient := mocks.MockHttpClient{
			DoFunc: func(req *http.Request) (*http.Response, error) {
				return nil, errors.New("http")
			},
		}

		_, err := NewRadioBrowserWithDependencies(&mockDNSLookupService, &mockHttpClient)

		assert.Error(t, err)
		assert.Equal(t, "parse \"http://[&!@\": net/url: invalid userinfo", err.Error())

	})

}

func TestBrowserImplGetStations(t *testing.T) {

	// Note: Search term set to "searchTerm" in all test cases

	testCases := []struct {
		name             string
		queryType        common.StationQuery
		expectedEndpoint string
	}{
		{
			name:             "builds the correct URL for StationQueryAll",
			queryType:        common.StationQueryAll,
			expectedEndpoint: "/json/stations",
		},
		{
			name:             "builds the correct URL for StationQueryByUUID",
			queryType:        common.StationQueryByUuid,
			expectedEndpoint: "/json/stations/byuuid/searchTerm",
		},
		{
			name:             "builds the correct URL for StationQueryByName",
			queryType:        common.StationQueryByName,
			expectedEndpoint: "/json/stations/byname/searchTerm",
		},
		{
			name:             "builds the correct URL for StationQueryByNameExact",
			queryType:        common.StationQueryByNameExact,
			expectedEndpoint: "/json/stations/bynameexact/searchTerm",
		},
		{
			name:             "builds the correct URL for StationQueryByCodec",
			queryType:        common.StationQueryByCodec,
			expectedEndpoint: "/json/stations/bycodec/searchTerm",
		},
		{
			name:             "builds the correct URL for StationQueryByCodecExact",
			queryType:        common.StationQueryByCodecExact,
			expectedEndpoint: "/json/stations/bycodecexact/searchTerm",
		},
		{
			name:             "builds the correct URL for StationQueryByCountry",
			queryType:        common.StationQueryByCountry,
			expectedEndpoint: "/json/stations/bycountry/searchTerm",
		},
		{
			name:             "builds the correct URL for StationQueryByCountryExact",
			queryType:        common.StationQueryByCountryExact,
			expectedEndpoint: "/json/stations/bycountryexact/searchTerm",
		},
		{
			name:             "builds the correct URL for StationQueryByCountryCodeExact",
			queryType:        common.StationQueryByCountryCodeExact,
			expectedEndpoint: "/json/stations/bycountrycodeexact/searchTerm",
		},
		{
			name:             "builds the correct URL for StationQueryByState",
			queryType:        common.StationQueryByState,
			expectedEndpoint: "/json/stations/bystate/searchTerm",
		},
		{
			name:             "builds the correct URL for StationQueryByStateExact",
			queryType:        common.StationQueryByStateExact,
			expectedEndpoint: "/json/stations/bystateexact/searchTerm",
		},
		{
			name:             "builds the correct URL for StationQueryByLanguage",
			queryType:        common.StationQueryByLanguage,
			expectedEndpoint: "/json/stations/bylanguage/searchTerm",
		},
		{
			name:             "builds the correct URL for StationQueryByLanguageExact",
			queryType:        common.StationQueryByLanguageExact,
			expectedEndpoint: "/json/stations/bylanguageexact/searchTerm",
		},
		{
			name:             "builds the correct URL for StationQueryByTag",
			queryType:        common.StationQueryByTag,
			expectedEndpoint: "/json/stations/bytag/searchTerm",
		},
		{
			name:             "builds the correct URL for StationQueryByTagExact",
			queryType:        common.StationQueryByTagExact,
			expectedEndpoint: "/json/stations/bytagexact/searchTerm",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			mockDNSLookupService := mocks.MockDNSLookupService{
				LookupIPFunc: func(host string) ([]string, error) {
					return []string{"127.0.0.1"}, nil
				},
			}

			mockHttpClient := mocks.MockHttpClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					assert.Equal(t, tc.expectedEndpoint, req.URL.Path)
					assert.Equal(t, "GET", req.Method)
					assert.Equal(t, "application/json", req.Header.Get("Accept"))
					assert.Equal(t, data.UserAgent, req.Header.Get("User-Agent"))
					responseBody := io.NopCloser(bytes.NewReader([]byte(`[]`)))
					return &http.Response{
						StatusCode: 200,
						Body:       responseBody,
					}, nil
				},
			}

			browser, err := NewRadioBrowserWithDependencies(&mockDNSLookupService, &mockHttpClient)

			assert.NoError(t, err)

			_, err = browser.GetStations(tc.queryType, "searchTerm", "name", false, 0, 10, true)

			assert.NoError(t, err)

		})
	}
}
func TestBrowserImplClickStation(t *testing.T) {

	station := common.Station{
		StationUuid: uuid.MustParse("941ef6f1-0699-4821-95b1-2b678e3ff62e"),
	}

	mockDNSLookupService := mocks.MockDNSLookupService{
		LookupIPFunc: func(host string) ([]string, error) {
			return []string{"127.0.0.1"}, nil
		},
	}

	mockHttpClient := mocks.MockHttpClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			expectedUrl := "http://127.0.0.1/json/url/941ef6f1-0699-4821-95b1-2b678e3ff62e"
			assert.Equal(t, "POST", req.Method)
			assert.Equal(t, expectedUrl, req.URL.String())
			assert.Equal(t, "application/json", req.Header.Get("Accept"))
			assert.Equal(t, data.UserAgent, req.Header.Get("User-Agent"))

			responseBody := io.NopCloser(bytes.NewReader([]byte(`
			{
				"ok": true,
				"message": "retrieved station url",
				"stationuuid": "9617a958-0601-11e8-ae97-52543be04c81",
				"name": "Station name",
				"url": "http://this.is.an.url"
			}
			`)))
			return &http.Response{
				StatusCode: 200,
				Body:       responseBody,
			}, nil
		},
	}

	radioBrowser, err := NewRadioBrowserWithDependencies(&mockDNSLookupService, &mockHttpClient)
	assert.NoError(t, err)

	response, err := radioBrowser.ClickStation(station)
	assert.NoError(t, err)

	assert.Equal(t, true, response.Ok)
}
