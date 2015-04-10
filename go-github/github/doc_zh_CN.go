// Copyright The go-github Authors. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// Package github provides a client for using the GitHub API.
//
// Construct a new GitHub client, then use the various services on the client to
// access different parts of the GitHub API. For example:
//
//	client := github.NewClient(nil)
//
//	// list all organizations for user "willnorris"
//	orgs, _, err := client.Organizations.List("willnorris", nil)
//
// Set optional parameters for an API method by passing an Options object.
//
//	// list recently updated repositories for org "github"
//	opt := &github.RepositoryListByOrgOptions{Sort: "updated"}
//	repos, _, err := client.Repositories.ListByOrg("github", opt)
//
// The services of a client divide the API into logical chunks and correspond to
// the structure of the GitHub API documentation at
// http://developer.github.com/v3/.
//
//
// Authentication
//
// The go-github library does not directly handle authentication. Instead, when
// creating a new client, pass an http.Client that can handle authentication for
// you. The easiest and recommended way to do this is using the goauth2 library,
// but you can always use any other library that provides an http.Client. If you
// have an OAuth2 access token (for example, a personal API token), you can use it
// with the goauth2 using:
//
//	import "code.google.com/p/goauth2/oauth"
//
//	// simple OAuth transport if you already have an access token;
//	// see goauth2 library for full usage
//	t := &oauth.Transport{
//		Token: &oauth.Token{AccessToken: "..."},
//	}
//
//	client := github.NewClient(t.Client())
//
//	// list all repositories for the authenticated user
//	repos, _, err := client.Repositories.List("", nil)
//
// Note that when using an authenticated Client, all calls made by the client will
// include the specified OAuth token. Therefore, authenticated clients should
// almost never be shared between different users.
//
//
// Rate Limiting
//
// GitHub imposes a rate limit on all API clients. Unauthenticated clients are
// limited to 60 requests per hour, while authenticated clients can make up to
// 5,000 requests per hour. To receive the higher rate limit when making calls that
// are not issued on behalf of a user, use the UnauthenticatedRateLimitedTransport.
//
// The Rate field on a client tracks the rate limit information based on the most
// recent API call. This is updated on every call, but may be out of date if it's
// been some time since the last API call and other clients have made subsequent
// requests since then. You can always call RateLimit() directly to get the most
// up-to-date rate limit data for the client.
//
// Learn more about GitHub rate limiting at
// http://developer.github.com/v3/#rate-limiting.
//
//
// Conditional Requests
//
// The GitHub API has good support for conditional requests which will help prevent
// you from burning through your rate limit, as well as help speed up your
// application. go-github does not handle conditional requests directly, but is
// instead designed to work with a caching http.Transport. We recommend using
// https://github.com/gregjones/httpcache, which can be used in conjuction with
// https://github.com/sourcegraph/apiproxy to provide additional flexibility and
// control of caching rules.
//
// Learn more about GitHub conditional requests at
// https://developer.github.com/v3/#conditional-requests.
//
//
// Creating and Updating Resources
//
// All structs for GitHub resources use pointer values for all non-repeated fields.
// This allows distinguishing between unset fields and those set to a zero-value.
// Helper functions have been provided to easily create these pointers for string,
// bool, and int values. For example:
//
//	// create a new private repository named "foo"
//	repo := &github.Repository{
//		Name:    github.String("foo"),
//		Private: github.Bool(true),
//	}
//	client.Repositories.Create("", repo)
//
// Users who have worked with protocol buffers should find this pattern familiar.
//
//
// Pagination
//
// All requests for resource collections (repos, pull requests, issues, etc)
// support pagination. Pagination options are described in the ListOptions struct
// and passed to the list methods directly or as an embedded type of a more
// specific list options struct (for example PullRequestListOptions). Pages
// information is available via Response struct.
//
//	opt := &github.RepositoryListByOrgOptions{
//		ListOptions: github.ListOptions{PerPage: 10},
//	}
//	// get all pages of results
//	var allRepos []github.Repository
//	for {
//		repos, resp, err := client.Repositories.ListByOrg("github", opt)
//		if err != nil {
//			return err
//		}
//		allRepos = append(allRepos, repos...)
//		if resp.NextPage == 0 {
//			break
//		}
//		opt.ListOptions.Page = resp.NextPage
//	}

// Package github 提供客户端使用 GitHub API.
//
// 构造一个新的 GitHub 客户端, 然后使用客户端的各种服务去
// 访问不同的 GitHub API. 例子:
//
//	client := github.NewClient(nil)
//
//	// 罗列用户 "willnorris" 所有的组织
//	orgs, _, err := client.Organizations.List("willnorris", nil)
//
// 通过 Options 对象为 API 方法设置可选参数.
//
//	// 罗列 "github" 组织最近更新的仓库
//	opt := &github.RepositoryListByOrgOptions{Sort: "updated"}
//	repos, _, err := client.Repositories.ListByOrg("github", opt)
//
// 客户端的服务以 API 分割成块, 并对应 GitHub API 结构文档
// http://developer.github.com/v3/.
//
//
// 授权认证
//
// go-github 不直接处理授权认证. 作为替代, 创建新的 Client 时,
// 通过一个 http.Client 来为你处理授权认证.
// 要做到这一点, 最简单的和推荐的方法是使用 goauth2 库,
// 你可以随时使用其他任何一个提供了 http.Client 的库. If you
// 如果你有一个 OAuth2 访问 token (比如, a personal API token),
// 你可以与 goauth2 一起使用它:
//
//	import "code.google.com/p/goauth2/oauth"
//
//	// 简单示例 OAuth transport, 如果你已经有了一个访问 token;
//	// 详见 goauth2 库使用
//	t := &oauth.Transport{
//		Token: &oauth.Token{AccessToken: "..."},
//	}
//
//	client := github.NewClient(t.Client())
//
//	// 罗列已授权用户所有的组织
//	repos, _, err := client.Repositories.List("", nil)
//
// Note: 当使用已认证的 Client 时, 所有客户端的调用都会
// 包含特定的 OAuth token. 因此, 不同用户的认证客户端不能共享.
//
//
// 频次限制
//
// GitHub 对所有客户端 API 实施了频次限制.
// 未认证客户端限制 60 次请求/小时, 已认证客户端限制 5,000 次请求/小时.
// 使用 UnauthenticatedRateLimitedTransport 可让授权用户获得较高的请求频率.
//
// 客户端的 Rate 字段基于最近的 API 调用跟踪频次限制. 每次调用都它会更新,
// 但可能不准, 如果自上次 API 调用一段时间内其它客户做出后续请求.
// 您可以随时调用 RateLimit() 直接得到最新频率限额.
//
// 学习 GitHub 频次限制请至
// http://developer.github.com/v3/#rate-limiting.
//
//
// 条件请求
//
// GitHub API 对条件请求有良好的支持, 助于你防止过快消耗频率限额以及加速应用.
// go-github 不直接处理条件请求, 不过具有缓存设计的 http.Transport 可以工作.
// 我们推荐
// https://github.com/gregjones/httpcache
// https://github.com/sourcegraph/apiproxy
// 一起使用可提供灵活的控制缓存规则.
//
// 学习 GitHub 条件请求请至
// https://developer.github.com/v3/#conditional-requests.
//
//
// 创建和更新资源
//
// 所有 GitHub 资源结构体的字段使用指针值. 这让未设置字段和零值字段加以区分.
// 辅助函数提供轻松创建 string, bool, 和 int 值指针. 例子:
//
//	// 新建名为 "foo" 的私有仓库
//	repo := &github.Repository{
//		Name:    github.String("foo"),
//		Private: github.Bool(true),
//	}
//	client.Repositories.Create("", repo)
//
// 有制作 protocol buffers 的用户会发现这个熟悉的模式.
//
//
// 分页
//
// 所有集合资源请求 (repos, pull requests, issues, etc) 都支持分页.
// 分页选项在 ListOptions 结构中声明, 直接传递给列表方法,
// 或嵌入到特定选项结构中 (例如 PullRequestListOptions).
// Response 结构体中有可用的分页信息.
//
//	opt := &github.RepositoryListByOrgOptions{
//		ListOptions: github.ListOptions{PerPage: 10},
//	}
//	// get all pages of results
//	var allRepos []github.Repository
//	for {
//		repos, resp, err := client.Repositories.ListByOrg("github", opt)
//		if err != nil {
//			return err
//		}
//		allRepos = append(allRepos, repos...)
//		if resp.NextPage == 0 {
//			break
//		}
//		opt.ListOptions.Page = resp.NextPage
//	}
package github

const (
	// Tarball specifies an archive in gzipped tar format.

	// Tarball 指定采用 gzip 压缩 tar 文档格式.
	Tarball archiveFormat = "tarball"

	// Zipball specifies an archive in zip format.

	// Zipball 指定 zip 压缩文档格式.
	Zipball archiveFormat = "zipball"
)

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.

// Bool 辅助函数分配一个保存 v 的 bool 值并返回它的指针.
func Bool(v bool) *bool

// CheckResponse checks the API response for errors, and returns them if
// present.  A response is considered an error if it has a status code outside
// the 200 range.  API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse.  Any other
// response body will be silently ignored.

// CheckResponse 检查 API 响应错误, 如果存在返回它们.
// 状态码为 200 以外的响应被为认为是个错误响应.
// API 错误响应预计要么没有响应体, 或者是含有 ErrorResponse 的 JSON.
// 任何其它响应身体会忽略掉.
func CheckResponse(r *http.Response) error

// Int is a helper routine that allocates a new int32 value
// to store v and returns a pointer to it, but unlike Int32
// its argument value is an int.

// Int 辅助函数分配一个保存 v 的 Int32 值并返回它的指针,
// 参数类型是 int 而不是 Int32.
func Int(v int) *int

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.

// 辅助函数分配一个保存 v 的 string 值并返回它的指针.
func String(v string) *string

// Stringify attempts to create a reasonable string representation of types in
// the GitHub library.  It does things like resolve pointers to their values
// and omits struct fields with nil values.

// Stringify 尝试创建一个字符串表示的 GitHub 库类型.
// 它解决类似结构体指针字段值为 nil 的问题.
func Stringify(message interface{}) string

// APIMeta represents metadata about the GitHub API.

// APIMeta 表示 GitHub API 元数据.
type APIMeta struct {
	// An Array of IP addresses in CIDR format specifying the addresses
	// that incoming service hooks will originate from on GitHub.com.

	// CIDR 格式 IP 地址数组, 指定源自 GitHub.com 的呼入服务钩子地址.
	Hooks []string `json:"hooks,omitempty"`

	// An Array of IP addresses in CIDR format specifying the Git servers
	// for GitHub.com.

	// CIDR 格式 IP 地址数组, 为 GitHub.com 指定 Git 服务器.
	Git []string `json:"git,omitempty"`

	// Whether authentication with username and password is supported.
	// (GitHub Enterprise instances using CAS or OAuth for authentication
	// will return false. Features like Basic Authentication with a
	// username and password, sudo mode, and two-factor authentication are
	// not supported on these servers.)

	// 无论认证是否支持含用户名和密码.
	// (使用 CAS 或 OAuth 认证的 GitHub 企业实例将返回 false.
	// 特性如使用用户名和密码的基本认证, sudo 模式, 双因素身份验证不支持此服务.)
	VerifiablePasswordAuthentication *bool `json:"verifiable_password_authentication,omitempty"`
}

// ActivityListStarredOptions specifies the optional parameters to the
// ActivityService.ListStarred method.

// ActivityListStarredOptions 定义 ActivityService.ListStarred 可选参数.
type ActivityListStarredOptions struct {
	// How to sort the repository list.  Possible values are: created, updated,
	// pushed, full_name.  Default is "full_name".

	// 仓库列表排序方式. 可选值: created, updated, pushed, full_name.
	// 缺省为 "full_name".
	Sort string `url:"sort,omitempty"`

	// Direction in which to sort repositories.  Possible values are: asc, desc.
	// Default is "asc" when sort is "full_name", otherwise default is "desc".

	// 仓库排序方向. 可选值: asc, desc.
	// 缺省时, 当 Sort 为 "full_name" 时为 "asc", 其它为 "desc".
	Direction string `url:"direction,omitempty"`

	ListOptions
}

// ActivityService handles communication with the activity related methods of the
// GitHub API.
//
// GitHub API docs: http://developer.github.com/v3/activity/

// ActivityService 处理 GitHub API 中与活动相关的通信方法.
//
// GitHub API docs: http://developer.github.com/v3/activity/
type ActivityService struct {
	// contains filtered or unexported fields
}

// DeleteRepositorySubscription deletes the subscription for the specified
// repository for the authenticated user.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/watching/#delete-a-repository-subscription

// DeleteRepositorySubscription 删除认证用户指定的仓库订阅.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/watching/#delete-a-repository-subscription
func (s *ActivityService) DeleteRepositorySubscription(owner, repo string) (*Response, error)

// DeleteThreadSubscription deletes the subscription for the specified thread for
// the authenticated user.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/notifications/#delete-a-thread-subscription

// DeleteThreadSubscription 删除认证用户指定的主题订阅.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/notifications/#delete-a-thread-subscription
func (s *ActivityService) DeleteThreadSubscription(id string) (*Response, error)

// GetRepositorySubscription returns the subscription for the specified repository
// for the authenticated user. If the authenticated user is not watching the
// repository, a nil Subscription is returned.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/watching/#get-a-repository-subscription

// GetRepositorySubscription 返回认证用户指定的仓库订阅.
// 如果用户没有收听仓库, 返回值 Subscription 为 nil.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/watching/#get-a-repository-subscription
func (s *ActivityService) GetRepositorySubscription(owner, repo string) (*Subscription, *Response, error)

// GetThread gets the specified notification thread.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/notifications/#view-a-single-thread

// GetThread 返回指定的主题订阅通知.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/notifications/#view-a-single-thread
func (s *ActivityService) GetThread(id string) (*Notification, *Response, error)

// GetThreadSubscription checks to see if the authenticated user is subscribed to a
// thread.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/notifications/#get-a-thread-subscription

// GetThreadSubscription 检查查看认证用户所订阅主题.
//
// GitHub API 文档:
// https://developer.github.com/v3/activity/notifications/#get-a-thread-subscription
func (s *ActivityService) GetThreadSubscription(id string) (*Subscription, *Response, error)

// IsStarred checks if a repository is starred by authenticated user.
//
// GitHub API docs:
// https://developer.github.com/v3/activity/starring/#check-if-you-are-starring-a-repository

// IsStarred 检查被认证用户标星的仓库.
//
// GitHub API 文档:
// https://developer.github.com/v3/activity/starring/#check-if-you-are-starring-a-repository
func (s *ActivityService) IsStarred(owner, repo string) (bool, *Response, error)

// ListEvents drinks from the firehose of all public events across GitHub.
//
// GitHub API docs:
// http://developer.github.com/v3/activity/events/#list-public-events

// ListEvents 罗列(喝掉)所有来自 GitHub 公共事件流水.
//
// GitHub API 文档:
// http://developer.github.com/v3/activity/events/#list-public-events
func (s *ActivityService) ListEvents(opt *ListOptions) ([]Event, *Response, error)

// ListEventsForOrganization lists public events for an organization.
//
// GitHub API docs:
// http://developer.github.com/v3/activity/events/#list-public-events-for-an-organization

// ListEventsForOrganization 罗列某组织的公共事件.
//
// GitHub API 文档:
// http://developer.github.com/v3/activity/events/#list-public-events-for-an-organization
func (s *ActivityService) ListEventsForOrganization(org string, opt *ListOptions) ([]Event, *Response, error)

// ListEventsForRepoNetwork lists public events for a network of repositories.
//
// GitHub API docs:
// http://developer.github.com/v3/activity/events/#list-public-events-for-a-network-of-repositories

// ListEventsForRepoNetwork 罗列某仓库的网络公共事件.
//
// GitHub API 文档:
// http://developer.github.com/v3/activity/events/#list-public-events-for-a-network-of-repositories
func (s *ActivityService) ListEventsForRepoNetwork(owner, repo string, opt *ListOptions) ([]Event, *Response, error)

// ListEventsPerformedByUser lists the events performed by a user. If publicOnly is
// true, only public events will be returned.
//
// GitHub API docs:
// http://developer.github.com/v3/activity/events/#list-events-performed-by-a-user

// ListEventsPerformedByUser 罗列某用户执行的事件. 如果 publicOnly 为 true,
// 只返回公共事件.
//
// GitHub API 文档:
// http://developer.github.com/v3/activity/events/#list-events-performed-by-a-user
func (s *ActivityService) ListEventsPerformedByUser(user string, publicOnly bool, opt *ListOptions) ([]Event, *Response, error)

// ListEventsRecievedByUser lists the events recieved by a user. If publicOnly is
// true, only public events will be returned.
//
// GitHub API docs:
// http://developer.github.com/v3/activity/events/#list-events-that-a-user-has-received

// ListEventsRecievedByUser 罗列某用户收到的事件. 如果 publicOnly 为 true,
// 只返回公共事件.
//
// GitHub API 文档:
// http://developer.github.com/v3/activity/events/#list-events-that-a-user-has-received
func (s *ActivityService) ListEventsRecievedByUser(user string, publicOnly bool, opt *ListOptions) ([]Event, *Response, error)

// ListIssueEventsForRepository lists issue events for a repository.
//
// GitHub API docs:
// http://developer.github.com/v3/activity/events/#list-issue-events-for-a-repository

// ListIssueEventsForRepository 罗列某仓库的 issue 事件.
//
// GitHub API 文档:
// http://developer.github.com/v3/activity/events/#list-issue-events-for-a-repository
func (s *ActivityService) ListIssueEventsForRepository(owner, repo string, opt *ListOptions) ([]Event, *Response, error)

// ListNotifications lists all notifications for the authenticated user.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/notifications/#list-your-notifications

// ListNotifications 罗列授权用户所有通知.
//
// GitHub API 文档:
// https://developer.github.com/v3/activity/notifications/#list-your-notifications
func (s *ActivityService) ListNotifications(opt *NotificationListOptions) ([]Notification, *Response, error)

// ListRepositoryEvents lists events for a repository.
//
// GitHub API docs:
// http://developer.github.com/v3/activity/events/#list-repository-events

// ListRepositoryEvents 罗列某仓库事件.
//
// GitHub API 文档:
// http://developer.github.com/v3/activity/events/#list-repository-events
func (s *ActivityService) ListRepositoryEvents(owner, repo string, opt *ListOptions) ([]Event, *Response, error)

// ListRepositoryNotifications lists all notifications in a given repository for
// the authenticated user.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/notifications/#list-your-notifications-in-a-repository

// ListRepositoryNotifications 罗列授权用户的给定仓库所有通知.
//
// GitHub API 文档:
// https://developer.github.com/v3/activity/notifications/#list-your-notifications-in-a-repository
func (s *ActivityService) ListRepositoryNotifications(owner, repo string, opt *NotificationListOptions) ([]Notification, *Response, error)

// ListStargazers lists people who have starred the specified repo.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/starring/#list-stargazers

// ListStargazers 罗列给指定仓库加星标的人.
//
// GitHub API 文档:
// https://developer.github.com/v3/activity/starring/#list-stargazers
func (s *ActivityService) ListStargazers(owner, repo string, opt *ListOptions) ([]User, *Response, error)

// ListStarred lists all the repos starred by a user. Passing the empty string will
// list the starred repositories for the authenticated user.
//
// GitHub API docs:
// http://developer.github.com/v3/activity/starring/#list-repositories-being-starred

// ListStarred 罗列某用户加星标的所有仓库. 传递空字符串将罗列授权用户加星标的仓库.
//
// GitHub API 文档:
// http://developer.github.com/v3/activity/starring/#list-repositories-being-starred
func (s *ActivityService) ListStarred(user string, opt *ActivityListStarredOptions) ([]Repository, *Response, error)

// ListUserEventsForOrganization provides the user’s organization dashboard. You
// must be authenticated as the user to view this.
//
// GitHub API docs:
// http://developer.github.com/v3/activity/events/#list-events-for-an-organization

// ListUserEventsForOrganization 提供用户的组织 dashboard. 你必须被该用户授权才能看到它.
//
// GitHub API 文档:
// http://developer.github.com/v3/activity/events/#list-events-for-an-organization
func (s *ActivityService) ListUserEventsForOrganization(org, user string, opt *ListOptions) ([]Event, *Response, error)

// ListWatched lists the repositories the specified user is watching. Passing the
// empty string will fetch watched repos for the authenticated user.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/watching/#list-repositories-being-watched

// ListWatched 罗列指定用户监视的仓库. 传递空字符串将获取授权用户监视仓库.
//
// GitHub API 文档:
// https://developer.github.com/v3/activity/watching/#list-repositories-being-watched
func (s *ActivityService) ListWatched(user string) ([]Repository, *Response, error)

// ListWatchers lists watchers of a particular repo.
//
// GitHub API Docs: http://developer.github.com/v3/activity/watching/#list-watchers

// ListWatchers 罗列特定仓库的监视者.
//
// GitHub API 文档:
// http://developer.github.com/v3/activity/watching/#list-watchers
func (s *ActivityService) ListWatchers(owner, repo string, opt *ListOptions) ([]User, *Response, error)

// MarkNotificationsRead marks all notifications up to lastRead as read.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/notifications/#mark-as-read

// MarkNotificationsRead 标记所有上至 lastRead 时间的通知为已读.
//
// GitHub API 文档:
// https://developer.github.com/v3/activity/notifications/#mark-as-read
func (s *ActivityService) MarkNotificationsRead(lastRead time.Time) (*Response, error)

// MarkRepositoryNotificationsRead marks all notifications up to lastRead in the
// specified repository as read.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/notifications/#mark-notifications-as-read-in-a-repository

// MarkRepositoryNotificationsRead 标记指定仓库中所有上至 lastRead 时间的通知为已读.
//
// GitHub API 文档:
// https://developer.github.com/v3/activity/notifications/#mark-notifications-as-read-in-a-repository
func (s *ActivityService) MarkRepositoryNotificationsRead(owner, repo string, lastRead time.Time) (*Response, error)

// MarkThreadRead marks the specified thread as read.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/notifications/#mark-a-thread-as-read

// MarkThreadRead 标记指定主题为已读.
//
// GitHub API 文档:
// https://developer.github.com/v3/activity/notifications/#mark-a-thread-as-read
func (s *ActivityService) MarkThreadRead(id string) (*Response, error)

// SetRepositorySubscription sets the subscription for the specified repository for
// the authenticated user.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/watching/#set-a-repository-subscription

// SetRepositorySubscription 设置授权用户下指定仓库的订阅.
//
// GitHub API 文档:
// https://developer.github.com/v3/activity/watching/#set-a-repository-subscription
func (s *ActivityService) SetRepositorySubscription(owner, repo string, subscription *Subscription) (*Subscription, *Response, error)

// SetThreadSubscription sets the subscription for the specified thread for the
// authenticated user.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/notifications/#set-a-thread-subscription

// SetThreadSubscription 设置授权用户下指定主题的订阅.
//
// GitHub API Docs:
// https://developer.github.com/v3/activity/notifications/#set-a-thread-subscription
func (s *ActivityService) SetThreadSubscription(id string, subscription *Subscription) (*Subscription, *Response, error)

// Star a repository as the authenticated user.
//
// GitHub API docs:
// https://developer.github.com/v3/activity/starring/#star-a-repository

// Star 星标授权用户的一个仓库.
//
// GitHub API 文档:
// https://developer.github.com/v3/activity/starring/#star-a-repository
func (s *ActivityService) Star(owner, repo string) (*Response, error)

// Unstar a repository as the authenticated user.
//
// GitHub API docs:
// https://developer.github.com/v3/activity/starring/#unstar-a-repository

// Unstar 取消授权用户一个仓库的星标.
//
// GitHub API 文档:
// https://developer.github.com/v3/activity/starring/#unstar-a-repository
func (s *ActivityService) Unstar(owner, repo string) (*Response, error)

// Blob represents a blob object.

// Blob 表示一个 blob 对象.
type Blob struct {
	Content  *string `json:"content,omitempty"`
	Encoding *string `json:"encoding,omitempty"`
	SHA      *string `json:"sha,omitempty"`
	Size     *int    `json:"size,omitempty"`
	URL      *string `json:"url,omitempty"`
}

// Branch represents a repository branch

// Branch 表示一个仓库分支.
type Branch struct {
	Name   *string `json:"name,omitempty"`
	Commit *Commit `json:"commit,omitempty"`
}

// A Client manages communication with the GitHub API.

