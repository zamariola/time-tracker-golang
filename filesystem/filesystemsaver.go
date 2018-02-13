package filesystem

import (
	"github.com/zamariola/time-tracker-golang/entity"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"bufio"
	"github.com/zamariola/time-tracker-golang/util"
	"strings"
	"time"
)

const (
	LINE_SEPARATOR = ","
	MARSHALLING_DATE_FORMAT = "2006-01-02 15:04"
)

type FileSystemHandler struct {
	trackingPath string;
}

func (fsh *FileSystemHandler) TrackingPath() string {
	return fsh.trackingPath;
}

func NewFileSystemHandler(path string) *FileSystemHandler {
	return &FileSystemHandler{path}
}

func (fsh FileSystemHandler) Write(task *entity.Task) error {

	_, err := CreateIfNotExists(fsh.TrackingPath());
	if err != nil {
		log.Errorf("Error while writing task %s on file %s", task, fsh.TrackingPath())
		return err;
	}

	return WriteStringToFile(fsh.TrackingPath(), fsh.Format(task) + "\n");
}

func (fsh FileSystemHandler) ReadLast() *entity.Task {

	lines, err := ReadFile(fsh.TrackingPath());
	util.CheckError(err);
	keys := util.GetIntMapKeys(&lines);
	lastTaskPtr := Unmarshall(lines[len(keys) - 1]);
	return lastTaskPtr;

}

func (fsh FileSystemHandler) Format(task *entity.Task) string {

	log.Debugf("Formatting %s %s %s", task.Message(), task.Start(), task.End())
	return fmt.Sprint(task.Message(), LINE_SEPARATOR,
		task.Start().Format(MARSHALLING_DATE_FORMAT), LINE_SEPARATOR,
		task.End().Format(MARSHALLING_DATE_FORMAT))
}

func Unmarshall(text string) *entity.Task {

	log.Debugf("Unmarshalling %s", text);
	columns := strings.Split(text, LINE_SEPARATOR);

	message := columns[0];
	startTime, err := time.Parse(MARSHALLING_DATE_FORMAT, columns[1]);
	util.CheckError(err);

	endTime, err := time.Parse(MARSHALLING_DATE_FORMAT, columns[2]);
	util.CheckError(err);

	return entity.NewTask(message, startTime, endTime);
}

func WriteStringToFile(path, text string) error {

	f, err := os.OpenFile(path, os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}

func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			log.Debugf("File %s not exists", path)
			return false
		}
	}
	log.Debugf("File %s already exists", path)
	return true
}

func CreateIfNotExists(path string) (*os.File, error) {

	if !Exists(path) {
		return os.Create(path)
	}
	return os.Open(path);
}


//TODO: Watch for performance issues
func ReadFile(path string) (map[int]string, error) {

	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if (err != nil) {
		return make(map[int]string), err;
	}
	defer file.Close()

	readFile := make(map[int]string);
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		readFile[i] = scanner.Text();
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return readFile, nil;

}

func ReadEndOfFile(path string, buf *[]byte) error {

	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if (err != nil) {
		return err;
	}
	defer file.Close()

	_, err = file.Seek(int64(-cap(*buf)), 2)
	util.CheckError(err)

	file.Read(*buf);
	return nil
}

func NewFileSystemHandlerFromDefaultConfig() *FileSystemHandler {
	config, _ := LoadConfig("")
	path := config[KEY_CONFIG_TRACKING_PATH];
	return NewFileSystemHandler(path);
}

