/*
Copyright (C) GMO GlobalSign, Inc. 2019 - All Rights Reserved.

Unauthorized copying of this file, via any medium is strictly prohibited.
No distribution/modification of whole or part thereof is allowed.

Proprietary and confidential.
*/

package config_test

import (
	"path/filepath"
	"testing"

	"github.com/globalsign/hvclient/internal/config"
)

func TestConfigNewFromFile(t *testing.T) {
	t.Parallel()

	var testcases = []struct {
		filename string
		want     config.Config
	}{
		{
			"testdata/test.conf",
			config.Config{
				URL:           "https://emea.api.hvca.globalsign.com:8443/v2",
				APIKey:        "api key goes here",
				APISecret:     "api secret goes here",
				CertFile:      "/home/jdoe/fully/qualified/path/to/certfile.pem",
				KeyFile:       "/home/jdoe/fully/qualified/path/to/keyfile.pem",
				KeyPassphrase: "",
				Timeout:       30,
			},
		},
		{
			"testdata/test_enc.conf",
			config.Config{
				URL:           "https://emea.api.hvca.globalsign.com:8443/v2",
				APIKey:        "api key goes here",
				APISecret:     "api secret goes here",
				CertFile:      "/home/jdoe/fully/qualified/path/to/certfile.pem",
				KeyFile:       "/home/jdoe/fully/qualified/path/to/keyfile.pem",
				KeyPassphrase: "mypassphrase",
				Timeout:       30,
			},
		},
		{
			"testdata/test_insecure.conf",
			config.Config{
				URL:                "https://emea.api.hvca.globalsign.com:8443/v2",
				APIKey:             "api key goes here",
				APISecret:          "api secret goes here",
				CertFile:           "/home/jdoe/fully/qualified/path/to/certfile.pem",
				KeyFile:            "/home/jdoe/fully/qualified/path/to/keyfile.pem",
				KeyPassphrase:      "",
				InsecureSkipVerify: true,
				Timeout:            30,
			},
		},
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(filepath.Base(tc.filename), func(t *testing.T) {
			t.Parallel()

			var got *config.Config
			var err error
			if got, err = config.NewFromFile(tc.filename); err != nil {
				t.Fatalf("couldn't get configuration from file: %v", err)
			}

			if *got != tc.want {
				t.Errorf("got %v, want %v", *got, tc.want)
			}
		})
	}
}

func TestConfigNewFromFileError(t *testing.T) {
	t.Parallel()

	var testcases = []string{
		"there/is/no_such_file.conf",
		"testdata/malformed.conf",
	}

	for _, tc := range testcases {
		var tc = tc

		t.Run(filepath.Base(tc), func(t *testing.T) {
			t.Parallel()

			if _, err := config.NewFromFile(tc); err == nil {
				t.Errorf("unexpectedly got configuration from file")
			}
		})
	}
}