// Client 管理一个 GitHub API 通信.
type Client struct {

	// Base URL for API requests.  Defaults to the public GitHub API, but can be
	// set to a domain endpoint to use with GitHub Enterprise.  BaseURL should
	// always be specified with a trailing slash.

	// API 请求的基本 URL. 默认为 GitHub 公共 API, 但可以设置一个使用
	// GitHub 企业的域节点. 规定 BaseURL 总是以反斜线结尾.
	BaseURL *url.URL

	// Base URL for uploading files.

	// 上传文件的基本 RUL.
	UploadURL *url.URL

	// User agent used when communicating with the GitHub API.

	// 用户角色用于 GitHub API 通讯时.
	UserAgent string

	// Rate specifies the current rate limit for the client as determined by the
	// most recent API call.  If the client is used in a multi-user application,
	// this rate may not always be up-to-date.  Call RateLimit() to check the
	// current rate.

	// Rate 规定默认的客户端频次限制, 被最新的 API 调用确定.
	Rate Rate

	// Services used for talking to different parts of the GitHub API.

	// 不同 GitHub API 服务所涉及的部分.
	Activity      *ActivityService
	Gists         *GistsService
	Git           *GitService
	Gitignores    *GitignoresService
	Issues        *IssuesService
	Organizations *OrganizationsService
	PullRequests  *PullRequestsService
	Repositories  *RepositoriesService
	Search        *SearchService
	Users         *UsersService
	// contains filtered or unexported fields
}

// NewClient returns a new GitHub API client. If a nil httpClient is provided,
// http.DefaultClient will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication for
// you (such as that provided by the goauth2 library).

// NewClient 返回一个新的 GitHub API 客户端. 如果 HttpClient 为 nil,
// http.DefaultClient 将被使用. 要使用 API 的方法需要
// 认证, 提供的 http.Client 将用于验证你的身份
// ( 比如 goauth2 库提供的 )
func NewClient(httpClient *http.Client) *Client

// APIMeta returns information about GitHub.com, the service. Or, if you access
// this endpoint on your organization’s GitHub Enterprise installation, this
// endpoint provides information about that installation.
//
// GitHub API docs: https://developer.github.com/v3/meta/

// APIMeta 返回 GitHub.com 服务信息. 或者你安装的 GitHub 企业组织存取节点,
// 该安装的节点提供那些信息.
//
// GitHub API 文档: https://developer.github.com/v3/meta/
func (c *Client) APIMeta() (*APIMeta, *Response, error)

// Do sends an API request and returns the API response. The API response is JSON
// decoded and stored in the value pointed to by v, or returned as an error if an
// API error has occurred. If v implements the io.Writer interface, the raw
// response body will be written to v, without attempting to first decode it.

// Do 发送 API 请求并返回 API 响应. 该 API 响应为 JSON, 解码并按 v 指向的值排序,
// 如果发生 API 错误, 返回一个错误. 如果 v 实现了 io.Writer 接口,
// 未曾尝试解码的原始响应体被写入 w.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error)

// ListEmojis returns the emojis available to use on GitHub.
//
// GitHub API docs: https://developer.github.com/v3/emojis/

// ListEmojis 返回用于 GitHub 的有效表情符号.
//
// GitHub API 文档: https://developer.github.com/v3/emojis/
func (c *Client) ListEmojis() (map[string]string, *Response, error)

// Markdown renders an arbitrary Markdown document.
//
// GitHub API docs: https://developer.github.com/v3/markdown/

// Markdown 渲染一个 Markdown 文档.
//
// GitHub API docs: https://developer.github.com/v3/markdown/
func (c *Client) Markdown(text string, opt *MarkdownOptions) (string, *Response, error)

// NewRequest creates an API request. A relative URL can be provided in urlStr, in
// which case it is resolved relative to the BaseURL of the Client. Relative URLs
// should always be specified without a preceding slash. If specified, the value
// pointed to by body is JSON encoded and included as the request body.

// NewRequest 创建一个 API 请求. urlStr 中可提供一个相对 URL, 在这种情况下,
// 它被相对到客户端 BaseURL. 相对 URLs 应该总是定为不含前导反斜线.
// 如果指定了, 该值指向包含 JSON 编码的请求体.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error)

// NewUploadRequest creates an upload request. A relative URL can be provided in
// urlStr, in which case it is resolved relative to the UploadURL of the Client.
// Relative URLs should always be specified without a preceding slash.

// NewUploadRequest 创建一个上传请求. urlStr 中可提供一个相对 URL, 在这种情况下,
// 它被相对到客户端 BaseURL. 相对 URLs 应该总是定为不含前导反斜线.
func (c *Client) NewUploadRequest(urlStr string, reader io.Reader, size int64, mediaType string) (*http.Request, error)

// Octocat returns an ASCII art octocat with the specified message in a speech
// bubble. If message is empty, a random zen phrase is used.

// Octocat 返回章鱼猫艺术 ASCII 消息气泡. 如果 message 为空使用随机禅语.
func (c *Client) Octocat(message string) (string, *Response, error)

// RateLimit is deprecated. Use RateLimits instead.

// RateLimit 已经过时. 使用 RateLimits 替代.
func (c *Client) RateLimit() (*Rate, *Response, error)

// RateLimits returns the rate limits for the current client.

// RateLimits 返回当前客户端频次限制.
func (c *Client) RateLimits() (*RateLimits, *Response, error)

// Zen returns a random line from The Zen of GitHub.
//
// see also: http://warpspire.com/posts/taste/

// Zen 从 Zen of GitHub 随机返回一行.
//
// 参阅: http://warpspire.com/posts/taste/
func (c *Client) Zen() (string, *Response, error)

// CodeResult represents a single search result.

// CodeResult 表示搜索返回项.
type CodeResult struct {
	Name        *string     `json:"name,omitempty"`
	Path        *string     `json:"path,omitempty"`
	SHA         *string     `json:"sha,omitempty"`
	HTMLURL     *string     `json:"html_url,omitempty"`
	Repository  *Repository `json:"repository,omitempty"`
	TextMatches []TextMatch `json:"text_matches,omitempty"`
}

func (c CodeResult) String() string

// CodeSearchResult represents the result of an code search.

// CodeSearchResult 表示代码搜索的结果.
type CodeSearchResult struct {
	Total       *int         `json:"total_count,omitempty"`
	CodeResults []CodeResult `json:"items,omitempty"`
}

// CombinedStatus represents the combined status of a repository at a particular
// reference.

// CombinedStatus 表示某个仓库特定引用的综合状态.
type CombinedStatus struct {
	// State is the combined state of the repository.  Possible values are:
	// failture, pending, or success.

	// State 是仓库的综合状态. 可能值为:
	// failture, pending, 或 success.
	State *string `json:"state,omitempty"`

	Name       *string      `json:"name,omitempty"`
	SHA        *string      `json:"sha,omitempty"`
	TotalCount *int         `json:"total_count,omitempty"`
	Statuses   []RepoStatus `json:"statuses,omitempty"`

	CommitURL     *string `json:"commit_url,omitempty"`
	RepositoryURL *string `json:"repository_url,omitempty"`
}

func (s CombinedStatus) String() string

// Commit represents a GitHub commit.

// Commit 表示一个 GitHub 提交.
type Commit struct {
	SHA       *string       `json:"sha,omitempty"`
	Author    *CommitAuthor `json:"author,omitempty"`
	Committer *CommitAuthor `json:"committer,omitempty"`
	Message   *string       `json:"message,omitempty"`
	Tree      *Tree         `json:"tree,omitempty"`
	Parents   []Commit      `json:"parents,omitempty"`
	Stats     *CommitStats  `json:"stats,omitempty"`
	URL       *string       `json:"url,omitempty"`

	// CommentCount is the number of GitHub comments on the commit.  This
	// is only populated for requests that fetch GitHub data like
	// Pulls.ListCommits, Repositories.ListCommits, etc.

	// CommentCount 是 GitHub 提交注释的数量. 这只是为请求填充 GitHub 数据,
	// 如 Pulls.ListCommits, Repositories.ListCommits 等的请求.
	CommentCount *int `json:"comment_count,omitempty"`
}

func (c Commit) String() string

// CommitAuthor represents the author or committer of a commit. The commit author
// may not correspond to a GitHub User.

// CommitAuthor 表示作者或者提交者. 提交的作者可能没有响应的 GitHub 用户.
type CommitAuthor struct {
	Date  *time.Time `json:"date,omitempty"`
	Name  *string    `json:"name,omitempty"`
	Email *string    `json:"email,omitempty"`
}

func (c CommitAuthor) String() string

// CommitFile represents a file modified in a commit.

// CommitFile 表示提交中的某文件变更.
type CommitFile struct {
	SHA       *string `json:"sha,omitempty"`
	Filename  *string `json:"filename,omitempty"`
	Additions *int    `json:"additions,omitempty"`
	Deletions *int    `json:"deletions,omitempty"`
	Changes   *int    `json:"changes,omitempty"`
	Status    *string `json:"status,omitempty"`
	Patch     *string `json:"patch,omitempty"`
}

func (c CommitFile) String() string

// CommitStats represents the number of additions / deletions from a file in a
// given RepositoryCommit.

// CommitStats 表示给定 RepositoryCommit 中某文件的添加/删除的数量.
type CommitStats struct {
	Additions *int `json:"additions,omitempty"`
	Deletions *int `json:"deletions,omitempty"`
	Total     *int `json:"total,omitempty"`
}

func (c CommitStats) String() string

// CommitsComparison is the result of comparing two commits. See CompareCommits()
// for details.

// CommitsComparison 是两个提交的比较结果. 细节参见 CompareCommits().
type CommitsComparison struct {
	BaseCommit *RepositoryCommit `json:"base_commit,omitempty"`

	// Head can be 'behind' or 'ahead'

	// 头可为 'behind' 或 'ahead'
	Status       *string `json:"status,omitempty"`
	AheadBy      *int    `json:"ahead_by,omitempty"`
	BehindBy     *int    `json:"behind_by,omitempty"`
	TotalCommits *int    `json:"total_commits,omitempty"`

	Commits []RepositoryCommit `json:"commits,omitempty"`

	Files []CommitFile `json:"files,omitempty"`
}

func (c CommitsComparison) String() string

// CommitsListOptions specifies the optional parameters to the
// RepositoriesService.ListCommits method.

// CommitsListOptions 为 RepositoriesService.ListCommits 方法指定可选参数.
type CommitsListOptions struct {
	// SHA or branch to start listing Commits from.

	// SHA 或分支, 指向 Commits 列表起始处.
	SHA string `url:"sha,omitempty"`

	// Path that should be touched by the returned Commits.

	// Path 可能被返回的 Commits 所影响.
	Path string `url:"path,omitempty"`

	// Author of by which to filter Commits.

	// 过滤 Commits 所需要的作者.
	Author string `url:"author,omitempty"`

	// Since when should Commits be included in the response.

	// Since 可在提交时包含在请求中.
	Since time.Time `url:"since,omitempty"`

	// Until when should Commits be included in the response.

	// Until 可在提交时包含在请求中.
	Until time.Time `url:"until,omitempty"`

	ListOptions
}

// Contributor represents a repository contributor

// Contributor 表示某仓库的贡献者
type Contributor struct {
	Login             *string `json:"login,omitempty"`
	ID                *int    `json:"id,omitempty"`
	AvatarURL         *string `json:"avatar_url,omitempty"`
	GravatarID        *string `json:"gravatar_id,omitempty"`
	URL               *string `json:"url,omitempty"`
	HTMLURL           *string `json:"html_url,omitempty"`
	FollowersURL      *string `json:"followers_url,omitempty"`
	FollowingURL      *string `json:"following_url,omitempty"`
	GistsURL          *string `json:"gists_url,omitempty"`
	StarredURL        *string `json:"starred_url,omitempty"`
	SubscriptionsURL  *string `json:"subscriptions_url,omitempty"`
	OrganizationsURL  *string `json:"organizations_url,omitempty"`
	ReposURL          *string `json:"repos_url,omitempty"`
	EventsURL         *string `json:"events_url,omitempty"`
	ReceivedEventsURL *string `json:"received_events_url,omitempty"`
	Type              *string `json:"type,omitempty"`
	SiteAdmin         *bool   `json:"site_admin"`
	Contributions     *int    `json:"contributions,omitempty"`
}

// ContributorStats represents a contributor to a repository and their weekly
// contributions to a given repo.

// ContributorStats 表示某贡献者对给定仓库的每周贡献.
type ContributorStats struct {
	Author *Contributor  `json:"author,omitempty"`
	Total  *int          `json:"total,omitempty"`
	Weeks  []WeeklyStats `json:"weeks,omitempty"`
}

func (c ContributorStats) String() string

// Deployment represents a deployment in a repo

// Deployment 表示某仓库的部署信息
type Deployment struct {
	URL         *string         `json:"url,omitempty"`
	ID          *int            `json:"id,omitempty"`
	SHA         *string         `json:"sha,omitempty"`
	Ref         *string         `json:"ref,omitempty"`
	Task        *string         `json:"task,omitempty"`
	Payload     json.RawMessage `json:"payload,omitempty"`
	Environment *string         `json:"environment,omitempty"`
	Description *string         `json:"description,omitempty"`
	Creator     *User           `json:"creator,omitempty"`
	CreatedAt   *Timestamp      `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp      `json:"pushed_at,omitempty"`
}

// DeploymentRequest represents a deployment request

// DeploymentRequest 表示一个部署请求
type DeploymentRequest struct {
	Ref              *string  `json:"ref,omitempty"`
	Task             *string  `json:"task,omitempty"`
	AutoMerge        *bool    `json:"auto_merge,omitempty"`
	RequiredContexts []string `json:"required_contexts,omitempty"`
	Payload          *string  `json:"payload,omitempty"`
	Environment      *string  `json:"environment,omitempty"`
	Description      *string  `json:"description,omitempty"`
}

// DeploymentStatus represents the status of a particular deployment.

// DeploymentStatus 表示特定部署的状态
type DeploymentStatus struct {
	ID          *int       `json:"id,omitempty"`
	State       *string    `json:"state,omitempty"`
	Creator     *User      `json:"creator,omitempty"`
	Description *string    `json:"description,omitempty"`
	TargetURL   *string    `json:"target_url,omitempty"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp `json:"pushed_at,omitempty"`
}

// DeploymentStatusRequest represents a deployment request

// DeploymentStatusRequest 表示一个部署状态请求
type DeploymentStatusRequest struct {
	State       *string `json:"state,omitempty"`
	TargetURL   *string `json:"target_url,omitempty"`
	Description *string `json:"description,omitempty"`
}

// DeploymentsListOptions specifies the optional parameters to the
// RepositoriesService.ListDeployments method.

// DeploymentsListOptions 指定 RepositoriesService.ListDeployments 方法的可选参数.
type DeploymentsListOptions struct {
	// SHA of the Deployment.

	// 部署的 SHA 值.
	SHA string `url:"sha,omitempty"`

	// List deployments for a given ref.

	// 罗列给定引用的部署
	Ref string `url:"ref,omitempty"`

	// List deployments for a given task.

	// 罗列给定任务的部署
	Task string `url:"task,omitempty"`

	// List deployments for a given environment.

	// 罗列给定环境的部署
	Environment string `url:"environment,omitempty"`

	ListOptions
}

// An Error reports more details on an individual error in an ErrorResponse. These
// are the possible validation error codes:
//
//	missing:
//	    resource does not exist
//	missing_field:
//	    a required field on a resource has not been set
//	invalid:
//	    the formatting of a field is invalid
//	already_exists:
//	    another resource has the same valid as this field
//
// GitHub API docs: http://developer.github.com/v3/#client-errors

// Error 呈报 ErrorResponse 中某个单独错误的细节.
// 这些都是可能的验证错误代码:
//
//	missing:
//	    资源不存在
//	missing_field:
//	    请求字段在资源中未设置
//	invalid:
//	    无效的字段格式
//	already_exists:
//	    另外一个资源也有效具有同样的字段
//
// GitHub API 文档: http://developer.github.com/v3/#client-errors
type Error struct {
	Resource string `json:"resource"` // resource on which the error occurred
	Field    string `json:"field"`    // field on which the error occurred
	Code     string `json:"code"`     // validation error code
}

func (e *Error) Error() string

// An ErrorResponse reports one or more errors caused by an API request.
//
// GitHub API docs: http://developer.github.com/v3/#client-errors

// ErrorResponse 呈报一个或多个 API 请求引发的错误.
//
// GitHub API 文档: http://developer.github.com/v3/#client-errors
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
	Errors   []Error        `json:"errors"`  // more detail on individual errors
}

func (r *ErrorResponse) Error() string

// Event represents a GitHub event.

// Event 表示一个 GitHub 事件.
type Event struct {
	Type       *string          `json:"type,omitempty"`
	Public     *bool            `json:"public"`
	RawPayload *json.RawMessage `json:"payload,omitempty"`
	Repo       *Repository      `json:"repo,omitempty"`
	Actor      *User            `json:"actor,omitempty"`
	Org        *Organization    `json:"org,omitempty"`
	CreatedAt  *time.Time       `json:"created_at,omitempty"`
	ID         *string          `json:"id,omitempty"`
}

// Payload returns the parsed event payload. For recognized event types
// (PushEvent), a value of the corresponding struct type will be returned.

// Payload 返回解析的事件有效负载. 对于公认的事件类型 (PushEvent),
// 返回一个相应结构类型的值.
func (e *Event) Payload() (payload interface{})

func (e Event) String() string

// Gist represents a GitHub's gist.

// Gist 表示一个 GitHub's gist.
type Gist struct {
	ID          *string                   `json:"id,omitempty"`
	Description *string                   `json:"description,omitempty"`
	Public      *bool                     `json:"public,omitempty"`
	Owner       *User                     `json:"owner,omitempty"`
	Files       map[GistFilename]GistFile `json:"files,omitempty"`
	Comments    *int                      `json:"comments,omitempty"`
	HTMLURL     *string                   `json:"html_url,omitempty"`
	GitPullURL  *string                   `json:"git_pull_url,omitempty"`
	GitPushURL  *string                   `json:"git_push_url,omitempty"`
	CreatedAt   *time.Time                `json:"created_at,omitempty"`
	UpdatedAt   *time.Time                `json:"updated_at,omitempty"`
}

func (g Gist) String() string

// GistComment represents a Gist comment.

