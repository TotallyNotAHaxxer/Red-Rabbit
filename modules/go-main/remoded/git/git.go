package GITHUB_OSINT_Utilities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	c "main/modg/colors"
	SUPER_git_Types "main/modules/go-main/SUPER_TYPES"
	SUPER_git "main/modules/go-main/remoded/git/constants"
	"net/http"
	"time"
)

func Get_USER_INFO(username string) {
	data := fmt.Sprintf(SUPER_git.GITHUB_USERNAME_API_URL, username)
	r, x := http.Get(data)
	if x != nil {
		log.Fatal(x)
	} else {
		defer r.Body.Close()
		b, e := ioutil.ReadAll(r.Body)
		if e != nil {
			log.Fatal(e)
		} else {
			var results SUPER_git_Types.Account_Information
			if err := json.Unmarshal(b, &results); err != nil {
				log.Fatal(err)
			}
			fmt.Println("------------------------------------------------------")
			fmt.Println(">> Avatar URL            \t| ", results.AvatarURL)
			fmt.Println(">> Recieved Events URL   \t| ", results.ReceivedEventsURL)
			fmt.Println(">> Followers URL         \t| ", results.FollowersURL)
			fmt.Println(">> Events URL            \t| ", results.EventsURL)
			fmt.Println(">> Following URL         \t| ", results.FollowingURL)
			fmt.Println(">> Organization URL      \t| ", results.OrganizationsURL)
			fmt.Println(">> Gists URL             \t| ", results.GistsURL)
			fmt.Println(">> Starred URL           \t| ", results.StarredURL)
			fmt.Println(">> Repos URL             \t| ", results.ReposURL)
			fmt.Println(">> Subscriptions URL     \t| ", results.SubscriptionsURL)
			fmt.Println(">> URL                   \t| ", results.URL)
			fmt.Println(">> User BIO              \t| ", results.Bio)
			fmt.Println(">> User Blog             \t| ", results.Blog)
			fmt.Println(">> User Company          \t| ", results.Company)
			fmt.Println(">> User Account created  \t| ", results.CreatedAt)
			fmt.Println(">> User Email            \t| ", results.Email)
			fmt.Println(">> User Follower Count   \t| ", results.Followers)
			fmt.Println(">> User following Count  \t| ", results.Following)
			fmt.Println(">> User Gravitar URL     \t| ", results.GravatarID)
			fmt.Println(">> User HTML URL         \t| ", results.HTMLURL)
			fmt.Println(">> User is hireable?     \t| ", results.Hireable)
			fmt.Println(">> User ID               \t| ", results.ID)
			fmt.Println(">> User Location         \t| ", results.Location)
			fmt.Println(">> User login name       \t| ", results.Login)
			fmt.Println(">> User name             \t| ", results.Name)
			fmt.Println(">> User Node ID          \t| ", results.NodeID)
			fmt.Println(">> User Public GISTS     \t| ", results.PublicGists)
			fmt.Println(">> User Public REPOS     \t| ", results.PublicRepos)
			fmt.Println(">> User Site Admin       \t| ", results.SiteAdmin)
			fmt.Println(">> User Twitter Username \t| ", results.TwitterUsername)
			fmt.Println(">> Type                  \t| ", results.Type)
		}
	}
}

