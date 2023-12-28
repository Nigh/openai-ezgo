package openaiezgo

import "time"

func init() {
	go chatTimeoutTimer()
}

func chatTimeoutTimer() {
	timer := time.NewTicker(1 * time.Second)
	for range timer.C {
		for k, v := range Chats {
			if v.Timeout > 0 {
				v.Timeout--
			} else {
				if config.TimeoutCallback != nil {
					go config.TimeoutCallback(k, v.TokenUsed)
				}
				delete(Chats, k)
			}
		}
	}
}
