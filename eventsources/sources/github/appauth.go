/*

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

package github

import (
	"github.com/bradleyfalzon/ghinstallation/v2"
	"net/http"
)

type AppsAuthStrategy struct {
	AppID          int64
	InstallationID int64
	PrivateKey     string
	GithubBaseURL  string
	Transport      http.RoundTripper
}

// AuthTransport implements the AuthStrategy interface.
func (t *AppsAuthStrategy) AuthTransport() (http.RoundTripper, error) {
	ght, e := ghinstallation.New(t.transport(), t.AppID, t.InstallationID, []byte(t.PrivateKey))

	if e != nil {
		return nil, e
	}

	if len(t.GithubBaseURL) > 0 {
		ght.BaseURL = t.GithubBaseURL
	}

	return ght, nil
}

func (t *AppsAuthStrategy) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}

	return http.DefaultTransport
}
