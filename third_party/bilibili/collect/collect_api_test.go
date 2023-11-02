package collect

import (
	"context"
	"fmt"
	bg "github.com/iyear/biligo"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestBilibiliClientImpl_QueryMidTotalAidList(t *testing.T) {
	client := NewBilibiliClient()
	mid := int64(35847683) // 峰哥亡命天涯
	totalAidList, totalCount, err := client.QueryMidTotalAidList(context.TODO(), mid, proto.Int64(60))
	if err != nil {
		t.Fatalf("get err %+v", err)
	}
	t.Logf("total aid size=%v, aid list=%+v", len(totalAidList), totalAidList)
	if len(totalAidList) == 0 {
		t.Fatalf("total aid list size is 0")
	}
	if totalCount <= 0 {
		t.Fatalf("total_count=%v < 0", totalCount)
	}
}

func TestSpaceSearchVideo(t *testing.T) {
	commClient := bg.NewCommClient(&bg.CommSetting{Client: nil, DebugMode: true,
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
	})
	searchVideo, err := commClient.SpaceSearchVideo(35847683, "", 0, "", 1, 10)
	fmt.Println(err)
	fmt.Println(searchVideo)
}

func TestBilibiliClientImpl_QueryVideoInfoByAid(t *testing.T) {
	client := NewBilibiliClient()
	videoInfo, err := client.QueryVideoInfoByAid(context.TODO(), 478629139)
	if err != nil {
		t.Fatalf("QueryVideoInfoByAid failed, err=%+v", err)
	}
	t.Logf("QueryVideoInfoByAid get video info=%+v", videoInfo)
	if videoInfo == nil {
		t.Fatalf("QueryVideoInfoByAid get nil")
	}
}
