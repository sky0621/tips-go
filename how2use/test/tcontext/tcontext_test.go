package tcontext

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTContext_Wait(t *testing.T) {
	ctx := t.Context()

	var wg sync.WaitGroup

	// クリーンナップ（＝テスト終了）時に
	// WaitGroupのWaitを呼び出す（＝すべてのゴルーチンの終了を待機する）
	// これにより、このテスト終了後に別のテストが起動しても、
	// このテストで起動したすべてのゴルーチンが終了している（＝別のテストに影響が出ない）ことを保証できる。
	t.Cleanup(wg.Wait)

	wg.Add(1)
	go worker(ctx, &wg, "worker1")

	wg.Add(1)
	go worker(ctx, &wg, "worker2")

	wg.Add(1)
	go worker(ctx, &wg, "worker3")
}

func worker(ctx context.Context, wg *sync.WaitGroup, workerName string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: context canceled: %v\n", workerName, ctx.Err())
			return
		default:
			fmt.Printf("%s: working...\n", workerName)
			time.Sleep(100 * time.Millisecond) // 何かしら重い処理に相当
		}
	}
}