func Get_users_followers(username string) {
	data_url := fmt.Sprintf("https://api.github.com/users/%s/followers?per_page=100", username)
	f, x := http.Get(data_url)
	if x != nil {
		log.Fatal(x)
	} else {
		defer f.Body.Close()
		b, x := ioutil.ReadAll(f.Body)
		if x != nil {
			log.Fatal(x)
		}
		var results SUPER_git_Types.Github_user_followers
		if err := json.Unmarshal(b, &results); err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(results); i++ {
			fmt.Printf("---------------------------------------------------\n")
			fmt.Println(c.HIGH_BLUE, "|Got user  \t| ", results[i].Login)
			fmt.Println(c.HIGH_PINK, "|>> Avatar URL            \t| ", results[i].AvatarURL)
			fmt.Println("|>> Recieved Events URL   \t| ", results[i].ReceivedEventsURL)
			fmt.Println("|>> Followers URL         \t| ", results[i].FollowersURL)
			fmt.Println("|>> Events URL            \t| ", results[i].EventsURL)
			fmt.Println("|>> Following URL         \t| ", results[i].FollowingURL)
			fmt.Println("|>> Organization URL      \t| ", results[i].OrganizationsURL)
			fmt.Println("|>> Gists URL             \t| ", results[i].GistsURL)
			fmt.Println("|>> Starred URL           \t| ", results[i].StarredURL)
			fmt.Println("|>> Repos URL             \t| ", results[i].ReposURL)
			fmt.Println("|>> Subscriptions URL     \t| ", results[i].SubscriptionsURL)
			fmt.Println("|>> URL                   \t| ", results[i].URL)
			fmt.Println("|>> User Gravitar URL     \t| ", results[i].GravatarID)
			fmt.Println("|>> User HTML URL         \t| ", results[i].HTMLURL)
			fmt.Println("|>> User ID               \t| ", results[i].ID)
			fmt.Println("|>> User login name       \t| ", results[i].Login)
			fmt.Println("|>> User Node ID          \t| ", results[i].NodeID)
			fmt.Println("|>> User Site Admin       \t| ", results[i].SiteAdmin)
			fmt.Println("|>> Type                  \t| ", results[i].Type)
			fmt.Printf("---------------------------------------------------")
			time.Sleep(90 * time.Millisecond)
		}
	}
}

