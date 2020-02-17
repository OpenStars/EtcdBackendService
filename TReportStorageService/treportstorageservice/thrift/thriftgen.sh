rm -rf gen*
thrift -r -gen cpp treportstorage.thrift
thrift -r -gen go  treportstorage.thrift

