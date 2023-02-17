package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/rewarder/models"
	"log"
	"strings"
	"time"

	"go.etcd.io/etcd/client/v3"
)

const (
	dialTimeout = 2 * time.Second
)

var ETCD *clientv3.Client
var CampaignsEtcd = make(map[string]models.Campaign)
var ExclusionsEtcd = make(map[string]models.Exclusion)

// we should consider whether these two global variables should have a mutex to prevent race conditions/dirty reads

func InitEtcdClient(endpointsString string) {
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

func WatchEtcd() {
	ctx := context.Background()

	// initial setup, copy from etcd
	err := etcdGetCampaigns()
	if err != nil {
		log.Fatalln(err)
	}
	err = etcdGetExclusions()
	if err != nil {
		log.Fatalln(err)
	}

	//testPrint()

	exclusionsWatchCh := ETCD.Watch(ctx, "exclusion", clientv3.WithPrefix())
	campaignsWatchCh := ETCD.Watch(ctx, "campaign", clientv3.WithPrefix())

	go handleWatchEvents(exclusionsWatchCh, "exclusion")
	go handleWatchEvents(campaignsWatchCh, "campaign")
}

func handleWatchEvents(watchCh clientv3.WatchChan, key string) {
	for watchResp := range watchCh {
		for _, event := range watchResp.Events {
			switch event.Type {
			case clientv3.EventTypePut:
				//testPrint()
				if key == "campaign" {
					var campaign models.Campaign
					err := json.Unmarshal(event.Kv.Value, &campaign)
					if err != nil {
						log.Println(err)
					}
					CampaignsEtcd[string(event.Kv.Key)] = campaign
				} else if key == "exclusion" {
					var exclusion models.Exclusion
					err := json.Unmarshal(event.Kv.Value, &exclusion)
					if err != nil {
						log.Println(err)
					}
					ExclusionsEtcd[string(event.Kv.Key)] = exclusion
				}
				//testPrint()
			case clientv3.EventTypeDelete:
				//testPrint()
				if key == "campaign" {
					delete(CampaignsEtcd, string(event.Kv.Key))
				} else if key == "exclusion" {
					delete(ExclusionsEtcd, string(event.Kv.Key))
				}
				//testPrint()
			}
		}
	}
}

func etcdGetCampaigns() (err error) {
	response, err := ETCD.Get(context.Background(), "campaign", clientv3.WithPrefix())
	if err != nil {
		log.Println(err)
		return err
	}

	for _, ev := range response.Kvs {
		var campaign models.Campaign
		err := json.Unmarshal(ev.Value, &campaign)
		if err != nil {
			return err
		}

		CampaignsEtcd[string(ev.Key)] = campaign
	}

	return nil
}

func etcdGetExclusions() (err error) {
	response, err := ETCD.Get(context.Background(), "exclusion", clientv3.WithPrefix())
	if err != nil {
		log.Println(err)
		return err
	}

	for _, ev := range response.Kvs {
		var exclusion models.Exclusion
		err := json.Unmarshal(ev.Value, &exclusion)
		if err != nil {
			return err
		}

		ExclusionsEtcd[string(ev.Key)] = exclusion
	}

	return nil
}

func testPrint() {
	fmt.Println("campaigns:")
	for _, campaign := range CampaignsEtcd {
		res, _ := json.Marshal(campaign)
		fmt.Println(string(res))
	}

	fmt.Println("exclusions:")
	for _, exclusion := range ExclusionsEtcd {
		res, _ := json.Marshal(exclusion)
		fmt.Println(string(res))
	}
}
