curl -X PUT "http://localhost:9200/video_info_index" -H "Content-Type: application/json" --data-binary "@../deploy/es/video_info_mapping.json"
