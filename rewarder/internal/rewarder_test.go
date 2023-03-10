package internal

import (
	"reflect"
	"testing"
	"time"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/rewarder/models"
	"github.com/google/uuid"
)

func Test_isExcluded(t *testing.T) {
	etcdAddSeedData()
	
	type args struct {
		transactionDate time.Time
		mcc             int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"exclusion_none", args{time.Now(), 4}, false},
		{"exclusion_seed_started_today", args{time.Now(), 4125}, true},
		{"exclusion_seed_starts_tomorrow", args{time.Now(), 5001}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isExcluded(tt.args.transactionDate, tt.args.mcc); got != tt.want {
				t.Errorf("isExcluded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMatchingCampaigns(t *testing.T) {
	type args struct {
		transaction models.Transaction
	}
	tests := []struct {
		name         string
		args         args
		wantCampaign [][]models.Campaign
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCampaign := getMatchingCampaigns(tt.args.transaction); !reflect.DeepEqual(gotCampaign, tt.wantCampaign) {
				t.Errorf("getMatchingCampaigns() = %v, want %v", gotCampaign, tt.wantCampaign)
			}
		})
	}
}

func etcdAddSeedData() {
	ExclusionsEtcd["001"] = models.Exclusion{
		uuid.MustParse("e38adb10-a96a-4b55-aebd-7cdc9b973e7b"),
		4125,
		time.Now(),
	}
	ExclusionsEtcd["002"] = models.Exclusion{
		uuid.MustParse("e38adb10-a96a-4b55-aebd-7cdc9b973e7b"),
		5001,
		time.Now().AddDate(0, 0, 1),
	}
}