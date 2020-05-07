rm gen-cpp/*
../../../contribs/ApacheThrift/bin/thrift -r -gen cpp tmarketplace.thrift
../../../contribs/ApacheThrift/bin/thrift -r -gen go  tmarketplace.thrift
rm gen-cpp/*skele*
