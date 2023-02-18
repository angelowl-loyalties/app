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
        "reward_program": "Points",
        "reward_amount":  500,
        "mcc":           4495,
        "merchant":      "Best Buy",
    },
    {
        "name":           "Winter Warmup",
        "min_spend":      100.0,
        "start_date": "2023-12-01T00:00:00Z",
        "end_date": "2024-02-28T23:59:59Z",
        "reward_program": "Cashback",
        "reward_amount":  25,
        "mcc":            8371,
        "merchant":       "Home Depot",
    },
    # {
    #     "name":          "Spring Fling",
    #     "min_spend":     75.0,
    #     "start_date":   datetime.datetime(2023, 3, 1, 0, 0, 0, 0, tzinfo=datetime.timezone.utc),
    #     "end_date":    datetime.datetime(2023, 5, 31, 23, 59, 59, 0, tzinfo=datetime.timezone.utc),
    #     "reward_program": "Discount",
    #     "reward_amount": 10,
    #     "mcc":           9311,
    #     "merchant":      "Petco",
    # },
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
    # {
    #     "mcc":        4125,
    #     "valid_from": datetime.datetime(2023, 3, 1, 0, 0, 0, 0, tzinfo=datetime.timezone.utc),
    # },
    # {
    #     "mcc":        9311,
    #     "valid_from": datetime.datetime(2023, 4, 1, 0, 0, 0, 0, tzinfo=datetime.timezone.utc), # test to see the most commonly seen MCC marked excluded from next month
    # },
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
    

    res = requests.post("http://localhost:9000/2015-03-31/functions/function/invocations", data="{}")
    print(f"Lambda Status code: {res.status_code}")

if __name__ == "__main__":
    main()
    time.sleep(30)
    res_informer = requests.get("http://localhost:8083/reward")
    print("informer all rewards rn:")
    print(res_informer.json())