func Get_All_User_Repos_Information(username string) {
	data_url := fmt.Sprintf(SUPER_git.GITHUB_USERNAME_TOTAL_REPOS, username)
	f, x := http.Get(data_url)
	if x != nil {
		log.Fatal(x)
	} else {
		defer f.Body.Close()
		b, x := ioutil.ReadAll(f.Body)
		if x != nil {
			log.Fatal(x)
		}
		var results SUPER_git_Types.Github_Username_Repos_OSINT
		if err := json.Unmarshal(b, &results); err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(results); i++ {
			time.Sleep(90 * time.Millisecond)
			fmt.Printf("--------------------------------------------------------------\n")
			fmt.Printf("\033[38;5;21m| Repo ID           \t|%v\n", results[i].ID)
			fmt.Printf("\033[38;5;21m| Repo Node ID      \t|%v\n", results[i].NodeID)
			fmt.Printf("\033[38;5;21m| Repo Name         \t|%s\n", results[i].Name)
			fmt.Printf("\033[38;5;21m| Repo Full Name    \t|%s\n", results[i].FullName)
			fmt.Printf("\033[38;5;21m| Repo Private?     \t|%v\n", results[i].Private)
			fmt.Println(c.HIGH_PINK, "--------------------------------------------------------------")
			fmt.Printf("| Repo HTML URL    \t|%s\t|\n", results[i].HTMLURL)
			fmt.Printf("| Repo Descript    \t|%s\t|\n", results[i].Description)
			fmt.Printf("| Repo Fork        \t|%v\t|\n", results[i].Fork)
			fmt.Printf("| Repo URL         \t|%s\t|\n", results[i].URL)
			fmt.Printf("| Repo Forks URL   \t|%s\t|\n", results[i].ForksURL)
			fmt.Printf("| Repo Colab URL   \t|%s\t|\n", results[i].CollaboratorsURL)
			fmt.Printf("| Repo Teams URL   \t|%s\t|\n", results[i].TeamsURL)
			fmt.Printf("| Repo Hooks URL   \t|%s\t|\n", results[i].HooksURL)
			fmt.Printf("| Repo Event URL   \t|%s\t|\n", results[i].EventsURL)
			fmt.Printf("| Repo Assign URL  \t|%s\t|\n", results[i].AssigneesURL)
			fmt.Printf("| Repo Branch URL  \t|%s\t|\n", results[i].BranchesURL)
			fmt.Printf("| Repo Tags   URL  \t|%s\t|\n", results[i].TagsURL)
			fmt.Printf("| Repo Blogs  URL  \t|%s\t|\n", results[i].BlobsURL)
			fmt.Printf("| Repo GitT URL    \t|%s\t|\n", results[i].GitTagsURL)
			fmt.Printf("| Repo Trees URL   \t|%s\t|\n", results[i].TreesURL)
			fmt.Printf("| Repo Status URL  \t|%s\t|\n", results[i].StatusesURL)
			fmt.Printf("| Repo LANG URL    \t|%s\t|\n", results[i].LanguagesURL)
			fmt.Printf("| Repo STARG URL   \t|%s\t|\n", results[i].StargazersURL)
			fmt.Printf("| Repo CONTRIB URL \t|%s\t|\n", results[i].ContributorsURL)
			fmt.Printf("| Repo SUBSCRB URL \t|%s\t|\n", results[i].SubscribersURL)
			fmt.Printf("| Repo SUBSRTN URL \t|%s\t|\n", results[i].SubscriptionURL)
			fmt.Printf("| Repo COMMITS URL \t|%s\t|\n", results[i].CommitsURL)
			fmt.Printf("| Repo GITCOMM URL \t|%s\t|\n", results[i].GitCommitsURL)
			fmt.Printf("| Repo COMMENT URL \t|%s\t|\n", results[i].CommentsURL)
			fmt.Printf("| Repo ISSUES URL  \t|%s\t|\n", results[i].IssueCommentURL)
			fmt.Printf("| Repo CONTENT URL \t|%s\t|\n", results[i].ContentsURL)
			fmt.Printf("| Repo Compare URL \t|%s\t|\n", results[i].CompareURL)
			fmt.Printf("| Repo Merges URL  \t|%s\t|\n", results[i].MergesURL)
			fmt.Printf("| Repo Archive URL \t|%s\t|\n", results[i].ArchiveURL)
			fmt.Printf("| Repo Downloa URL \t|%s\t|\n", results[i].DownloadsURL)
			fmt.Printf("| Repo ISSUES URL  \t|%s\t|\n", results[i].IssuesURL)
			fmt.Printf("| Repo PULLS URL   \t|%s\t|\n", results[i].PullsURL)
			fmt.Printf("| Repo MILES URL   \t|%s\t|\n", results[i].MilestonesURL)
			fmt.Printf("| Repo NOTIF URL   \t|%s\t|\n", results[i].NotificationsURL)
			fmt.Printf("| Repo LABEL URL   \t|%s\t|\n", results[i].LabelsURL)
			fmt.Printf("| Repo RELEA URL   \t|%s\t|\n", results[i].ReleasesURL)
			fmt.Printf("| Repo DEPLO URL   \t|%s\t|\n", results[i].DeploymentsURL)
			fmt.Printf("|--------------------|INFO|--------------------------------")
			fmt.Printf("| Repo Homepage \t|%s\t|\n", results[i].Homepage)
			fmt.Printf("| Repo Size     \t|%v\t|\n", results[i].Size)
			fmt.Printf("| Repo Stargaze \t|%v\t|\n", results[i].StargazersCount)
			fmt.Printf("| Repo Watchers \t|%v\t|\n", results[i].WatchersCount)
			fmt.Printf("| Repo Language \t|%s\t|\n", results[i].Language)
			fmt.Printf("| Repo Issues   \t|%v\t|\n", results[i].HasIssues)
			fmt.Printf("| Repo Projects \t|%v\t|\n", results[i].HasProjects)
			fmt.Printf("| Repo Download \t|%v\t|\n", results[i].HasDownloads)
			fmt.Printf("| Repo Has Wiki \t|%v\t|\n", results[i].HasWiki)
			fmt.Printf("| Repo Has Page \t|%v\t|\n", results[i].HasPages)
			fmt.Printf("| Repo Forks    \t|%v\t|\n", results[i].ForksCount)
			fmt.Printf("| Repo Mirror   \t|%s\t|\n", results[i].MirrorURL)
			fmt.Printf("| Repo Archived \t|%v\t|\n", results[i].Archived)
			fmt.Printf("| Repo Disabled \t|%v\t|\n", results[i].Disabled)
			fmt.Printf("| Repo Open Issu\t|%v\t|\n", results[i].OpenIssuesCount)
			fmt.Printf("| Repo License K\t|%s\t|\n", results[i].License.Key)
			fmt.Printf("| Repo License  \t|%s\t|\n", results[i].License.Name)
			fmt.Printf("| Repo License s\t|%s\t|\n", results[i].License.SpdxID)
			fmt.Printf("| License URL   \t|%s\t|\n", results[i].License.URL)
			fmt.Printf("| Repo NODE ID  \t|%s\t|\n", results[i].License.NodeID)
		}
	}
}

