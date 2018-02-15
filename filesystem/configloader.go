package filesystem

import (
	"io/ioutil"
	"os"
	"log"
	"os/user"
	"bufio"
	"strings"
	"io"
	"github.com/zamariola/time-tracker-golang/util"
)

const (
	DEFAULT_CONFIG_PATH = "/.time-tracker/.config"
	KEY_CONFIG_TRACKING_PATH = "tracking.path"
)

var (
	DEFAULT_CONFIG_HEADER = []byte("### Insert here the pair of key=value ###")
)

type Config map[string]string

func LoadConfig(customPath string) (Config, error) {

	path := getHomeFolder() + DEFAULT_CONFIG_PATH;
	if len(customPath) > 0 {
		path = customPath;
	}

	if !Exists(path) {
		log.Print("File not exists, creating it ", path);
		err := ioutil.WriteFile(path, DEFAULT_CONFIG_HEADER, 0644);
		util.CheckError(err)
	}

	return ReadConfig(path);
}

func getHomeFolder() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir;
}

func ReadConfig(filename string) (Config, error) {

	config := Config{}
	if len(filename) == 0 {
		return config, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		//Skipping comment line
		if strings.HasPrefix(line,"#") {
			continue;
		}

		// check if the line has = sign
		// and process the line. Ignore the rest.
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal + 1:])
				}
				// assign the config map
				config[key] = value
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}

