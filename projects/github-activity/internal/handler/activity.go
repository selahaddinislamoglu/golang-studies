package handler

import (
	"fmt"

	"github.com/selahaddinislamoglu/golang-studies/projects/github-activity/internal/service"
)

type ActivityHandler struct {
	service service.ServiceInterface
}

func NewActivityHandler(service service.ServiceInterface) HandlerInterface {
	return &ActivityHandler{
		service: service,
	}
}

func (h *ActivityHandler) GetUserActivity(username string) ([]string, error) {

	events, err := h.service.GetUserActivity(username)
	if err != nil {
		return nil, err
	}

	var activity []string
	for _, event := range events {
		var action string
		switch event.Type {
		case "CommitCommentEvent":
			action = fmt.Sprintf("Commented on a commit in %s", event.Repo.Name)
		case "CreateEvent":
			action = fmt.Sprintf("Created %s %s in %s", event.Payload.RefType, event.Payload.Ref, event.Repo.Name)
		case "DeleteEvent":
			action = fmt.Sprintf("Deleted %s %s in %s", event.Payload.RefType, event.Payload.Ref, event.Repo.Name)
		case "ForkEvent":
			action = fmt.Sprintf("Forked %s from %s", event.Repo.Name, event.Repo.Name)
		case "GollumEvent":
			action = fmt.Sprintf("Updated a wiki in %s", event.Repo.Name)
		case "IssueCommentEvent":
			action = fmt.Sprintf("Commented on an issue in %s", event.Repo.Name)
		case "IssuesEvent":
			action = fmt.Sprintf("Created or edited an issue in %s", event.Repo.Name)
		case "MemberEvent":
			action = fmt.Sprintf("Added a new member to %s", event.Repo.Name)
		case "PublicEvent":
			action = fmt.Sprintf("Made %s public", event.Repo.Name)
		case "PullRequestEvent":
			action = fmt.Sprintf("Opened a pull request in %s", event.Repo.Name)
		case "PullRequestReviewEvent":
			action = fmt.Sprintf("Reviewed a pull request in %s", event.Repo.Name)
		case "PullRequestReviewCommentEvent":
			action = fmt.Sprintf("Commented on a pull request review in %s", event.Repo.Name)
		case "PullRequestReviewThreadEvent":
			action = fmt.Sprintf("Commented on a pull request review thread in %s", event.Repo.Name)
		case "PushEvent":
			if len(event.Payload.Commits) > 0 {
				action = fmt.Sprintf("Pushed %d commits to %s in %s", len(event.Payload.Commits), event.Payload.Ref, event.Repo.Name)
			} else {
				action = fmt.Sprintf("Pushed to %s in %s", event.Payload.Ref, event.Repo.Name)
			}
		case "ReleaseEvent":
			action = fmt.Sprintf("Released a new version in %s", event.Repo.Name)
		case "SponsorshipEvent":
			action = fmt.Sprintf("Sponsored a user in %s", event.Repo.Name)
		case "WatchEvent":
			action = fmt.Sprintf("Starred %s", event.Repo.Name)
		default:
			action = fmt.Sprintf("%s in %s", event.Type, event.Repo.Name)
		}
		activity = append(activity, action)
	}

	return activity, nil
}
