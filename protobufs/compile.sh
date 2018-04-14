# Remove the dead weeds, if they exist
rm -rf ./dist > /dev/null 2>&1

# Make our paths
mkdir dist > /dev/null 2>&1
mkdir dist/conversation > /dev/null 2>&1

# Compile the files
protoc --go_out=plugins=grpc:./dist/conversation \
       ./conversation.proto