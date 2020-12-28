package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfigFilePath(t *testing.T) {
	homeDir, _ := os.UserHomeDir()
	filePath := homeDir + "/git-users.json"

	assert.Equal(t, configFilePath(), filePath)
}
