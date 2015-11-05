##npm (node package manager)

	npm makes easy for javascript developers to share and reuse the code,it is easy to update the code that you are sharing

##Node installation 
		
	I. first download the nodejs latest version from the nodejs.org(node-v0.12.6.tar.gz)

###Move tar folder to home dir	   

	II.move the node.tar file from Downloads folder to current directory(home directory)
	  
	 $mv Downloads/node-v0.12.6.tar.gz .   (last . Represent the home directory or current directory) 

###How to run tar folder
	
	III.run the below cmd on terminal.It create the folder of node with version(eg:node-v0.12.6)

	 $tar -xzvf node-v0.12.6.tar.gz

			x-extract

			z-compressed
			
			v-verbose(what happened in the background)
			
			f-filename

###How to configure and install node

	IV.cd nodefolder(node-v0.12.6) and enter the follwing cmds
		
		$cd node-v0.12.6
		
		$./configure && make && sudo make install

			./configure --this cmd configures the all the cmds

			make --compile and build the tool
			
			$sudo make install  --load into appropriate folder
														 
							OR			

####If you'd like to install from the repository

		$git clone https://github.com/joyent/node.git

		$ cd node
	
		$ git checkout v0.6.18 #Try checking nodejs.org for what the stable version is

		$./configure && make && sudo make install


###If some reasons npm is not installed use npm ,run below cmd

		$sudo apt-get install npm

###To update the latest version  of npm 

		$sudo npm install npm -g 

