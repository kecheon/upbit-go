package upbit

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/kecheon/upbit-go/util"
)

var client *Client

func setUp() {
	client = NewClient(Config.KeyPair.AccessKey, Config.KeyPair.SecretKey)
}

func ExampleGetMarkets() {
	setUp()

	markets, err := client.Markets()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(markets[0].Market)

	// Output:
	// KRW-BTC
}

func ExampleGetMinuteCandles() {
	setUp()

	candles, err := client.MinuteCandles(1, "KRW-BTC", map[string]string{
		"count": "1",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(candles))
	fmt.Println(candles[0].Market)

	// Output:
	// 1
	// KRW-BTC
}

func ExampleWrongUnitGetMinuteCandles() {
	setUp()

	_, err := client.MinuteCandles(2, "KRW-BTC")

	fmt.Println(err)

	// Output:
	// Invalid unit
}

func ExampleDayCandles() {
	setUp()

	candles, err := client.DayCandles("KRW-BTC", map[string]string{
		"count": "5",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(candles))
	fmt.Println(candles[0].Market)

	// Output:
	// 5
	// KRW-BTC
}

func ExampleWeekCandles() {
	setUp()

	candles, err := client.WeekCandles("KRW-BTC", map[string]string{
		"count": "5",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(candles))
	fmt.Println(candles[0].Market)

	// Output:
	// 5
	// KRW-BTC
}

func ExampleMonthCandles() {
	setUp()

	candles, err := client.MonthCandles("KRW-BTC", map[string]string{
		"count": "5",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(candles))
	fmt.Println(candles[0].Market)

	// Output:
	// 5
	// KRW-BTC
}

func ExampleTradeTicks() {
	setUp()

	tradeTicks, err := client.TradeTicks("KRW-BTC", map[string]string{
		"count": "5",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(tradeTicks))
	fmt.Println(tradeTicks[0].Market)

	// Output:
	// 5
	// KRW-BTC
}

func ExampleGetTickers() {
	setUp()

	ticks, err := client.Ticker("KRW-BTC, KRW-TRX")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(ticks))
	fmt.Println(ticks[0].Market, ticks[1].Market)

	// Output:
	// 2
	// KRW-BTC KRW-TRX
}

func ExampleOrderbooks() {
	setUp()

	ticks, err := client.Orderbooks("KRW-BTC, KRW-TRX")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(len(ticks))
	fmt.Println(ticks[0].Market, ticks[1].Market)
	fmt.Println(len(ticks[0].OrderbookUnits))

	// Output:
	// 2
	// KRW-BTC KRW-TRX
	// 15
}

func ExampleAccounts() {
	setUp()

	accounts, err := client.Accounts()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(accounts[0].Currency)

	// Output:
	// KRW
}

func ExampleOrderChange() {
	setUp()

	orderChance, err := client.OrderChance("KRW-BTC")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(orderChance.Market.Id)

	// Output:
	// KRW-BTC
}

func TestSellAndCancel(t *testing.T) {
	setUp()
	order, err := client.Order(
		strconv.Itoa(int(util.TimeStamp())),
		"ask",
		"KRW-BTC",
		"80000000",
		"0.0001",
		"limit",
	)
	if err != nil {
		t.Errorf("?????? ?????? ?????? %q", order)
	}

	fmt.Println(order)

	time.Sleep(time.Second * 1)
	cancel, err := client.CancelOrder(order.Uuid)
	fmt.Println(cancel)
	if err != nil {
		t.Errorf("?????? ?????? ?????? %s", order.Uuid)
	}
}
