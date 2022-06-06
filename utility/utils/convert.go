package utils

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"runtime"
)

// ConvertToPDF 转pdf
func ConvertToPDF(filePath string) (string, error) {
	commandName := ""
	var params []string
	switch runtime.GOOS {
	case "windows":
		commandName = "cmd"
		params = []string{"/c", "soffice", "--headless", "--invisible", "--convert-to", "pdf", "--outdir", "cache/pdf/", filePath}
	case "linux":
		commandName = "libreoffice7.3"
		params = []string{"--invisible", "--headless", "--convert-to", "pdf", "--outdir", "cache/pdf/", filePath}
	case "darwin":
		commandName = "/Applications/LibreOffice.app/Contents/MacOS/soffice"
		params = []string{"--headless", "--invisible", "--convert-to", "pdf", "--outdir", "cache/pdf/", filePath}
	default:
		return "", gerror.Newf("ConvertToPDF Nonsupport OS: %s", runtime.GOOS)
	}

	err := interactiveToexec(commandName, params)
	if err != nil {
		return "", err
	}
	resultPath := "cache/pdf/" + gfile.Name(filePath) + ".pdf"
	if gfile.Exists(resultPath) {
		return resultPath, nil
	} else {
		return "", gerror.Newf("ConvertToPDF resultPath NotExists: : %s", resultPath)
	}
}
