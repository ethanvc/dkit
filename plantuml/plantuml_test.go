package plantuml

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateSequenceScript(t *testing.T) {
	content := `
{
  "ParticipantInfo": [
    {
      "Name": "user",
      "AsName": "user",
      "Type": "actor"
    },
    {
      "Name": "app",
      "AsName": "app",
      "Type": "participant"
    },
    {
      "Name": "core",
      "AsName": "core",
      "Type": "participant"
    },
    {
      "Name": "db",
      "AsName": "db",
      "Type": "database"
    }
  ],
  "Items": [
    {
      "Caller": "user",
      "Api": {
        "SvrName": "app",
        "Path": "/createOrder"
      }
    },
    {
      "Caller": "app",
      "Api": {
        "SvrName": "core",
        "Path": "/processPayment"
      }
    },
    {
      "Caller": "core",
      "Api": {
        "SvrName": "db",
        "Path": "insertOrder"
      }
    }
  ]
}
`
	var req GenerateSequenceScriptReq
	err := json.Unmarshal([]byte(content), &req)
	require.NoError(t, err)
	s := GenerateSequenceScript(&req)
	fmt.Println(s)
}
