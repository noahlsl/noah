package versionx

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/noahlsl/noah/tools/ipx"
	"github.com/noahlsl/noah/tools/strx"
)

type Version struct {
	Server    string `json:"server"`
	BuildTime string `json:"build_time"`
	CommitId  string `json:"commit_id"`
	Branch    string `json:"branch"`
	Listen    string `json:"listen"`
}

func NewVersion(server, buildTime, commitId, branch string, port int) *Version {

	return &Version{
		Server:    server,
		BuildTime: buildTime,
		CommitId:  commitId,
		Branch:    branch,
		Listen:    fmt.Sprintf("%s:%d", ipx.GetClientIp(), port),
	}
}

func (r *Version) ToStr() string {
	marshal, _ := json.Marshal(r)
	return strx.B2s(marshal)
}

func (r *Version) ToBytes() []byte {
	marshal, _ := json.Marshal(r)
	return marshal
}
