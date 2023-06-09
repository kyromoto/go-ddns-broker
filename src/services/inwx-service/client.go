package inwxservice

import (
	"net/netip"

	pubsubservice "github.com/kyromoto/go-ddns-broker/src/services/pubsub-service"
	"github.com/nrdcg/goinwx"
	"github.com/rs/zerolog/log"
)

// type Inwx struct {
// 	client goinwx.Client
// }

// func NewInwx() *Inwx {
// 	client := goinwx.NewClient(envservice.InwxUsername(), envservice.InwxPassword(), &goinwx.ClientOptions{})

// 	return &Inwx{
// 		client: *client,
// 	}
// }

type ClientRepo interface {
}

func HandleInwxDnsUpdates(clientRepo ClientRepo) {
	subscriber := pubsubservice.NewSubscriber(pubsubservice.TopicUrl)

	subscriber.Listen(func(username string, ip netip.Addr) error {
		client := goinwx.NewClient("", "", &goinwx.ClientOptions{})
		_, err := client.Account.Login()

		if err != nil {
			return err
		}

		defer func() {
			err := client.Account.Logout()

			if err != nil {
				log.Err(err)
			}
		}()

		var recordType string

		if ip.Is4() {
			recordType = "A"
		} else {
			recordType = "AAAA"
		}

		req := goinwx.NameserverInfoRequest{
			Name: "k-rz-28.ddns.kyro.space",
			Type: recordType,
		}

		res, err := client.Nameservers.Info(&req)

		if err != nil {
			return err
		}

		// client.Nameservers.UpdateRecord(res.RoID, &goinwx.NameserverRecordRequest{
		// 	Content: ,
		// })

		log.Printf("%v", res)

		return nil
	})
}
