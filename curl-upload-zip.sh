#!/bin/bash


# curl -i -v  \
#      -F metadata='{"parent_folders":"./", "filename":"main.zip", "filetype":"video/mp4", "channel":"imagechannel", "encrypt_method":"drm", "content_id":"1sdfjapwef", "preview_start":"00:00:00", "preview_end":"00:00:30", "transcode_profile": "dance"}' \
#      -F file=@./main.zip \
#      'http://localhost:6000/file' 

curl -i -v  -F file=@./main.zip   'http://localhost:6000/file' 
