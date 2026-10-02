/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package httpserver

import (
	"net/http"
	"testing"
	"time"
)

// The server's idle timeout must exceed the 90s idle timeout of
// http.DefaultTransport (used by e.g. the kube-apiserver webhook client), so
// that clients close idle connections before the server does.
func TestNewIdleTimeoutExceedsDefaultTransport(t *testing.T) {
	const defaultTransportIdleTimeout = 90 * time.Second

	srv := New(t.Context(), http.NewServeMux())
	if srv.IdleTimeout <= defaultTransportIdleTimeout {
		t.Fatalf("IdleTimeout = %v, want > %v", srv.IdleTimeout, defaultTransportIdleTimeout)
	}
}