func Get_All_REPO_Stargazers_Information(username, repo string) {
	data_url := fmt.Sprintf(SUPER_git.GITHUB_USERNAME_REPO_AND_STARGAZER_BASED_INFORMATION_OF_REPO, username, repo)
	f, x := http.Get(data_url)
	if x != nil {
		log.Fatal(x)
	} else {
		defer f.Body.Close()
		b, x := ioutil.ReadAll(f.Body)
		if x != nil {
			log.Fatal(x)
		}
		var results SUPER_git_Types.Stargazers_Information
		if err := json.Unmarshal(b, &results); err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(results); i++ {
			fmt.Println(c.MAG)
			fmt.Print("##################################\n", c.BBLU)
			fmt.Printf("Stargazer #     \t|%v\t\n", i)
			fmt.Printf("Stargazer Login \t|%s\t\n", results[i].Login)
			fmt.Printf("Stargazer Admin \t|%v\t\n", results[i].SiteAdmin)
		}
	}
}

func Get_All_Contributor_Information(username, repo string) {
	data_url := fmt.Sprintf(SUPER_git.GITHUB_USERNAME_REPO_AND_CONTRIB_BASED_INFORMATION_URL_REPO, username, repo)
	f, x := http.Get(data_url)
	if x != nil {
		log.Fatal(x)
	} else {
		defer f.Body.Close()
		b, x := ioutil.ReadAll(f.Body)
		if x != nil {
			log.Fatal(x)
		}
		var results SUPER_git_Types.Contributors_Information
		if err := json.Unmarshal(b, &results); err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(results); i++ {
			fmt.Println(c.HIGH_BLUE)
			fmt.Printf(">> Contributor Number   |%v\n", i)
			fmt.Printf("---------------------------------\n")
			fmt.Println(c.RED)
			fmt.Printf("\033[0;31m>> User ID           \033[0;37m|%v\n", results[i].ID)
			fmt.Printf("\033[0;31m>> Login Name        \033[0;37m|%s\n", results[i].Login)
			fmt.Printf("\033[0;31m>> Node ID           \033[0;37m|%v\n", results[i].NodeID)
			fmt.Printf("\033[0;31m>> Avatar URL        \033[0;37m|%s\n", results[i].AvatarURL)
			fmt.Printf("\033[0;31m>> Contributions     \033[0;37m|%v\n", results[i].Contributions)
			fmt.Printf("\033[0;31m>> Events URL        \033[0;37m|%s\n", results[i].EventsURL)
			fmt.Printf("\033[0;31m>> Followers URL     \033[0;37m|%s\n", results[i].FollowersURL)
			fmt.Printf("\033[0;31m>> Following URL     \033[0;37m|%s\n", results[i].FollowingURL)
			fmt.Printf("\033[0;31m>> Gists URL         \033[0;37m|%s\n", results[i].GistsURL)
			fmt.Printf("\033[0;31m>> Gravatar ID       \033[0;37m|%v\n", results[i].GravatarID)
			fmt.Printf("\033[0;31m>> HTML URL          \033[0;37m|%s\n", results[i].HTMLURL)
			fmt.Printf("\033[0;31m>> Organizations URL \033[0;37m|%s\n", results[i].OrganizationsURL)
			fmt.Printf("\033[0;31m>> Received Events   \033[0;37m|%s\n", results[i].ReceivedEventsURL)
		}
	}
}

