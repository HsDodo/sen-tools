package k8s

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/retry"
)

func TestRetry(t *testing.T) {

	// 记录重试耗时
	startTime := time.Now()
	step := 0
	retry.OnError(wait.Backoff{
		Steps:    5,
		Duration: 20 * time.Millisecond,
		Factor:   1.0,
		Jitter:   0.1,
	}, func(err error) bool {
		step++
		return err != nil
	}, func() error {
		ipList, err := net.DefaultResolver.LookupIP(context.Background(), "ip", "www.cnblogs.com")
		fmt.Println(ipList[0].To4())
		return err
	})
	elapsedTime := time.Since(startTime)
	t.Logf("重试次数：%v \n 重试耗时: %v", step, elapsedTime)

}
