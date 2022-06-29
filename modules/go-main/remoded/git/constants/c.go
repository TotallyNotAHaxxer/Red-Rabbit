package GIT_OSINT_Constants

const (
	//https://api.github.com/users/{username}/orgs
	GITHUB_USERNAME_API_URL                                              = "https://api.github.com/users/%s"
	GITHUB_USERNAME_BASED_ORGANIZATION_URL                               = "https://api.github.com/orgs/%s"
	GITHUB_USERNAME_BASED_REPO_PROFILE_AND_INFORMATION                   = "https://api.github.com/repos/%s/%s"
	GITHUB_USERNAME_REPO_AND_PATH_BASED_PROFILE_REPO_PATH                = "https://api.github.com/repos/%s/%s/contents/%s"
	GITHUB_USERNAME_REPO_AND_CONTRIB_BASED_INFORMATION_URL_REPO          = "https://api.github.com/repos/%s/%s/contributors"
	GITHUB_USERNAME_REPO_AND_CONTRIB_BASED_INFORMATION_LANGUAGES_OF_REPO = "https://api.github.com/repos/%s/%s/languages"
	GITHUB_USERNAME_REPO_AND_STARGAZER_BASED_INFORMATION_OF_REPO         = "https://api.github.com/repos/%s/%s/stargazers"
	GITHUB_USERNAME_REPO_AND_FORKS_OF_REPO                               = "https://api.github.com/repos/%s/%s/forks"
	GITHUB_USERNAME_REPO_AND_RELEASES                                    = "https://api.github.com/repos/%s/%s/releases"
	GITHUB_USERNAME_ORGANIZATION_REPOS                                   = "https://api.github.com/orgs/%s/repos?per_page=100"
	GITHUB_ORGANIZATION_EVENTS                                           = "https://api.github.com/orgs/%s/events"
	GITHUB_ORGANIZATION_MEMBER                                           = "https://api.github.com/orgs/%s/public_members/%s"
	GITHUB_USERNAME_TOTAL_REPOS                                          = "https://api.github.com/users/%s/repos?per_page=100"
	GITHUB_USERNAME_TOTAL_GISTS                                          = "https://api.github.com/users/%s/gists"
	GITHUB_USERNAME_TOTAL_ORGS                                           = "https://api.github.com/users/%s/orgs"
	GITHUB_USERNAME_TOTAL_EVENTS                                         = "https://api.github.com/users/%s/events/public"
	GITHUB_USERNAME_TOTAL_SUBSCRIPTIONS                                  = "https://api.github.com/users/%s/subscriptions"
	GITHUB_USERNAME_TOTAL_FOLLOWERS                                      = "https://api.github.com/users/%s/followers?per_page=100"
	GITHUB_USERNAME_DOES_USER_FOLLOW_USER_A                              = "https://api.github.com/users/%s/following/%s"
	GITHUB_USERNAME_DOES_REPOA_EXIST                                     = "https://api.github.com/search/repositories?q=%s&per_page=100" // etc stuff all search is etc
	GITHUB_USERNAME_DOES_TOPIC_EXIST                                     = "https://api.github.com/search/topics?q=%s&per_page=100"
	GITHUB_LOOKUP_DOES_ISSUE_EXIST                                       = "https://api.github.com/search/issues?q=%s&per_page=100"
	GITHUB_LOOKUP_DOES_COMMIT_EXIST                                      = "https://api.github.com/search/commits?q=%s&per_page=100"
)

var Test_URLS = []string{
	"https://api.github.com/orgs/Google",
	"https://api.github.com/repos/ArkAngeL43/Red-Rabbit-V4",
	"https://api.github.com/repos/ArkAngeL43/Red-Rabbit-V4/contents/main",
	"https://api.github.com/repos/ArkAngeL43/Red-Rabbit-V4/contributors",
	"https://api.github.com/repos/ArkAngeL43/Red-Rabbit-V4/languages",
	"https://api.github.com/repos/ArkAngeL43/Red-Rabbit-V4/stargazers",
	"https://api.github.com/repos/ArkAngeL43/Red-Rabbit-V4/forks",
	"https://api.github.com/repos/ArkAngeL43/Red-Rabbit-V4/releases",
	"https://api.github.com/orgs/Google/repos?per_page=100",
	"https://api.github.com/orgs/Google/events",
	"https://api.github.com/orgs/Google/public_members/nyaxt",
	"https://api.github.com/users/ArkAngeL43/repos?per_page=100",
	"https://api.github.com/users/ArkAngeL43/gists",
	"https://api.github.com/users/ArkAngeL43/orgs",
	"https://api.github.com/users/ArkAngeL43/events/public",
	"https://api.github.com/users/ArkAngeL43/subscriptions",
	"https://api.github.com/users/ArkAngeL43/followers?per_page=100",
	"https://api.github.com/users/ArkAngeL43/following/nyaxt",
}
