package config

import (
	"encoding/json"
	"os"
	"path"
	"runtime"

	"github.com/zew/logx"
	"github.com/zew/util"
)

type SQLHost struct {
	User             string            `json:"user"`
	Host             string            `json:"host"`
	Port             string            `json:"port"`
	DBName           string            `json:"db_name"`
	ConnectionParams map[string]string `json:"connection_params"`
}

type ConfigT struct {
	Email                string             `json:"email"`
	VersionMajor         int                `json:"version_major"`
	VersionMinor         int                `json:"version_minor"`
	AppName              string             `json:"app_name"`
	GoogleCustomSearchId string             `json:"google_custom_search_id"` // searchEngineId for google custom search engine "cse"
	AppEngineServerKey   string             `json:"appengine_server_key"`    // "Server key 1" from an app engine app
	SQLite               bool               `json:"sql_lite"`
	SQLHosts             map[string]SQLHost `json:"sql_hosts"`
	CredentialFileNames  []string           `json:"credential_file_names"`
}

var Config ConfigT

var credentialFileIdx = 0

func CredentialFileName(revolve bool) string {
	if revolve {
		credentialFileIdx++
		if credentialFileIdx > len(Config.CredentialFileNames)-1 {
			credentialFileIdx = 0
		}
	}
	return Config.CredentialFileNames[credentialFileIdx]
}

func init() {

	for _, v := range []string{"SQL_PW"} {
		util.EnvVar(v)
	}

	workDir, err := os.Getwd()
	util.CheckErr(err)
	logx.Println("workDir: ", workDir)

	_, srcFile, _, ok := runtime.Caller(1)
	if !ok {
		logx.Fatalf("runtime caller not found")
	}

	{
		fullP := path.Join(path.Dir(srcFile), "config.json")
		file, err := os.Open(fullP)
		if err != nil {
			logx.Printf("could not find %v: %v", fullP, err)
			file, err = os.Open("config.json")
			util.CheckErr(err)
		}

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Config)
		util.CheckErr(err)
		logx.Printf("\n%#s", util.IndentedDump(Config))
	}

}
