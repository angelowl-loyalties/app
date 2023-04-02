package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/rewarder/models"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	dialTimeout = 2 * time.Second
)

var ETCD *clientv3.Client

var BaseCampaignsEtcd = make(map[string]models.Campaign)
var baseCampaignMutex sync.RWMutex

var PromoCampaignsEtcd = make(map[string]models.Campaign)
var promoCampaignMutex sync.RWMutex

var ExclusionsEtcd = make(map[string]models.Exclusion)
var exclusionsMutex sync.RWMutex

func InitEtcdClient(endpointsCsv string, username string, password string) {
	endpoints := strings.Split(endpointsCsv, ",")

	var err error
	ETCD, err = clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
		Username:    username,
		Password:    password,
	})

	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Connected to: " + endpointsCsv)
	}
}

func RefreshFromEtcd() {
	err := etcdGetCampaigns()
	if err != nil {
		log.Fatalln(err)
	}
	err = etcdGetExclusions()
	if err != nil {
		log.Fatalln(err)
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
					baseCampaignMutex.Lock()
					BaseCampaignsEtcd[string(event.Kv.Key)] = campaign
					baseCampaignMutex.Unlock()
				} else if key == "promo_campaign" {
					var campaign models.Campaign
					err := json.Unmarshal(event.Kv.Value, &campaign)
					if err != nil {
						log.Println(err)
					}
					promoCampaignMutex.Lock()
					PromoCampaignsEtcd[string(event.Kv.Key)] = campaign
					promoCampaignMutex.Unlock()
				} else if key == "exclusion" {
					var exclusion models.Exclusion
					err := json.Unmarshal(event.Kv.Value, &exclusion)
					if err != nil {
						log.Println(err)
					}
					exclusionsMutex.Lock()
					ExclusionsEtcd[string(event.Kv.Key)] = exclusion
					exclusionsMutex.Unlock()
				}
				// testPrint()
			case clientv3.EventTypeDelete:
				//testPrint()
				if key == "base_campaign" {
					baseCampaignMutex.Lock()
					delete(BaseCampaignsEtcd, string(event.Kv.Key))
					baseCampaignMutex.Unlock()
				} else if key == "promo_campaign" {
					promoCampaignMutex.Lock()
					delete(PromoCampaignsEtcd, string(event.Kv.Key))
					promoCampaignMutex.Unlock()
				} else if key == "exclusion" {
					exclusionsMutex.Lock()
					delete(ExclusionsEtcd, string(event.Kv.Key))
					exclusionsMutex.Unlock()
				}
				//testPrint()
			}
		}
	}
}

// for initialisation of global variable
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
		baseCampaignMutex.Lock()
		BaseCampaignsEtcd[string(ev.Key)] = campaign
		baseCampaignMutex.Unlock()
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
		promoCampaignMutex.Lock()
		PromoCampaignsEtcd[string(ev.Key)] = campaign
		promoCampaignMutex.Unlock()
	}

	return nil
}

// for initialisation of global variable
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
		exclusionsMutex.Lock()
		ExclusionsEtcd[string(ev.Key)] = exclusion
		exclusionsMutex.Unlock()
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
