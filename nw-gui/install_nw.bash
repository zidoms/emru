NODE_URL=http://dl.node-webkit.org/v0.10.5/node-webkit-v0.10.5-linux-x64.tar.gz
NODE_ZIP=node-webkit-v0.10.5-linux-x64.tar.gz
NODE_WEBKIT=bin/node-webkit-v0.10.5-linux-x64

if [ ! -f "$NODE_WEBKIT" ]; then
	cd bin
	if [ ! -f "$NODE_ZIP" ]; then
		wget $NODE_URL
	fi
	tar -xv -f $NODE_ZIP
fi