// GistComment 表示一个 Gist 注释.
type GistComment struct {
	ID        *int       `json:"id,omitempty"`
	URL       *string    `json:"url,omitempty"`
	Body      *string    `json:"body,omitempty"`
	User      *User      `json:"user,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

func (g GistComment) String() string

// GistFile represents a file on a gist.

// GistFile 表示 gist 上的某个文件.
type GistFile struct {
	Size     *int    `json:"size,omitempty"`
	Filename *string `json:"filename,omitempty"`
	RawURL   *string `json:"raw_url,omitempty"`
	Content  *string `json:"content,omitempty"`
}

func (g GistFile) String() string

// GistFilename represents filename on a gist.

// GistFilename 表示 gist 上的文件名.
type GistFilename string

// GistListOptions specifies the optional parameters to the GistsService.List,
// GistsService.ListAll, and GistsService.ListStarred methods.

// GistListOptions 指定 GistsService.List, GistsService.ListAll,
// 和 GistsService.ListStarred 方法的可选参数.
type GistListOptions struct {
	// Since filters Gists by time.

	// 过滤 Gists 起始时间.
	Since time.Time `url:"since,omitempty"`

	ListOptions
}

// GistsService handles communication with the Gist related methods of the GitHub
// API.
//
// GitHub API docs: http://developer.github.com/v3/gists/

// GistsService 处理与 Gist 相关的 GitHub API 通信方法.
//
// GitHub API 文档: http://developer.github.com/v3/gists/
type GistsService struct {
	// contains filtered or unexported fields
}

// Create a gist for authenticated user.
//
// GitHub API docs: http://developer.github.com/v3/gists/#create-a-gist

// Create 为授权用户创建一个 gist.
//
// GitHub API 文档: http://developer.github.com/v3/gists/#create-a-gist
func (s *GistsService) Create(gist *Gist) (*Gist, *Response, error)

// CreateComment creates a comment for a gist.
//
// GitHub API docs: http://developer.github.com/v3/gists/comments/#create-a-comment

// CreateComment 为一个 gist 创建注释.
//
// GitHub API 文档: http://developer.github.com/v3/gists/comments/#create-a-comment
func (s *GistsService) CreateComment(gistID string, comment *GistComment) (*GistComment, *Response, error)

// Delete a gist.
//
// GitHub API docs: http://developer.github.com/v3/gists/#delete-a-gist

// Delete 删除一个 gist.
//
// GitHub API 文档: http://developer.github.com/v3/gists/#delete-a-gist
func (s *GistsService) Delete(id string) (*Response, error)

// DeleteComment deletes a gist comment.
//
// GitHub API docs: http://developer.github.com/v3/gists/comments/#delete-a-comment

// DeleteComment 删除一个 gist 注释.
//
// GitHub API 文档: http://developer.github.com/v3/gists/comments/#delete-a-comment
func (s *GistsService) DeleteComment(gistID string, commentID int) (*Response, error)

// Edit a gist.
//
// GitHub API docs: http://developer.github.com/v3/gists/#edit-a-gist

// Edit 编辑一个 gist.
//
// GitHub API 文档: http://developer.github.com/v3/gists/#edit-a-gist
func (s *GistsService) Edit(id string, gist *Gist) (*Gist, *Response, error)

// EditComment edits an existing gist comment.
//
// GitHub API docs: http://developer.github.com/v3/gists/comments/#edit-a-comment

// EditComment edits an existing gist comment.
//
// GitHub API 文档: http://developer.github.com/v3/gists/comments/#edit-a-comment
func (s *GistsService) EditComment(gistID string, commentID int, comment *GistComment) (*GistComment, *Response, error)

// Fork a gist.
//
// GitHub API docs: http://developer.github.com/v3/gists/#fork-a-gist

// Fork 一个 gist.
//
// GitHub API 文档: http://developer.github.com/v3/gists/#fork-a-gist
func (s *GistsService) Fork(id string) (*Gist, *Response, error)

// Get a single gist.
//
// GitHub API docs: http://developer.github.com/v3/gists/#get-a-single-gist

// Get 获取单一 gist.
//
// GitHub API 文档: http://developer.github.com/v3/gists/#get-a-single-gist
func (s *GistsService) Get(id string) (*Gist, *Response, error)

// GetComment retrieves a single comment from a gist.
//
// GitHub API docs:
// http://developer.github.com/v3/gists/comments/#get-a-single-comment

// GetComment 检索某 gist 的单一注释.
//
// GitHub API 文档:
// http://developer.github.com/v3/gists/comments/#get-a-single-comment
func (s *GistsService) GetComment(gistID string, commentID int) (*GistComment, *Response, error)

// IsStarred checks if a gist is starred by authenticated user.
//
// GitHub API docs:
// http://developer.github.com/v3/gists/#check-if-a-gist-is-starred

// IsStarred 检查一个 gist 是否被认证用户星标.
//
// GitHub API 文档:
// http://developer.github.com/v3/gists/#check-if-a-gist-is-starred
func (s *GistsService) IsStarred(id string) (bool, *Response, error)

// List gists for a user. Passing the empty string will list all public gists if
// called anonymously. However, if the call is authenticated, it will returns all
// gists for the authenticated user.
//
// GitHub API docs: http://developer.github.com/v3/gists/#list-gists

// List 罗列某用户的 gists. 传递空字符串将罗列所有 anonymously 的公共 gists.
// 无论如何, 如果是授权调用, 它将返回授权用户的所有 gists.
//
// GitHub API 文档: http://developer.github.com/v3/gists/#list-gists
func (s *GistsService) List(user string, opt *GistListOptions) ([]Gist, *Response, error)

// ListAll lists all public gists.
//
// GitHub API docs: http://developer.github.com/v3/gists/#list-gists

// ListAll 罗列所有公共 gists.
//
// GitHub API 文档: http://developer.github.com/v3/gists/#list-gists
func (s *GistsService) ListAll(opt *GistListOptions) ([]Gist, *Response, error)

// ListComments lists all comments for a gist.
//
// GitHub API docs:
// http://developer.github.com/v3/gists/comments/#list-comments-on-a-gist

// ListComments 罗列某 gist 的所有注释.
//
// GitHub API 文档:
// http://developer.github.com/v3/gists/comments/#list-comments-on-a-gist
func (s *GistsService) ListComments(gistID string, opt *ListOptions) ([]GistComment, *Response, error)

// ListStarred lists starred gists of authenticated user.
//
// GitHub API docs: http://developer.github.com/v3/gists/#list-gists

// ListStarred 罗列授权用户被星标的 gists.
//
// GitHub API 文档: http://developer.github.com/v3/gists/#list-gists
func (s *GistsService) ListStarred(opt *GistListOptions) ([]Gist, *Response, error)

// Star a gist on behalf of authenticated user.
//
// GitHub API docs: http://developer.github.com/v3/gists/#star-a-gist

// Star 星标授权用户的某个 gist.
//
// GitHub API 文档: http://developer.github.com/v3/gists/#star-a-gist
func (s *GistsService) Star(id string) (*Response, error)

// Unstar a gist on a behalf of authenticated user.
//
// Github API docs: http://developer.github.com/v3/gists/#unstar-a-gist

// Unstar 取消授权用户的某个 gist 星标.
//
// Github API 文档: http://developer.github.com/v3/gists/#unstar-a-gist
func (s *GistsService) Unstar(id string) (*Response, error)

// GitObject represents a Git object.

// GitObject 表示一个 Git 对象.
type GitObject struct {
	Type *string `json:"type"`
	SHA  *string `json:"sha"`
	URL  *string `json:"url"`
}

func (o GitObject) String() string

// GitService handles communication with the git data related methods of the GitHub
// API.
//
// GitHub API docs: http://developer.github.com/v3/git/

// GitService 处理与 git 数据相关的 GitHub API 通信方法.
//
// GitHub API 文档: http://developer.github.com/v3/git/
type GitService struct {
	// contains filtered or unexported fields
}

// CreateBlob creates a blob object.
//
// GitHub API docs: http://developer.github.com/v3/git/blobs/#create-a-blob

// CreateBlob 创建一个 blob 对象.
//
// GitHub API 文档: http://developer.github.com/v3/git/blobs/#create-a-blob
func (s *GitService) CreateBlob(owner string, repo string, blob *Blob) (*Blob, *Response, error)

// CreateCommit creates a new commit in a repository.
//
// The commit.Committer is optional and will be filled with the commit.Author data
// if omitted. If the commit.Author is omitted, it will be filled in with the
// authenticated user’s information and the current date.
//
// GitHub API docs: http://developer.github.com/v3/git/commits/#create-a-commit

// CreateCommit 在某仓库新建一个提交.
//
// commit.Committer 是可选的, 如果省略会用 commit.Author 的数据填充.
// 如果 commit.Author 省略, 用授权用户信息和当前日期填充它.
//
// GitHub API 文档: http://developer.github.com/v3/git/commits/#create-a-commit
func (s *GitService) CreateCommit(owner string, repo string, commit *Commit) (*Commit, *Response, error)

// CreateRef creates a new ref in a repository.
//
// GitHub API docs: http://developer.github.com/v3/git/refs/#create-a-reference

// CreteRef 在某仓库新建一个引用.
//
// GitHub API 文档: http://developer.github.com/v3/git/refs/#create-a-reference
func (s *GitService) CreateRef(owner string, repo string, ref *Reference) (*Reference, *Response, error)

// CreateTag creates a tag object.
//
// GitHub API docs: http://developer.github.com/v3/git/tags/#create-a-tag-object

// CreateTag 创建一个标签对象.
//
// GitHub API 文档: http://developer.github.com/v3/git/tags/#create-a-tag-object
func (s *GitService) CreateTag(owner string, repo string, tag *Tag) (*Tag, *Response, error)

// CreateTree creates a new tree in a repository. If both a tree and a nested path
// modifying that tree are specified, it will overwrite the contents of that tree
// with the new path contents and write a new tree out.
//
// GitHub API docs: http://developer.github.com/v3/git/trees/#create-a-tree

// CreateTree 在某仓库新建一个树. 如果定义的树和嵌套的路径都更改了,
// 它会用新的路径内容写一个新树覆盖原树的内容.
//
// GitHub API 文档: http://developer.github.com/v3/git/trees/#create-a-tree
func (s *GitService) CreateTree(owner string, repo string, baseTree string, entries []TreeEntry) (*Tree, *Response, error)

// DeleteRef deletes a ref from a repository.
//
// GitHub API docs: http://developer.github.com/v3/git/refs/#delete-a-reference

// DeleteRef 从某仓库删除一个引用.
//
// GitHub API 文档: http://developer.github.com/v3/git/refs/#delete-a-reference
func (s *GitService) DeleteRef(owner string, repo string, ref string) (*Response, error)

// GetBlob fetchs a blob from a repo given a SHA.
//
// GitHub API docs: http://developer.github.com/v3/git/blobs/#get-a-blob

// GetBlob 以给定的 SHA 从某仓库提取一个 Blob.
//
// GitHub API 文档: http://developer.github.com/v3/git/blobs/#get-a-blob
func (s *GitService) GetBlob(owner string, repo string, sha string) (*Blob, *Response, error)

// GetCommit fetchs the Commit object for a given SHA.
//
// GitHub API docs: http://developer.github.com/v3/git/commits/#get-a-commit

// GetCommit 以给定的 SHA 从某仓库提取一个 Commit.
//
// GitHub API 文档: http://developer.github.com/v3/git/commits/#get-a-commit
func (s *GitService) GetCommit(owner string, repo string, sha string) (*Commit, *Response, error)

// GetRef fetches the Reference object for a given Git ref.
//
// GitHub API docs: http://developer.github.com/v3/git/refs/#get-a-reference

// GetRef 以给定的 Git 引用提取 Reference 对象.
//
// GitHub API 文档: http://developer.github.com/v3/git/refs/#get-a-reference
func (s *GitService) GetRef(owner string, repo string, ref string) (*Reference, *Response, error)

// GetTag fetchs a tag from a repo given a SHA.
//
// GitHub API docs: http://developer.github.com/v3/git/tags/#get-a-tag

// GetTag 以给定的 SHA 从某仓库提取一个 Tag.
//
// GitHub API 文档: http://developer.github.com/v3/git/tags/#get-a-tag
func (s *GitService) GetTag(owner string, repo string, sha string) (*Tag, *Response, error)

// GetTree fetches the Tree object for a given sha hash from a repository.
//
// GitHub API docs: http://developer.github.com/v3/git/trees/#get-a-tree

// GetTree 以给定的 SHA 从某仓库提取 Tree 对象.
//
// GitHub API 文档: http://developer.github.com/v3/git/trees/#get-a-tree
func (s *GitService) GetTree(owner string, repo string, sha string, recursive bool) (*Tree, *Response, error)

// ListRefs lists all refs in a repository.
//
// GitHub API docs: http://developer.github.com/v3/git/refs/#get-all-references

// ListRefs 罗列某仓库的所有引用.
//
// GitHub API 文档: http://developer.github.com/v3/git/refs/#get-all-references
func (s *GitService) ListRefs(owner, repo string, opt *ReferenceListOptions) ([]Reference, *Response, error)

// UpdateRef updates an existing ref in a repository.
//
// GitHub API docs: http://developer.github.com/v3/git/refs/#update-a-reference

// UpdateRef 更新某仓库的一个引用.
//
// GitHub API 文档: http://developer.github.com/v3/git/refs/#update-a-reference
func (s *GitService) UpdateRef(owner string, repo string, ref *Reference, force bool) (*Reference, *Response, error)

// Gitignore represents a .gitignore file as returned by the GitHub API.

// Gitignore 表示 GitHub API 所返回的一个 .gitignore 文件.
type Gitignore struct {
	Name   *string `json:"name,omitempty"`
	Source *string `json:"source,omitempty"`
}

func (g Gitignore) String() string

// GitignoresService provides access to the gitignore related functions in the
// GitHub API.
//
// GitHub API docs: http://developer.github.com/v3/gitignore/

// GitignoresService 提供访问 gitignore 相关 GitHub API 功能.
//
// GitHub API 文档: http://developer.github.com/v3/gitignore/
type GitignoresService struct {
	// contains filtered or unexported fields
}

// Get a Gitignore by name.
//
// http://developer.github.com/v3/gitignore/#get-a-single-template

// Get 通过 name 获取一个 Gitignore.
//
// http://developer.github.com/v3/gitignore/#get-a-single-template
func (s GitignoresService) Get(name string) (*Gitignore, *Response, error)

// List all available Gitignore templates.
//
// http://developer.github.com/v3/gitignore/#listing-available-templates

// List 罗列所有可用的 Gitignore 模版.
//
// http://developer.github.com/v3/gitignore/#listing-available-templates
func (s GitignoresService) List() ([]string, *Response, error)

// Hook represents a GitHub (web and service) hook for a repository.

// Hook 表示某仓库的一个 GitHub (web 服务) 钩子.
type Hook struct {
	CreatedAt *time.Time             `json:"created_at,omitempty"`
	UpdatedAt *time.Time             `json:"updated_at,omitempty"`
	Name      *string                `json:"name,omitempty"`
	Events    []string               `json:"events,omitempty"`
	Active    *bool                  `json:"active,omitempty"`
	Config    map[string]interface{} `json:"config,omitempty"`
	ID        *int                   `json:"id,omitempty"`
}

func (h Hook) String() string

// Issue represents a GitHub issue on a repository.

// Issue 表示某仓库的一个 GitHub 问题.
type Issue struct {
	Number           *int              `json:"number,omitempty"`
	State            *string           `json:"state,omitempty"`
	Title            *string           `json:"title,omitempty"`
	Body             *string           `json:"body,omitempty"`
	User             *User             `json:"user,omitempty"`
	Labels           []Label           `json:"labels,omitempty"`
	Assignee         *User             `json:"assignee,omitempty"`
	Comments         *int              `json:"comments,omitempty"`
	ClosedAt         *time.Time        `json:"closed_at,omitempty"`
	CreatedAt        *time.Time        `json:"created_at,omitempty"`
	UpdatedAt        *time.Time        `json:"updated_at,omitempty"`
	URL              *string           `json:"url,omitempty"`
	HTMLURL          *string           `json:"html_url,omitempty"`
	Milestone        *Milestone        `json:"milestone,omitempty"`
	PullRequestLinks *PullRequestLinks `json:"pull_request,omitempty"`

	// TextMatches is only populated from search results that request text matches
	// See: search.go and https://developer.github.com/v3/search/#text-match-metadata

	// TextMatches 只是填入了搜索文本匹配请求的结果.
	// See: search.go and https://developer.github.com/v3/search/#text-match-metadata
	TextMatches []TextMatch `json:"text_matches,omitempty"`
}

func (i Issue) String() string

// IssueActivityEvent represents the payload delivered by Issue webhook

// IssueActivityEvent 表示 Issue webhook 交付的有效负载
type IssueActivityEvent struct {
	Action *string     `json:"action,omitempty"`
	Issue  *Issue      `json:"issue,omitempty"`
	Repo   *Repository `json:"repository,omitempty"`
	Sender *User       `json:"sender,omitempty"`
}

// IssueComment represents a comment left on an issue.

// IssueComment 表示一个 issue 之中的一个评论.
type IssueComment struct {
	ID        *int       `json:"id,omitempty"`
	Body      *string    `json:"body,omitempty"`
	User      *User      `json:"user,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	URL       *string    `json:"url,omitempty"`
	HTMLURL   *string    `json:"html_url,omitempty"`
	IssueURL  *string    `json:"issue_url,omitempty"`
}

func (i IssueComment) String() string

// IssueCommentEvent represents the payload delivered by IssueComment webhook
//
// This webhook also gets fired for comments on pull requests

// IssueCommentEvent 表示 IssueComment webhook 交付的有效负载
//
// 此 webhook 也触发上拉请求
type IssueCommentEvent struct {
	Action  *string       `json:"action,omitempty"`
	Issue   *Issue        `json:"issue,omitempty"`
	Comment *IssueComment `json:"comment,omitempty"`
	Repo    *Repository   `json:"repository,omitempty"`
	Sender  *User         `json:"sender,omitempty"`
}

// IssueEvent represents an event that occurred around an Issue or Pull Request.

// IssueEvent 表示围绕着一个问题或上拉请求时发生的事件.
type IssueEvent struct {
	ID  *int    `json:"id,omitempty"`
	URL *string `json:"url,omitempty"`

	// The User that generated this event.

	// 产生该事件的用户
	Actor *User `json:"actor,omitempty"`

	// Event identifies the actual type of Event that occurred.  Possible
	// values are:
	//
	//     closed
	//       The issue was closed by the actor. When the commit_id is
	//       present, it identifies the commit that closed the issue using
	//       “closes / fixes #NN” syntax.
	//
	//     reopened
	//       The issue was reopened by the actor.
	//
	//     subscribed
	//       The actor subscribed to receive notifications for an issue.
	//
	//     merged
	//       The issue was merged by the actor. The commit_id attribute is the SHA1 of the HEAD commit that was merged.
	//
	//     referenced
	//       The issue was referenced from a commit message. The commit_id attribute is the commit SHA1 of where that happened.
	//
	//     mentioned
	//       The actor was @mentioned in an issue body.
	//
	//     assigned
	//       The issue was assigned to the actor.
	//
	//     head_ref_deleted
	//       The pull request’s branch was deleted.
	//
	//     head_ref_restored
	//       The pull request’s branch was restored.

	// Event 标识所发生事件的实际类型. 可能的值有:
	//
	//     closed
	//       该问题被 actor 关闭. 当 commit_id 出现, 它用
	//       "closes / fixes #NN" 语法标识提交该问题被关闭.
	//
	//     reopened
	//       该问题被 actor 重新开放.
	//
	//     subscribed
	//       该问题被 actor 订阅, 接收通知.
	//
	//     merged
	//       该问题被 actor 合并. commit_id 属性为该合并提交 HEAD 的 SHA1.
	//
	//     referenced
	//       该问题被某提交消息引用. commit_id 属性为发生提交的 SHA1.
	//
	//     mentioned
	//       actor 被问题信息主体提到 @mentioned.
	//
	//     assigned
	//       该问题被指派给 actor.
	//
	//     head_ref_deleted
	//       上拉请求分支被删除了.
	//
	//     head_ref_restored
	//       上拉请求分支被恢复了.
	Event *string `json:"event,omitempty"`

	// The SHA of the commit that referenced this commit, if applicable.

	// 如果有, 该提交引用了此提交(SHA).
	CommitID *string `json:"commit_id,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	Issue     *Issue     `json:"issue,omitempty"`
}

// IssueListByRepoOptions specifies the optional parameters to the
// IssuesService.ListByRepo method.

// IssueListByRepoOptions 指定 IssuesService.ListByRepo 方法的可选参数.
type IssueListByRepoOptions struct {
	// Milestone limits issues for the specified milestone.  Possible values are
	// a milestone number, "none" for issues with no milestone, "*" for issues
	// with any milestone.

	// Milestone 限定指定里程碑的问题. 可能的值为里程碑号,
	// "none" 为无里程碑的问题, "*" 为任意里程碑的问题.
	Milestone string `url:"milestone,omitempty"`

	// State filters issues based on their state.  Possible values are: open,
	// closed.  Default is "open".

	// State 过滤问题, 基于他们状态. 可能的值有: open, closed. 缺省为 "open".
	State string `url:"state,omitempty"`

	// Assignee filters issues based on their assignee.  Possible values are a
	// user name, "none" for issues that are not assigned, "*" for issues with
	// any assigned user.

	// Assignee 过滤问题, 基于受理人. 可能的值为一个用户名,
	// "none" 为无指派用户的问题, "*" 为任意指派用户的问题.
	Assignee string `url:"assignee,omitempty"`

	// Assignee filters issues based on their creator.

	// 过滤问题, 基于受理人的创建者
	Creator string `url:"creator,omitempty"`

	// Assignee filters issues to those mentioned a specific user.

	// 过滤问题, 所提到的那些特定的受理人用户.
	Mentioned string `url:"mentioned,omitempty"`

	// Labels filters issues based on their label.

	// Labels 过滤问题, 基于他们的标记.
	Labels []string `url:"labels,omitempty,comma"`

	// Sort specifies how to sort issues.  Possible values are: created, updated,
	// and comments.  Default value is "assigned".

	// Sort 指定如何排序问题. 可能的值有: created, updated, comments.
	// 缺省为 "assigned".
	Sort string `url:"sort,omitempty"`

	// Direction in which to sort issues.  Possible values are: asc, desc.
	// Default is "asc".

	// 问题排序的方向. 可能的只有: asc, desc. 缺省为 "asc".
	Direction string `url:"direction,omitempty"`

	// Since filters issues by time.

	// 以起始时间过滤问题.
	Since time.Time `url:"since,omitempty"`

	ListOptions
}

// IssueListCommentsOptions specifies the optional parameters to the
// IssuesService.ListComments method.

// IssueListCommentsOptions 指定 IssuesService.ListComments 方法的可选参数.
type IssueListCommentsOptions struct {
	// Sort specifies how to sort comments.  Possible values are: created, updated.

	// 指定如何排序评论. 可能的值有: created, updated.
	Sort string `url:"sort,omitempty"`

	// Direction in which to sort comments.  Possible values are: asc, desc.

	// 评论排序的方向. 可能的只有: asc, desc.
	Direction string `url:"direction,omitempty"`

	// Since filters comments by time.

	// 以起始时间过滤评论.
	Since time.Time `url:"since,omitempty"`

	ListOptions
}

// IssueListOptions specifies the optional parameters to the IssuesService.List and
// IssuesService.ListByOrg methods.

// IssueListOptions 指定 IssuesService.List 和 IssuesService.ListByOrg
// 方法的可选参数.
type IssueListOptions struct {
	// Filter specifies which issues to list.  Possible values are: assigned,
	// created, mentioned, subscribed, all.  Default is "assigned".

	// Filter 指定罗列那些问题. 可能的值有:
	// assigned, created, mentioned, subscribed, all. 缺省为 "assigned".
	Filter string `url:"filter,omitempty"`

	// State filters issues based on their state.  Possible values are: open,
	// closed.  Default is "open".

	// State 过滤问题, 基于他们的状态. 可能的值有: open, closed. 缺省为 "open".
	State string `url:"state,omitempty"`

	// Labels filters issues based on their label.

	// Labels 过滤问题, 基于他们的标记.
	Labels []string `url:"labels,comma,omitempty"`

	// Sort specifies how to sort issues.  Possible values are: created, updated,
	// and comments.  Default value is "assigned".

	// Sort 指定如何排序问题. 可能的值有: created, updated, comments.
	// 缺省为 "assigned".
	Sort string `url:"sort,omitempty"`

	// Direction in which to sort issues.  Possible values are: asc, desc.
	// Default is "asc".

	// 问题排序的方向. 可能的只有: asc, desc.
	Direction string `url:"direction,omitempty"`

	// Since filters issues by time.

	// 以起始时间过滤问题.
	Since time.Time `url:"since,omitempty"`

	ListOptions
}

// IssueRequest represents a request to create/edit an issue. It is separate from
// Issue above because otherwise Labels and Assignee fail to serialize to the
// correct JSON.

// IssueRequest 表示建立/编辑一个问题的请求. 它在 Issue 上是独立的,
// 不然的话 Labels 和 Assignee 不能序列化到正确的 JSON.
type IssueRequest struct {
	Title     *string  `json:"title,omitempty"`
	Body      *string  `json:"body,omitempty"`
	Labels    []string `json:"labels,omitempty"`
	Assignee  *string  `json:"assignee,omitempty"`
	State     *string  `json:"state,omitempty"`
	Milestone *int     `json:"milestone,omitempty"`
}

// IssuesSearchResult represents the result of an issues search.

// IssuesSearchResult 表示问题搜索的结果.
type IssuesSearchResult struct {
	Total  *int    `json:"total_count,omitempty"`
	Issues []Issue `json:"items,omitempty"`
}

// IssuesService handles communication with the issue related methods of the GitHub
// API.
//
// GitHub API docs: http://developer.github.com/v3/issues/

// IssuesService 处理 GitHub API 问题相关的通信方法.
//
// GitHub API 文档: http://developer.github.com/v3/issues/
type IssuesService struct {
	// contains filtered or unexported fields
}

// AddLabelsToIssue adds labels to an issue.
//
// GitHub API docs:
// http://developer.github.com/v3/issues/labels/#list-all-labels-for-this-repository

// AddLabelsToIssue 给问题添加标记.
//
// GitHub API 文档:
// http://developer.github.com/v3/issues/labels/#list-all-labels-for-this-repository
func (s *IssuesService) AddLabelsToIssue(owner string, repo string, number int, labels []string) ([]Label, *Response, error)

// Create a new issue on the specified repository.
//
// GitHub API docs: http://developer.github.com/v3/issues/#create-an-issue

// Create 在指定仓库上新建一个问题.
//
// GitHub API 文档: http://developer.github.com/v3/issues/#create-an-issue
func (s *IssuesService) Create(owner string, repo string, issue *IssueRequest) (*Issue, *Response, error)

// CreateComment creates a new comment on the specified issue.
//
// GitHub API docs:
// http://developer.github.com/v3/issues/comments/#create-a-comment

// CreateComment 在指定问题上新建一个评论.
//
// GitHub API 文档:
// http://developer.github.com/v3/issues/comments/#create-a-comment
func (s *IssuesService) CreateComment(owner string, repo string, number int, comment *IssueComment) (*IssueComment, *Response, error)

// CreateLabel creates a new label on the specified repository.
//
// GitHub API docs: http://developer.github.com/v3/issues/labels/#create-a-label

// CreateLabel 在指定仓库上新建一个标记.
//
// GitHub API 文档: http://developer.github.com/v3/issues/labels/#create-a-label
func (s *IssuesService) CreateLabel(owner string, repo string, label *Label) (*Label, *Response, error)

// CreateMilestone creates a new milestone on the specified repository.
//
// GitHub API docs:
// https://developer.github.com/v3/issues/milestones/#create-a-milestone

// CreateMilestone 在指定仓库上新建一个里程碑.
//
// GitHub API 文档:
// https://developer.github.com/v3/issues/milestones/#create-a-milestone
func (s *IssuesService) CreateMilestone(owner string, repo string, milestone *Milestone) (*Milestone, *Response, error)

// DeleteComment deletes an issue comment.
//
// GitHub API docs:
// http://developer.github.com/v3/issues/comments/#delete-a-comment

// DeleteComment 删除一个问题评论.
//
// GitHub API 文档:
// http://developer.github.com/v3/issues/comments/#delete-a-comment
func (s *IssuesService) DeleteComment(owner string, repo string, id int) (*Response, error)

// DeleteLabel deletes a label.
//
// GitHub API docs: http://developer.github.com/v3/issues/labels/#delete-a-label

// DeleteLabel 删除一个标记.
//
// GitHub API 文档: http://developer.github.com/v3/issues/labels/#delete-a-label
func (s *IssuesService) DeleteLabel(owner string, repo string, name string) (*Response, error)

// DeleteMilestone deletes a milestone.
//
// GitHub API docs:
// https://developer.github.com/v3/issues/milestones/#delete-a-milestone

// DeleteMilestone 删除一个里程碑.
//
// GitHub API 文档:
// https://developer.github.com/v3/issues/milestones/#delete-a-milestone
func (s *IssuesService) DeleteMilestone(owner string, repo string, number int) (*Response, error)

// Edit an issue.
//
// GitHub API docs: http://developer.github.com/v3/issues/#edit-an-issue

// Edit 编辑一个问题.
//
// GitHub API 文档: http://developer.github.com/v3/issues/#edit-an-issue
func (s *IssuesService) Edit(owner string, repo string, number int, issue *IssueRequest) (*Issue, *Response, error)

// EditComment updates an issue comment.
//
// GitHub API docs: http://developer.github.com/v3/issues/comments/#edit-a-comment

// EditComment 更新一个问题评论.
//
// GitHub API 文档: http://developer.github.com/v3/issues/comments/#edit-a-comment
func (s *IssuesService) EditComment(owner string, repo string, id int, comment *IssueComment) (*IssueComment, *Response, error)

// EditLabel edits a label.
//
// GitHub API docs: http://developer.github.com/v3/issues/labels/#update-a-label

// EditLabel 编辑一个标记.
//
// GitHub API 文档: http://developer.github.com/v3/issues/labels/#update-a-label
func (s *IssuesService) EditLabel(owner string, repo string, name string, label *Label) (*Label, *Response, error)

// EditMilestone edits a milestone.
//
// GitHub API docs:
// https://developer.github.com/v3/issues/milestones/#update-a-milestone

// EditMilestone 编辑一个里程碑.
//
// GitHub API 文档:
// https://developer.github.com/v3/issues/milestones/#update-a-milestone
func (s *IssuesService) EditMilestone(owner string, repo string, number int, milestone *Milestone) (*Milestone, *Response, error)

// Get a single issue.
//
// GitHub API docs: http://developer.github.com/v3/issues/#get-a-single-issue

// Get 获取单个问题.
//
// GitHub API 文档: http://developer.github.com/v3/issues/#get-a-single-issue
func (s *IssuesService) Get(owner string, repo string, number int) (*Issue, *Response, error)

// GetComment fetches the specified issue comment.
//
// GitHub API docs:
// http://developer.github.com/v3/issues/comments/#get-a-single-comment

// GetComment 获取指定的问题评论.
//
// GitHub API 文档:
// http://developer.github.com/v3/issues/comments/#get-a-single-comment
func (s *IssuesService) GetComment(owner string, repo string, id int) (*IssueComment, *Response, error)

// GetEvent returns the specified issue event.
//
// GitHub API docs:
// https://developer.github.com/v3/issues/events/#get-a-single-event

// GetEvent 返回指定的问题事件.
//
// GitHub API 文档:
// https://developer.github.com/v3/issues/events/#get-a-single-event
func (s *IssuesService) GetEvent(owner, repo string, id int) (*IssueEvent, *Response, error)

// GetLabel gets a single label.
//
// GitHub API docs:
// http://developer.github.com/v3/issues/labels/#get-a-single-label

// GetLabel 获取单个标记.
//
// GitHub API 文档:
// http://developer.github.com/v3/issues/labels/#get-a-single-label
func (s *IssuesService) GetLabel(owner string, repo string, name string) (*Label, *Response, error)

// GetMilestone gets a single milestone.
//
// GitHub API docs:
// https://developer.github.com/v3/issues/milestones/#get-a-single-milestone

// GetMilestone 获取单个里程碑.
//
// GitHub API 文档:
// https://developer.github.com/v3/issues/milestones/#get-a-single-milestone
func (s *IssuesService) GetMilestone(owner string, repo string, number int) (*Milestone, *Response, error)

// IsAssignee checks if a user is an assignee for the specified repository.
//
// GitHub API docs: http://developer.github.com/v3/issues/assignees/#check-assignee

// IsAssignee 检查用户是否是指定仓库的受理人.
//
// GitHub API 文档: http://developer.github.com/v3/issues/assignees/#check-assignee
func (s *IssuesService) IsAssignee(owner string, repo string, user string) (bool, *Response, error)

// List the issues for the authenticated user. If all is true, list issues across
// all the user's visible repositories including owned, member, and organization
// repositories; if false, list only owned and member repositories.
//
// GitHub API docs: http://developer.github.com/v3/issues/#list-issues

// List 罗列授权用户的问题. 如果 all 为 true, 罗列横跨用户所有的可见仓库,
// 包括自有的, 成员的, 和组织的仓库; 如果为 false, 仅罗列自有的和成员的仓库.
//
// GitHub API 文档: http://developer.github.com/v3/issues/#list-issues
func (s *IssuesService) List(all bool, opt *IssueListOptions) ([]Issue, *Response, error)

// ListAssignees fetches all available assignees (owners and collaborators) to
// which issues may be assigned.
//
// GitHub API docs: http://developer.github.com/v3/issues/assignees/#list-assignees

// ListAssignees 获取所有那些有效被指派问题的受理人 (所有者和合作者).
//
// GitHub API 文档: http://developer.github.com/v3/issues/assignees/#list-assignees
func (s *IssuesService) ListAssignees(owner string, repo string, opt *ListOptions) ([]User, *Response, error)

// ListByOrg fetches the issues in the specified organization for the authenticated
// user.
//
// GitHub API docs: http://developer.github.com/v3/issues/#list-issues

// ListByOrg 获取授权用户的指定组织的问题.
//
// GitHub API 文档: http://developer.github.com/v3/issues/#list-issues
func (s *IssuesService) ListByOrg(org string, opt *IssueListOptions) ([]Issue, *Response, error)

// ListByRepo lists the issues for the specified repository.
//
// GitHub API docs:
// http://developer.github.com/v3/issues/#list-issues-for-a-repository

// ListByRepo 罗列指定仓库的问题.
//
// GitHub API 文档:
// http://developer.github.com/v3/issues/#list-issues-for-a-repository
func (s *IssuesService) ListByRepo(owner string, repo string, opt *IssueListByRepoOptions) ([]Issue, *Response, error)

// ListComments lists all comments on the specified issue. Specifying an issue
// number of 0 will return all comments on all issues for the repository.
//
// GitHub API docs:
// http://developer.github.com/v3/issues/comments/#list-comments-on-an-issue

// ListComments 罗列指定问题所有的评论. 指定问题号 0 将返回仓库所有问题评论.
//
// GitHub API 文档:
// http://developer.github.com/v3/issues/comments/#list-comments-on-an-issue
func (s *IssuesService) ListComments(owner string, repo string, number int, opt *IssueListCommentsOptions) ([]IssueComment, *Response, error)

// ListIssueEvents lists events for the specified issue.
//
// GitHub API docs:
// https://developer.github.com/v3/issues/events/#list-events-for-an-issue

// ListIssueEvents 罗列指定问题的事件.
//
// GitHub API 文档:
// https://developer.github.com/v3/issues/events/#list-events-for-an-issue
func (s *IssuesService) ListIssueEvents(owner, repo string, number int, opt *ListOptions) ([]IssueEvent, *Response, error)

// ListLabels lists all labels for a repository.
//
// GitHub API docs:
// http://developer.github.com/v3/issues/labels/#list-all-labels-for-this-repository

// ListLabels 罗列某仓库所有的标记.
//
// GitHub API 文档:
// http://developer.github.com/v3/issues/labels/#list-all-labels-for-this-repository
func (s *IssuesService) ListLabels(owner string, repo string, opt *ListOptions) ([]Label, *Response, error)

// ListLabelsByIssue lists all labels for an issue.
//
// GitHub API docs:
// http://developer.github.com/v3/issues/labels/#list-all-labels-for-this-repository

// ListLabelsByIssue 罗列某问题所有的标记.
//
// GitHub API 文档:
// http://developer.github.com/v3/issues/labels/#list-all-labels-for-this-repository
func (s *IssuesService) ListLabelsByIssue(owner string, repo string, number int, opt *ListOptions) ([]Label, *Response, error)

// ListLabelsForMilestone lists labels for every issue in a milestone.
//
// GitHub API docs:
// http://developer.github.com/v3/issues/labels/#get-labels-for-every-issue-in-a-milestone

// ListLabelsForMilestone 罗列某里程碑每个问题的标记.
//
// GitHub API 文档:
// http://developer.github.com/v3/issues/labels/#get-labels-for-every-issue-in-a-milestone
func (s *IssuesService) ListLabelsForMilestone(owner string, repo string, number int, opt *ListOptions) ([]Label, *Response, error)

// ListMilestones lists all milestones for a repository.
//
// GitHub API docs:
// https://developer.github.com/v3/issues/milestones/#list-milestones-for-a-repository

// ListMilestones 罗列某仓库所有的里程碑.
//
// GitHub API 文档:
// https://developer.github.com/v3/issues/milestones/#list-milestones-for-a-repository
func (s *IssuesService) ListMilestones(owner string, repo string, opt *MilestoneListOptions) ([]Milestone, *Response, error)

// ListRepositoryEvents lists events for the specified repository.
//
// GitHub API docs:
// https://developer.github.com/v3/issues/events/#list-events-for-a-repository

// ListRepositoryEvents 罗列指定仓库的事件.
//
// GitHub API 文档:
// https://developer.github.com/v3/issues/events/#list-events-for-a-repository
func (s *IssuesService) ListRepositoryEvents(owner, repo string, opt *ListOptions) ([]IssueEvent, *Response, error)

// RemoveLabelForIssue removes a label for an issue.
//
// GitHub API docs:
// http://developer.github.com/v3/issues/labels/#remove-a-label-from-an-issue

// RemoveLabelForIssue 删除某问题的一个标记.
//
// GitHub API 文档:
// http://developer.github.com/v3/issues/labels/#remove-a-label-from-an-issue
func (s *IssuesService) RemoveLabelForIssue(owner string, repo string, number int, label string) (*Response, error)

// RemoveLabelForIssue removes a label for an issue.
//
// GitHub API docs:
// http://developer.github.com/v3/issues/labels/#remove-a-label-from-an-issue

// RemoveLabelForIssue 删除某问题的所有标记.
//
// GitHub API 文档:
// http://developer.github.com/v3/issues/labels/#remove-a-label-from-an-issue
func (s *IssuesService) RemoveLabelsForIssue(owner string, repo string, number int) (*Response, error)

// ReplaceLabelsForIssue replaces all labels for an issue.
//
// GitHub API docs:
// http://developer.github.com/v3/issues/labels/#replace-all-labels-for-an-issue

// ReplaceLabelsForIssue 替换某问题的所有标记.
//
// GitHub API 文档:
// http://developer.github.com/v3/issues/labels/#replace-all-labels-for-an-issue
func (s *IssuesService) ReplaceLabelsForIssue(owner string, repo string, number int, labels []string) ([]Label, *Response, error)

// Key represents a public SSH key used to authenticate a user or deploy script.

// Key 表示公共 SSH 授权密钥, 用于用户或部署脚本.
type Key struct {
	ID    *int    `json:"id,omitempty"`
	Key   *string `json:"key,omitempty"`
	URL   *string `json:"url,omitempty"`
	Title *string `json:"title,omitempty"`
}

func (k Key) String() string

// Label represents a GitHib label on an Issue

// Label 表示某问题的 GitHub 标记.
type Label struct {
	URL   *string `json:"url,omitempty"`
	Name  *string `json:"name,omitempty"`
	Color *string `json:"color,omitempty"`
}

func (l Label) String() string

// ListContributorsOptions specifies the optional parameters to the
// RepositoriesService.ListContributors method.

// ListContributorsOptions 指定 RepositoriesService.ListContributors 方法的可选参数.
type ListContributorsOptions struct {
	// Include anonymous contributors in results or not

	// 结果是否包含匿名贡献者
	Anon string `url:"anon,omitempty"`

	ListOptions
}

// ListMembersOptions specifies optional parameters to the
// OrganizationsService.ListMembers method.

// ListMembersOptions 指定 OrganizationsService.ListMembers 方法的可选参数.
type ListMembersOptions struct {
	// If true (or if the authenticated user is not an owner of the
	// organization), list only publicly visible members.

	// 如果为 true (或者授权用户不是组织的拥有者), 仅罗列公共可见成员.
	PublicOnly bool `url:"-"`

	// Filter members returned in the list.  Possible values are:
	// 2fa_disabled, all.  Default is "all".

	// 过滤返回成员列表. 可能的值为: 2fa_disabled, all.  缺省为 "all".
	Filter string `url:"filter,omitempty"`

	ListOptions
}

// ListOptions specifies the optional parameters to various List methods that
// support pagination.

// ListOptions 指定列表方法支持分页的可选参数.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.

	// 对于分页结果集, 结果索引页码.
	Page int `url:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.

	// 对于分页结果集, 每页结果包含数量.
	PerPage int `url:"per_page,omitempty"`
}

