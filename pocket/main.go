package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var sites = []Info{
	// HTTP
	{URL: "http://liqi.io", Handler2: handleLiQi},
	{URL: "http://wanqu.co", Handler2: handleWanQu},
	{URL: "http://johng.cn", Handler2: handleJohng},
	{URL: "http://fangkc.cn", Handler2: handleFangkc},
	{URL: "http://cizixs.com", Handler2: handleCizixs},
	{URL: "http://xiaorui.cc", Handler2: handleXiaoRui},
	{URL: "http://allenguo.com", Handler2: handleAllenGuo},
	{URL: "http://foocoder.com", Handler2: handleFoocoder},
	{URL: "http://blog.xiayf.cn", Handler2: handleXiayf},
	{URL: "http://air.googol.im", Handler2: handleAirOnG},
	{URL: "http://lihaoquan.me/", Handler2: handleLiHaoQuan},
	{URL: "http://sanyuesha.com", Handler2: handleSanYueSha},
	{URL: "http://xiaqunfeng.cc", Handler2: handleXiaQunFeng},
	{URL: "http://www.ibtsat.com", Handler2: handleIBTSAT}, // TODO: not yet
	{URL: "http://blog.yfgeek.com", Handler2: handleYfGeek},
	{URL: "http://www.gtdlife.com", Handler2: handleGTDLife},
	{URL: "http://www.readern.com", Handler2: handleReadErn},
	{URL: "http://www.rosipov.com", Handler2: handleROsipov},
	{URL: "http://www.flysnow.org", Handler2: handleFlySnow},
	{URL: "http://openmymind.net/", Handler2: handleOpenMyMind},
	{URL: "http://blog.mgechev.com", Handler2: handleMgechev},
	{URL: "http://gravitysworm.com", Handler2: handleGravitySworm},
	{URL: "http://nladuo.github.io/", Handler2: handleNladuo},
	{URL: "http://popozhu.github.io", Handler2: handlePopoZhu},
	{URL: "http://forrestchang.com/", Handler2: handleForrestChang},
	{URL: "http://www.mikewootc.com/", Handler2: handleMikeWootc},
	{URL: "http://www.carlpullein.com", Handler2: handleCarlPullein},
	{URL: "http://yinwang0.lofter.com", Handler2: handleYinWangLofter},
	{URL: "http://misscaffeinated.com", Handler2: handleMissCaffeinated},
	{URL: "http://appshere.bitcron.com", Handler2: handleAppShere},
	{URL: "http://mypaper.pchome.com.tw/zerojudge", Handler2: handleZeroJudge},
	{URL: "http://www.csie.ntnu.edu.tw/~u91029/index.html", Handler2: handleAlgorithmNote},
	// HTTPS
	{URL: "https://jesor.me", Handler2: handleJesor},
	{URL: "https://wzyboy.im", Handler2: handleWzyBoy},
	{URL: "https://wdxtub.com", Handler2: handleWdxtub},
	{URL: "https://lepture.com", Handler2: handleLepture},
	{URL: "https://tonybai.com", Handler2: handleTonyBai},
	{URL: "https://xiaozhou.net", Handler2: handleXiaoZhou},
	{URL: "https://liyang85.com", Handler2: handleLiYang85},
	{URL: "https://cindysss.com", Handler2: handleCindysss},
	{URL: "https://golangbot.com", Handler2: handleGolangBot},
	{URL: "https://adamdrake.com", Handler2: handleAdamDrake},
	{URL: "https://luvletter.moe", Handler2: handleLuvLetter},
	{URL: "https://zfanw.com/blog", Handler2: handleZFanW},
	{URL: "https://blog.zfanw.com", Handler2: handleZFanW2},
	{URL: "https://www.4async.com", Handler2: handle4Async},
	{URL: "https://www.calhoun.io", Handler2: handleCalhoun},
	{URL: "https://blog.chewxy.com", Handler2: handleChewxy},
	{URL: "https://dave.cheney.net", Handler2: handleCheney},
	{URL: "https://blog.ropnop.com", Handler2: handleRopnop},
	{URL: "https://blog.golang.org", Handler2: handleGolangBlog},
	{URL: "https://jamesstuber.com", Handler2: handleJamesStuber},
	{URL: "https://tech.meituan.com", Handler2: handleMeiTuan},
	{URL: "https://blog.merovius.de", Handler2: handleMerovius},
	{URL: "https://mycodesmells.com", Handler2: handleMyCodeSmells},
	{URL: "https://bitsplitting.org", Handler2: handleBitsplitting},
	{URL: "https://blog.jimmylv.info", Handler2: handleJimmyLv},
	{URL: "https://marcjenkins.co.uk", Handler2: handleMarcJenkins},
	{URL: "https://joecieplinski.com", Handler2: handleJoecieplinski},
	{URL: "https://hamberg.no/erlend", Handler2: handleErlendHamberg},
	{URL: "https://sheepbao.github.io", Handler2: handleSheepBao},
	{URL: "https://blog.just4fun.site", Handler2: handleJust4Fun},
	{URL: "https://blog.agilebits.com", Handler2: handleAgilebits},
	{URL: "https://www.playpcesor.com", Handler2: handlePlayPCEsor},
	{URL: "https://www.kawabangga.com", Handler2: handleKawabangga},
	{URL: "https://www.macstories.net", Handler2: handleMacStories},
	{URL: "https://www.discoverdev.io", Handler2: handleDiscoverDev},
	{URL: "https://hiberabyss.github.io", Handler2: handleHiberAbyss},
	{URL: "https://wikimatze.de/articles", Handler2: handleWikiMatze},
	{URL: "https://www.dataquest.io/blog", Handler2: handleDataQuest},
	{URL: "https://blog.yitianshijie.net", Handler2: handleYiTianShiJie},
	{URL: "https://www.everythingcli.org", Handler2: handleEverythingCLI},
	{URL: "https://leetcode.com/articles", Handler2: handleLeetcodeArticle},
	{URL: "https://www.ardanlabs.com/blog", Handler2: handleArdanLabs},
	{URL: "https://kingdomhe.wordpress.com", Handler2: handleKingdomhe},
	{URL: "https://itimetraveler.github.io", Handler2: handleITimeTraveler},
	{URL: "https://nathanleclaire.com/post", Handler2: handleNathanLeclaire},
	{URL: "https://illidiumq36.wordpress.com", Handler2: handleIllidiumq36},
	{URL: "https://blog.stevenocchipinti.com", Handler2: handleStevenOcchipinti},
	{URL: "https://go101.org/article/101.html", Handler2: handleGo101},
	{URL: "https://matthewphillips.info/programming/index.html", Handler2: handleMatthewPhillips},
	{URL: "https://www.douban.com/people/4553925/notes?start=0&type=note", Handler2: handleMissPower},
	{
		URL:      "http://haohailong.net",
		ListPath: "div.posts div[id] div.post-inner div.post-header",
		URLPath:  "h2.post-title a",
		NextPath: "div.archive-nav a.post-nav-older",
	},
	{
		URL:      "https://scomper.me",
		ListPath: "div.content div.post.animated.fadeInDown",
		URLPath:  "h2 a",
		NextPath: "div.paginator.pager.pagination a.btn.next.older-posts.older_posts",
	},
	{
		URL:      "https://zapier.com/blog",
		ListPath: "div.entries div.excerpt",
		URLPath:  "h2.title a",
		NextPath: "div.pagination div.page-nav a.next",
	},
	{
		URL:       "http://www.markwk.com",
		URLSuffix: "/blog/archives/",
		ListPath:  "div#blog-archives article",
		URLPath:   "h1 a",
	},
	{
		URL:      "https://vickylai.com/verbose",
		ListPath: "div#blog-link-list li",
		URLPath:  "a.blog-post-link",
	},
	{
		URL:      "https://liudanking.com",
		ListPath: "div#content article[id]",
		URLPath:  "h1.entry-title a",
		NextPath: "nav#nav-below div.nav-previous a",
	},
	{
		URL:       "https://www.lifesuccessengineer.com",
		URLSuffix: "/blog/",
		ListPath:  "section.bSe.right div.awr",
		URLPath:   "h2.entry-title a",
		NextPath:  "div.pgn.clearfix a.next.page-numbers",
	},
	{
		URL:      "https://unclutterer.com",
		ListPath: "div.content.row article[class]",
		URLPath:  "h2.entry-title a",
		NextPath: "nav.post-nav li.previous a",
	},
	{
		URL:      "http://www.leyafo.com",
		ListPath: "article.post.post",
		URLPath:  "h3.post-title a",
		NextPath: "nav.pagination a.older-posts",
	},
	{
		URL:       "http://www.geekpreneur.com",
		URLSuffix: "/archives-index",
		ListPath:  "div.azindex ul li",
		URLPath:   "a",
	},
	{
		URL:      "https://www.do1618.com",
		ListPath: "div#content article[id]",
		URLPath:  "header.entry-header h1.entry-title a",
		NextPath: "nav.navigation div.nav-previous a",
	},
	{
		URL:      "https://www.iplaysoft.com",
		ListPath: "div#postlist div[class][itemtype]",
		URLPath:  "div.entry-head h2.entry-title a",
		NextPath: "div.pagenavi-simple a",
		// NextCondition: "i.ipsicon.ipsicon-next.ipsicon-lspace",
	},
	{
		URL:       "https://www.stevepavlina.com",
		URLSuffix: "/archives",
		ListPath:  "div.sya_container ul li",
		URLPath:   "div.sya_postcontent a",
	},
	{
		URL:      "http://cyhsu.xyz",
		ListPath: "div.content div.post-title",
		URLPath:  "h3 a",
		NextPath: "div.pagination li.next.pagbuttons a",
	},
	{
		URL:       "http://x-wei.github.io",
		URLSuffix: "/archives.html",
		ListPath:  "section#content div#archives p",
		URLPath:   "a[class]",
	},
	{
		URL:       "https://mymorningroutine.com",
		URLSuffix: "/routines/all/#continue-routine",
		ListPath:  "div#js-archive-list div.card-img.card-img--archive",
		URLPath:   "a.u-block",
	},
	{
		URL:      "https://miao.hu",
		ListPath: "li.mv2",
		URLPath:  "a", // remove time in title
	},
	{
		URL:      "http://yuezhu.org",
		ListPath: "section.entryTypePostExcerptContainer article[class]",
		URLPath:  "h2.entryTitle a",
		NextPath: "div.posts-pagination a",
		// NextCondition: "span.previous-posts-link",
	},
	{
		URL:       "http://blog.leanote.com",
		URLSuffix: "/archives/carlking5019",
		ListPath:  "div#posts div.each-post ul li",
		URLPath:   "a",
	},
	{
		URL:       "http://www.monkeyuser.com",
		URLSuffix: "/toc/",
		ListPath:  "div.toc div.toc-entry",
		URLPath:   "div.et div a[href]",
	},
	{
		URL:      "https://blog.brickgao.com",
		ListPath: "div.post-summary",
		URLPath:  "div.post-title a",
		NextPath: "div.paginator a.extend.next",
	},
	{
		URL:      "http://blog.yuelong.info",
		ListPath: "section#posts article[class]",
		URLPath:  "h2 a",
		NextPath: "div.alignleft a",
	},
	{
		URL:       "http://onespiece.strikingly.com",
		URLSuffix: "/blog/df60b9b6a7b",
		NextPath:  "span.s-blog-footer-btn.s-blog-footer-previous a",
	},
	{
		URL:       "http://gank.io",
		URLSuffix: "/history",
		ListPath:  "li div.row",
		URLPath:   "a",
	},
	{
		URL:      "https://www.appinn.com",
		ListPath: "div#spost div.spost.post",
		URLPath:  "h2.entry-title a",
		NextPath: "div.navigation a.nextpostslink",
	},
	{
		URL:       "https://productivityist.com",
		URLSuffix: "/category/blog/",
		ListPath:  "article[class]",
		URLPath:   "h2.entry-title a",
		NextPath:  "li.pagination-next a",
	},
	{
		URL:      "https://maqmodo.com",
		ListPath: "div.blogpostcategory",
		// TitlePath: "h2.title",
		URLPath:  "h2.title a",
		NextPath: "div.wp-pagenavi a.nextpostslink",
	},
	{
		URL:      "http://wsfdl.com", // 网址有中文
		ListPath: "ul.post-list li",
		URLPath:  "h2 a",
	},
	{
		URL:      "https://blog.jez.io",
		ListPath: "article.hentry div.entry-wrapper",
		URLPath:  "h3.entry-title a",
	},
	{
		URL:      "https://usesthis.com",
		ListPath: "article.interviewee.h-card",
		URLPath:  "h3 a",
		NextPath: "nav#paginator a#next",
	},
	{
		URL:       "http://www.asianefficiency.com", // something wrong
		URLSuffix: "/blog",
		ListPath:  "article[class]",
		// TitlePath: "h1",
		URLPath:  "h1 a",
		NextPath: "nav.archive.pagination div.next a",
	},
	{
		URL:      "https://deans.site",
		ListPath: "div.post_list div.post",
		URLPath:  "div.post_title a",
		NextPath: "div.paginator.pager.pagination a.btn.next.older-posts.older_posts",
	},
	// now in descending order
	{
		URL:      "http://blog.josui.me",
		ListPath: "div.blog div.content article",
		NextPath: "nav.pagination a.pagination-next",
		// Handler:  handleJosuiWritings,
	},
	{
		URL:      "http://www.yinwang.org",
		ListPath: "li.list-group-item.title",
		// Handler:  handleYinWang,
	},
	{
		URL:      "http://jannerchang.bitcron.com",
		ListPath: "div.post_in_list",
		NextPath: "div.paginator.pager.pagination a.btn.next.older-posts.older_posts",
		// Handler:  handleJannerChang,
	},
	{
		URL:      "https://blog.todoist.com",
		ListPath: "article.tdb-article-slat",
		NextPath: "div.tdb-pagination-holder a.next.page-numbers.nav__action",
		// Handler:  handleTodoist,
	},
	{
		URL:      "https://ulyssesapp.com/blog",
		ListPath: "main#main article[id]",
		NextPath: "nav.navigation.paging-navigation div.nav-previous a",
		// Handler:  handleUlysses,
	},
	{
		URL:       "https://www.scotthyoung.com/blog",
		URLSuffix: "/articles",
		ListPath:  "div#date-block ul li",
		// Handler:   handleScottHYoung,
	},
	{
		URL:      "https://startupnextdoor.com",
		ListPath: "main#content article[class]",
		NextPath: "nav.pagination a.older-posts",
		// Handler:  handleStartUpNextDoor,
	},
	{
		URL:      "https://blog.trello.com",
		ListPath: "div.post-item",
		NextPath: "div.blog-pagination a.next-posts-link",
		// Handler:  handleTrello,
	},
	{
		URL:      "http://www.catcoder.com",
		ListPath: "section#posts article",
		NextPath: "nav.pagination a.extend.next",
		// Handler:  handleCatCoder,
	},
	{
		URL:      "https://www.waerfa.com",
		ListPath: "main#main article[id]",
		NextPath: "nav.navigation.posts-navigation div.nav-previous a",
		// Handler:  handleWaerfa,
	},
	{
		URL:      "https://unee.wang",
		ListPath: "div.post-list div.mod-post",
		NextPath: "div.paginator.pager.pagination a.btn.next.older-posts.older_posts",
		// Handler:  handleUneeWang,
	},
	{
		URL:      "https://xiaomu.bitcron.com",
		ListPath: "div.post",
		NextPath: "div.paginator.pager.pagination a.btn.next.older-posts.older_posts",
		// Handler:  handleXiaomu,
	},
}

