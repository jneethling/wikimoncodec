package wikimon

var AvroSchema = `
{
    "type":"record",
    "name":"wikimon",
    "fields": [
        {"name":"action","type":"string","default":""},
        {"name":"change_size","type":"long","default":0},
        {"name":"geo_ip",
            "type":{
                "type":"record",
                "name":"geo_ip",
                "fields": [
                    {"name":"city","type":"string","default":""},
                    {"name":"country_name","type":"string","default":""},
                    {"name":"latitude","type":"float","default":0},
                    {"name":"longitude","type":"float","default":0},
                    {"name":"region_name","type":"string","default":""}
                ]
            },"default":{}
        },
        {"name":"hashtags","type":{"type":"array","items":"string"},"default":[]},
        {"name":"is_anon", "type":"boolean","default":false},
        {"name":"is_bot", "type":"boolean","default":false},
        {"name":"is_minor", "type":"boolean","default":false},
        {"name":"is_new", "type":"boolean","default":false},
        {"name":"is_unpatrolled", "type":"boolean","default":false},
        {"name":"mentions","type":{"type":"array","items":"string"},"default":[]},
        {"name":"ns","type":"string","default":""},
        {"name":"page_title","type":"string","default":""},
        {"name":"parent_rev_id","type":"string","default":""},
        {"name":"rev_id","type":"string","default":""},
        {"name":"summary","type":"string","default":""},
        {"name":"url","type":"string","default":""},
        {"name":"user","type":"string","default":""}
    ]
}
`