// ListOrgMembershipsOptions specifies optional parameters to the
// OrganizationsService.ListOrgMemberships method.

// ListOrgMembershipsOptions 指定 OrganizationsService.ListOrgMemberships 方法的可选参数.
type ListOrgMembershipsOptions struct {
	// Filter memberships to include only those withe the specified state.
	// Possible values are: "active", "pending".

	// 过滤仅含有指定状态的成员. 可能的值有: "active", "pending".
	State string `url:"state,omitempty"`

	ListOptions
}

// MarkdownOptions specifies optional parameters to the Markdown method.

// MarkdownOptions 为 Markdown 方法指定可选参数.
type MarkdownOptions struct {
	// Mode identifies the rendering mode.  Possible values are:
	//   markdown - render a document as plain Markdown, just like
	//   README files are rendered.
	//
	//   gfm - to render a document as user-content, e.g. like user
	//   comments or issues are rendered. In GFM mode, hard line breaks are
	//   always taken into account, and issue and user mentions are linked
	//   accordingly.
	//
	// Default is "markdown".

	// Mode 标识渲染模式. 可能的值有:
	//   markdown - 以纯 Markdown 渲染一个文档, 就像渲染 README 文件那样.
	//
	//   gfm - 以 user-content  渲染一个文档, 例如用户评论或问题会被渲染.
	//   在 GFM 模式下, 总是考虑到硬分行以及用户提及的响应连接.
	Mode string

	// Context identifies the repository context.  Only taken into account
	// when rendering as "gfm".

	// Context 标记表示内容. 仅当以 "gfm" 渲染时才考虑.
	Context string
}

// Match represents a single text match.

// Match 表示单个的文本匹配.
type Match struct {
	Text    *string `json:"text,omitempty"`
	Indices []int   `json:"indices,omitempty"`
}

// Membership represents the status of a user's membership in an organization or
// team.

// Membership 表示一个组织或团队中的一个成员用户状态.
type Membership struct {
	URL *string `json:"url,omitempty"`

	// State is the user's status within the organization or team.
	// Possible values are: "active", "pending"

	// State 是组织或团队中的一个用户状态. 可能的值有: "active", "pending".
	State *string `json:"state,omitempty"`

	// TODO(willnorris): add docs
	Role *string `json:"role,omitempty"`

	// For organization membership, the API URL of the organization.

	// 对于组织成员, 该组织的 API URL.
	OrganizationURL *string `json:"organization_url,omitempty"`

	// For organization membership, the organization the membership is for.

	// 对于组织成员, 该成员具体组织对象.
	Organization *Organization `json:"organization,omitempty"`

	// For organization membership, the user the membership is for.

	// 对于组织成员, 该成员具体用户对象.
	User *User `json:"user,omitempty"`
}

func (m Membership) String() string

// Milestone represents a Github repository milestone.

// Milestone 表示一个 Github 仓库里程碑.
type Milestone struct {
	URL          *string    `json:"url,omitempty"`
	Number       *int       `json:"number,omitempty"`
	State        *string    `json:"state,omitempty"`
	Title        *string    `json:"title,omitempty"`
	Description  *string    `json:"description,omitempty"`
	Creator      *User      `json:"creator,omitempty"`
	OpenIssues   *int       `json:"open_issues,omitempty"`
	ClosedIssues *int       `json:"closed_issues,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	DueOn        *time.Time `json:"due_on,omitempty"`
}

func (m Milestone) String() string

// MilestoneListOptions specifies the optional parameters to the
// IssuesService.ListMilestones method.

// MilestoneListOptions 指定 IssuesService.ListMilestones 方法的可选参数.
type MilestoneListOptions struct {
	// State filters milestones based on their state. Possible values are:
	// open, closed. Default is "open".

	// State 过滤里程碑, 基于他们的状态.
	State string `url:"state,omitempty"`

	// Sort specifies how to sort milestones. Possible values are: due_date, completeness.
	// Default value is "due_date".

	// 指定如何排序里程碑. 可能的值有: due_date, completeness. 缺省值为 "due_date".
	Sort string `url:"sort,omitempty"`

	// Direction in which to sort milestones. Possible values are: asc, desc.
	// Default is "asc".

	// 里程碑排序方向. 可能的值有: asc, desc. 缺省值为 "asc".
	Direction string `url:"direction,omitempty"`
}

// NewPullRequest represents a new pull request to be created.

// NewPullRequest 表示新建立一个上拉请求.
type NewPullRequest struct {
	Title *string `json:"title,omitempty"`
	Head  *string `json:"head,omitempty"`
	Base  *string `json:"base,omitempty"`
	Body  *string `json:"body,omitempty"`
	Issue *int    `json:"issue,omitempty"`
}

// Notification identifies a GitHub notification for a user.

// Notification 标识某用户的一个 GitHub 通知.
type Notification struct {
	ID         *string              `json:"id,omitempty"`
	Repository *Repository          `json:"repository,omitempty"`
	Subject    *NotificationSubject `json:"subject,omitempty"`

	// Reason identifies the event that triggered the notification.
	//
	// GitHub API Docs: https://developer.github.com/v3/activity/notifications/#notification-reasons

	// Reason 标识触发通知事件的原因.
	//
	// GitHub API 文档: https://developer.github.com/v3/activity/notifications/#notification-reasons
	Reason *string `json:"reason,omitempty"`

	Unread     *bool      `json:"unread,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	LastReadAt *time.Time `json:"last_read_at,omitempty"`
	URL        *string    `json:"url,omitempty"`
}

// NotificationListOptions specifies the optional parameters to the
// ActivityService.ListNotifications method.

// NotificationListOptions 指定 ActivityService.ListNotifications 方法的可选参数.
type NotificationListOptions struct {
	All           bool      `url:"all,omitempty"`
	Participating bool      `url:"participating,omitempty"`
	Since         time.Time `url:"since,omitempty"`
}

// NotificationSubject identifies the subject of a notification.

// NotificationSubject 标识通知的主题.
type NotificationSubject struct {
	Title            *string `json:"title,omitempty"`
	URL              *string `json:"url,omitempty"`
	LatestCommentURL *string `json:"latest_comment_url,omitempty"`
	Type             *string `json:"type,omitempty"`
}

// Organization represents a GitHub organization account.

