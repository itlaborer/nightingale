package cconf

const EVENT_EXAMPLE = `
{
    "id": 1000000,
    "cate": "prometheus",
    "datasource_id": 1,
    "group_id": 1,
    "group_name": "Default Busi Group",
    "hash": "2cb966f9ba1cdc7af94c3796e855955a",
    "rule_id": 23,
    "rule_name": "测试告警",
    "rule_note": "测试告警",
    "rule_prod": "metric",
    "rule_algo": "",
    "severity": 0,
    "prom_for_duration": 60,
    "prom_ql": "",
    "prom_eval_interval": 30,
    "callbacks": [],
    "runbook_url": "",
    "notify_recovered": 1,
    "notify_channels": [],
    "notify_groups": [],
    "notify_groups_obj": null,
    "target_ident": "",
    "target_note": "",
    "trigger_time": 1677229517,
    "trigger_value": "2273533952",
    "tags": [
        "__name__=disk_free",
        "dc=qcloud-dev",
        "device=vda1",
        "fstype=ext4",
        "ident=tt-fc-dev00.nj"
    ],
    "is_recovered": false,
    "notify_users_obj": null,
    "last_eval_time": 1677229517,
    "last_sent_time": 1677229517,
    "notify_cur_number": 1,
    "first_trigger_time": 1677229517
}
`
