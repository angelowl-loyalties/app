package internal

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/campaignex/models"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	dialTimeout = 5 * time.Second
)

var ETCD *clientv3.Client

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

// Campaign etcd functions

//func etcdGetCampaign(id string) (err error) {
//	response, err := ETCD.Get(context.Background(), "campaign_"+id)
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//
//	var campaign models.Campaign
//	for _, ev := range response.Kvs {
//		err := json.Unmarshal(ev.Value, &campaign)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

func etcdPutCampaign(campaign *models.Campaign) (err error) {
	jsonCampaign, err := json.Marshal(campaign)
	if err != nil {
		log.Println(err)
		return err
	}

	if campaign.IsBaseReward {
		_, err = ETCD.Put(context.Background(), "base_campaign_"+campaign.ID.String(), string(jsonCampaign))
	} else {
		_, err = ETCD.Put(context.Background(), "promo_campaign_"+campaign.ID.String(), string(jsonCampaign))
	}
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func etcdDeleteCampaign(isBase bool, id string) (deleted int, err error) {
	var response *clientv3.DeleteResponse
	if isBase {
		response, err = ETCD.Delete(context.Background(), "base_campaign_"+id)
	} else {
		response, err = ETCD.Delete(context.Background(), "promo_campaign_"+id)
	}
	if err != nil {
		return 0, err
	}

	return int(response.Deleted), nil
}

// Exclusion etcd functions

//func etcdGetExclusion(id string) (err error) {
//    response, err := ETCD.Get(context.Background(), "exclusion_"+id)
//    if err != nil {
//        log.Println(err)
//        return err
//    }
//
//    var exclusion models.Exclusion
//    for _, ev := range response.Kvs {
//        err := json.Unmarshal(ev.Value, &exclusion)
//        if err != nil {
//            return err
//        }
//    }
//
//    return nil
//}

func etcdPutExclusion(exclusion *models.Exclusion) (err error) {
	jsonExclusion, err := json.Marshal(exclusion)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = ETCD.Put(context.Background(), "exclusion_"+exclusion.ID.String(), string(jsonExclusion))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func etcdDeleteExclusion(id string) (deleted int, err error) {
	response, err := ETCD.Delete(context.Background(), "exclusion_"+id)
	if err != nil {
		return 0, err
	}

	return int(response.Deleted), nil
}