// Organization 表示一个 GitHub 组织账户.
type Organization struct {
	Login             *string    `json:"login,omitempty"`
	ID                *int       `json:"id,omitempty"`
	AvatarURL         *string    `json:"avatar_url,omitempty"`
	HTMLURL           *string    `json:"html_url,omitempty"`
	Name              *string    `json:"name,omitempty"`
	Company           *string    `json:"company,omitempty"`
	Blog              *string    `json:"blog,omitempty"`
	Location          *string    `json:"location,omitempty"`
	Email             *string    `json:"email,omitempty"`
	PublicRepos       *int       `json:"public_repos,omitempty"`
	PublicGists       *int       `json:"public_gists,omitempty"`
	Followers         *int       `json:"followers,omitempty"`
	Following         *int       `json:"following,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	TotalPrivateRepos *int       `json:"total_private_repos,omitempty"`
	OwnedPrivateRepos *int       `json:"owned_private_repos,omitempty"`
	PrivateGists      *int       `json:"private_gists,omitempty"`
	DiskUsage         *int       `json:"disk_usage,omitempty"`
	Collaborators     *int       `json:"collaborators,omitempty"`
	BillingEmail      *string    `json:"billing_email,omitempty"`
	Type              *string    `json:"type,omitempty"`
	Plan              *Plan      `json:"plan,omitempty"`

	// API URLs
	URL              *string `json:"url,omitempty"`
	EventsURL        *string `json:"events_url,omitempty"`
	MembersURL       *string `json:"members_url,omitempty"`
	PublicMembersURL *string `json:"public_members_url,omitempty"`
	ReposURL         *string `json:"repos_url,omitempty"`
}

func (o Organization) String() string

// OrganizationsService provides access to the organization related functions in
// the GitHub API.
//
// GitHub API docs: http://developer.github.com/v3/orgs/

// OrganizationsService 提供访问组织 GitHub API 相关功能.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/
type OrganizationsService struct {
	// contains filtered or unexported fields
}

// AddTeamMember adds a user to a team.
//
// GitHub API docs: http://developer.github.com/v3/orgs/teams/#add-team-member

// AddTeamMember 添加一个用户到团队.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/teams/#add-team-member
func (s *OrganizationsService) AddTeamMember(team int, user string) (*Response, error)

// AddTeamMembership adds or invites a user to a team.
//
// In order to add a membership between a user and a team, the authenticated user
// must have 'admin' permissions to the team or be an owner of the organization
// that the team is associated with.
//
// If the user is already a part of the team's organization (meaning they're on at
// least one other team in the organization), this endpoint will add the user to
// the team.
//
// If the user is completely unaffiliated with the team's organization (meaning
// they're on none of the organization's teams), this endpoint will send an
// invitation to the user via email. This newly-created membership will be in the
// "pending" state until the user accepts the invitation, at which point the
// membership will transition to the "active" state and the user will be added as a
// member of the team.
//
// GitHub API docs: https://developer.github.com/v3/orgs/teams/#add-team-membership

// AddTeamMembership 添加或邀请一个用户到团队.
//
// 为了增加用户和团队之间的成员, 与授权用户必须具有关联团队的 'admin' 权限,
// 或者是组织的拥有者.
//
// 如果用户已经是组织团队的一份子(也就是说他至少在组织中的一个团队), 这将添加用户到团队.
//
// 如果用户完全不属于组织团队(也就是说他不在组织的团队中), 这将给用户发送邀请邮件.
// 新成员为 "pending" 状态, 直到该用户接受邀请, 此时用户添加到团队成员并转为 "active" 状态.
//
// GitHub API 文档: https://developer.github.com/v3/orgs/teams/#add-team-membership
func (s *OrganizationsService) AddTeamMembership(team int, user string) (*Membership, *Response, error)

// AddTeamRepo adds a repository to be managed by the specified team. The specified
// repository must be owned by the organization to which the team belongs, or a
// direct fork of a repository owned by the organization.
//
// GitHub API docs: http://developer.github.com/v3/orgs/teams/#add-team-repo

// AddTeamRepo 添加一个仓库给指定的团队管理. 指定的仓库必须由该团队所属的组织所拥有,
// 或者组织拥有直接 fork 的仓库.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/teams/#add-team-repo
func (s *OrganizationsService) AddTeamRepo(team int, owner string, repo string) (*Response, error)

// ConcealMembership conceals a user's membership in an organization.
//
// GitHub API docs:
// http://developer.github.com/v3/orgs/members/#conceal-a-users-membership

// ConcealMembership 隐藏组织中的成员用户.
//
// GitHub API 文档:
// http://developer.github.com/v3/orgs/members/#conceal-a-users-membership
func (s *OrganizationsService) ConcealMembership(org, user string) (*Response, error)

// CreateTeam creates a new team within an organization.
//
// GitHub API docs: http://developer.github.com/v3/orgs/teams/#create-team

// CreateTeam 新建一个组织团队.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/teams/#create-team
func (s *OrganizationsService) CreateTeam(org string, team *Team) (*Team, *Response, error)

// DeleteTeam deletes a team.
//
// GitHub API docs: http://developer.github.com/v3/orgs/teams/#delete-team

// DeleteTeam 删除一个团队.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/teams/#delete-team
func (s *OrganizationsService) DeleteTeam(team int) (*Response, error)

// Edit an organization.
//
// GitHub API docs: http://developer.github.com/v3/orgs/#edit-an-organization

// Edit 一个组织.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/#edit-an-organization
func (s *OrganizationsService) Edit(name string, org *Organization) (*Organization, *Response, error)

// EditOrgMembership edits the membership for the authenticated user for the
// specified organization.
//
// GitHub API docs:
// https://developer.github.com/v3/orgs/members/#edit-your-organization-membership

// EditOrgMembership 编辑授权用户指定组织的一个成员.
//
// GitHub API 文档:
// https://developer.github.com/v3/orgs/members/#edit-your-organization-membership
func (s *OrganizationsService) EditOrgMembership(org string, membership *Membership) (*Membership, *Response, error)

// EditTeam edits a team.
//
// GitHub API docs: http://developer.github.com/v3/orgs/teams/#edit-team

// EditTeam 编辑一个团队.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/teams/#edit-team
func (s *OrganizationsService) EditTeam(id int, team *Team) (*Team, *Response, error)

// Get fetches an organization by name.
//
// GitHub API docs: http://developer.github.com/v3/orgs/#get-an-organization

// Get 以 name 获取一个组织.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/#get-an-organization
func (s *OrganizationsService) Get(org string) (*Organization, *Response, error)

// GetOrgMembership gets the membership for the authenticated user for the
// specified organization.
//
// GitHub API docs:
// https://developer.github.com/v3/orgs/members/#get-your-organization-membership

// GetOrgMembership 获取授权用户在指定组织的成员对象.
//
// GitHub API 文档:
// https://developer.github.com/v3/orgs/members/#get-your-organization-membership
func (s *OrganizationsService) GetOrgMembership(org string) (*Membership, *Response, error)

// GetTeam fetches a team by ID.
//
// GitHub API docs: http://developer.github.com/v3/orgs/teams/#get-team

// GetTeam 以 ID 获取一个团队.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/teams/#get-team
func (s *OrganizationsService) GetTeam(team int) (*Team, *Response, error)

// GetTeamMembership returns the membership status for a user in a team.
//
// GitHub API docs: https://developer.github.com/v3/orgs/teams/#get-team-membership

// GetTeamMembership 返回某用户在团队中的成员状态.
//
// GitHub API 文档: https://developer.github.com/v3/orgs/teams/#get-team-membership
func (s *OrganizationsService) GetTeamMembership(team int, user string) (*Membership, *Response, error)

// IsMember checks if a user is a member of an organization.
//
// GitHub API docs: http://developer.github.com/v3/orgs/members/#check-membership

// IsMember 检查某用户是否为一个组织的成员.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/members/#check-membership
func (s *OrganizationsService) IsMember(org, user string) (bool, *Response, error)

// IsPublicMember checks if a user is a public member of an organization.
//
// GitHub API docs:
// http://developer.github.com/v3/orgs/members/#check-public-membership

// IsPublicMember 检查某用户是否为一个组织的公开成员.
//
// GitHub API 文档:
// http://developer.github.com/v3/orgs/members/#check-public-membership
func (s *OrganizationsService) IsPublicMember(org, user string) (bool, *Response, error)

// IsTeamMember checks if a user is a member of the specified team.
//
// GitHub API docs: http://developer.github.com/v3/orgs/teams/#get-team-member

// IsTeamMember 检查某用户是否为一个团队成员.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/teams/#get-team-member
func (s *OrganizationsService) IsTeamMember(team int, user string) (bool, *Response, error)

// IsTeamRepo checks if a team manages the specified repository.
//
// GitHub API docs: http://developer.github.com/v3/orgs/teams/#get-team-repo

// IsTeamRepo 检查指定仓库是否被某团队管理.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/teams/#get-team-repo
func (s *OrganizationsService) IsTeamRepo(team int, owner string, repo string) (bool, *Response, error)

// List the organizations for a user. Passing the empty string will list
// organizations for the authenticated user.
//
// GitHub API docs: http://developer.github.com/v3/orgs/#list-user-organizations

// List 罗列某用户所在的组织. 传递空字符串将罗列授权用户所在的组织.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/#list-user-organizations
func (s *OrganizationsService) List(user string, opt *ListOptions) ([]Organization, *Response, error)

// ListMembers lists the members for an organization. If the authenticated user is
// an owner of the organization, this will return both concealed and public
// members, otherwise it will only return public members.
//
// GitHub API docs: http://developer.github.com/v3/orgs/members/#members-list

// ListMembers 罗列某组织成员. 如果授权用户是该组织的拥有者,
// 它将返回隐藏的和公开的成员, 否则只返回公开成员.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/members/#members-list
func (s *OrganizationsService) ListMembers(org string, opt *ListMembersOptions) ([]User, *Response, error)

// ListOrgMemberships lists the organization memberships for the authenticated
// user.
//
// GitHub API docs:
// https://developer.github.com/v3/orgs/members/#list-your-organization-memberships

// ListOrgMemberships l罗列授权用户的组织成员.
//
// GitHub API 文档:
// https://developer.github.com/v3/orgs/members/#list-your-organization-memberships
func (s *OrganizationsService) ListOrgMemberships(opt *ListOrgMembershipsOptions) ([]Membership, *Response, error)

// ListTeamMembers lists all of the users who are members of the specified team.
//
// GitHub API docs: http://developer.github.com/v3/orgs/teams/#list-team-members

// ListTeamMembers 罗列指定团队所有成员的用户.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/teams/#list-team-members
func (s *OrganizationsService) ListTeamMembers(team int, opt *ListOptions) ([]User, *Response, error)

// ListTeamRepos lists the repositories that the specified team has access to.
//
// GitHub API docs: http://developer.github.com/v3/orgs/teams/#list-team-repos

// ListTeamRepos 罗列指定团队可存取的仓库.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/teams/#list-team-repos
func (s *OrganizationsService) ListTeamRepos(team int, opt *ListOptions) ([]Repository, *Response, error)

// ListTeams lists all of the teams for an organization.
//
// GitHub API docs: http://developer.github.com/v3/orgs/teams/#list-teams

// ListTeams 罗列某组织所有的团队.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/teams/#list-teams
func (s *OrganizationsService) ListTeams(org string, opt *ListOptions) ([]Team, *Response, error)

// ListUserTeams lists a user's teams GitHub API docs:
// https://developer.github.com/v3/orgs/teams/#list-user-teams

// ListUserTeams 罗列用户所在的团队. GitHub API 文档:
// https://developer.github.com/v3/orgs/teams/#list-user-teams
func (s *OrganizationsService) ListUserTeams(opt *ListOptions) ([]Team, *Response, error)

// PublicizeMembership publicizes a user's membership in an organization.
//
// GitHub API docs:
// http://developer.github.com/v3/orgs/members/#publicize-a-users-membership

// PublicizeMembership 公开组织中的某成员用户.
//
// GitHub API 文档:
// http://developer.github.com/v3/orgs/members/#publicize-a-users-membership
func (s *OrganizationsService) PublicizeMembership(org, user string) (*Response, error)

// RemoveMember removes a user from all teams of an organization.
//
// GitHub API docs: http://developer.github.com/v3/orgs/members/#remove-a-member

// RemoveMember 从某组织所有团队中删除一个用户.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/members/#remove-a-member
func (s *OrganizationsService) RemoveMember(org, user string) (*Response, error)

// RemoveTeamMember removes a user from a team.
//
// GitHub API docs: http://developer.github.com/v3/orgs/teams/#remove-team-member

// RemoveTeamMember 从某团队中删除一个用户(过时的).
//
// GitHub API 文档: http://developer.github.com/v3/orgs/teams/#remove-team-member
func (s *OrganizationsService) RemoveTeamMember(team int, user string) (*Response, error)

// RemoveTeamMembership removes a user from a team.
//
// GitHub API docs:
// https://developer.github.com/v3/orgs/teams/#remove-team-membership

// RemoveTeamMembership 从某团队中删除一个用户.
//
// GitHub API 文档:
// https://developer.github.com/v3/orgs/teams/#remove-team-membership
func (s *OrganizationsService) RemoveTeamMembership(team int, user string) (*Response, error)

// RemoveTeamRepo removes a repository from being managed by the specified team.
// Note that this does not delete the repository, it just removes it from the team.
//
// GitHub API docs: http://developer.github.com/v3/orgs/teams/#remove-team-repo

// RemoveTeamRepo 删除指定团队所管理的某个仓库.
// 注意这不是删除该仓库, 只是从团队中移除.
//
// GitHub API 文档: http://developer.github.com/v3/orgs/teams/#remove-team-repo
func (s *OrganizationsService) RemoveTeamRepo(team int, owner string, repo string) (*Response, error)

// Pages represents a GitHub Pages site configuration.

// Pages 表示一个 GitHub Pages 站点配置.
type Pages struct {
	URL       *string `json:"url,omitempty"`
	Status    *string `json:"status,omitempty"`
	CNAME     *string `json:"cname,omitempty"`
	Custom404 *bool   `json:"custom_404,omitempty"`
}

// PagesBuild represents the build information for a GitHub Pages site.

// PagesBuild 表示一个 GitHub Pages 站点的构建信息.
type PagesBuild struct {
	URL       *string     `json:"url,omitempty"`
	Status    *string     `json:"status,omitempty"`
	Error     *PagesError `json:"error,omitempty"`
	Pusher    *User       `json:"pusher,omitempty"`
	Commit    *string     `json:"commit,omitempty"`
	Duration  *int        `json:"duration,omitempty"`
	CreatedAt *Timestamp  `json:"created_at,omitempty"`
	UpdatedAt *Timestamp  `json:"created_at,omitempty"`
}

// PagesError represents a build error for a GitHub Pages site.

// PagesError 表示一个 GitHub Pages 站点的构建错误.
type PagesError struct {
	Message *string `json:"message,omitempty"`
}

// Plan represents the payment plan for an account. See plans at
// https://github.com/plans.

// Plan 表示一个账户的付款计划. 计划参见
// https://github.com/plans.
type Plan struct {
	Name          *string `json:"name,omitempty"`
	Space         *int    `json:"space,omitempty"`
	Collaborators *int    `json:"collaborators,omitempty"`
	PrivateRepos  *int    `json:"private_repos,omitempty"`
}

func (p Plan) String() string

// PullRequest represents a GitHub pull request on a repository.

// PullRequest 表示一个 GitHub 仓库的上拉请求.
type PullRequest struct {
	Number       *int       `json:"number,omitempty"`
	State        *string    `json:"state,omitempty"`
	Title        *string    `json:"title,omitempty"`
	Body         *string    `json:"body,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	ClosedAt     *time.Time `json:"closed_at,omitempty"`
	MergedAt     *time.Time `json:"merged_at,omitempty"`
	User         *User      `json:"user,omitempty"`
	Merged       *bool      `json:"merged,omitempty"`
	Mergeable    *bool      `json:"mergeable,omitempty"`
	MergedBy     *User      `json:"merged_by,omitempty"`
	Comments     *int       `json:"comments,omitempty"`
	Commits      *int       `json:"commits,omitempty"`
	Additions    *int       `json:"additions,omitempty"`
	Deletions    *int       `json:"deletions,omitempty"`
	ChangedFiles *int       `json:"changed_files,omitempty"`
	URL          *string    `json:"url,omitempty"`
	HTMLURL      *string    `json:"html_url,omitempty"`
	IssueURL     *string    `json:"issue_url,omitempty"`
	StatusesURL  *string    `json:"statuses_url,omitempty"`

	Head *PullRequestBranch `json:"head,omitempty"`
	Base *PullRequestBranch `json:"base,omitempty"`
}

func (p PullRequest) String() string

// PullRequestBranch represents a base or head branch in a GitHub pull request.

// PullRequestBranch 表示一个 GitHub 基础或分支的上拉请求.
type PullRequestBranch struct {
	Label *string     `json:"label,omitempty"`
	Ref   *string     `json:"ref,omitempty"`
	SHA   *string     `json:"sha,omitempty"`
	Repo  *Repository `json:"repo,omitempty"`
	User  *User       `json:"user,omitempty"`
}

// PullRequestComment represents a comment left on a pull request.

