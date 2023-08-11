curl -XPUT "http://localhost:9200/userservice" -H 'Content-Type: application/json' -d '{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  },
  "mappings": {
    "properties": {
      "timestamp": {
        "type": "date"
      },
      "data": {
        "type": "object"
      }
    }
  }
}'

curl -XPUT "http://localhost:9200/loggerservice" -H 'Content-Type: application/json' -d '{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  },
  "mappings": {
    "properties": {
      "timestamp": {
        "type": "date"
      },
      "data": {
        "type": "object"
      }
    }
  }
}'

curl -XPUT "http://localhost:9200/mailerservice" -H 'Content-Type: application/json' -d '{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  },
  "mappings": {
    "properties": {
      "timestamp": {
        "type": "date"
      },
      "data": {
        "type": "object"
      }
    }
  }
}'

curl -XPUT "http://localhost:9200/cronitor" -H 'Content-Type: application/json' -d '{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1
  },
  "mappings": {
    "properties": {
      "timestamp": {
        "type": "date"
      },
      "data": {
        "type": "object"
      }
    }
  }
}'