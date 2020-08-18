// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package safehttp

// StatusCode contains HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
type StatusCode int

const (
	StatusOK                   StatusCode = 200 // RFC 7231, 6.3.1
	StatusNoContent                       = 204 // RFC 7231, 6.3.5
	StatusMovedPermanently                = 301 // RFC 7231, 6.4.2
	StatusBadRequest                      = 400 // RFC 7231, 6.5.1
	StatusUnauthorized                    = 401 // RFC 7231, 3.1
	StatusForbidden                       = 403 // RFC 7231, 6.5.3
	StatusMethodNotAllowed                = 405 // RFC 7231, 6.5.5
	StatusUnsupportedMediaType            = 415 // RFC 7231, 6.5.13. RFC 7694, 3.
	StatusInternalServerError             = 500 // RFC 7231, 6.6.1
)
