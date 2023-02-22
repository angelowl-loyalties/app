package internal

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/campaigner/models"
	"go.etcd.io/etcd/client/v3"
)

const (
	dialTimeout = 5 * time.Second
)

var ETCD *clientv3.Client

func InitClient(endpointsString string) {
	endpoints := strings.Split(endpointsString, ",")

	var err error
	ETCD, err = clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})

	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Connected to: " + endpointsString)
	}
}

func etcdGetCampaign(id string) (err error) {
	response, err := ETCD.Get(context.Background(), "campaign_"+id)
	if err != nil {
		log.Println(err)
		return err
	}

	var campaign models.Campaign
	for _, ev := range response.Kvs {
		err := json.Unmarshal(ev.Value, &campaign)
		if err != nil {
			return err
		}
	}

	return nil
}

func etcdPutCampaign(campaign *models.Campaign) (err error) {
	jsonCampaign, err := json.Marshal(campaign)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = ETCD.Put(context.Background(), "campaign_"+campaign.ID.String(), string(jsonCampaign))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func etcdDeleteCampaign(id string) (deleted int, err error) {
	response, err := ETCD.Delete(context.Background(), "campaign_"+id)
	if err != nil {
		return 0, err
	}

	return int(response.Deleted), nil
}
