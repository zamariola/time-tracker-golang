package filesystem

import (
	"testing"
	"github.com/zamariola/time-tracker-golang/entity"
	"flag"
	"path/filepath"
	"io/ioutil"
	"fmt"
	"os"
)

const (
	TESTDATA_PATH = "../testdata"
	TEST_WRITE_GOLDEN_PATH = "written_task.golden"
)

var update = flag.Bool("update", false, "update .golden files")

func TestShouldAppendTaskOnEndOfFile(t *testing.T) {

	filePath := filepath.Join(TESTDATA_PATH, t.Name() + ".log");

	os.Remove(filePath);

	fsh := NewFileSystemHandler(filePath);

	task := entity.NewTask(MESSAGE, START_TIME, END_TIME);

	golden := filepath.Join(TESTDATA_PATH, TEST_WRITE_GOLDEN_PATH);

	if *update || !Exists(golden) {
		fmt.Println("Writing goldenfile " + golden);
		ioutil.WriteFile(golden, []byte(fsh.Format(task) + "\n"), 0644)
	}

	fsh.Write(task)

	expectedFile, _ := ioutil.ReadFile(golden)
	actualFile, _ := ioutil.ReadFile(filePath)

	if (string(expectedFile) != string(actualFile)) {
		t.Errorf("Error while writing file, expected: %s, got: %s", string(expectedFile), string(actualFile))
	}
}






