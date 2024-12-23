package entmediasaver

import (
	"context"
	"fmt"
	"saas/gen/ent"
	"saas/gen/ent/mediable"
)

type SaveMediaProp struct {
	Tag            string
	Order          int
	AppID          string
	MediableType   string
	MediableID     string
	AddMediaIDs    []string
	RemoveMediaIDs []string
	ClearAll       bool
}

func Hello() {
	fmt.Println("hello saver")
}

func SaveMedia(client *ent.Client, props *SaveMediaProp) {
	ctx := context.Background()

	if props.ClearAll {
		client.Mediable.Delete().Where(mediable.AppID(props.AppID), mediable.MediableType(props.MediableType), mediable.MediaID(props.MediableID)).Exec(ctx)
	}

	if props.RemoveMediaIDs != nil {
		client.Mediable.Delete().Where(mediable.AppID(props.AppID), mediable.MediableType(props.MediableType), mediable.MediaID(props.MediableID), mediable.MediaIDIn(props.RemoveMediaIDs...)).Exec(ctx)
	}

	if props.AddMediaIDs != nil {
		for _, mediaID := range props.AddMediaIDs {
			client.Mediable.Create().
				SetAppID(props.AppID).
				SetTag(props.Tag).
				SetSortOrder(0).
				SetMediableType(props.MediableType).
				SetMediableID(props.MediableID).
				SetMediaID(mediaID).
				Save(ctx)
		}
	}
}
