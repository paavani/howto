## git installation

   git for all platforms
  
   http://git-scm.com
  
## installation on ubuntu

   $ sudo apt-get install git
   
## Create a local repository
  
  $git init

## After instalation of git set username and email only once

  $ git config --global user.name "your_name"
  
  $ git config --global user.email "user@example.com"
  
## Checking your settings

  $git config --list
  
##Status

  $git status  
  
  List the files you have changed and those you still need to add or commit
  
  
##git add
	
	$git add <list-of-files>

	to add list

###Different ways to add 

	$git add --all         -->Add all files

	$git add *.txt         -->Add all txt files in current directory

	$git add mydir/*.txt   -->Add all txt files in mydir directory

	$git add mydir/		  -->Add all files in mydir directory

	$git add "*.txt"		  -->Add all txt files in the whole project

	$git add .				  -->Add all(new,modified and deleted)

## Commit

	$ git commit -m "<content>"

####Directly running commit instead of git add for modified files 

	$ git commit -a

	Note: It is only for modified files not for new files
	

##push

	$git push -u origin master

##clone from exit  remote repository to local 

	$git clone <repository-url>

###TO know Project history

	go to the project folder and give cmd as

	$git log

	It displays the history of what you change

	$git log -p

	It displays the history of  complete diffs at each step



##branch

	 single git repository by default branch is master.we can create and maintain multiple branches.

###To create a new branch

	$git branch <branchname>
	
	$git branch subbranch

### To check the list of existing branches

	$git branch

	eg: 
		subbranch
		* master

		note:* denotes you are currently on which branch

### To change from one branch to another branch

	$git checkout subbranch

	
