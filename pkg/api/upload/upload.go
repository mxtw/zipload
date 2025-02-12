package upload

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path"

	"github.com/mxtw/zipload/pkg/api"
)

type uploadResponse struct {
	Files []string `json:"files"`
}

func Upload(client *api.Client, filename string, options Options) ([]string, error) {

	endpoint, err := url.JoinPath(client.Host, "/api/upload")
	if err != nil {
		log.Println(err)
		return []string{}, err
	}

	buf, contentType, err := readFileToMulitpart(filename)

	req, err := http.NewRequest("POST", endpoint, &buf)
	if err != nil {
		log.Println(err)
		return []string{}, err
	}

	headers := options.toHeaders()

	req.Header = headers
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", client.Token)

	hc := http.Client{}

	resp, err := hc.Do(req)
	if err != nil {
		log.Println("Error:", err)
		return []string{}, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Status)
		return []string{}, err
	}

	jsonResponse := uploadResponse{}
	json.NewDecoder(resp.Body).Decode(&jsonResponse)
	defer resp.Body.Close()

	log.Printf("uploaded file '%s' successfully", filename)

	return jsonResponse.Files, nil
}

func readFileToMulitpart(filename string) (bytes.Buffer, string, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return buf, "", err
	}

	mimeType, err := detectContentType(file)
	if err != nil {
		log.Println(err)
		return buf, "", err
	}
	defer file.Close()

	log.Printf(`Detected Content-Type '%s' in file '%s'`, mimeType, filename)

	part, err := w.CreatePart(textproto.MIMEHeader{
		"Content-Disposition": []string{fmt.Sprintf(`form-data; name="file"; filename="%s"`, path.Base(filename))},
		"Content-Type":        []string{mimeType},
	})
	if err != nil {
		log.Println(err)
		return buf, "", err
	}

	io.Copy(part, file)
	w.Close()

	return buf, w.FormDataContentType(), nil
}

func detectContentType(file *os.File) (string, error) {
	header := make([]byte, 512)
	n, err := file.Read(header)
	if err != nil {
		log.Println(err)
		return "", err
	}
	file.Seek(0, io.SeekStart) // is this actually safe to do here?
	mimeType := http.DetectContentType(header[:n])

	return mimeType, nil
}
