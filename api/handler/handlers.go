package handler

import (
	"encoding/json"
	"fmt"
	"github.com/sharma1612harshit/docker.ui/api/docker"
	"log"
	"net/http"
)

// handler for images api
func ImagesHandler(w http.ResponseWriter, r *http.Request) {
	images, err := docker.GetImages(r.Header.Get("all"), r.Header.Get("filters"))

	if err != nil {
		log.Print(err)

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("500 - Something bad happened!"))
	} else {
		imageResponse, err := json.Marshal(images)

		if err != nil {
			log.Print(err)

			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("500 - Something bad happened!"))
		} else {
			_, _ = w.Write(imageResponse)
		}
	}
}

// handler for pull image api
func PullImageHandler(w http.ResponseWriter, r *http.Request) {
	status, err := docker.PullImage(r.Header.Get("all"), r.Header.Get("imageref"), r.Header.Get("username"), r.Header.Get("password"))

	if err != nil {
		log.Print(err)

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("%s - %s", status, err)))
	} else {
		_, _ = w.Write([]byte(status))
	}
}

// handler for containers api
func ContainersHandler(w http.ResponseWriter, r *http.Request) {
	containers, err := docker.GetContainers()

	if err != nil {
		log.Print(err)

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("%s", err)))
	} else {
		containerResponse, err := json.Marshal(containers)

		if err != nil {
			log.Print(err)

			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("%s", err)))
		} else {
			_, _ = w.Write(containerResponse)
		}
	}
}

func LaunchContainerHandler(w http.ResponseWriter, r *http.Request) {
	launchResponse, err := docker.RunContainer(r.Header.Get("name"), r.Header.Get("image_name"), "", "", "", []string{}, []string{}, map[string]string{})

	if err != nil {
		log.Print(err)

		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("%s", err)))
	} else {
		response, err := json.Marshal(launchResponse)

		if err != nil {
			log.Print(err)

			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("%s", err)))
		} else {
			_, _ = w.Write(response)
		}
	}
}
