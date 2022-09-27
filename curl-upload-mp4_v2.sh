#!/bin/bash

# https://stackoverflow.com/questions/29231926/curl-how-to-post-multipart-form-data-data-and-how-to-read-multipart-form-data-in


# curl -i -H "Authorization":"eyJhbGciOiJIUzI1NiIsImV4cCI6MTQyNjcwNTY4NiwiaWF0IjoxNDI2NzAyMDg2fQ.eyJpZCI6MTc3fQ.yBwLFez2RnxTojLniL8YLItWVvBb90HF_yfhcsyg3lY" \
#     -F user_data='{"user data": {"preferred_city":"Newyork","within_radious":"5"}}' \
#     -F uploaded_documents=@mydocument.pdf \
#     http://127.0.0.1:5000/api/city


#echo '{"parent_folders":"/tmp/up-video/", "filename":"cc-video_001_Canada-Landscape.mp4", "filetype":"video/mp4", "channel":"imagechannel", "encrypt_method":"drm", "content_id":"1sdfjapwef", "preview_start":"00:00:00", "preview_end":"00:00:30", "transcode_profile": "dance"}' > metadata.json


# https://stackoverflow.com/questions/51484022/how-to-upload-a-file-via-curl-binary-data-and-webkit-form-boundary
#  there's pretty much never any reason to craft the multipart/form-data request manually, stop doing that. just use the -F parameter, in your case

# -F upload_file=@abcd.png -F task=bov -F xst=tmvattachaddfiledrop -F windowid=56 -F xcpwinid=57 -F actionmoniker='de..metamodel.LinkRelationAction-1337#' -F xsrftoken=myxsrftoken , and curl will create the multipart/form-data stuff for you.

# also remove the Content-Type: multipart/form-data; header, curl will add it, and set the appropriate boundary, wheras you will set the wrong boundary (and if you set the header manually, curl won't override your custom header, but if you don't set it, curl will create it for you.)


# -i : 

curl -i -v  \
     -F metadata='{"parent_folders":"/tmp/up-video/", "filename":"cc-video_001_Canada-Landscape.mp4", "filetype":"video/mp4", "channel":"imagechannel", "encrypt_method":"drm", "content_id":"1sdfjapwef", "preview_start":"00:00:00", "preview_end":"00:00:30", "transcode_profile": "dance"}' \
     -F file=@/tmp/up-video/cc-video_001_Canada-Landscape.mp4 \
     'http://localhost:6000/file' 


curl -i -v  \
     -F metadata='{"parent_folders":"/tmp/up-video/", "filename":"cc-video_001_Canada-Landscape.mp4", "filetype":"video/mp4", "channel":"imagechannel", "encrypt_method":"drm", "content_id":"1sdfjapwef", "preview_start":"00:00:00", "preview_end":"00:00:30", "transcode_profile": "dance"}' \
     -F file=@/tmp/up-video/cc-video_001_Canada-Landscape.mp4 \
     'http://localhost:6000/remote' 

