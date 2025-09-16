package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"image_processor/metrics"
	"net/http"
	"strconv"
	"time"
	"user_backend/domain"

	"github.com/disintegration/imaging"
)

func proccessImage(task *domain.Task) error {
	base64image := task.Base64Image
	imgData, err := base64.StdEncoding.DecodeString(base64image)
	if err != nil {
		fmt.Println("error decoding base64 image:", err)
		fmt.Println(base64image)
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
		if !ok {
			fmt.Println("there is no needed parameters")
			return fmt.Errorf("there is no needed parameters")
		}

		switch sigmaType := sigmaAny.(type) {
		case float64:
			img = imaging.Blur(img, sigmaType)

		case string:
			sigma, err := strconv.ParseFloat(sigmaType, 64)
			if err != nil {
				fmt.Println("invalid parameters:", err)
				return fmt.Errorf("invalid parameters: %v", err)
			}
			img = imaging.Blur(img, sigma)

		default:
			fmt.Printf("unsupported type for sigma: %T\n", sigmaType)
			return fmt.Errorf("unsupported type for sigma: %T", sigmaType)
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

func ProcessTask(task *domain.Task) error {
	start := time.Now()

	err := proccessImage(task)

	duration := time.Since(start).Seconds()
	filter := task.Filter.Name
	metrics.ProcessDuration.WithLabelValues(filter).Observe(duration)
	metrics.ProcessedImages.WithLabelValues(filter).Inc()

	return err
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
