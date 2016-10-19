package main

import "log"

type talkroom struct {
	// 入室中の話者を管理
	talkersInOutMap map[*talker]bool
	// 入室しようとしている話者（チャネルでの通知）
	joinning chan *talker
	// 退室しようとしている話者（チャネルでの通知）
	leaving chan *talker
	// 入室中の話者からのトーク
	talk chan string
}

func newTalkroom() *talkroom {
	return &talkroom{
		talkersInOutMap: make(map[*talker]bool),
		joinning:        make(chan *talker),
		leaving:         make(chan *talker),
		talk:            make(chan string),
	}
}

// １つ（に固定）のトークルームが起動し、以後、話者の入退室とトークを待ち受ける
func runTalkRoom(r *talkroom) {
	log.Println("<talkroom.go>[runTalkRoom]START")
	// 起動後は無限ループ突入
	for {
		// チャネルの種類に応じた制御開始！
		// この中では言語の仕組みとして勝手に同期が取られている。
		select {
		case talker := <-r.joinning:
			log.Printf("<talkroom.go>[runTalkRoom]「%s」が入室しようとしている。\n", talker.name)
			r.talkersInOutMap[talker] = true
			log.Printf("<talkroom.go>[runTalkRoom]「%s」が入室が完了した。\n", talker.name)

		case talker := <-r.leaving:
			log.Printf("<talkroom.go>[runTalkRoom]「%s」が退室しようとしている。\n", talker.name)
			delete(r.talkersInOutMap, talker)
			// close(talker.talk) // トーク用のチャネルを閉じる
			log.Printf("<talkroom.go>[runTalkRoom]「%s」が退室が完了した。\n", talker.name)

		case talk := <-r.talk:
			log.Printf("<talkroom.go>[runTalkRoom]話者からトーク[%s]の通知が来た。\n", talk)
			// 全ての話者にトークを通知
			for talker := range r.talkersInOutMap {
				select {
				case talker.recv <- talk:
					log.Printf("<talkroom.go>[runTalkRoom]先ほど来たトークを話者[%s]へ通知した。\n", talker.name)
				default:
					log.Println("<talkroom.go>[runTalkRoom]デフォルトルート！")
					delete(r.talkersInOutMap, talker)
					// close(talker.talk) // トーク用のチャネルを閉じる
				}
			}
		}
	}
}
