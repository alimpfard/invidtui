package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// VideoResult stores the video data.
type VideoResult struct {
	Title           string `json:"title"`
	Author          string `json:"author"`
	VideoID         string `json:"videoId"`
	LengthSeconds   int    `json:"lengthSeconds"`
	AdaptiveFormats []struct {
		Type       string `json:"type"`
		URL        string `json:"url"`
		Itag       string `json:"itag"`
		Resolution string `json:"resolution,omitempty"`
	} `json:"adaptiveFormats"`
}

const videoFields = "?fields=title,videoId,author,publishedText,lengthSeconds,adaptiveFormats"

// Video gets the video with the given ID and returns a VideoResult.
func (c *Client) Video(id string) (VideoResult, error) {
	var result VideoResult

	res, err := c.ClientRequest(context.Background(), "videos/"+id+videoFields)
	if err != nil {
		return VideoResult{}, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return VideoResult{}, err
	}

	return result, nil
}

// LoadVideo takes a VideoResult, determines whether to play
// video or just audio (according to the audio parameter), and
// appropriately loads the URLs into mpv.
func LoadVideo(video VideoResult, audio bool) error {
	var err error
	var audioUrl, videoUrl string

	videoUrl, audioUrl = getVideoByItag(video, audio)

	if audio && audioUrl == "" {
		return fmt.Errorf("Could not find an audio stream")
	}

	if !audio && videoUrl == "" {
		return fmt.Errorf("Could not find a video stream")
	}

	// A data parameter is appended to audioUrl/videoUrl so that
	// updatePlaylist() can display media data.
	// MPV does not return certain track data like author and duration.
	titleparam := "&title=" + url.QueryEscape(video.Title)
	titleparam += "&author=" + url.QueryEscape(video.Author)
	titleparam += "&length=" + url.QueryEscape(FormatDuration(video.LengthSeconds))

	if audio {
		_, err = IsValidURL(audioUrl + titleparam)
		if err != nil {
			return fmt.Errorf("Could not find an audio stream")
		}

		audioUrl += titleparam

		err = GetMPV().LoadFile(
			video.Title,
			video.LengthSeconds,
			audioUrl)

	} else {
		_, err = IsValidURL(videoUrl + titleparam)
		if err != nil {
			return fmt.Errorf("Could not find a video stream")
		}

		videoUrl += titleparam

		err = GetMPV().LoadFile(
			video.Title,
			video.LengthSeconds,
			videoUrl, audioUrl)
	}
	if err != nil {
		return err
	}

	return nil
}

// getVideoByItag gets the appropriate itag of the video format, and
// returns a video and audio url using getLatestURL().
func getVideoByItag(video VideoResult, audio bool) (string, string) {
	var ftype, videoUrl, audioUrl string

	// For video streams, itag 22 is 720p and itag 18 is 360p
	// as of now in most invidious instances, may change.
	if !audio && (*videoResolution == "720p" || *videoResolution == "360p") {
		switch *videoResolution {
		case "720p":
			videoUrl = getLatestURL(video.VideoID, "22")
		case "360p":
			videoUrl = getLatestURL(video.VideoID, "18")
		}

		// audioUrl is blank since the audio stream is
		// is already merged along with the video in
		// videoUrl.
		return videoUrl, audioUrl
	}

	for _, format := range video.AdaptiveFormats {
		v := strings.Split(format.Type, ";")
		p := strings.Split(v[0], "/")

		if (audio && audioUrl != "") || (!audio && videoUrl != "") {
			break
		}

		if ftype == "" {
			ftype = p[1]
		}

		if p[1] == ftype {
			if p[0] == "audio" {
				audioUrl = getLatestURL(video.VideoID, format.Itag)
			} else if p[0] == "video" {
				videoUrl = videoWithResolution(video, "itag")
			}
		}
	}

	return videoUrl, audioUrl
}

// getVideoByFormatURL returns a URL from a VideoResult's AdaptiveFormats.
func getVideoByFormatURL(video VideoResult, audio bool) (string, string) {
	var ftype, audioUrl, videoUrl string

	for _, format := range video.AdaptiveFormats {
		v := strings.Split(format.Type, ";")
		p := strings.Split(v[0], "/")

		if (audio && audioUrl != "") || (!audio && videoUrl != "") {
			break
		}

		if ftype == "" {
			ftype = p[1]
		}

		if p[1] == ftype {
			if p[0] == "audio" {
				audioUrl = format.URL
			} else if p[0] == "video" {
				videoUrl = videoWithResolution(video, "url")
			}
		}
	}

	return videoUrl, audioUrl
}

// videoWithResolution returns a video URL that corresponds to the
// videoResolution setting (passed via command line option --video-res=).
func videoWithResolution(video VideoResult, vtype string) string {
	var prevData string

	vq := *videoResolution

	for _, format := range video.AdaptiveFormats {
		q := format.Resolution
		if len(q) <= 0 {
			continue
		}

		switch vtype {
		case "url":
			if q == vq {
				return format.URL
			}

			prevData = format.URL

		case "itag":
			if q == vq {
				return getLatestURL(video.VideoID, format.Itag)
			}

			prevData = getLatestURL(video.VideoID, format.Itag)
		}
	}

	return prevData
}

// getLatestURL appends the latest_version query to the current client's host URL.
// For example: https://invidious.snopyta.org/latest_version?id=mWDOxRWcoPE&itag=22&local=true
func getLatestURL(id, itag string) string {
	host := GetClient().host

	idstr := "id=" + id
	itagstr := "&itag=" + itag

	return host + "/latest_version?" + idstr + itagstr + "&local=true"
}
