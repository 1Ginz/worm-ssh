package history

import (
	"encoding/json"
	"fmt"
	"github.com/JcKendo/worm-ssh/internal/config"
	"github.com/charmbracelet/bubbles/table"
	"os"
	"slices"
	"strings"
	"time"
)

func AddHistoryFromArgs(args []string, mode string) {
	if mode == config.TSHMode {
		addHistoryTSH(args)
	} else {
		addHistorySSH(args)
	}
}

func addHistorySSH(args []string) {
	if len(args) == 1 && !strings.Contains(args[0], "@") {
		localConfig, err := config.GetConfig(args[0])
		if err != nil || localConfig.Name == "" {
			return
		}

		AddHistory(localConfig)
		return
	}

	generatedConfig := config.SSHConfig{Mode: config.SSHMode}

	skipNext := false
	for i, arg := range args {
		if skipNext {
			skipNext = false
			continue
		}

		switch {
		case strings.HasPrefix(arg, "-p"):
			if arg == "-p" {
				generatedConfig.Port = args[i+1]
				skipNext = true
			} else {
				generatedConfig.Port = args[i][2:]
			}
		case arg == "-i":
			generatedConfig.Key = args[i+1]
			skipNext = true
		case strings.Contains(arg, "@"):
			values := strings.Split(arg, "@")
			generatedConfig.User = values[0]
			generatedConfig.Host = values[1]
		}
	}
	AddHistory(generatedConfig)
}
func addHistoryTSH(args []string) {
	if len(args) == 1 && !strings.Contains(args[0], "@") {
		localConfig, err := config.GetConfig(args[0])
		if err != nil || localConfig.Name == "" {
			return
		}
		AddHistory(localConfig)
		return
	}

	generatedConfig := config.SSHConfig{Mode: config.TSHMode}

	for _, arg := range args {

		switch {
		case strings.Contains(arg, "@"):
			values := strings.Split(arg, "@")
			generatedConfig.User = values[0]
			generatedConfig.Host = values[1][3:]
		}
	}
	AddHistory(generatedConfig)
}

func AddHistory(c config.SSHConfig) {
	if c.Host == "" {
		return
	}

	list, err := Fetch(getFile())

	if err != nil {
		fmt.Println("error getting file")
		return
	}

	err = saveFile(SSHHistory{Connection: c, Date: time.Now()}, list)
	if err != nil {
		fmt.Println("error saving file")
		return
	}
}

func RemoveByIP(row table.Row) {
	list, err := Fetch(getFile())

	if err != nil {
		fmt.Println("error getting file")
		return
	}

	ip := row[0]

	saving := make([]SSHHistory, 0, len(list)-1)

	for _, item := range list {
		if item.Connection.Host == ip {
			continue
		}

		saving = append(saving, item)
	}

	err = saveFile(SSHHistory{}, saving)
	if err != nil {
		panic("error saving file")
	}

}

func saveFile(n SSHHistory, l []SSHHistory) error {
	file := getFileLocation()
	fileContent := stringify(n, l)

	err := os.WriteFile(file, []byte(fileContent), 0644)

	return err
}

func stringify(n SSHHistory, l []SSHHistory) string {
	history := make([]SSHHistory, 0)

	for i, sshHistory := range l {
		if sshHistory.Connection.Host == n.Connection.Host &&
			sshHistory.Connection.Name == n.Connection.Name {
			l = slices.Delete(l, i, i+1)
		}
	}

	if n.Connection.Host != "" {
		history = append(history, n)
	}

	history = append(history, l...)
	content, err := json.Marshal(history)

	if err != nil {
		return ""
	}

	return string(content)
}
