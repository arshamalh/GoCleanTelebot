package telegram

import (
	"bytes"
	"context"
	"coryptex.com/bot/vip-signal/internal/data/repositories"
	"coryptex.com/bot/vip-signal/internal/domain/entities"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"os"
)

type telegram struct {
	client *http.Client
}

// NewTelegram help tp port new instance of telegram httpclient to other layers
func NewTelegram(c *http.Client) repositories.SignalPub {
	return &telegram{c}
}

func (sl *telegram) Publish(ctx context.Context, s entities.Signal) error {
	// Publish to telegram
	query := "https://api.telegram.org/bot" + os.Getenv("TOKEN") + "/sendphoto"
	caption :=
		"#" + s.Pair + "\n\n" +
			"💰 Price:               " + s.EntryPrice + " $\n" +
			"✅ Take Profit:    " + s.TargetPrice + " $\n" +
			"❌ Stop-Loss:      " + s.StopLoss + " $\n" +
			"📈 Reward/Risk: " + s.Risk2Reward + "\n" +
			"➡️ Indicators:      " + s.SignalDirection() + "\n" +
			"-------------------\n" +
			"📅 " + s.DateRef("/") + "\n" +
			"🔰 Tahlil Crypto Vip"

	postBody, _ := json.Marshal(map[string]string{
		"photo":      s.ImageURL,
		"caption":    caption,
		"chat_id":    os.Getenv("CHNL_ID"),
		"parse_mode": "HTML",
	})
	responseBody := bytes.NewBuffer(postBody)

	//_, err := http.Post(query, "application/json", responseBody)
	_, err := sl.client.Post(query, "application/json", responseBody)

	if err != nil {
		return errors.Wrap(err, "Error Publishing to channel, internal.data.datasource.telegram")
	}
	/// TODO: Read the response and try to return better result, maybe like a struct
	//defer resp.Body.Close()
	////Read the response body
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//sb := string(body)
	//log.Printf(sb)
	return nil
}
