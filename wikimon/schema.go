package wikimon

var AvroSchema = `
{
    "type":"record",
    "name":"wikimon",
    "fields": [
        {"name":"action","type":"string"},
        {"name":"change_size","type":"long"},
        {"name":"flags","type":["string", "null"]},
        {"name":"geo_ip",
            "type":{
                "type":"record",
                "name":"geo_ip",
                "fields": [
                    {"name":"city","type":"string"},
                    {"name":"country_name","type":"string"},
                    {"name":"latitude","type":"float"},
                    {"name":"longitude","type":"float"},
                    {"name":"region_name","type":"string"}
                ]
            }
        },
        {"name":"hashtags","type": { 
            "type": "array",
            "items": "string"
            }
        },
        {"name":"is_anon", "type":"boolean"},
        {"name":"is_bot", "type":"boolean"},
        {"name":"is_minor", "type":"boolean"},
        {"name":"is_new", "type":"boolean"},
        {"name":"is_unpatrolled", "type":"boolean"},
        {"name":"mentions","type": { 
            "type": "array",
            "items": "string"
            }
        },
        {"name":"ns","type":"string"},
        {"name":"page_title","type":"string"},
        {"name":"parent_rev_id","type":"string"},
        {"name":"rev_id","type":"string"},
        {"name":"summary","type":"string"},
        {"name":"url","type":"string"},
        {"name":"user","type":"string"}
    ]
}
`
