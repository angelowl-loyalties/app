"""
Before you run this wait for all containers to be up, rewarder and informer need to be restarted when cassandra is up.

When cassandra is up, its container logs will show that a default superuser has been created. Restart rewarder and informer when this is seen.

This script simulates the following
- Add in 2 base exclusions and campaigns to etcd via campaignex
- Trigger the lambda to add transactions to production Kafka (needs to be accessed via a VPN)

Endpoints:
- ingestor: localhost:9000
- campaignex: localhost:8081
- informer: localhost:8082
"""

import requests
import json

seed_campaigns = [
    {
        "name":          "Summer Sale",
        "min_spend":      50.0,
        "start_date": "2023-01-01T00:00:00Z",
        "end_date": "2023-08-31T23:59:59Z",
        "reward_program": "scis_platinummiles",
        "reward_amount":  500,
        "mcc":           4495,
        "merchant":      "Best Buy",
    },
    {
        "name":           "Winter Warmup",
        "min_spend":      100.0,
        "start_date": "2023-12-01T00:00:00Z",
        "end_date": "2024-02-28T23:59:59Z",
        "reward_program": "scis_freedom",
        "reward_amount":  25,
        "mcc":            8371,
        "merchant":       "Home Depot",
    },
]

seed_exclusions = [
    {
        "mcc":       5774,
        "valid_from": "2023-01-01T00:00:00Z",
    },
    {
        "mcc":        7080,
        "valid_from": "2023-03-01T00:00:00Z",
    },
]

def main():
    # Add in 2 Campaigns and 2 Exclusions
    print("Adding 2 Exclusions and Campaigns")
    for i in range(2):
        res_exclusion = requests.post("http://localhost:8081/exclusion", data=json.dumps(seed_exclusions[i]))
        if res_exclusion.status_code != 201:
            print("Add seed exclusion failed")
            return

        res_campaign = requests.post("http://localhost:8081/campaign", data=json.dumps(seed_campaigns[i]))
        if res_campaign.status_code != 201:
            print("Add seed campaign failed")
            return

        print(res_exclusion.json())
        print(res_campaign.json())
    
    # This rquest blocks for 2 minutes
    res = requests.post("http://localhost:9000/2015-03-31/functions/function/invocations", data="{}")
    print(f"Lambda Status code: {res.status_code}")

if __name__ == "__main__":
    main()