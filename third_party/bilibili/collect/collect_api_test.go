package collect

import (
	"context"
	"testing"
)

func TestBilibiliClientImpl_QueryMidTotalAidList(t *testing.T) {
	client := NewBilibiliClient()
	mid := int64(35847683) // 峰哥亡命天涯
	totalAidList, err := client.QueryMidTotalAidList(context.TODO(), mid)
	if err != nil {
		t.Fatalf("get err %+v", err)
	}
	t.Logf("total aid size=%v, aid list=%+v", len(totalAidList), totalAidList)
	if len(totalAidList) == 0 {
		t.Fatalf("total aid list size is 0")
	}
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