func Get_All_Organization_Repos(orgname string) {
	parser := fmt.Sprintf(SUPER_git.GITHUB_USERNAME_ORGANIZATION_REPOS, orgname)
	f, x := http.Get(parser)
	if x != nil {
		fmt.Println("<RR6> OSINT Module -> Requests: Could not make a request to the URL -> ", x)
	} else {
		defer f.Body.Close()
		b, x := ioutil.ReadAll(f.Body)
		if x != nil {
			fmt.Println("<RR6> OSINT Module -> I/O Utilities: Could not read the response body -> ", x)
		} else {
			var results SUPER_git_Types.Organization_Repos_information
			if err := json.Unmarshal(b, &results); err != nil {
				fmt.Println("<RR6> OSINT Module -> JSON module: Could not unmarshal json, got error -> ", x)
			} else {
				for i := 0; i < len(results); i++ {
					fmt.Println("______________________________________________")
					fmt.Printf("\033[0;37mOrganization          \033[0;34m: %s\n", orgname)
					fmt.Printf("\033[0;37mRepository            \033[0;34m: %s\n", results[i].FullName)
					fmt.Printf("\033[0;37mRepo Name             \033[0;34m: %s\n", results[i].Name)
					fmt.Printf("\033[0;37mRepo Owner            \033[0;34m: %s\n", results[i].Owner.Login)
					fmt.Printf("\033[0;37mRepo Owner ID         \033[0;34m: %v\n", results[i].Owner.ID)
					fmt.Printf("\033[0;37mRepo Owner Node ID    \033[0;34m: %v\n", results[i].Owner.NodeID)
					fmt.Printf("\033[0;37mRepo Owner Admin?     \033[0;34m: %v\n", results[i].Permissions.Admin)
					fmt.Printf("\033[0;37mRepo Owner Maintain?  \033[0;34m: %v\n", results[i].Permissions.Maintain)
					fmt.Printf("\033[0;37mRepo Owner Push?      \033[0;34m: %v\n", results[i].Permissions.Push)
					fmt.Printf("\033[0;37mRepo Owner Triage?    \033[0;34m: %v\n", results[i].Permissions.Triage)
					fmt.Printf("\033[0;37mRepo Owner Pull?      \033[0;34m: %v\n", results[i].Permissions.Pull)
					time.Sleep(90 * time.Millisecond)
				}
			}
		}
	}
}

func Get_ALl_Org_Events(organization string) {
	parser := fmt.Sprintf(SUPER_git.GITHUB_USERNAME_TOTAL_ORGS, organization)
	f, x := http.Get(parser)
	if x != nil {
		fmt.Println("<RR6> OSINT Module -> Requests: Could not make a request to the URL -> ", x)
	} else {
		defer f.Body.Close()
		b, x := ioutil.ReadAll(f.Body)
		var results SUPER_git_Types.Github_User_To_Organization
		if x != nil {
			fmt.Println("<RR6> OSINT Module -> I/O Utilities: Could not read the response body -> ", x)
		} else {
			if err := json.Unmarshal(b, &results); err != nil {
				fmt.Println("<RR6> OSINT Module -> JSON module: Could not unmarshal json, got error -> ", x)
			} else {
				for i := 0; i < len(results); i++ {
					fmt.Println("_________________________________________________")
					fmt.Println(c.WHT)
					fmt.Printf("\033[38;5;198mOrganization Name               \033[0;34m:%s\n", results[i].Login)
					fmt.Printf("\033[38;5;198mOrganization ID                 \033[0;34m:%v\n", results[i].ID)
					fmt.Printf("\033[38;5;198mOrganization Node ID            \033[0;34m:%s\n", results[i].NodeID)
					fmt.Printf("\033[38;5;198mOrganization URL                \033[0;34m:%s\n", results[i].URL)
					fmt.Printf("\033[38;5;198mOrganization Repos URL          \033[0;34m:%s\n", results[i].ReposURL)
					fmt.Printf("\033[38;5;198mOrganization Events URL         \033[0;34m:%s\n", results[i].EventsURL)
					fmt.Printf("\033[38;5;198mOrganization Hooks URL          \033[0;34m:%s\n", results[i].HooksURL)
					fmt.Printf("\033[38;5;198mOrganization Issues URL         \033[0;34m:%s\n", results[i].IssuesURL)
					fmt.Printf("\033[38;5;198mOrganization Members URL        \033[0;34m:%s\n", results[i].MembersURL)
					fmt.Printf("\033[38;5;198mOrganization Public URL         \033[0;34m:%s\n", results[i].PublicMembersURL)
					fmt.Printf("\033[38;5;198mOrganization Avatar URL         \033[0;34m:%s\n", results[i].AvatarURL)
					fmt.Printf("\033[38;5;198mOrganization Description        \033[0;34m:%s\n", results[i].Description)
				}
			}
		}
	}
}