// Info stores some basic info for one site
type Info struct {
	URL       string               `json:"url"`
	URLSuffix string               `json:"url_suffix"`
	ListPath  string               `json:"list_path"`
	URLPath   string               `json:"url_path"`
	NextPath  string               `json:"next_path"`
	Handle    bool                 `json:"handle"`
	Handler2  func(string, Pocket) `json:"handler_2"`
}

// read the article about how to get access token:
// http://www.cnblogs.com/febwave/p/4242333.html

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// Pocket holds some keys
type Pocket struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
}

// NewPocket creates a Pocket structure for operations
func NewPocket() Pocket {
	f, err := os.Open("auth.json")
	handleError(err)
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	handleError(err)

	p := Pocket{}
	if err := json.Unmarshal(bs, &p); err != nil {
		panic(err)
	}
	return p
}

type action struct {
	Action string `json:"action"`
	URL    string `json:"url"`
}

// AddMultiple adds multiple urls at one time
func (p Pocket) AddMultiple(urls []string) {
	actions := []action{}
	for _, url := range urls {
		actions = append(actions, action{
			Action: "add",
			URL:    url,
		})
	}
	body := struct {
		ConsumerKey string   `json:"consumer_key"`
		AccessToken string   `json:"access_token"`
		Actions     []action `json:"actions"`
	}{
		ConsumerKey: p.ConsumerKey,
		AccessToken: p.AccessToken,
		Actions:     actions,
	}
	bs, err := json.Marshal(body)
	handleError(err)
	req, err := http.Post(
		"https://getpocket.com/v3/send",
		"application/json",
		bytes.NewReader(bs),
	)
	handleError(err)
	if req.StatusCode != 200 {
		panic(req.Status + " fail to save articles: " + strings.Join(urls, "\n"))
	}
	fmt.Printf("[%s] Added %d articles done.\n",
		time.Now().Format(time.RFC3339), len(urls))
}

// Usage: go run *.go
func main() {
	p := NewPocket()

	for i, site := range sites {
		if !site.Handle {
			fmt.Printf("%3d: [skip] %s\n", i+1, site.URL)
			continue
		}
		fmt.Println("Started:", site.URL)
		if site.Handler2 == nil {
			panic(fmt.Sprintf("missing handler for site: %s\n", site.URL))
		}
		site.Handler2(site.URL, p)
		break
	}
}
