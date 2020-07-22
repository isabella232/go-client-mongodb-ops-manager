// Copyright 2019 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package opsmngr

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/go-test/deep"
)

func TestAutomation_GetMonitoringAgentConfig(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/groups/%s/automationConfig/monitoringAgentConfig", projectID), func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		_, _ = fmt.Fprint(w, `{
	"logPath": "/var/log/mongodb-mms-automation/monitoring-agent.log",
	"logPathWindows": "%SystemDrive%\\MMSAutomation\\log\\mongodb-mms-automation\\monitoring-agent.log",
	"logRotate": {
		"sizeThresholdMB": 1000,
		"timeThresholdHrs": 24
	},
	"username": "mms-monitoring-agent",
	"password": "7EeGTtFWbXc4Y26tMbZedMEN"
}`)
	})

	config, _, err := client.Automation.GetMonitoringAgentConfig(ctx, projectID)
	if err != nil {
		t.Fatalf("Automation.GetMonitoringAgentConfig returned error: %v", err)
	}

	expected := &AgentConfig{
		LogRotate: &LogRotate{
			SizeThresholdMB:  1000,
			TimeThresholdHrs: 24,
		},
		LogPath:        "/var/log/mongodb-mms-automation/monitoring-agent.log",
		LogPathWindows: "%SystemDrive%\\MMSAutomation\\log\\mongodb-mms-automation\\monitoring-agent.log",
		Username:       "mms-monitoring-agent",
		Password:       "7EeGTtFWbXc4Y26tMbZedMEN",
	}
	if diff := deep.Equal(config, expected); diff != nil {
		t.Error(diff)
	}
}