// PullRequestComment 表示一个上拉请求上的评论.
type PullRequestComment struct {
	ID        *int       `json:"id,omitempty"`
	Body      *string    `json:"body,omitempty"`
	Path      *string    `json:"path,omitempty"`
	Position  *int       `json:"position,omitempty"`
	CommitID  *string    `json:"commit_id,omitempty"`
	User      *User      `json:"user,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (p PullRequestComment) String() string

// PullRequestEvent represents the payload delivered by PullRequestEvent webhook

// PullRequestEvent 表示由 PullRequestEvent webhook 交付的有效负载.
type PullRequestEvent struct {
	Action      *string      `json:"action,omitempty"`
	Number      *int         `json:"number,omitempty"`
	PullRequest *PullRequest `json:"pull_request,omitempty"`
	Repo        *Repository  `json:"repository,omitempty"`
	Sender      *User        `json:"sender,omitempty"`
}

// PullRequestLinks object is added to the Issue object when it's an issue included
// in the IssueCommentEvent webhook payload, if the webhooks is fired by a comment
// on a PR

// PullRequestLinks 对象添加到 Issue 对象中, 当问题含有 IssueCommentEvent
// 交付的有效负载, 如果该 webhooks 被上拉请求中的某评论触发.
type PullRequestLinks struct {
	URL      *string `json:"url,omitempty"`
	HTMLURL  *string `json:"html_url,omitempty"`
	DiffURL  *string `json:"diff_url,omitempty"`
	PatchURL *string `json:"patch_url,omitempty"`
}

// PullRequestListCommentsOptions specifies the optional parameters to the
// PullRequestsService.ListComments method.

// PullRequestListCommentsOptions 指定 PullRequestsService.ListComments 方法的可选参数.
type PullRequestListCommentsOptions struct {
	// Sort specifies how to sort comments.  Possible values are: created, updated.

	// Sort 指定如何排序评论. 可能的值有: created, updated.
	Sort string `url:"sort,omitempty"`

	// Direction in which to sort comments.  Possible values are: asc, desc.

	// 评论的排序方向. 可能的值有: asc, desc.
	Direction string `url:"direction,omitempty"`

	// Since filters comments by time.

	// 过滤评论的起始时间.
	Since time.Time `url:"since,omitempty"`

	ListOptions
}

// PullRequestListOptions specifies the optional parameters to the
// PullRequestsService.List method.

// PullRequestListOptions 指定 PullRequestsService.List 方法的可选参数.
type PullRequestListOptions struct {
	// State filters pull requests based on their state.  Possible values are:
	// open, closed.  Default is "open".

	// State 过滤上拉请求, 基于他们的状态. 可能的值有:
	// open, closed. 缺省为 "open".
	State string `url:"state,omitempty"`

	// Head filters pull requests by head user and branch name in the format of:
	// "user:ref-name".

	// Head 过滤上拉请求, 通过首用户和分支名, 格式为:
	// "user:ref-name".
	Head string `url:"head,omitempty"`

	// Base filters pull requests by base branch name.

	// Base 过滤上拉请求, 通过分支名.
	Base string `url:"base,omitempty"`

	ListOptions
}

// PullRequestMergeResult represents the result of merging a pull request.

// PullRequestMergeResult 表示上拉合并请求的结果.
type PullRequestMergeResult struct {
	SHA     *string `json:"sha,omitempty"`
	Merged  *bool   `json:"merged,omitempty"`
	Message *string `json:"message,omitempty"`
}

// PullRequestsService handles communication with the pull request related methods
// of the GitHub API.
//
// GitHub API docs: http://developer.github.com/v3/pulls/

// PullRequestsService 处理 GitHub API 上拉请求相关通信方法.
//
// GitHub API 文档: http://developer.github.com/v3/pulls/
type PullRequestsService struct {
	// contains filtered or unexported fields
}

// Create a new pull request on the specified repository.
//
// GitHub API docs: https://developer.github.com/v3/pulls/#create-a-pull-request

// Create 在指定仓库新建一个上拉请求.
//
// GitHub API 文档: https://developer.github.com/v3/pulls/#create-a-pull-request
func (s *PullRequestsService) Create(owner string, repo string, pull *NewPullRequest) (*PullRequest, *Response, error)

// CreateComment creates a new comment on the specified pull request.
//
// GitHub API docs:
// https://developer.github.com/v3/pulls/comments/#get-a-single-comment

// CreateComment 在指定的上拉请求上新建一个评论.
//
// GitHub API 文档:
// https://developer.github.com/v3/pulls/comments/#get-a-single-comment
func (s *PullRequestsService) CreateComment(owner string, repo string, number int, comment *PullRequestComment) (*PullRequestComment, *Response, error)

// DeleteComment deletes a pull request comment.
//
// GitHub API docs:
// https://developer.github.com/v3/pulls/comments/#delete-a-comment

// DeleteComment 删除一个上拉请求评论.
//
// GitHub API 文档:
// https://developer.github.com/v3/pulls/comments/#delete-a-comment
func (s *PullRequestsService) DeleteComment(owner string, repo string, number int) (*Response, error)

// Edit a pull request.
//
// GitHub API docs: https://developer.github.com/v3/pulls/#update-a-pull-request

// Edit 编辑一个上拉请求.
//
// GitHub API 文档: https://developer.github.com/v3/pulls/#update-a-pull-request
func (s *PullRequestsService) Edit(owner string, repo string, number int, pull *PullRequest) (*PullRequest, *Response, error)

// EditComment updates a pull request comment.
//
// GitHub API docs: https://developer.github.com/v3/pulls/comments/#edit-a-comment

// EditComment 更新一个上拉请求评论.
//
// GitHub API 文档: https://developer.github.com/v3/pulls/comments/#edit-a-comment
func (s *PullRequestsService) EditComment(owner string, repo string, number int, comment *PullRequestComment) (*PullRequestComment, *Response, error)

// Get a single pull request.
//
// GitHub API docs:
// https://developer.github.com/v3/pulls/#get-a-single-pull-request

// Get 获取单个上拉请求.
//
// GitHub API 文档:
// https://developer.github.com/v3/pulls/#get-a-single-pull-request
func (s *PullRequestsService) Get(owner string, repo string, number int) (*PullRequest, *Response, error)

// GetComment fetches the specified pull request comment.
//
// GitHub API docs:
// https://developer.github.com/v3/pulls/comments/#get-a-single-comment

// GetComment 获取指定上拉请求的评论.
//
// GitHub API 文档:
// https://developer.github.com/v3/pulls/comments/#get-a-single-comment
func (s *PullRequestsService) GetComment(owner string, repo string, number int) (*PullRequestComment, *Response, error)

// IsMerged checks if a pull request has been merged.
//
// GitHub API docs:
// https://developer.github.com/v3/pulls/#get-if-a-pull-request-has-been-merged

// IsMerged 检查一个上拉请求是否被合并.
//
// GitHub API 文档:
// https://developer.github.com/v3/pulls/#get-if-a-pull-request-has-been-merged
func (s *PullRequestsService) IsMerged(owner string, repo string, number int) (bool, *Response, error)

// List the pull requests for the specified repository.
//
// GitHub API docs: http://developer.github.com/v3/pulls/#list-pull-requests

// List 罗列指定仓库的上拉请求.
//
// GitHub API 文档: http://developer.github.com/v3/pulls/#list-pull-requests
func (s *PullRequestsService) List(owner string, repo string, opt *PullRequestListOptions) ([]PullRequest, *Response, error)

// ListComments lists all comments on the specified pull request. Specifying a pull
// request number of 0 will return all comments on all pull requests for the
// repository.
//
// GitHub API docs:
// https://developer.github.com/v3/pulls/comments/#list-comments-on-a-pull-request

// ListComments 罗列指定上拉请求的所有评论. 指定上拉请求号为 0 将返回该仓库所有
// 上拉请求的所有评论.
//
// GitHub API 文档:
// https://developer.github.com/v3/pulls/comments/#list-comments-on-a-pull-request
func (s *PullRequestsService) ListComments(owner string, repo string, number int, opt *PullRequestListCommentsOptions) ([]PullRequestComment, *Response, error)

// ListCommits lists the commits in a pull request.
//
// GitHub API docs:
// https://developer.github.com/v3/pulls/#list-commits-on-a-pull-request

// ListCommits 罗列一个上拉请求的提交.
//
// GitHub API 文档:
// https://developer.github.com/v3/pulls/#list-commits-on-a-pull-request
func (s *PullRequestsService) ListCommits(owner string, repo string, number int, opt *ListOptions) ([]RepositoryCommit, *Response, error)

// ListFiles lists the files in a pull request.
//
// GitHub API docs: https://developer.github.com/v3/pulls/#list-pull-requests-files

// ListFiles 罗列一个上拉请求中的文件.
//
// GitHub API 文档: https://developer.github.com/v3/pulls/#list-pull-requests-files
func (s *PullRequestsService) ListFiles(owner string, repo string, number int, opt *ListOptions) ([]CommitFile, *Response, error)

// Merge a pull request (Merge Button™).
//
// GitHub API docs:
// https://developer.github.com/v3/pulls/#merge-a-pull-request-merge-buttontrade

// Merge 合并一个上拉请求 (Merge Button™).
//
// GitHub API 文档:
// https://developer.github.com/v3/pulls/#merge-a-pull-request-merge-buttontrade
func (s *PullRequestsService) Merge(owner string, repo string, number int, commitMessage string) (*PullRequestMergeResult, *Response, error)

// PunchCard respresents the number of commits made during a given hour of a day of
// thew eek.

// PunchCard 表示每天特定时间段内的提交数量.
type PunchCard struct {
	Day     *int // Day of the week (0-6: =Sunday - Saturday).
	Hour    *int // Hour of day (0-23).
	Commits *int // Number of commits.
}

// PushEvent represents a git push to a GitHub repository.
//
// GitHub API docs: http://developer.github.com/v3/activity/events/types/#pushevent

// PushEvent 表示一个 GitHub 仓库推送.
//
// GitHub API 文档: http://developer.github.com/v3/activity/events/types/#pushevent
type PushEvent struct {
	PushID  *int              `json:"push_id,omitempty"`
	Head    *string           `json:"head,omitempty"`
	Ref     *string           `json:"ref,omitempty"`
	Size    *int              `json:"size,omitempty"`
	Commits []PushEventCommit `json:"commits,omitempty"`
	Repo    *Repository       `json:"repository,omitempty"`
}

func (p PushEvent) String() string

// PushEventCommit represents a git commit in a GitHub PushEvent.

// PushEventCommit 表示 GitHub PushEvent 中的一个 git 提交.
type PushEventCommit struct {
	SHA      *string       `json:"sha,omitempty"`
	Message  *string       `json:"message,omitempty"`
	Author   *CommitAuthor `json:"author,omitempty"`
	URL      *string       `json:"url,omitempty"`
	Distinct *bool         `json:"distinct,omitempty"`
	Added    []string      `json:"added,omitempty"`
	Removed  []string      `json:"removed,omitempty"`
	Modified []string      `json:"modified,omitempty"`
}

func (p PushEventCommit) String() string

// Rate represents the rate limit for the current client.

// Rate 表示当前客户端的频次限制.
type Rate struct {
	// The number of requests per hour the client is currently limited to.

	// 该客户端当前没小时限制请求数.
	Limit int `json:"limit"`

	// The number of remaining requests the client can make this hour.

	// 客户端在这个小时内剩余的请求数.
	Remaining int `json:"remaining"`

	// The time at which the current rate limit will reset.

	// 当前频次限制在何时将被重置.
	Reset Timestamp `json:"reset"`
}

func (r Rate) String() string

// RateLimits represents the rate limits for the current client.

// RateLimits 表示当前客户端的频次限制.
type RateLimits struct {
	// The rate limit for non-search API requests.  Unauthenticated
	// requests are limited to 60 per hour.  Authenticated requests are
	// limited to 5,000 per hour.

	// 非搜索 API 请求的频次限制. 未授权请求每小时限制 60 次.
	// 授权请求每小时限制 5,000 次.
	Core *Rate `json:"core"`

	// The rate limit for search API requests.  Unauthenticated requests
	// are limited to 5 requests per minutes.  Authenticated requests are
	// limited to 20 per minute.
	//
	// GitHub API docs: https://developer.github.com/v3/search/#rate-limit

	// 搜索 API 请求的频次限制. 未授权请求每分钟限制 5 次请求.
	// 授权请求每分钟限制 20 次请求.
	//
	// GitHub API 文档: https://developer.github.com/v3/search/#rate-limit
	Search *Rate `json:"search"`
}

func (r RateLimits) String() string

// Reference represents a GitHub reference.

// Reference 表示一个 GitHub 引用.
type Reference struct {
	Ref    *string    `json:"ref"`
	URL    *string    `json:"url"`
	Object *GitObject `json:"object"`
}

func (r Reference) String() string

// ReferenceListOptions specifies optional parameters to the GitService.ListRefs
// method.

// ReferenceListOptions 指定 GitService.ListRefs 方法的可选参数.
type ReferenceListOptions struct {
	Type string `url:"-"`

	ListOptions
}

// ReleaseAsset represents a Github release asset in a repository.

// ReleaseAsset 表示某仓库的一个 Github 正式版本资源.
type ReleaseAsset struct {
	ID                 *int       `json:"id,omitempty"`
	URL                *string    `json:"url,omitempty"`
	Name               *string    `json:"name,omitempty"`
	Label              *string    `json:"label,omitempty"`
	State              *string    `json:"state,omitempty"`
	ContentType        *string    `json:"content_type,omitempty"`
	Size               *int       `json:"size,omitempty"`
	DownloadCount      *int       `json:"download_count,omitempty"`
	CreatedAt          *Timestamp `json:"created_at,omitempty"`
	UpdatedAt          *Timestamp `json:"updated_at,omitempty"`
	BrowserDownloadURL *string    `json:"browser_download_url,omitempty"`
	Uploader           *User      `json:"uploader,omitempty"`
}

func (r ReleaseAsset) String() string

// RepoStatus represents the status of a repository at a particular reference.

// RepoStatus 表示某仓库中的一个特定引用状态.
type RepoStatus struct {
	ID  *int    `json:"id,omitempty"`
	URL *string `json:"url,omitempty"`

	// State is the current state of the repository.  Possible values are:
	// pending, success, error, or failure.

	// State 为该仓库当前状态. 可能的值有: pending, success, error,或 failure.
	State *string `json:"state,omitempty"`

	// TargetURL is the URL of the page representing this status.  It will be
	// linked from the GitHub UI to allow users to see the source of the status.

	// TargetURL 为代表该状态页面的 URL. 它将让用户看到来自 GitHub UI 链接的状态源码.
	TargetURL *string `json:"target_url,omitempty"`

	// Description is a short high level summary of the status.

	// Description 为该状态的高级简短摘要.
	Description *string `json:"description,omitempty"`

	// A string label to differentiate this status from the statuses of other systems.

	// 字符串标记以区分该状态来自其它系统的状态.
	Context *string `json:"context,omitempty"`

	Creator   *User      `json:"creator,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (r RepoStatus) String() string

// RepositoriesSearchResult represents the result of a repositories search.

// RepositoriesSearchResult 表示仓库搜索的结果.
type RepositoriesSearchResult struct {
	Total        *int         `json:"total_count,omitempty"`
	Repositories []Repository `json:"items,omitempty"`
}

// RepositoriesService handles communication with the repository related methods of
// the GitHub API.
//
// GitHub API docs: http://developer.github.com/v3/repos/

// RepositoriesService 处理与 GitHub API 仓库相关的通信方法.
//
// GitHub API 文档: http://developer.github.com/v3/repos/
type RepositoriesService struct {
	// contains filtered or unexported fields
}

// AddCollaborator adds the specified Github user as collaborator to the given
// repo.
//
// GitHub API docs:
// http://developer.github.com/v3/repos/collaborators/#add-collaborator

// AddCollaborator 添加指定的 Github 用户为给定仓库合作者.
//
// GitHub API 文档:
// http://developer.github.com/v3/repos/collaborators/#add-collaborator
func (s *RepositoriesService) AddCollaborator(owner, repo, user string) (*Response, error)

// CompareCommits compares a range of commits with each other. todo: support media
// formats - https://github.com/google/go-github/issues/6
//
// GitHub API docs:
// http://developer.github.com/v3/repos/commits/index.html#compare-two-commits

// CompareCommits 在彼此提交范围间进行对比.
// todo: 支持媒体格式 - https://github.com/google/go-github/issues/6
//
// GitHub API 文档:
// http://developer.github.com/v3/repos/commits/index.html#compare-two-commits
func (s *RepositoriesService) CompareCommits(owner, repo string, base, head string) (*CommitsComparison, *Response, error)

// Create a new repository. If an organization is specified, the new repository
// will be created under that org. If the empty string is specified, it will be
// created for the authenticated user.
//
// GitHub API docs: http://developer.github.com/v3/repos/#create

// Create 新建一个仓库. 如果指定了组织, 新仓库建于该组织之下.
// 如果为空字符串,  它建于授权用户.
//
// GitHub API 文档: http://developer.github.com/v3/repos/#create
func (s *RepositoriesService) Create(org string, repo *Repository) (*Repository, *Response, error)

// CreateComment creates a comment for the given commit. Note: GitHub allows for
// comments to be created for non-existing files and positions.
//
// GitHub API docs:
// http://developer.github.com/v3/repos/comments/#create-a-commit-comment

// CreateComment 为提交创建评论.
// Note: GitHub 上允许为不存在的文件和位置创建评论.
//
// GitHub API 文档:
// http://developer.github.com/v3/repos/comments/#create-a-commit-comment
func (s *RepositoriesService) CreateComment(owner, repo, sha string, comment *RepositoryComment) (*RepositoryComment, *Response, error)

// CreateDeployment creates a new deployment for a repository.
//
// GitHub API docs:
// https://developer.github.com/v3/repos/deployments/#create-a-deployment

// CreateDeployment 为仓库新建部署.
//
// GitHub API 文档:
// https://developer.github.com/v3/repos/deployments/#create-a-deployment
func (s *RepositoriesService) CreateDeployment(owner, repo string, request *DeploymentRequest) (*Deployment, *Response, error)

// CreateDeploymentStatus creates a new status for a deployment.
//
// GitHub API docs:
// https://developer.github.com/v3/repos/deployments/#create-a-deployment-status

// CreateDeploymentStatus 为部署新建状态.
//
// GitHub API 文档:
// https://developer.github.com/v3/repos/deployments/#create-a-deployment-status
func (s *RepositoriesService) CreateDeploymentStatus(owner, repo string, deployment int, request *DeploymentStatusRequest) (*DeploymentStatus, *Response, error)

// CreateFile creates a new file in a repository at the given path and returns the
// commit and file metadata.
//
// GitHub API docs: http://developer.github.com/v3/repos/contents/#create-a-file

// CreateFile 以给定路径在某仓库新建一个文件并返回该提交和文件元数据.
//
// GitHub API 文档: http://developer.github.com/v3/repos/contents/#create-a-file
func (s *RepositoriesService) CreateFile(owner, repo, path string, opt *RepositoryContentFileOptions) (*RepositoryContentResponse, *Response, error)

// CreateFork creates a fork of the specified repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/forks/#list-forks

// CreateFork 创建指定仓库的 fork.
//
// GitHub API 文档: http://developer.github.com/v3/repos/forks/#list-forks
func (s *RepositoriesService) CreateFork(owner, repo string, opt *RepositoryCreateForkOptions) (*Repository, *Response, error)

// CreateHook creates a Hook for the specified repository. Name and Config are
// required fields.
//
// GitHub API docs: http://developer.github.com/v3/repos/hooks/#create-a-hook

// CreateHook 为指定仓库创建 Hook. Name 和 Config 为必填字段.
//
// GitHub API 文档: http://developer.github.com/v3/repos/hooks/#create-a-hook
func (s *RepositoriesService) CreateHook(owner, repo string, hook *Hook) (*Hook, *Response, error)

// CreateKey adds a deploy key for a repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/keys/#create

// CreateKey 为某仓库添加部署密匙.
//
// GitHub API 文档: http://developer.github.com/v3/repos/keys/#create
func (s *RepositoriesService) CreateKey(owner string, repo string, key *Key) (*Key, *Response, error)

// CreateRelease adds a new release for a repository.
//
// GitHub API docs :
// http://developer.github.com/v3/repos/releases/#create-a-release

// CreateRelease 为某仓库添加正式版.
//
// GitHub API 文档 :
// http://developer.github.com/v3/repos/releases/#create-a-release
func (s *RepositoriesService) CreateRelease(owner, repo string, release *RepositoryRelease) (*RepositoryRelease, *Response, error)

// CreateStatus creates a new status for a repository at the specified reference.
// Ref can be a SHA, a branch name, or a tag name.
//
// GitHub API docs: http://developer.github.com/v3/repos/statuses/#create-a-status

// CreateStatus 以指定引用为某仓库新建状态.
// Ref 可以是 SHA, 分支名, 或标签名.
//
// GitHub API 文档: http://developer.github.com/v3/repos/statuses/#create-a-status
func (s *RepositoriesService) CreateStatus(owner, repo, ref string, status *RepoStatus) (*RepoStatus, *Response, error)

// Delete a repository.
//
// GitHub API docs: https://developer.github.com/v3/repos/#delete-a-repository

// Delete 删除一个仓库.
//
// GitHub API 文档: https://developer.github.com/v3/repos/#delete-a-repository
func (s *RepositoriesService) Delete(owner, repo string) (*Response, error)

// DeleteComment deletes a single comment from a repository.
//
// GitHub API docs:
// http://developer.github.com/v3/repos/comments/#delete-a-commit-comment

// DeleteComment 从仓库删除单个评论.
//
// GitHub API 文档:
// http://developer.github.com/v3/repos/comments/#delete-a-commit-comment
func (s *RepositoriesService) DeleteComment(owner, repo string, id int) (*Response, error)

// DeleteFile deletes a file from a repository and returns the commit. Requires the
// blob SHA of the file to be deleted.
//
// GitHub API docs: http://developer.github.com/v3/repos/contents/#delete-a-file

// DeleteFile 从仓库删除一个文件并返回该提交. 规定 SHA 相应的文件被删除.
//
// GitHub API 文档: http://developer.github.com/v3/repos/contents/#delete-a-file
func (s *RepositoriesService) DeleteFile(owner, repo, path string, opt *RepositoryContentFileOptions) (*RepositoryContentResponse, *Response, error)

// DeleteHook deletes a specified Hook.
//
// GitHub API docs: http://developer.github.com/v3/repos/hooks/#delete-a-hook

// DeleteHook 删除指定的 Hook.
//
// GitHub API 文档: http://developer.github.com/v3/repos/hooks/#delete-a-hook
func (s *RepositoriesService) DeleteHook(owner, repo string, id int) (*Response, error)

// DeleteKey deletes a deploy key.
//
// GitHub API docs: http://developer.github.com/v3/repos/keys/#delete

// DeleteKey 删除部署密匙.
//
// GitHub API 文档: http://developer.github.com/v3/repos/keys/#delete
func (s *RepositoriesService) DeleteKey(owner string, repo string, id int) (*Response, error)

// DeleteRelease delete a single release from a repository.
//
// GitHub API docs :
// http://developer.github.com/v3/repos/releases/#delete-a-release

// DeleteRelease 从仓库删除单个正式版.
//
// GitHub API 文档 :
// http://developer.github.com/v3/repos/releases/#delete-a-release
func (s *RepositoriesService) DeleteRelease(owner, repo string, id int) (*Response, error)

// DeleteReleaseAsset delete a single release asset from a repository.
//
// GitHub API docs :
// http://developer.github.com/v3/repos/releases/#delete-a-release-asset

// DeleteReleaseAsset 从仓库删除单个正式版资源.
//
// GitHub API 文档 :
// http://developer.github.com/v3/repos/releases/#delete-a-release-asset
func (s *RepositoriesService) DeleteReleaseAsset(owner, repo string, id int) (*Response, error)

// Edit updates a repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/#edit

// Edit 更新仓库.
//
// GitHub API 文档: http://developer.github.com/v3/repos/#edit
func (s *RepositoriesService) Edit(owner, repo string, repository *Repository) (*Repository, *Response, error)

// EditHook updates a specified Hook.
//
// GitHub API docs: http://developer.github.com/v3/repos/hooks/#edit-a-hook

// EditHook 更新指定的 Hook.
//
// GitHub API 文档: http://developer.github.com/v3/repos/hooks/#edit-a-hook
func (s *RepositoriesService) EditHook(owner, repo string, id int, hook *Hook) (*Hook, *Response, error)

// EditKey edits a deploy key.
//
// GitHub API docs: http://developer.github.com/v3/repos/keys/#edit

// EditKey 编辑部署密匙.
//
// GitHub API 文档: http://developer.github.com/v3/repos/keys/#edit
func (s *RepositoriesService) EditKey(owner string, repo string, id int, key *Key) (*Key, *Response, error)

// EditRelease edits a repository release.
//
// GitHub API docs : http://developer.github.com/v3/repos/releases/#edit-a-release

// EditRelease 编辑仓库正式版.
//
// GitHub API 文档 : http://developer.github.com/v3/repos/releases/#edit-a-release
func (s *RepositoriesService) EditRelease(owner, repo string, id int, release *RepositoryRelease) (*RepositoryRelease, *Response, error)

// EditReleaseAsset edits a repository release asset.
//
// GitHub API docs :
// http://developer.github.com/v3/repos/releases/#edit-a-release-asset

// EditReleaseAsset 编辑仓库正式版资源.
//
// GitHub API 文档 :
// http://developer.github.com/v3/repos/releases/#edit-a-release-asset
func (s *RepositoriesService) EditReleaseAsset(owner, repo string, id int, release *ReleaseAsset) (*ReleaseAsset, *Response, error)

// Get fetches a repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/#get

// Get 获取仓库.
//
// GitHub API 文档: http://developer.github.com/v3/repos/#get
func (s *RepositoriesService) Get(owner, repo string) (*Repository, *Response, error)

// GetArchiveLink returns an URL to download a tarball or zipball archive for a
// repository. The archiveFormat can be specified by either the github.Tarball or
// github.Zipball constant.
//
// GitHub API docs: http://developer.github.com/v3/repos/contents/#get-archive-link

// GetArchiveLink 返回某仓库 tarball 或 zipball 压缩文档下载 URL.
// archiveFormat 可指定为常量 github.Tarball 或 github.Zipball 之一.
//
// GitHub API 文档: http://developer.github.com/v3/repos/contents/#get-archive-link
func (s *RepositoriesService) GetArchiveLink(owner, repo string, archiveformat archiveFormat, opt *RepositoryContentGetOptions) (*url.URL, *Response, error)

// GetBranch gets the specified branch for a repository.
//
// GitHub API docs: https://developer.github.com/v3/repos/#get-branch

// GetBranch 获取仓库指定分支.
//
// GitHub API 文档: https://developer.github.com/v3/repos/#get-branch
func (s *RepositoriesService) GetBranch(owner, repo, branch string) (*Branch, *Response, error)

// GetCombinedStatus returns the combined status of a repository at the specified
// reference. ref can be a SHA, a branch name, or a tag name.
//
// GitHub API docs:
// https://developer.github.com/v3/repos/statuses/#get-the-combined-status-for-a-specific-ref

// GetCombinedStatus 返回指定仓库引用的组合状态.
// ref 可为 SHA, 分支名, 或标签名.
//
// GitHub API 文档:
// https://developer.github.com/v3/repos/statuses/#get-the-combined-status-for-a-specific-ref
func (s *RepositoriesService) GetCombinedStatus(owner, repo, ref string, opt *ListOptions) (*CombinedStatus, *Response, error)

// GetComment gets a single comment from a repository.
//
// GitHub API docs:
// http://developer.github.com/v3/repos/comments/#get-a-single-commit-comment

// GetComment 获取仓库单个评论.
//
// GitHub API 文档:
// http://developer.github.com/v3/repos/comments/#get-a-single-commit-comment
func (s *RepositoriesService) GetComment(owner, repo string, id int) (*RepositoryComment, *Response, error)

// GetCommit fetches the specified commit, including all details about it. todo:
// support media formats - https://github.com/google/go-github/issues/6
//
// GitHub API docs:
// http://developer.github.com/v3/repos/commits/#get-a-single-commit See also:
// http://developer.github.com//v3/git/commits/#get-a-single-commit provides the
// same functionality

// GetCommit 获取指定提交, 包括其所有细节. todo:
// 支持媒体格式 - https://github.com/google/go-github/issues/6
//
// GitHub API 文档:
// http://developer.github.com/v3/repos/commits/#get-a-single-commit 参见:
// http://developer.github.com//v3/git/commits/#get-a-single-commit 提供了相同功能
func (s *RepositoriesService) GetCommit(owner, repo, sha string) (*RepositoryCommit, *Response, error)

// GetContents can return either the metadata and content of a single file (when
// path references a file) or the metadata of all the files and/or subdirectories
// of a directory (when path references a directory). To make it easy to
// distinguish between both result types and to mimic the API as much as possible,
// both result types will be returned but only one will contain a value and the
// other will be nil.
//
// GitHub API docs: http://developer.github.com/v3/repos/contents/#get-contents

// GetContents 可返回单个文件的元数据或者内容 (当 path 引用一个文件) 或
// 某目录的所有文件和子目录元数据 (当 path 引用一个目录).
// 为方便区分类型和模仿 API 两者的结果, 两者的结果的都被返回,
// 但是只有一个有内容, 另外一个为 nil.
//
// GitHub API 文档: http://developer.github.com/v3/repos/contents/#get-contents
func (s *RepositoriesService) GetContents(owner, repo, path string, opt *RepositoryContentGetOptions) (fileContent *RepositoryContent,
	directoryContent []*RepositoryContent, resp *Response, err error)

// GetHook returns a single specified Hook.
//
// GitHub API docs: http://developer.github.com/v3/repos/hooks/#get-single-hook

// GetHook 返回单个指定的 Hook.
//
// GitHub API 文档: http://developer.github.com/v3/repos/hooks/#get-single-hook
func (s *RepositoriesService) GetHook(owner, repo string, id int) (*Hook, *Response, error)

// GetKey fetches a single deploy key.
//
// GitHub API docs: http://developer.github.com/v3/repos/keys/#get

// GetKey 获取单个部署密匙.
//
// GitHub API 文档: http://developer.github.com/v3/repos/keys/#get
func (s *RepositoriesService) GetKey(owner string, repo string, id int) (*Key, *Response, error)

// GetLatestPagesBuild fetches the latest build information for a GitHub pages
// site.
//
// GitHub API docs:
// https://developer.github.com/v3/repos/pages/#list-latest-pages-build

// GetLatestPagesBuild 获取 GitHub pages 站点最后构建信息.
//
// GitHub API 文档:
// https://developer.github.com/v3/repos/pages/#list-latest-pages-build
func (s *RepositoriesService) GetLatestPagesBuild(owner string, repo string) (*PagesBuild, *Response, error)

// GetPagesInfo fetches information about a GitHub Pages site.
//
// GitHub API docs:
// https://developer.github.com/v3/repos/pages/#get-information-about-a-pages-site

// GetPagesInfo 获取 GitHub Pages 站点信息.
//
// GitHub API 文档:
// https://developer.github.com/v3/repos/pages/#get-information-about-a-pages-site
func (s *RepositoriesService) GetPagesInfo(owner string, repo string) (*Pages, *Response, error)

// GetReadme gets the Readme file for the repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/contents/#get-the-readme

// GetReadme 获取某仓库的 Readme 文件.
//
// GitHub API 文档: http://developer.github.com/v3/repos/contents/#get-the-readme
func (s *RepositoriesService) GetReadme(owner, repo string, opt *RepositoryContentGetOptions) (*RepositoryContent, *Response, error)

// GetRelease fetches a single release.
//
// GitHub API docs:
// http://developer.github.com/v3/repos/releases/#get-a-single-release

// GetRelease 获取单个正式版.
//
// GitHub API 文档:
// http://developer.github.com/v3/repos/releases/#get-a-single-release
func (s *RepositoriesService) GetRelease(owner, repo string, id int) (*RepositoryRelease, *Response, error)

// GetReleaseAsset fetches a single release asset.
//
// GitHub API docs :
// http://developer.github.com/v3/repos/releases/#get-a-single-release-asset

// GetReleaseAsset 获取单个正式版资源.
//
// GitHub API 文档 :
// http://developer.github.com/v3/repos/releases/#get-a-single-release-asset
func (s *RepositoriesService) GetReleaseAsset(owner, repo string, id int) (*ReleaseAsset, *Response, error)

// IsCollaborator checks whether the specified Github user has collaborator access
// to the given repo. Note: This will return false if the user is not a
// collaborator OR the user is not a GitHub user.
//
// GitHub API docs: http://developer.github.com/v3/repos/collaborators/#get

// IsCollaborator 检查指定的 Github 用户是否有存取给定仓库的合作者.
// Note: 如果用户不是合作者或者不是 GitHub 用户将返回 false.
//
// GitHub API 文档: http://developer.github.com/v3/repos/collaborators/#get
func (s *RepositoriesService) IsCollaborator(owner, repo, user string) (bool, *Response, error)

// List the repositories for a user. Passing the empty string will list
// repositories for the authenticated user.
//
// GitHub API docs: http://developer.github.com/v3/repos/#list-user-repositories

// List 罗列某用户的仓库. 传递空字符串将罗列授权用户的仓库.
//
// GitHub API 文档: http://developer.github.com/v3/repos/#list-user-repositories
func (s *RepositoriesService) List(user string, opt *RepositoryListOptions) ([]Repository, *Response, error)

// ListAll lists all GitHub repositories in the order that they were created.
//
// GitHub API docs:
// http://developer.github.com/v3/repos/#list-all-public-repositories

// ListAll 罗列所有 GitHub 仓库, 按照他们创建的顺序.
//
// GitHub API 文档:
// http://developer.github.com/v3/repos/#list-all-public-repositories
func (s *RepositoriesService) ListAll(opt *RepositoryListAllOptions) ([]Repository, *Response, error)

// ListBranches lists branches for the specified repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/#list-branches

// ListBranches 罗列指定仓库的分支.
//
// GitHub API 文档: http://developer.github.com/v3/repos/#list-branches
func (s *RepositoriesService) ListBranches(owner string, repo string, opt *ListOptions) ([]Branch, *Response, error)

// ListByOrg lists the repositories for an organization.
//
// GitHub API docs:
// http://developer.github.com/v3/repos/#list-organization-repositories

// ListByOrg 罗列某组织的仓库.
//
// GitHub API 文档:
// http://developer.github.com/v3/repos/#list-organization-repositories
func (s *RepositoriesService) ListByOrg(org string, opt *RepositoryListByOrgOptions) ([]Repository, *Response, error)

// ListCodeFrequency returns a weekly aggregate of the number of additions and
// deletions pushed to a repository. Returned WeeklyStats will contain additiona
// and deletions, but not total commits.
//
// GitHub API Docs:
// https://developer.github.com/v3/repos/statistics/#code-frequency

// ListCodeFrequency 返回每周推送到仓库的增加和删除的合计数.
// 返回的 WeeklyStats 会包含增加的和删除的, 但没有提交总和.
//
// GitHub API Docs:
// https://developer.github.com/v3/repos/statistics/#code-frequency
func (s *RepositoriesService) ListCodeFrequency(owner, repo string) ([]WeeklyStats, *Response, error)

// ListCollaborators lists the Github users that have access to the repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/collaborators/#list

// ListCollaborators 罗列可存取某仓库的 Github 用户.
//
// GitHub API 文档: http://developer.github.com/v3/repos/collaborators/#list
func (s *RepositoriesService) ListCollaborators(owner, repo string, opt *ListOptions) ([]User, *Response, error)

// ListComments lists all the comments for the repository.
//
// GitHub API docs:
// http://developer.github.com/v3/repos/comments/#list-commit-comments-for-a-repository

// ListComments 罗列某仓库的所有评论.
//
// GitHub API 文档:
// http://developer.github.com/v3/repos/comments/#list-commit-comments-for-a-repository
func (s *RepositoriesService) ListComments(owner, repo string, opt *ListOptions) ([]RepositoryComment, *Response, error)

// ListCommitActivity returns the last year of commit activity grouped by week. The
// days array is a group of commits per day, starting on Sunday.
//
// If this is the first time these statistics are requested for the given
// repository, this method will return a non-nil error and a status code of 202.
// This is because this is the status that github returns to signify that it is now
// computing the requested statistics. A follow up request, after a delay of a
// second or so, should result in a successful request.
//
// GitHub API Docs:
// https://developer.github.com/v3/repos/statistics/#commit-activity

// ListCommitActivity 返回最后一年的提交活动, 按周进行分组.
// Days 数组是每日提交分组, 开始于星期天.
//
// 如果是给定仓库的首次统计请求, 该方法会返回状态码为 202 的 non-nil 错误.
// 因为这表示 github 现在要计算该请求的统计信息.
// 随后延迟一秒左右的请求将成功返回.
//
// GitHub API 文档:
// https://developer.github.com/v3/repos/statistics/#commit-activity
func (s *RepositoriesService) ListCommitActivity(owner, repo string) ([]WeeklyCommitActivity, *Response, error)

// ListCommitComments lists all the comments for a given commit SHA.
//
// GitHub API docs:
// http://developer.github.com/v3/repos/comments/#list-comments-for-a-single-commit

// ListCommitComments 罗列给定 SHA 提交的所有评论.
//
// GitHub API 文档:
// http://developer.github.com/v3/repos/comments/#list-comments-for-a-single-commit
func (s *RepositoriesService) ListCommitComments(owner, repo, sha string, opt *ListOptions) ([]RepositoryComment, *Response, error)

// ListCommits lists the commits of a repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/commits/#list

// ListCommits 罗列某仓库的提交.
//
// GitHub API 文档: http://developer.github.com/v3/repos/commits/#list
func (s *RepositoriesService) ListCommits(owner, repo string, opt *CommitsListOptions) ([]RepositoryCommit, *Response, error)

// ListContributors lists contributors for a repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/#list-contributors

// ListContributors 罗列某仓库的贡献者.
//
// GitHub API 文档: http://developer.github.com/v3/repos/#list-contributors
func (s *RepositoriesService) ListContributors(owner string, repository string, opt *ListContributorsOptions) ([]Contributor, *Response, error)

// ListContributorsStats gets a repo's contributor list with additions, deletions
// and commit counts.
//
// If this is the first time these statistics are requested for the given
// repository, this method will return a non-nil error and a status code of 202.
// This is because this is the status that github returns to signify that it is now
// computing the requested statistics. A follow up request, after a delay of a
// second or so, should result in a successful request.
//
// GitHub API Docs: https://developer.github.com/v3/repos/statistics/#contributors

// ListContributorsStats 获取某仓库贡献者列表, 包含增加, 删除和提交计数值.
//
// 如果是给定仓库的首次统计请求, 该方法会返回状态码为 202 的 non-nil 错误.
// 因为这表示 github 现在要计算该请求的统计信息.
// 随后延迟一秒左右的请求将成功返回.
//
// GitHub API 文档: https://developer.github.com/v3/repos/statistics/#contributors
func (s *RepositoriesService) ListContributorsStats(owner, repo string) ([]ContributorStats, *Response, error)

// ListDeploymentStatuses lists the statuses of a given deployment of a repository.
//
// GitHub API docs:
// https://developer.github.com/v3/repos/deployments/#list-deployment-statuses

// ListDeploymentStatuses 罗列给定仓库的部署状态.
//
// GitHub API 文档:
// https://developer.github.com/v3/repos/deployments/#list-deployment-statuses
func (s *RepositoriesService) ListDeploymentStatuses(owner, repo string, deployment int, opt *ListOptions) ([]DeploymentStatus, *Response, error)

// ListDeployments lists the deployments of a repository.
//
// GitHub API docs:
// https://developer.github.com/v3/repos/deployments/#list-deployments

// ListDeployments 罗列给定仓库的部署.
//
// GitHub API 文档:
// https://developer.github.com/v3/repos/deployments/#list-deployments
func (s *RepositoriesService) ListDeployments(owner, repo string, opt *DeploymentsListOptions) ([]Deployment, *Response, error)

// ListForks lists the forks of the specified repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/forks/#list-forks

// ListForks 罗列指定仓库的 forks.
//
// GitHub API 文档: http://developer.github.com/v3/repos/forks/#list-forks
func (s *RepositoriesService) ListForks(owner, repo string, opt *RepositoryListForksOptions) ([]Repository, *Response, error)

// ListHooks lists all Hooks for the specified repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/hooks/#list

// ListHooks 罗列指定仓库所有的 Hooks.
//
// GitHub API 文档: http://developer.github.com/v3/repos/hooks/#list
func (s *RepositoriesService) ListHooks(owner, repo string, opt *ListOptions) ([]Hook, *Response, error)

// ListKeys lists the deploy keys for a repository.
//
// GitHub API docs: http://developer.github.com/v3/repos/keys/#list

// ListKeys 罗列某仓库的部署密匙.
//
// GitHub API 文档: http://developer.github.com/v3/repos/keys/#list
func (s *RepositoriesService) ListKeys(owner string, repo string, opt *ListOptions) ([]Key, *Response, error)

// ListLanguages lists languages for the specified repository. The returned map
// specifies the languages and the number of bytes of code written in that
// language. For example:
//
//	{
//	  "C": 78769,
//	  "Python": 7769
//	}
//
// GitHub API Docs: http://developer.github.com/v3/repos/#list-languages

// ListLanguages 罗列指定仓库的语言. 返回的 map 指定了语言和该语言书写的代码字节数.
// 例子:
//
//	{
//	  "C": 78769,
//	  "Python": 7769
//	}
//
// GitHub API 文档: http://developer.github.com/v3/repos/#list-languages
func (s *RepositoriesService) ListLanguages(owner string, repo string) (map[string]int, *Response, error)

// ListPagesBuilds lists the builds for a GitHub Pages site.
//
// GitHub API docs: https://developer.github.com/v3/repos/pages/#list-pages-builds

// ListPagesBuilds 罗列 GitHub Pages 站点的构建信息.
//
// GitHub API 文档: https://developer.github.com/v3/repos/pages/#list-pages-builds
func (s *RepositoriesService) ListPagesBuilds(owner string, repo string) ([]PagesBuild, *Response, error)

// ListParticipation returns the total commit counts for the 'owner' and total
// commit counts in 'all'. 'all' is everyone combined, including the 'owner' in the
// last 52 weeks. If you’d like to get the commit counts for non-owners, you can
// subtract 'all' from 'owner'.
//
// The array order is oldest week (index 0) to most recent week.
//
// If this is the first time these statistics are requested for the given
// repository, this method will return a non-nil error and a status code of 202.
// This is because this is the status that github returns to signify that it is now
// computing the requested statistics. A follow up request, after a delay of a
// second or so, should result in a successful request.
//
// GitHub API Docs: https://developer.github.com/v3/repos/statistics/#participation

// ListParticipation 返回 'owner' 提交计数总和以及 'all' 提交计数总和.
// 'all' 为所有人在最后 52 周内的, 包括 'owner'.
// 如果你想获得非拥有者的提交计数, 你可以从 'all' 中减去 'owner'.
//
// The array order is oldest week (index 0) to most recent week.
//
// 如果是给定仓库的首次统计请求, 该方法会返回状态码为 202 的 non-nil 错误.
// 因为这表示 github 现在要计算该请求的统计信息.
// 随后延迟一秒左右的请求将成功返回.
//
// GitHub API 文档: https://developer.github.com/v3/repos/statistics/#participation
func (s *RepositoriesService) ListParticipation(owner, repo string) (*RepositoryParticipation, *Response, error)

// ListPunchCard returns the number of commits per hour in each day.
//
// GitHub API Docs: https://developer.github.com/v3/repos/statistics/#punch-card

// ListPunchCard 返回每天每个小时的提交数量.
//
// GitHub API 文档: https://developer.github.com/v3/repos/statistics/#punch-card
func (s *RepositoriesService) ListPunchCard(owner, repo string) ([]PunchCard, *Response, error)

// ListReleaseAssets lists the release's assets.
//
// GitHub API docs :
// http://developer.github.com/v3/repos/releases/#list-assets-for-a-release

// ListReleaseAssets 罗列正式版资源.
//
// GitHub API 文档 :
// http://developer.github.com/v3/repos/releases/#list-assets-for-a-release
func (s *RepositoriesService) ListReleaseAssets(owner, repo string, id int, opt *ListOptions) ([]ReleaseAsset, *Response, error)

// ListReleases lists the releases for a repository.
//
// GitHub API docs:
// http://developer.github.com/v3/repos/releases/#list-releases-for-a-repository

// ListReleases 罗列某仓库的正式版.
//
// GitHub API 文档:
// http://developer.github.com/v3/repos/releases/#list-releases-for-a-repository
func (s *RepositoriesService) ListReleases(owner, repo string, opt *ListOptions) ([]RepositoryRelease, *Response, error)

// ListServiceHooks lists all of the available service hooks.
//
// GitHub API docs: https://developer.github.com/webhooks/#services

// ListServiceHooks 罗列所有可用的服务钩子.
//
// GitHub API 文档: https://developer.github.com/webhooks/#services
func (s *RepositoriesService) ListServiceHooks() ([]ServiceHook, *Response, error)

// ListStatuses lists the statuses of a repository at the specified reference. ref
// can be a SHA, a branch name, or a tag name.
//
// GitHub API docs:
// http://developer.github.com/v3/repos/statuses/#list-statuses-for-a-specific-ref

// ListStatuses 罗列某仓库指定引用的状态. ref 可为 SHA, 分支名, 或标签名.
//
// GitHub API 文档:
// http://developer.github.com/v3/repos/statuses/#list-statuses-for-a-specific-ref
func (s *RepositoriesService) ListStatuses(owner, repo, ref string, opt *ListOptions) ([]RepoStatus, *Response, error)

// ListTags lists tags for the specified repository.
//
// GitHub API docs: https://developer.github.com/v3/repos/#list-tags

// ListTags 罗列指定仓库的标签.
//
// GitHub API 文档: https://developer.github.com/v3/repos/#list-tags
func (s *RepositoriesService) ListTags(owner string, repo string, opt *ListOptions) ([]RepositoryTag, *Response, error)

// ListTeams lists the teams for the specified repository.
//
// GitHub API docs: https://developer.github.com/v3/repos/#list-teams

// ListTeams 罗列指定仓库的团队.
//
// GitHub API 文档: https://developer.github.com/v3/repos/#list-teams
func (s *RepositoriesService) ListTeams(owner string, repo string, opt *ListOptions) ([]Team, *Response, error)

// Merge a branch in the specified repository.
//
// GitHub API docs: https://developer.github.com/v3/repos/merging/#perform-a-merge

// Merge 合并指定仓库的一个分支.
//
// GitHub API 文档: https://developer.github.com/v3/repos/merging/#perform-a-merge
func (s *RepositoriesService) Merge(owner, repo string, request *RepositoryMergeRequest) (*RepositoryCommit, *Response, error)

// RemoveCollaborator removes the specified Github user as collaborator from the
// given repo. Note: Does not return error if a valid user that is not a
// collaborator is removed.
//
// GitHub API docs:
// http://developer.github.com/v3/repos/collaborators/#remove-collaborator

// RemoveCollaborator 从给定仓库删除 Github 合作用户.
// Note: 如果有效用户不是合作者不会返回错误.
//
// GitHub API 文档:
// http://developer.github.com/v3/repos/collaborators/#remove-collaborator
func (s *RepositoriesService) RemoveCollaborator(owner, repo, user string) (*Response, error)

// TestHook triggers a test Hook by github.
//
// GitHub API docs: http://developer.github.com/v3/repos/hooks/#test-a-push-hook

// TestHook 由 github 触发测试 Hook.
//
// GitHub API 文档: http://developer.github.com/v3/repos/hooks/#test-a-push-hook
func (s *RepositoriesService) TestHook(owner, repo string, id int) (*Response, error)

// UpdateComment updates the body of a single comment.
//
// GitHub API docs:
// http://developer.github.com/v3/repos/comments/#update-a-commit-comment

// UpdateComment 更新单个注释主体.
//
// GitHub API 文档:
// http://developer.github.com/v3/repos/comments/#update-a-commit-comment
func (s *RepositoriesService) UpdateComment(owner, repo string, id int, comment *RepositoryComment) (*RepositoryComment, *Response, error)

// UpdateFile updates a file in a repository at the given path and returns the
// commit and file metadata. Requires the blob SHA of the file being updated.
//
// GitHub API docs: http://developer.github.com/v3/repos/contents/#update-a-file

// UpdateFile 更新某仓库给定路径的文件, 返回该提交和文件元数据. 更新必填文件 SHA.
//
// GitHub API 文档: http://developer.github.com/v3/repos/contents/#update-a-file
func (s *RepositoriesService) UpdateFile(owner, repo, path string, opt *RepositoryContentFileOptions) (*RepositoryContentResponse, *Response, error)

// UploadReleaseAsset creates an asset by uploading a file into a release
// repository. To upload assets that cannot be represented by an os.File, call
// NewUploadRequest directly.
//
// GitHub API docs :
// http://developer.github.com/v3/repos/releases/#upload-a-release-asset

// UploadReleaseAsset 以一个上传的文件创建仓库正式版资源.
// 不能用 os.File 表示上传的资源, 直接调用 NewUploadRequest.
//
// GitHub API 文档 :
// http://developer.github.com/v3/repos/releases/#upload-a-release-asset
func (s *RepositoriesService) UploadReleaseAsset(owner, repo string, id int, opt *UploadOptions, file *os.File) (*ReleaseAsset, *Response, error)

// Repository represents a GitHub repository.

// Repository 表示一个 GitHub 仓库.
type Repository struct {
	ID               *int             `json:"id,omitempty"`
	Owner            *User            `json:"owner,omitempty"`
	Name             *string          `json:"name,omitempty"`
	FullName         *string          `json:"full_name,omitempty"`
	Description      *string          `json:"description,omitempty"`
	Homepage         *string          `json:"homepage,omitempty"`
	DefaultBranch    *string          `json:"default_branch,omitempty"`
	MasterBranch     *string          `json:"master_branch,omitempty"`
	CreatedAt        *Timestamp       `json:"created_at,omitempty"`
	PushedAt         *Timestamp       `json:"pushed_at,omitempty"`
	UpdatedAt        *Timestamp       `json:"updated_at,omitempty"`
	HTMLURL          *string          `json:"html_url,omitempty"`
	CloneURL         *string          `json:"clone_url,omitempty"`
	GitURL           *string          `json:"git_url,omitempty"`
	MirrorURL        *string          `json:"mirror_url,omitempty"`
	SSHURL           *string          `json:"ssh_url,omitempty"`
	SVNURL           *string          `json:"svn_url,omitempty"`
	Language         *string          `json:"language,omitempty"`
	Fork             *bool            `json:"fork"`
	ForksCount       *int             `json:"forks_count,omitempty"`
	NetworkCount     *int             `json:"network_count,omitempty"`
	OpenIssuesCount  *int             `json:"open_issues_count,omitempty"`
	StargazersCount  *int             `json:"stargazers_count,omitempty"`
	SubscribersCount *int             `json:"subscribers_count,omitempty"`
	WatchersCount    *int             `json:"watchers_count,omitempty"`
	Size             *int             `json:"size,omitempty"`
	AutoInit         *bool            `json:"auto_init,omitempty"`
	Parent           *Repository      `json:"parent,omitempty"`
	Source           *Repository      `json:"source,omitempty"`
	Organization     *Organization    `json:"organization,omitempty"`
	Permissions      *map[string]bool `json:"permissions,omitempty"`

	// Additional mutable fields when creating and editing a repository

	// 当创建和编辑仓库是附加的可变字段
	Private      *bool `json:"private"`
	HasIssues    *bool `json:"has_issues"`
	HasWiki      *bool `json:"has_wiki"`
	HasDownloads *bool `json:"has_downloads"`
	// Creating an organization repository. Required for non-owners.

	// 创建一个组织仓库时. 要求非所有者.
	TeamID *int `json:"team_id"`

	// API URLs
	URL              *string `json:"url,omitempty"`
	ArchiveURL       *string `json:"archive_url,omitempty"`
	AssigneesURL     *string `json:"assignees_url,omitempty"`
	BlobsURL         *string `json:"blobs_url,omitempty"`
	BranchesURL      *string `json:"branches_url,omitempty"`
	CollaboratorsURL *string `json:"collaborators_url,omitempty"`
	CommentsURL      *string `json:"comments_url,omitempty"`
	CommitsURL       *string `json:"commits_url,omitempty"`
	CompareURL       *string `json:"compare_url,omitempty"`
	ContentsURL      *string `json:"contents_url,omitempty"`
	ContributorsURL  *string `json:"contributors_url,omitempty"`
	DownloadsURL     *string `json:"downloads_url,omitempty"`
	EventsURL        *string `json:"events_url,omitempty"`
	ForksURL         *string `json:"forks_url,omitempty"`
	GitCommitsURL    *string `json:"git_commits_url,omitempty"`
	GitRefsURL       *string `json:"git_refs_url,omitempty"`
	GitTagsURL       *string `json:"git_tags_url,omitempty"`
	HooksURL         *string `json:"hooks_url,omitempty"`
	IssueCommentURL  *string `json:"issue_comment_url,omitempty"`
	IssueEventsURL   *string `json:"issue_events_url,omitempty"`
	IssuesURL        *string `json:"issues_url,omitempty"`
	KeysURL          *string `json:"keys_url,omitempty"`
	LabelsURL        *string `json:"labels_url,omitempty"`
	LanguagesURL     *string `json:"languages_url,omitempty"`
	MergesURL        *string `json:"merges_url,omitempty"`
	MilestonesURL    *string `json:"milestones_url,omitempty"`
	NotificationsURL *string `json:"notifications_url,omitempty"`
	PullsURL         *string `json:"pulls_url,omitempty"`
	ReleasesURL      *string `json:"releases_url,omitempty"`
	StargazersURL    *string `json:"stargazers_url,omitempty"`
	StatusesURL      *string `json:"statuses_url,omitempty"`
	SubscribersURL   *string `json:"subscribers_url,omitempty"`
	SubscriptionURL  *string `json:"subscription_url,omitempty"`
	TagsURL          *string `json:"tags_url,omitempty"`
	TreesURL         *string `json:"trees_url,omitempty"`
	TeamsURL         *string `json:"teams_url,omitempty"`

	// TextMatches is only populated from search results that request text matches
	// See: search.go and https://developer.github.com/v3/search/#text-match-metadata

	// TextMatches 只是填入了搜索文本匹配请求的结果.
	// See: search.go and https://developer.github.com/v3/search/#text-match-metadata
	TextMatches []TextMatch `json:"text_matches,omitempty"`
}

func (r Repository) String() string

// RepositoryComment represents a comment for a commit, file, or line in a
// repository.

// RepositoryComment 表示仓库的提交注释, 文件的, 或行内的.
type RepositoryComment struct {
	HTMLURL   *string    `json:"html_url,omitempty"`
	URL       *string    `json:"url,omitempty"`
	ID        *int       `json:"id,omitempty"`
	CommitID  *string    `json:"commit_id,omitempty"`
	User      *User      `json:"user,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	// User-mutable fields
	Body *string `json:"body"`
	// User-initialized fields
	Path     *string `json:"path,omitempty"`
	Position *int    `json:"position,omitempty"`
}

func (r RepositoryComment) String() string

// RepositoryCommit represents a commit in a repo. Note that it's wrapping a
// Commit, so author/committer information is in two places, but contain different
// details about them: in RepositoryCommit "github details", in Commit - "git
// details".

// RepositoryCommit 表示一个仓库提交. 注意这是对 Commit 的包装,
// 所以 作者/提交者 分在两个地方, 且他们含有不同的细节:
// RepositoryCommit 包含 "github details", Commit 包含 "git details".
type RepositoryCommit struct {
	SHA       *string  `json:"sha,omitempty"`
	Commit    *Commit  `json:"commit,omitempty"`
	Author    *User    `json:"author,omitempty"`
	Committer *User    `json:"committer,omitempty"`
	Parents   []Commit `json:"parents,omitempty"`
	Message   *string  `json:"message,omitempty"`
	HTMLURL   *string  `json:"html_url,omitempty"`

	// Details about how many changes were made in this commit. Only filled in during GetCommit!
	Stats *CommitStats `json:"stats,omitempty"`
	// Details about which files, and how this commit touched. Only filled in during GetCommit!
	Files []CommitFile `json:"files,omitempty"`
}

func (r RepositoryCommit) String() string

// RepositoryContent represents a file or directory in a github repository.

// RepositoryContent 表示仓库中的一个文件或目录.
type RepositoryContent struct {
	Type     *string `json:"type,omitempty"`
	Encoding *string `json:"encoding,omitempty"`
	Size     *int    `json:"size,omitempty"`
	Name     *string `json:"name,omitempty"`
	Path     *string `json:"path,omitempty"`
	Content  *string `json:"content,omitempty"`
	SHA      *string `json:"sha,omitempty"`
	URL      *string `json:"url,omitempty"`
	GitURL   *string `json:"giturl,omitempty"`
	HTMLURL  *string `json:"htmlurl,omitempty"`
}

// Decode decodes the file content if it is base64 encoded.

// Decode 解码文件内容, 如果是以 base64 编码的话.
func (r *RepositoryContent) Decode() ([]byte, error)

func (r RepositoryContent) String() string

// RepositoryContentFileOptions specifies optional parameters for CreateFile,
// UpdateFile, and DeleteFile.

// RepositoryContentFileOptions 指定 CreateFile, UpdateFile, 和 DeleteFile 的可选参数.
type RepositoryContentFileOptions struct {
	Message   *string       `json:"message,omitempty"`
	Content   []byte        `json:"content,omitempty"`
	SHA       *string       `json:"sha,omitempty"`
	Branch    *string       `json:"branch,omitempty"`
	Author    *CommitAuthor `json:"author,omitempty"`
	Committer *CommitAuthor `json:"committer,omitempty"`
}

// RepositoryContentGetOptions represents an optional ref parameter, which can be a
// SHA, branch, or tag

// RepositoryContentGetOptions 表示一个可选引用参数, 可以是 SHA, 分支, 或标签之一.
type RepositoryContentGetOptions struct {
	Ref string `url:"ref,omitempty"`
}

// RepositoryContentResponse holds the parsed response from CreateFile, UpdateFile,
// and DeleteFile.

// RepositoryContentResponse 持有 CreateFile, UpdateFile 和 DeleteFile 响应解析.
type RepositoryContentResponse struct {
	Content *RepositoryContent `json:"content,omitempty"`
	Commit  `json:"commit,omitempty"`
}

// RepositoryCreateForkOptions specifies the optional parameters to the
// RepositoriesService.CreateFork method.

// RepositoryCreateForkOptions 指定 RepositoriesService.CreateFork 方法的可选参数.
type RepositoryCreateForkOptions struct {
	// The organization to fork the repository into.

	// 仓库 fork 到该组织.
	Organization string `url:"organization,omitempty"`
}

// RepositoryListAllOptions specifies the optional parameters to the
// RepositoriesService.ListAll method.

// RepositoryListAllOptions 指定 RepositoriesService.ListAll 方法的可选参数.
type RepositoryListAllOptions struct {
	// ID of the last repository seen

	// 仓库的最后一个可见 ID.
	Since int `url:"since,omitempty"`

	ListOptions
}

// RepositoryListByOrgOptions specifies the optional parameters to the
// RepositoriesService.ListByOrg method.

// RepositoryListByOrgOptions 指定 RepositoriesService.ListByOrg 方法的可选参数.
type RepositoryListByOrgOptions struct {
	// Type of repositories to list.  Possible values are: all, public, private,
	// forks, sources, member.  Default is "all".

	// 列表仓库的类型. 可能的值有: all, public, private, forks, sources, member.
	// 缺省为 "all".
	Type string `url:"type,omitempty"`

	ListOptions
}

// RepositoryListForksOptions specifies the optional parameters to the
// RepositoriesService.ListForks method.

// RepositoryListForksOptions 指定 RepositoriesService.ListForks 方法的可选参数.
type RepositoryListForksOptions struct {
	// How to sort the forks list.  Possible values are: newest, oldest,
	// watchers.  Default is "newest".

	// 如何排序 forks 列表. 可能的值有: newest, oldest, watchers.
	// 缺省为 "newest".
	Sort string `url:"sort,omitempty"`

	ListOptions
}

// RepositoryListOptions specifies the optional parameters to the
// RepositoriesService.List method.

// RepositoryListOptions 指定  RepositoriesService.List 方法的可选参数.
type RepositoryListOptions struct {
	// Type of repositories to list.  Possible values are: all, owner, public,
	// private, member.  Default is "all".

	// 列表仓库的类型. 可能的值有: all, public, private, forks, sources, member.
	// 缺省为 "all".
	Type string `url:"type,omitempty"`

	// How to sort the repository list.  Possible values are: created, updated,
	// pushed, full_name.  Default is "full_name".

	// 如何排序仓库列表. 可能的值有: created, updated, pushed, full_name.
	// 缺省为 "full_name".
	Sort string `url:"sort,omitempty"`

	// Direction in which to sort repositories.  Possible values are: asc, desc.
	// Default is "asc" when sort is "full_name", otherwise default is "desc".
	Direction string `url:"direction,omitempty"`

	ListOptions
}

// RepositoryMergeRequest represents a request to merge a branch in a repository.

// RepositoryMergeRequest 表示合并仓库分支的请求.
type RepositoryMergeRequest struct {
	Base          *string `json:"base,omitempty"`
	Head          *string `json:"head,omitempty"`
	CommitMessage *string `json:"commit_message,omitempty"`
}

// RepositoryParticipation is the number of commits by everyone who has contributed
// to the repository (including the owner) as well as the number of commits by the
// owner themself.

// RepositoryParticipation 是所有贡献者提交到仓库 (包括自有仓库) 的数量,
// 以及所有者自己提交的.
type RepositoryParticipation struct {
	All   []int `json:"all,omitempty"`
	Owner []int `json:"owner,omitempty"`
}

func (r RepositoryParticipation) String() string

// RepositoryRelease represents a GitHub release in a repository.

// RepositoryRelease 表示仓库的 GitHub 正式版.
type RepositoryRelease struct {
	ID              *int           `json:"id,omitempty"`
	TagName         *string        `json:"tag_name,omitempty"`
	TargetCommitish *string        `json:"target_commitish,omitempty"`
	Name            *string        `json:"name,omitempty"`
	Body            *string        `json:"body,omitempty"`
	Draft           *bool          `json:"draft,omitempty"`
	Prerelease      *bool          `json:"prerelease,omitempty"`
	CreatedAt       *Timestamp     `json:"created_at,omitempty"`
	PublishedAt     *Timestamp     `json:"published_at,omitempty"`
	URL             *string        `json:"url,omitempty"`
	HTMLURL         *string        `json:"html_url,omitempty"`
	AssetsURL       *string        `json:"assets_url,omitempty"`
	Assets          []ReleaseAsset `json:"assets,omitempty"`
	UploadURL       *string        `json:"upload_url,omitempty"`
	ZipballURL      *string        `json:"zipball_url,omitempty"`
	TarballURL      *string        `json:"tarball_url,omitempty"`
}

func (r RepositoryRelease) String() string

// RepositoryTag represents a repository tag.

// RepositoryTag 表示一个仓库标签.
type RepositoryTag struct {
	Name       *string `json:"name,omitempty"`
	Commit     *Commit `json:"commit,omitempty"`
	ZipballURL *string `json:"zipball_url,omitempty"`
	TarballURL *string `json:"tarball_url,omitempty"`
}

// Response is a GitHub API response. This wraps the standard http.Response
// returned from GitHub and provides convenient access to things like pagination
// links.

// Response 是 GitHub API 响应. 它包装自 GitHub 返回的标准 http.Response,
// 并提供分页链接之类的便捷访问.
type Response struct {
	*http.Response

	NextPage  int
	PrevPage  int
	FirstPage int
	LastPage  int

	Rate
}

// SearchOptions specifies optional parameters to the SearchService methods.

// SearchOptions 指定 SearchService 方法的可选参数.
type SearchOptions struct {
	// How to sort the search results.  Possible values are:
	//   - for repositories: stars, fork, updated
	//   - for code: indexed
	//   - for issues: comments, created, updated
	//   - for users: followers, repositories, joined
	//
	// Default is to sort by best match.

	// 如何排序搜索结果. 可能的值有:
	//   - 用于仓库: stars, fork, updated
	//   - 用于代码: indexed
	//   - 用于问题: comments, created, updated
	//   - 用于用户: followers, repositories, joined
	//
	// 缺省以最佳匹配排序.
	Sort string `url:"sort,omitempty"`

	// Sort order if sort parameter is provided. Possible values are: asc,
	// desc. Default is desc.

	// 排序顺序, 如果提供了 Sort 参数. 可能的值有: asc, desc.
	// 缺省为 desc.
	Order string `url:"order,omitempty"`

	// Whether to retrieve text match metadata with a query

	// 文本匹配是否检索查询元数据.
	TextMatch bool `url:"-"`

	ListOptions
}

// SearchService provides access to the search related functions in the GitHub API.
//
// GitHub API docs: http://developer.github.com/v3/search/

// SearchService 提供访问 GitHub API 中搜索相关的功能.
//
// GitHub API 文档: http://developer.github.com/v3/search/
type SearchService struct {
	// contains filtered or unexported fields
}

// Code searches code via various criteria.
//
// GitHub API docs: http://developer.github.com/v3/search/#search-code

// Code 搜索各种代码.
//
// GitHub API 文档: http://developer.github.com/v3/search/#search-code
func (s *SearchService) Code(query string, opt *SearchOptions) (*CodeSearchResult, *Response, error)

// Issues searches issues via various criteria.
//
// GitHub API docs: http://developer.github.com/v3/search/#search-issues

// Issues 搜索各种问题代码.
//
// GitHub API 文档: http://developer.github.com/v3/search/#search-issues
func (s *SearchService) Issues(query string, opt *SearchOptions) (*IssuesSearchResult, *Response, error)

// Repositories searches repositories via various criteria.
//
// GitHub API docs: http://developer.github.com/v3/search/#search-repositories

// Repositories 搜索各种问题仓库.
//
// GitHub API 文档: http://developer.github.com/v3/search/#search-repositories
func (s *SearchService) Repositories(query string, opt *SearchOptions) (*RepositoriesSearchResult, *Response, error)

// Users searches users via various criteria.
//
// GitHub API docs: http://developer.github.com/v3/search/#search-users

// Users 搜索各种问题用户.
//
// GitHub API 文档: http://developer.github.com/v3/search/#search-users
func (s *SearchService) Users(query string, opt *SearchOptions) (*UsersSearchResult, *Response, error)

// ServiceHook represents a hook that has configuration settings, a list of
// available events, and default events.

// ServiceHook 表示钩子的配置, 可用事件和缺省事件列表.
type ServiceHook struct {
	Name            *string    `json:"name,omitempty"`
	Events          []string   `json:"events,omitempty"`
	SupportedEvents []string   `json:"supported_events,omitempty"`
	Schema          [][]string `json:"schema,omitempty"`
}

func (s *ServiceHook) String() string

// Subscription identifies a repository or thread subscription.

// Subscription 标识仓库订阅或订阅线程.
type Subscription struct {
	Subscribed *bool      `json:"subscribed,omitempty"`
	Ignored    *bool      `json:"ignored,omitempty"`
	Reason     *string    `json:"reason,omitempty"`
	CreatedAt  *Timestamp `json:"created_at,omitempty"`
	URL        *string    `json:"url,omitempty"`

	// only populated for repository subscriptions

	// 只为仓库订阅
	RepositoryURL *string `json:"repository_url,omitempty"`

	// only populated for thread subscriptions

	// 只为订阅线程
	ThreadURL *string `json:"thread_url,omitempty"`
}

// Tag represents a tag object.

// Tag 表示标签对象.
type Tag struct {
	Tag     *string       `json:"tag,omitempty"`
	SHA     *string       `json:"sha,omitempty"`
	URL     *string       `json:"url,omitempty"`
	Message *string       `json:"message,omitempty"`
	Tagger  *CommitAuthor `json:"tagger,omitempty"`
	Object  *GitObject    `json:"object,omitempty"`
}

// Team represents a team within a GitHub organization. Teams are used to manage
// access to an organization's repositories.

// Team 表示 GitHub 组织中的团队. 团队用于管理存取组织仓库.
type Team struct {
	ID           *int          `json:"id,omitempty"`
	Name         *string       `json:"name,omitempty"`
	URL          *string       `json:"url,omitempty"`
	Slug         *string       `json:"slug,omitempty"`
	Permission   *string       `json:"permission,omitempty"`
	MembersCount *int          `json:"members_count,omitempty"`
	ReposCount   *int          `json:"repos_count,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
}

func (t Team) String() string

// TextMatch represents a text match for a SearchResult

// TextMatch 表示 SearchResult 的文本匹配.
type TextMatch struct {
	ObjectURL  *string `json:"object_url,omitempty"`
	ObjectType *string `json:"object_type,omitempty"`
	Property   *string `json:"property,omitempty"`
	Fragment   *string `json:"fragment,omitempty"`
	Matches    []Match `json:"matches,omitempty"`
}

func (tm TextMatch) String() string

// Timestamp represents a time that can be unmarshalled from a JSON string
// formatted as either an RFC3339 or Unix timestamp. This is necessary for some
// fields since the GitHub API is inconsistent in how it represents times. All
// exported methods of time.Time can be called on Timestamp.

// Timestamp 表示一个时间, 它可 unmarshalled 自以 RFC3339 或 Unix 时间戳格式化
// 的 JSON 字符串. 这对于 GitHub API 表示时间不一致的某些字段是非常必要的.
// 所有导出 time.Time 的方法可使用 Timestamp.
type Timestamp struct {
	time.Time
}

// Equal reports whether t and u are equal based on time.Equal

// Equal 报告 t 和 u 是否相等, 基于 time.Equal.
func (t Timestamp) Equal(u Timestamp) bool

func (t Timestamp) String() string

// UnmarshalJSON implements the json.Unmarshaler interface. Time is expected in
// RFC3339 or Unix format.

// UnmarshalJSON 实现了 json.Unmarshaler 接口. 期望 RFC3339 或 Unix 格式的时间.
func (t *Timestamp) UnmarshalJSON(data []byte) (err error)

// Tree represents a GitHub tree.

// Tree 表示 GitHub 树.
type Tree struct {
	SHA     *string     `json:"sha,omitempty"`
	Entries []TreeEntry `json:"tree,omitempty"`
}

func (t Tree) String() string

// TreeEntry represents the contents of a tree structure. TreeEntry can represent
// either a blob, a commit (in the case of a submodule), or another tree.

// TreeEntry 表示树结构的内容. TreeEntry 可表示 blob, 提交(子模块情况下), 或其它树.
type TreeEntry struct {
	SHA     *string `json:"sha,omitempty"`
	Path    *string `json:"path,omitempty"`
	Mode    *string `json:"mode,omitempty"`
	Type    *string `json:"type,omitempty"`
	Size    *int    `json:"size,omitempty"`
	Content *string `json:"content,omitempty"`
}

func (t TreeEntry) String() string

// UnauthenticatedRateLimitedTransport allows you to make unauthenticated calls
// that need to use a higher rate limit associated with your OAuth application.
//
//	t := &github.UnauthenticatedRateLimitedTransport{
//		ClientID:     "your app's client ID",
//		ClientSecret: "your app's client secret",
//	}
//	client := github.NewClient(t.Client())
//
// This will append the querystring params client_id=xxx&client_secret=yyy to all
// requests.
//
// See http://developer.github.com/v3/#unauthenticated-rate-limited-requests for
// more information.

// UnauthenticatedRateLimitedTransport 使你的授权应用以未授权身份访问,
// 以便更高效的使用频次限制.
//
// 	t := &github.UnauthenticatedRateLimitedTransport{
// 		ClientID:     "your app's client ID",
// 		ClientSecret: "your app's client secret",
// 	}
// 	client := github.NewClient(t.Client())
//
// 这回添加查询参数 client_id=xxx&client_secret=yyy 到所有请求.
//
// 更多信息参见 http://developer.github.com/v3/#unauthenticated-rate-limited-requests
type UnauthenticatedRateLimitedTransport struct {
	// ClientID is the GitHub OAuth client ID of the current application, which
	// can be found by selecting its entry in the list at
	// https://github.com/settings/applications.

	// ClientID 为当前 GitHub OAuth 应用客户端 ID, 可在
	// https://github.com/settings/applications 条目列表中找到.
	ClientID string

	// ClientSecret is the GitHub OAuth client secret of the current
	// application.

	// ClientSecret 为当前 GitHub OAuth 应用客户端密匙.
	ClientSecret string

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.

	// Transport 用来建立请求的 HTTP 底层传输.
	// 如果为 nil 使用用缺省值 http.DefaultTransport.
	Transport http.RoundTripper
}

// Client returns an *http.Client that makes requests which are subject to the rate
// limit of your OAuth application.

// Client 返回受你的授权应用频次限制的 *http.Client 请求.
func (t *UnauthenticatedRateLimitedTransport) Client() *http.Client

// RoundTrip implements the RoundTripper interface.

// RoundTrip 实现了 RoundTripper 接口.
func (t *UnauthenticatedRateLimitedTransport) RoundTrip(req *http.Request) (*http.Response, error)

// UploadOptions specifies the parameters to methods that support uploads.

// UploadOptions 指定支持上传方法的参数.
type UploadOptions struct {
	Name string `url:"name,omitempty"`
}

// User represents a GitHub user.

// User 表示一个 GitHub 用户.
type User struct {
	Login             *string    `json:"login,omitempty"`
	ID                *int       `json:"id,omitempty"`
	AvatarURL         *string    `json:"avatar_url,omitempty"`
	HTMLURL           *string    `json:"html_url,omitempty"`
	GravatarID        *string    `json:"gravatar_id,omitempty"`
	Name              *string    `json:"name,omitempty"`
	Company           *string    `json:"company,omitempty"`
	Blog              *string    `json:"blog,omitempty"`
	Location          *string    `json:"location,omitempty"`
	Email             *string    `json:"email,omitempty"`
	Hireable          *bool      `json:"hireable,omitempty"`
	Bio               *string    `json:"bio,omitempty"`
	PublicRepos       *int       `json:"public_repos,omitempty"`
	PublicGists       *int       `json:"public_gists,omitempty"`
	Followers         *int       `json:"followers,omitempty"`
	Following         *int       `json:"following,omitempty"`
	CreatedAt         *Timestamp `json:"created_at,omitempty"`
	UpdatedAt         *Timestamp `json:"updated_at,omitempty"`
	Type              *string    `json:"type,omitempty"`
	SiteAdmin         *bool      `json:"site_admin,omitempty"`
	TotalPrivateRepos *int       `json:"total_private_repos,omitempty"`
	OwnedPrivateRepos *int       `json:"owned_private_repos,omitempty"`
	PrivateGists      *int       `json:"private_gists,omitempty"`
	DiskUsage         *int       `json:"disk_usage,omitempty"`
	Collaborators     *int       `json:"collaborators,omitempty"`
	Plan              *Plan      `json:"plan,omitempty"`

	// API URLs
	URL               *string `json:"url,omitempty"`
	EventsURL         *string `json:"events_url,omitempty"`
	FollowingURL      *string `json:"following_url,omitempty"`
	FollowersURL      *string `json:"followers_url,omitempty"`
	GistsURL          *string `json:"gists_url,omitempty"`
	OrganizationsURL  *string `json:"organizations_url,omitempty"`
	ReceivedEventsURL *string `json:"received_events_url,omitempty"`
	ReposURL          *string `json:"repos_url,omitempty"`
	StarredURL        *string `json:"starred_url,omitempty"`
	SubscriptionsURL  *string `json:"subscriptions_url,omitempty"`

	// TextMatches is only populated from search results that request text matches
	// See: search.go and https://developer.github.com/v3/search/#text-match-metadata

	// TextMatches 只是填入了搜索文本匹配请求的结果.
	// See: search.go and https://developer.github.com/v3/search/#text-match-metadata
	TextMatches []TextMatch `json:"text_matches,omitempty"`
}

func (u User) String() string

// UserEmail represents user's email address

// UserEmail 表示用户的 email 地址.
type UserEmail struct {
	Email    *string `json:"email,omitempty"`
	Primary  *bool   `json:"primary,omitempty"`
	Verified *bool   `json:"verified,omitempty"`
}

// UserListOptions specifies optional parameters to the UsersService.List method.

// UserListOptions 指定 UsersService.List 方法的可选参数.
type UserListOptions struct {
	// ID of the last user seen
	Since int `url:"since,omitempty"`
}

// UsersSearchResult represents the result of an issues search.

// UsersSearchResult 表示问题搜索结果.
type UsersSearchResult struct {
	Total *int   `json:"total_count,omitempty"`
	Users []User `json:"items,omitempty"`
}

// UsersService handles communication with the user related methods of the GitHub
// API.
//
// GitHub API docs: http://developer.github.com/v3/users/

// UsersService 处理 GitHub API 中与用户相关的通信方法.
//
// GitHub API 文档: http://developer.github.com/v3/users/
type UsersService struct {
	// contains filtered or unexported fields
}

// AddEmails adds email addresses of the authenticated user.
//
// GitHub API docs:
// http://developer.github.com/v3/users/emails/#add-email-addresses

// AddEmails 添加授权用的 email 地址.
//
// GitHub API 文档:
// http://developer.github.com/v3/users/emails/#add-email-addresses
func (s *UsersService) AddEmails(emails []string) ([]UserEmail, *Response, error)

// CreateKey adds a public key for the authenticated user.
//
// GitHub API docs: http://developer.github.com/v3/users/keys/#create-a-public-key

// CreateKey 为授权用添加公钥.
//
// GitHub API 文档: http://developer.github.com/v3/users/keys/#create-a-public-key
func (s *UsersService) CreateKey(key *Key) (*Key, *Response, error)

// DeleteEmails deletes email addresses from authenticated user.
//
// GitHub API docs:
// http://developer.github.com/v3/users/emails/#delete-email-addresses

// DeleteEmails 删除授权用户的 email 地址.
//
// GitHub API 文档:
// http://developer.github.com/v3/users/emails/#delete-email-addresses
func (s *UsersService) DeleteEmails(emails []string) (*Response, error)

// DeleteKey deletes a public key.
//
// GitHub API docs: http://developer.github.com/v3/users/keys/#delete-a-public-key

// DeleteKey 删除公匙.
//
// GitHub API 文档: http://developer.github.com/v3/users/keys/#delete-a-public-key
func (s *UsersService) DeleteKey(id int) (*Response, error)

// DemoteSiteAdmin demotes a user from site administrator of a GitHub Enterprise
// instance.
//
// GitHub API docs:
// https://developer.github.com/v3/users/administration/#demote-a-site-administrator-to-an-ordinary-user

// DemoteSiteAdmin 降级一个 GitHub 企业实例网站管理员用户.
//
// GitHub API 文档:
// https://developer.github.com/v3/users/administration/#demote-a-site-administrator-to-an-ordinary-user
func (s *UsersService) DemoteSiteAdmin(user string) (*Response, error)

// Edit the authenticated user.
//
// GitHub API docs:
// http://developer.github.com/v3/users/#update-the-authenticated-user

// Edit 编辑授权用户.
//
// GitHub API 文档:
// http://developer.github.com/v3/users/#update-the-authenticated-user
func (s *UsersService) Edit(user *User) (*User, *Response, error)

// Follow will cause the authenticated user to follow the specified user.
//
// GitHub API docs: http://developer.github.com/v3/users/followers/#follow-a-user

// Follow 使授权用户关注指定用户.
//
// GitHub API 文档: http://developer.github.com/v3/users/followers/#follow-a-user
func (s *UsersService) Follow(user string) (*Response, error)

// Get fetches a user. Passing the empty string will fetch the authenticated user.
//
// GitHub API docs: http://developer.github.com/v3/users/#get-a-single-user

// Get 获取一个用户. 传递空字符串将获取授权用户.
//
// GitHub API 文档: http://developer.github.com/v3/users/#get-a-single-user
func (s *UsersService) Get(user string) (*User, *Response, error)

// GetKey fetches a single public key.
//
// GitHub API docs:
// http://developer.github.com/v3/users/keys/#get-a-single-public-key

// GetKey fetches a single public key.
//
// GitHub API 文档:
// http://developer.github.com/v3/users/keys/#get-a-single-public-key
func (s *UsersService) GetKey(id int) (*Key, *Response, error)

// IsFollowing checks if "user" is following "target". Passing the empty string for
// "user" will check if the authenticated user is following "target".
//
// GitHub API docs:
// http://developer.github.com/v3/users/followers/#check-if-you-are-following-a-user

// IsFollowing 检查 user 是否关注了 target.
// 传递空字符串给 user 将检查授权用户是否关注了 target.
//
// GitHub API 文档:
// http://developer.github.com/v3/users/followers/#check-if-you-are-following-a-user
func (s *UsersService) IsFollowing(user, target string) (bool, *Response, error)

// ListAll lists all GitHub users.
//
// GitHub API docs: http://developer.github.com/v3/users/#get-all-users

// ListAll 罗列所有的 GitHub 用户.
//
// GitHub API 文档: http://developer.github.com/v3/users/#get-all-users
func (s *UsersService) ListAll(opt *UserListOptions) ([]User, *Response, error)

// ListEmails lists all email addresses for the authenticated user.
//
// GitHub API docs:
// http://developer.github.com/v3/users/emails/#list-email-addresses-for-a-user

// ListEmails 罗列授权用户所有的 email 地址.
//
// GitHub API 文档:
// http://developer.github.com/v3/users/emails/#list-email-addresses-for-a-user
func (s *UsersService) ListEmails(opt *ListOptions) ([]UserEmail, *Response, error)

// ListFollowers lists the followers for a user. Passing the empty string will
// fetch followers for the authenticated user.
//
// GitHub API docs:
// http://developer.github.com/v3/users/followers/#list-followers-of-a-user

// ListFollowers 罗列用户的粉丝. 传递空字符串将罗列授权用户的粉丝.
//
// GitHub API 文档:
// http://developer.github.com/v3/users/followers/#list-followers-of-a-user
func (s *UsersService) ListFollowers(user string, opt *ListOptions) ([]User, *Response, error)

// ListFollowing lists the people that a user is following. Passing the empty
// string will list people the authenticated user is following.
//
// GitHub API docs:
// http://developer.github.com/v3/users/followers/#list-users-followed-by-another-user

// ListFollowing 罗列某用户关注的人. 传递空字符串将罗列授权用户关注的人.
//
// GitHub API 文档:
// http://developer.github.com/v3/users/followers/#list-users-followed-by-another-user
func (s *UsersService) ListFollowing(user string, opt *ListOptions) ([]User, *Response, error)

// ListKeys lists the verified public keys for a user. Passing the empty string
// will fetch keys for the authenticated user.
//
// GitHub API docs:
// http://developer.github.com/v3/users/keys/#list-public-keys-for-a-user

// ListKeys 罗列某用户验证过的公匙. 传递空字符串将罗列授权用户验证过的公匙.
//
// GitHub API 文档:
// http://developer.github.com/v3/users/keys/#list-public-keys-for-a-user
func (s *UsersService) ListKeys(user string, opt *ListOptions) ([]Key, *Response, error)

// PromoteSiteAdmin promotes a user to a site administrator of a GitHub Enterprise
// instance.
//
// GitHub API docs:
// https://developer.github.com/v3/users/administration/#promote-an-ordinary-user-to-a-site-administrator

// PromoteSiteAdmin promotes a user to a site administrator of a GitHub Enterprise
// instance.
//
// GitHub API 文档:
// https://developer.github.com/v3/users/administration/#promote-an-ordinary-user-to-a-site-administrator
func (s *UsersService) PromoteSiteAdmin(user string) (*Response, error)

// Suspend a user on a GitHub Enterprise instance.
//
// GitHub API docs:
// https://developer.github.com/v3/users/administration/#suspend-a-user

// Suspend 挂起某 GitHub 企业实例用户.
//
// GitHub API 文档:
// https://developer.github.com/v3/users/administration/#suspend-a-user
func (s *UsersService) Suspend(user string) (*Response, error)

// Unfollow will cause the authenticated user to unfollow the specified user.
//
// GitHub API docs: http://developer.github.com/v3/users/followers/#unfollow-a-user

// Unfollow 使得授权用户取消关注某用户.
//
// GitHub API 文档: http://developer.github.com/v3/users/followers/#unfollow-a-user
func (s *UsersService) Unfollow(user string) (*Response, error)

// Unsuspend a user on a GitHub Enterprise instance.
//
// GitHub API docs:
// https://developer.github.com/v3/users/administration/#unsuspend-a-user

// Unsuspend 取消挂起某 GitHub 企业实例用户.
//
// GitHub API 文档:
// https://developer.github.com/v3/users/administration/#unsuspend-a-user
func (s *UsersService) Unsuspend(user string) (*Response, error)

// WebHookAuthor represents the author or committer of a commit, as specified in a
// WebHookCommit. The commit author may not correspond to a GitHub User.

// WebHookAuthor 表示提交的作者或提交者, 如在 WebHookCommit 中指定.
// 该提交者可能未对应到 GitHub 用户.
type WebHookAuthor struct {
	Email    *string `json:"email,omitempty"`
	Name     *string `json:"name,omitempty"`
	Username *string `json:"username,omitempty"`
}

func (w WebHookAuthor) String() string

// WebHookCommit represents the commit variant we receive from GitHub in a
// WebHookPayload.

// WebHookCommit 表示来自 WebHookPayload 收到的 GitHub 提交变量.
type WebHookCommit struct {
	Added     []string       `json:"added,omitempty"`
	Author    *WebHookAuthor `json:"author,omitempty"`
	Committer *WebHookAuthor `json:"committer,omitempty"`
	Distinct  *bool          `json:"distinct,omitempty"`
	ID        *string        `json:"id,omitempty"`
	Message   *string        `json:"message,omitempty"`
	Modified  []string       `json:"modified,omitempty"`
	Removed   []string       `json:"removed,omitempty"`
	Timestamp *time.Time     `json:"timestamp,omitempty"`
}

func (w WebHookCommit) String() string

// WebHookPayload represents the data that is received from GitHub when a push
// event hook is triggered. The format of these payloads pre-date most of the
// GitHub v3 API, so there are lots of minor incompatibilities with the types
// defined in the rest of the API. Therefore, several types are duplicated here to
// account for these differences.
//
// GitHub API docs: https://help.github.com/articles/post-receive-hooks

// WebHookPayload 表示来自触发推送事件钩子收到的 GitHub 数据.
// 预期负载格式大部分为 GitHub v3 API, 所以为数不多的不兼容类型由其余 API 定义.
// 因此, 这里复制多种类型来解决这些差异.
//
// GitHub API 文档: https://help.github.com/articles/post-receive-hooks
type WebHookPayload struct {
	After      *string         `json:"after,omitempty"`
	Before     *string         `json:"before,omitempty"`
	Commits    []WebHookCommit `json:"commits,omitempty"`
	Compare    *string         `json:"compare,omitempty"`
	Created    *bool           `json:"created,omitempty"`
	Deleted    *bool           `json:"deleted,omitempty"`
	Forced     *bool           `json:"forced,omitempty"`
	HeadCommit *WebHookCommit  `json:"head_commit,omitempty"`
	Pusher     *User           `json:"pusher,omitempty"`
	Ref        *string         `json:"ref,omitempty"`
	Repo       *Repository     `json:"repository,omitempty"`
}

func (w WebHookPayload) String() string

// WeeklyCommitActivity represents the weekly commit activity for a repository. The
// days array is a group of commits per day, starting on Sunday.

// WeeklyCommitActivity 表示仓库一周的提交活动.
// Days 数组是每日提交分组, 开始于星期天.
type WeeklyCommitActivity struct {
	Days  []int      `json:"days,omitempty"`
	Total *int       `json:"total,omitempty"`
	Week  *Timestamp `json:"week,omitempty"`
}

func (w WeeklyCommitActivity) String() string

// WeeklyStats represents the number of additions, deletions and commits a
// Contributor made in a given week.

// WeeklyStats 表示给定周内某贡献者产生的增加的, 删除的和提交的数量.
type WeeklyStats struct {
	Week      *Timestamp `json:"w,omitempty"`
	Additions *int       `json:"a,omitempty"`
	Deletions *int       `json:"d,omitempty"`
	Commits   *int       `json:"c,omitempty"`
}

func (w WeeklyStats) String() string
