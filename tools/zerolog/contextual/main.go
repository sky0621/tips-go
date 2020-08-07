package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	//zerologを使用すると、キーと値のペアの形式でログメッセージにデータを追加できます。
	//メッセージに追加されたデータは、ログイベントに関する「コンテキスト」を追加します。
	//これは、デバッグやその他の無数の目的にとって重要な場合があります。この例は以下です
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	//{"level":"debug","Scale":"833 cents","Interval":833.09,"time":1596761149,"message":"Fibonacci is everywhere"}
	//{"level":"debug","Name":"Tom","time":1596761149}
	log.Debug().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")
	log.Debug().
		Str("Name", "Tom").
		Send()
}
