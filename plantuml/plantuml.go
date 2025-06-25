package plantuml

import (
	"bytes"
	"fmt"
)

type Api struct {
	SvrName string
	Path    string
}

type CallItem struct {
	Caller string
	Api    Api
}

type ParticipantInfo struct {
	Name   string
	AsName string
	Type   string
}

type GenerateSequenceScriptReq struct {
	Items           []CallItem
	ParticipantInfo []ParticipantInfo
}

func GenerateSequenceScript(req *GenerateSequenceScriptReq) string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("@startuml\n")
	participantMap := make(map[string]ParticipantInfo)
	for _, participant := range req.ParticipantInfo {
		participantMap[participant.Name] = participant
		fmt.Fprintf(buf, "%s %s as %s\n", participant.Type, participant.Name, participant.AsName)
	}
	buf.WriteRune('\n')

	for _, callItem := range req.Items {
		fmt.Fprintf(buf, "%s->%s: %s\n",
			participantMap[callItem.Caller].AsName,
			participantMap[callItem.Api.SvrName].AsName,
			callItem.Api.Path)
	}

	buf.WriteString("@enduml\n")
	return buf.String()
}
