package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/selahaddinislamoglu/golang-studies/projects/github-activity/internal/model"
)

type ActivityService struct {
}

func NewActivityService() ServiceInterface {
	return &ActivityService{}
}

func (s *ActivityService) GetUserActivity(username string) ([]model.Event, error) {

	var url string = "https://api.github.com/users/" + username + "/events"

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == 404 {
		return nil, fmt.Errorf("user not found. please check the username")
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching data: %d", response.StatusCode)
	}

	var events []model.Event
	if err := json.NewDecoder(response.Body).Decode(&events); err != nil {
		return nil, err
	}

	return events, nil
}
