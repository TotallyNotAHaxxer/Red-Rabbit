package OSINT_Engine

import (
	"regexp"
	"strings"

	maps "main/modules/go-main/SUPER-MAPS"
)

func Results() {
	for _, data := range Links_Crawled {
		if OPTIONS.Run_Twitter {
			if strings.Contains(data, "/i/events/") {
				maps.Twitter_map["event"] = append(maps.Twitter_map["event"], data)
			} else if strings.Contains(data, "/hashtag/") {
				maps.Twitter_map["hashtag"] = append(maps.Twitter_map["hashtag"], data)
			} else if strings.Contains(data, "/status/") {
				maps.Twitter_map["status"] = append(maps.Twitter_map["status"], data)
			} else {
				maps.Twitter_map["people"] = append(maps.Twitter_map["people"], data)
			}
			STDOUT(maps.Twitter_map)
		} else if OPTIONS.Run_Facebook {
			if strings.Contains(data, "/group/") {
				maps.Facebook_map["group"] = append(maps.Facebook_map["group"], data)
			} else if regexp.MustCompile("/videos/?$").MatchString(data) {
				maps.Facebook_map["video"] = append(maps.Facebook_map["video"], data)
			} else if regexp.MustCompile(`/posts/?[^\w\d]*`).MatchString(data) {
				maps.Facebook_map["post"] = append(maps.Facebook_map["post"], data)
			} else if regexp.MustCompile(`/photos/?[^\w\d]*`).MatchString(data) {
				maps.Facebook_map["photo"] = append(maps.Facebook_map["photo"], data)
			} else {
				maps.Facebook_map["people"] = append(maps.Facebook_map["people"], data)
			}
			STDOUT(maps.Facebook_map)
		} else if OPTIONS.Run_Linkedin {
			if strings.Contains(data, "/company/") {
				maps.LinkedIn_map["company"] = append(maps.LinkedIn_map["company"], data)
			} else if regexp.MustCompile(`/learning/`).MatchString(data) {
				maps.LinkedIn_map["learning"] = append(maps.LinkedIn_map["learning"], data)
			} else {
				maps.LinkedIn_map["people"] = append(maps.LinkedIn_map["people"], data)
			}
			STDOUT(maps.LinkedIn_map)
		} else {
			STDOUT(map[string][]string{"RAW - ": Links_Crawled})
		}
	}

}