func Get_All_Events_from_user(username string) {
	parser := fmt.Sprintf(SUPER_git.GITHUB_USERNAME_TOTAL_EVENTS, username)
	f, x := http.Get(parser)
	if x != nil {
		fmt.Println("<RR6> OSINT Module -> Requests: Could not make a request to the URL -> ", x)
	} else {
		defer f.Body.Close()
		b, x := ioutil.ReadAll(f.Body)
		var results SUPER_git_Types.User_Events
		if x != nil {
			fmt.Println("<RR6> OSINT Module -> I/O Utilities: Could not read the response body -> ", x)
		} else {
			if err := json.Unmarshal(b, &results); err != nil {
				fmt.Println("<RR6> OSINT Module -> JSON module: Could not unmarshal json, got error -> ", x)
			} else {
				for i := 0; i < len(results); i++ {
					fmt.Print("------------------------------------------\n")
					fmt.Printf("\033[0;31mEvent Type          \033[0;34m: %s\n", results[i].Type)
					fmt.Printf("\033[0;31mEvent ID            \033[0;34m: %s\n", results[i].ID)
					fmt.Printf("\033[0;31mEvent Host          \033[0;34m: %s\n", results[i].Actor.Login)
					fmt.Printf("\033[0;31mEvent Host Login    \033[0;34m: %s\n", results[i].Actor.DisplayLogin)
					fmt.Printf("\033[0;31mEvent repo name     \033[0;34m: %s\n", results[i].Repo.Name)
					fmt.Printf("\033[0;31mEvent Created At    \033[0;34m: %s\n", results[i].CreatedAt)
				}

			}
		}
	}
}

func Get_All_Events_from_organization(username string) {
	parser := fmt.Sprintf(SUPER_git.GITHUB_ORGANIZATION_EVENTS, username)
	f, x := http.Get(parser)
	if x != nil {
		fmt.Println("<RR6> OSINT Module -> Requests: Could not make a request to the URL -> ", x)
	} else {
		defer f.Body.Close()
		b, x := ioutil.ReadAll(f.Body)
		var results SUPER_git_Types.Github_Org_Events
		if x != nil {
			fmt.Println("<RR6> OSINT Module -> I/O Utilities: Could not read the response body -> ", x)
		} else {
			if err := json.Unmarshal(b, &results); err != nil {
				fmt.Println("<RR6> OSINT Module -> JSON module: Could not unmarshal json, got error -> ", x)
			} else {
				for i := 0; i < len(results); i++ {
					fmt.Print("------------------------------------------\n")
					fmt.Printf("\033[0;31mEvent Type          \033[0;34m: %s\n", results[i].Type)
					fmt.Printf("\033[0;31mEvent ID            \033[0;34m: %s\n", results[i].ID)
					fmt.Printf("\033[0;31mEvent Host          \033[0;34m: %s\n", results[i].Actor.Login)
					fmt.Printf("\033[0;31mEvent Host Login    \033[0;34m: %s\n", results[i].Actor.DisplayLogin)
					fmt.Printf("\033[0;31mEvent repo name     \033[0;34m: %s\n", results[i].Repo.Name)
					fmt.Printf("\033[0;31mEvent repo ID       \033[0;34m: %v\n", results[i].Repo.ID)
					fmt.Printf("\033[0;31mEvent repo URL      \033[0;34m: %s\n", results[i].Repo.URL)
					fmt.Printf("\033[0;31mEvent Created At    \033[0;34m: %s\n", results[i].CreatedAt)
					fmt.Printf("\033[0;31mEvent Org Gravatar  \033[0;34m: %s\n", results[i].Org.GravatarID)
					fmt.Printf("\033[0;31mEvent Org Login     \033[0;34m: %s\n", results[i].Org.Login)
					fmt.Printf("\033[0;31mEvent Action        \033[0;34m: %s\n", results[i].Payload.Action)
					fmt.Printf("\033[0;31mEvent Public?       \033[0;34m: %v\n", results[i].Public)
				}
			}
		}
	}
}
