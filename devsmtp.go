package devsmtp

import (
	"io/ioutil"
	"strings"

	"github.com/ernsheong/grand"
	"github.com/gobuffalo/buffalo/mail"
	"github.com/pkg/browser"
)

// DevSMTP is an in-memory implementation for buffalo `Sender`
// interface. It's intended to open sent messages in the browser for debugging purposes.
type DevSMTP struct {
}

// Send implements buffalo `Sender` interface, to open the mail in the browser.
func (s *DevSMTP) Send(m mail.Message) error {
	filehash := grand.GenerateRandomString(16)
	previewFilename := "tmp/preview_" + filehash + ".html"

	// Only care about HTML for now
	htmlContent := ""
	for _, body := range m.Bodies {
		// handle both "text/html" and "text/html; charset=utf-8"
		if strings.HasPrefix(body.ContentType, "text/html") {
			htmlContent += body.Content
		}
	}

	if err := ioutil.WriteFile(previewFilename, []byte(htmlContent), 0644); err != nil {
		return err
	}

	defer browser.OpenFile(previewFilename)
	return nil
}

// New constructs a new DevSMTP.
func New() *DevSMTP {
	return &DevSMTP{}
}
