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

  List the files you have changed and those you still need to add or commit.
  
  Note:Better to use this cmd very often,Sometimes things change and you don't notice it


##git add

	Take the snapshot of the content of the current file and saved in temporary staging area	
		

	$git add <list-of-files>

	to add list

###Different ways to add

	$git add --all        --> Add all files

	$git add * .txt       --> Add all txt files in current directory

	$git add mydir/* .txt --> Add all txt files in mydir directory

	$git add mydir/	    --> Add all files in mydir directory

	$git add '* .txt'	    --> Add all txt files in the whole project

	$git add .			    --> Add all(new,modified and deleted)

	$git add -A .         -->Add all 

## Commit

	using this cmd you can permanently stored content of index in the repository

	$ git commit -m "<content>"

	
	Note:Before using commit cmd must use git status cmd 

	-a it auto removes deleted files with commit

	$git commit -am "delete stuff"

####Directly running commit instead of git add for modified files

	$ git commit -a

	Note: It is only for modified files not for new files


##push

	push cmd tells git where to put  our commits when we are ready

	$git push -u origin master

	origin: name of our remote is origin

	master: default local branch name

	-u  :tells  git to remember parameters,so next time we simply run git push 

##clone from exit  remote repository to local

	$git clone <repository-url>

##TO know Project history

	go to the project folder and give cmd as

	$git log

	It displays the history of what you change

	$git log -p

	It displays the history of  complete diffs at each step


#Managing branches and Merge the changes

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

		Note: `*` denotes you are currently on which branch


###To list  all branches

	$git branch -a

### To change from one branch to another branch
	
	$git checkout subbranch

###To checkout and create branch at the same time 

	$git checkout -b mybranch

##Merge the changes
	
	switch to the subbranch,edit a file and commit the change and switch back to master branch 
		
		edit file
		$git commit -a
		git checkout master

	To merge the changes made in subbranch into master, run

	$git  merge subbranch
	
### To delete the branch

	$git branch -d subbranch

	Note:It won't work,when you delete something ,that hasn't been merged

###Force delete

	$git branch -D subbranch  ///here -D =  -d -f

			  or
	$git branch -d -f subbranch		  

### To Display the folder files in vertical manner

   $git ls-files

###To remove file/files from staging area

	$git reset <filename>
	
###To unstage the files

	$git rm --cached <filename>

###To remove folders

	$git rm -r <foldername>
