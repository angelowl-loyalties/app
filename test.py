"""
Before you run this wait for all 9 containers to be up, rewarder and informer will restart until cassandra is ready

This script simulates the following
- Add in 2 base exclusions and campaigns to etcd
- Trigger the lambda to add transactions to queue
- Sleep and print output of informer (should have some stuff)

Endpoints:
- ingestor: localhost:9000
- excluder: localhost:8081
- campaigner: localhost:8082
- informer: localhost:8083
"""

import time
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

        res_campaign = requests.post("http://localhost:8082/campaign", data=json.dumps(seed_campaigns[i]))
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
    res_informer = requests.get("http://localhost:8083/reward")
    print("Number of records on informer rn:")
    print(len(res_informer.json()))