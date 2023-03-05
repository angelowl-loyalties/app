package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/rewarder/models"
	"github.com/google/uuid"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	dialTimeout = 2 * time.Second
)

var ETCD *clientv3.Client
var BaseCampaignsEtcd = make(map[string]models.Campaign)
var PromoCampaignsEtcd = make(map[string]models.Campaign)
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

	baseCampaignsWatchCh := ETCD.Watch(ctx, "base_campaign", clientv3.WithPrefix())
	promoCampaignsWatchCh := ETCD.Watch(ctx, "promo_campaign", clientv3.WithPrefix())
	exclusionsWatchCh := ETCD.Watch(ctx, "exclusion", clientv3.WithPrefix())

	go handleWatchEvents(baseCampaignsWatchCh, "base_campaign")
	go handleWatchEvents(promoCampaignsWatchCh, "promo_campaign")
	go handleWatchEvents(exclusionsWatchCh, "exclusion")

	err = etcdAddInitial()
	if err != nil {
		log.Fatalln(err)
	}

	testPrint()
}

func handleWatchEvents(watchCh clientv3.WatchChan, key string) {
	for watchResp := range watchCh {
		for _, event := range watchResp.Events {
			switch event.Type {
			case clientv3.EventTypePut:
				// testPrint()
				if key == "base_campaign" {
					var campaign models.Campaign
					err := json.Unmarshal(event.Kv.Value, &campaign)
					if err != nil {
						log.Println(err)
					}
					BaseCampaignsEtcd[string(event.Kv.Key)] = campaign
				} else if key == "promo_campaign" {
					var campaign models.Campaign
					err := json.Unmarshal(event.Kv.Value, &campaign)
					if err != nil {
						log.Println(err)
					}
					PromoCampaignsEtcd[string(event.Kv.Key)] = campaign
				} else if key == "exclusion" {
					var exclusion models.Exclusion
					err := json.Unmarshal(event.Kv.Value, &exclusion)
					if err != nil {
						log.Println(err)
					}
					ExclusionsEtcd[string(event.Kv.Key)] = exclusion
				}
				testPrint()
			case clientv3.EventTypeDelete:
				//testPrint()
				if key == "base_campaign" {
					delete(BaseCampaignsEtcd, string(event.Kv.Key))
				} else if key == "promo_campaign" {
					delete(PromoCampaignsEtcd, string(event.Kv.Key))
				} else if key == "exclusion" {
					delete(ExclusionsEtcd, string(event.Kv.Key))
				}
				//testPrint()
			}
		}
	}
}

func etcdGetCampaigns() (err error) {
	response, err := ETCD.Get(context.Background(), "base_campaign", clientv3.WithPrefix())
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

		BaseCampaignsEtcd[string(ev.Key)] = campaign
	}

	response, err = ETCD.Get(context.Background(), "promo_campaign", clientv3.WithPrefix())
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

		PromoCampaignsEtcd[string(ev.Key)] = campaign
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

func etcdAddInitial() (err error) {
	var campaign = models.Campaign{
		ID:                 uuid.MustParse("ddb0a58f-6dca-41f3-a3a9-d40961670b5b"),
		Name:               "Spring Fling",
		MinSpend:           75.0,
		Start:              time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC),
		End:                time.Date(2023, 5, 31, 23, 59, 59, 0, time.UTC),
		RewardProgram:      "scis_shopping",
		RewardAmount:       10,
		MCC:                9311,
		Merchant:           "Petco",
		IsBaseReward:       false,
		ForForeignCurrency: false,
	}

	seed_campaign, err := json.Marshal(campaign)
	if err != nil {
		return err
	}

	_, err = ETCD.Put(context.Background(), "campaign_ddb0a58f-6dca-41f3-a3a9-d40961670b5b", string(seed_campaign))
	if err != nil {
		return err
	}

	var exclusion = models.Exclusion{
		ID:        uuid.MustParse("e38adb10-a96a-4b55-aebd-7cdc9b973e7b"),
		MCC:       4125,
		ValidFrom: time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
	}

	seed_exclusion, err := json.Marshal(exclusion)
	if err != nil {
		return err
	}

	_, err = ETCD.Put(context.Background(), "exclusion_e38adb10-a96a-4b55-aebd-7cdc9b973e7b", string(seed_exclusion))
	if err != nil {
		return err
	}
	return nil
}

func testPrint() {
	fmt.Println("base campaigns:")
	for _, campaign := range BaseCampaignsEtcd {
		res, _ := json.Marshal(campaign)
		fmt.Println(string(res))
	}

	fmt.Println("promo campaigns:")
	for _, campaign := range PromoCampaignsEtcd {
		res, _ := json.Marshal(campaign)
		fmt.Println(string(res))
	}

	fmt.Println("exclusions:")
	for _, exclusion := range ExclusionsEtcd {
		res, _ := json.Marshal(exclusion)
		fmt.Println(string(res))
	}
}
