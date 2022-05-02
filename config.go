package upbit

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	KeyPair struct {
		AccessKey string // 엑세스 키
		SecretKey string // 비밀 키
	}
	TradableBalanceRatio float64 `default:"0.1"` // 비중
	OrderRatio           float64 `default:"0.5"` // 주문 가격 비중
	// 선택적으로 사용할 수 있는 옵션들
	MaxTrackedMarket int      `default:"10"` // 추적할 코인의 갯수 (분산투자 비율하고도 관련있음)
	Whitelist        []string // 마켓을 전략을 실행할 마켓을 수동으로 설정한다.
	Blacklist        []string // 해당 마켓은 제외한다.
}{}

func init() {
	config := "config.yml"
	configor.New(&configor.Config{Silent: true, ErrorOnUnmatchedKeys: true}).Load(&Config, config)
}
