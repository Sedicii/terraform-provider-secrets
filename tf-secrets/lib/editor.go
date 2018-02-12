package lib

import (
	"os"
	"os/exec"
)

func getEditor() (string, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	return exec.LookPath(editor)
}

func EditFileWithEditor(fileName string) (err error) {
	editor, err := getEditor()
	if err != nil {
		return
	}
	subProcess := exec.Command(editor, fileName)

	subProcess.Stdout = os.Stdout
	subProcess.Stdin = os.Stdin
	subProcess.Stderr = os.Stderr
	err = subProcess.Run()
	return
}
