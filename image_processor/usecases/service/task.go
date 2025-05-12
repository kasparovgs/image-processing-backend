package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"net/http"
	"user_backend/domain"

	"github.com/disintegration/imaging"
)

func ProcessTask(task *domain.Task) error {
	base64image := task.Base64Image
	imgData, err := base64.StdEncoding.DecodeString(base64image)
	if err != nil {
		fmt.Println("error decoding base64 image:", err)
		return err
	}

	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		fmt.Println("error decoding image:", err)
		return err
	}
	switch task.Filter.Name {
	case "Blur":
		sigmaAny, ok := task.Filter.Parameters["sigma"]
		if ok {
			sigma, ok := sigmaAny.(float64)
			if ok {
				img = imaging.Blur(img, sigma)
			} else {
				return fmt.Errorf("invalid parametrs")
			}
		} else {
			return fmt.Errorf("there is no needed parameters")
		}
	default:
		fmt.Println("Unknown filter: ", task.Filter.Name)
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		fmt.Println("error encoding processed image:", err)
		return err
	}
	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	task.Base64Image = encoded
	task.Status = domain.Ready
	return nil
}

func CommitTask(task *domain.Task) error {
	json, err := json.Marshal(task)
	if err != nil {
		return err
	}

	resp, err := http.Post("http://app:8080/commit", "application/json", bytes.NewBuffer(json))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
