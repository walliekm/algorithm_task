//355. 设计推特
//https://leetcode-cn.com/problems/design-twitter/
//维护一个userMap，保存所有用户的信息
//每个用户的信息包括关注者集合，及其发布的推文列表，推文列表采用单链表结构
//用大根堆实现推文的优先队列，发布时间越近的优先级越高
//假设一个用户关注了N个人，每个人发布了M条推文，则获取timeline的时间复杂度为 O(MNLogN)
//关注、取消关注、发布推文，时间复杂度均为 O(1)

package week4

import (
	"container/heap"
	"time"
)

//推文结构
type tweet struct {
	tweetId   int       //推文id
	postedAt  time.Time //发布时间
	nextTweet *tweet    //指向下一条推文的指针
}

//时间线定义，实现heap接口，将发布时间作为优先队列的key
//将所有关注者(包含自己)的最新一条推文加入队列
//获取推文时，循环取出最新的一条，形成timeline
type timeline []*tweet

//heap默认为小根堆，要把时间较新的排在前面
func (tl timeline) Less(i, j int) bool  { return tl[i].postedAt.After(tl[j].postedAt) }
func (tl timeline) Len() int            { return len(tl) }
func (tl timeline) Swap(i, j int)       { tl[i], tl[j] = tl[j], tl[i] }
func (tl *timeline) Push(x interface{}) { *tl = append(*tl, x.(*tweet)) }
func (tl *timeline) Pop() interface{} {
	old := *tl
	n := len(old)
	item := old[n-1]
	old[n-1] = nil //防止内存泄漏
	*tl = old[:n-1]
	return item
}

//用户相关信息
type user struct {
	followed map[int]struct{} //关注者集合
	tweets   *tweet           //发布的推文列表
}

type Twitter struct {
	userMap map[int]*user //保存所有用户信息的map，key为userId
}

func Constructor() Twitter {
	return Twitter{
		userMap: map[int]*user{},
	}
}

//检查用户信息是否初始化，未初始化的进行初始化
func (this *Twitter) chkInitUser(userId int) {
	if _, ok := this.userMap[userId]; !ok {
		this.userMap[userId] = &user{
			followed: map[int]struct{}{},
			tweets:   &tweet{}, //链表表头保护结点
		}
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.chkInitUser(userId)
	newTweet := &tweet{
		tweetId:  tweetId,
		postedAt: time.Now(),
	}

	//在链表头部插入新推文
	newTweet.nextTweet = this.userMap[userId].tweets.nextTweet
	this.userMap[userId].tweets.nextTweet = newTweet
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	this.chkInitUser(userId)

	//初始化timeline列表，从自己和所有关注者里分别获取最新一条推文，加入列表
	var tl timeline
	if this.userMap[userId].tweets.nextTweet != nil {
		tl = append(tl, this.userMap[userId].tweets.nextTweet)
	}

	for fuid := range this.userMap[userId].followed {
		if this.userMap[fuid].tweets.nextTweet != nil {
			tl = append(tl, this.userMap[fuid].tweets.nextTweet)
		}
	}

	n := 10                  //获取推文的数量
	ans := make([]int, 0, n) //结果数据
	heap.Init(&tl)           //将timeline列表初始化为优先队列

	//从优先队列取出最新一条推文，并将其指向的下一条推文加入队列
	for len(ans) < n && len(tl) > 0 {
		t := heap.Pop(&tl).(*tweet)
		ans = append(ans, t.tweetId)
		if t.nextTweet != nil {
			heap.Push(&tl, t.nextTweet)
		}
	}

	return ans
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	this.chkInitUser(followerId)
	this.chkInitUser(followeeId)
	this.userMap[followerId].followed[followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	this.chkInitUser(followerId)
	this.chkInitUser(followeeId)
	delete(this.userMap[followerId].followed, followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
