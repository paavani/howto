### How to install go 

		https://golang.org/doc/install  --click on download go box.It has a version wise  list of different  OS files

		for ubuntu click on    go1.4.2.linux-amd64.tar.gz   (download the folder)

### Extract the archive folder into /usr/lib

		tar -C /usr/lib -xzf go$VERSION.$OS-$ARCH.tar.gz

			$VERSION  -->which version you want to download(suppose 1.4.2)

			$OS  --> linux

			$ARCH -->for 32 bit linux system --> 386
				          64 bit linux system --> amd64

		 tar -C /usr/lib -xzf go1.4.2.linux-amd64.tar.gz

		Note: run the command as in  root otherwise use sudo

### Create a folder name as goplay in HOME
		
		$mkdir -p goplay/bin


### Set the path in .bashrc

	we can add path in two ways.
	
	1.Directly open the bashrc file and enter as follows

		//go source
		export GOROOT=/usr/lib/go
		export PATH=$PATH:$GOROOT/bin


		//go workspace
		export GOPATH=$HOME/goplay
		export PATH=$PATH:$GOPATH/bin

	2.using echo

		echo export GOPATH=$HOME/goplay>> ~/.bashrc
		echo export GOROOT=/usr/lib/go>> ~/.bashrc
		echo export PATH=$PATH:$GOROOT/bin:$GOPATH/bin >> ~/.bashrc
	


### To check golang software  is installed or not

		$dpkg -s golang

### To check the version

		$go version

### To check only go  environment variables path settings

		$go env 

### TO check whole env path
	
		$env|grep go

##How to install golang doc

		$sudo apt-get install golang-doc

		Note:After installation Run  godoc -http= :6060 and open localhost:6060 in browser
