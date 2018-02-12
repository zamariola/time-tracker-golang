package filesystem

import (
	"github.com/zamariola/time-tracker-golang/input"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"bufio"
	"github.com/zamariola/time-tracker-golang/util"
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

func (fsh FileSystemHandler) Write(task *input.Task) error {

	_, err := CreateIfNotExists(fsh.TrackingPath());
	if err != nil {
		return err;
	}

	return WriteStringToFile(fsh.TrackingPath(), fsh.Format(task) + "\n");
}

func (fsh FileSystemHandler) ReadLast() *input.Task {


	lines,_ := ReadFile(fsh.TrackingPath());
	for line := range lines {
		fmt.Println(line)
	}

	return nil

}

func (fsh FileSystemHandler) Format(task *input.Task) string {

	log.Debugf("Formatting %s %s %s", task.Message(), task.Start(), task.End())
	return fmt.Sprint(task.Message(), ",",
		task.Start().Format("2006-01-02 15:04"), ",",
		task.End().Format("2006-01-02 15:04"))
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
			return false
		}
	}
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

	var readFile map[int]string;
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


