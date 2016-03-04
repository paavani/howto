## Linux basic commands

### To get help for each command

	$man commandname
	 
	 it displays help pages of given command
     
     note:press 'q' to quit from manual pages 

### Standard directory

	(few point about folders)
	/ 
	root directory, starting point of the directory tree.

	/home
		(private) directories of users.

	/dev
		device files that represent hardware components

	/etc
		important files for system configuration.

	/etc/init.d
		boot scripts

	/usr/bin
		generally accessible programs.

	/bin
		programs needed early in the boot process.

	/usr/sbin
		programs reserved for the system administrator.

	/sbin
		programs reserved for the system administrator and needed for booting.

	/usr/include
		header files for the C compiler

	/usr/include/g++
		header files for the C++ compiler.

	/usr/share/doc
		various documentation files.

	/usr/share/man
		system manual pages (man pages).

	/usr/src
		source code of system software

	/usr/src/linux
		kernel source code.

	/tmp, /var/tmp
		temporary files.

	/usr
		all application programs

	/var
		configuration files (e.g., those linked from /usr)

	/var/log
		system log files

	/var/adm
		system administration data

	/lib
		shared libraries (for dynamically linked programs)

	/proc
		process file system.

	/usr/local
		local, distribution-independent extensions.

	/opt
		optional software, larger add-on program packages (such as KDE, GNOME, Netscape).

##TO Know environment variable  path in linux ( case sensitive)

	$env --it display all env variables

	$env|grep GO	--it display GO path
	$env|grep go  --case sensitive nothing to display

	$env|grep postgres 


### Display current username
    
    $logname

### Present working directory path
    
    $pwd

### To display system date and time
    
    $date

### To clear screen
    
    clear or ctrl+l

### using cal
    
    1.to display current month calender
    
      $cal
    
    2.to display particular year
      
      $cal 2010
    
    3.to display specified month and year details
      
      $cal 09 2015

### To display operating system name
        $uname
  	
    to display kernel version 
        $uname -r

### To display servername
     
     $hostname
   
     ####To know Ip address
      
      $hostname -i

### who 
  	it displays the list of users who are connected to the server

### Open a terminal in Ubuntu
	
	ctrl+alt+t 
	
	or 
	
	ctrl+shift+n
	
#### Open a another  new terminal in same window	
	
	ctrl+shift+t

### Few terminal commands
	
	select what do you want copy and use below cmd
	Copy  -->  ctrl+shift+c
	
	Paste -->  ctrl+shift+v
	
	Move up &down -->ctrl+shift+up arrow or down arrow

	
	

### To know current terminal name
    
    $tty
    
## Working with files
 
### Creating a new file 
   
   I. we can create a new file in different ways using touch,echo,cat
     
     1.$touch file1 file2 file3
        
        we can create a multiple files at a time with zero byte size.It is used to change file time stamp
     
     2.$echo "file content data" > text.txt
     
     3.using cat 
        
        $cat> filename  press enter and write the content of  the  file  and press ctrl+d
   
   II To open a file 
      	
      	$cat filename
      
      To open multiple files
      	
      	$cat filename1 filename2 filename3 
      	
      Note:joins multiple files vertically without spaces
      	
   
   III To append data to the file
   	
   	$cat>>filename 
   			
    
   IV Deleting files
   		
   	1.delete a single file
   	
   	$rm filename
   			
   	2.delete a file with confirmation
   		
	 $rm -i filename
   		
   	3.deleting multiple files
   	
 	 $rm filename1 filename2 filename3 
   			
## Working with directories(dir)
   	
### To Create a new directory
   	  $mkdir dirname
   	
   	Creating multiple directories
   	  $mkdir dirname1 dirname2 dirname3
   	 
   	Create a directory  with full pathname at a time  
         $mkdir -p mydir/grandparent/parent/child
          -p  --parents
          if parent dir existing , no error otherwise it create a dir
    
 ### To change directory

	$cd dirname

### To come out from current working directory

	$cd ..

### changes to root directory

	$cd/

## Creating directories under sub directories	

	$mkdir abc  xyz   abc/a1  abc/a2  abc/a1/a11  abc/a1/a12   xyz/x1  xyz/x2

##Removing directories

	1.To remove a directory but directory must be empty

	$rmdir  mydir

	2.To remove entire directory structure recursively

	$rm -r abc

	3.To remove entire directory structure with confirmation

	$rm  -ri abc

	4.To remove recursively  entire directory structure with forcibly

	$rm -rf abc


## Copying files

	1.$cp: To copy a file

	 $cp  sourcefile   targetfile

	note : source file must be existing file and target file may be a new file or existing file

	2.Copying files from one directory to another directory

  	eg: copy a11 directory emp file into x1 directory
	
	$cp abc/a1/a11/emp     xyz/x1

	3. Copying multiples into one directory

	$cp file1 file2 file3 .... filen     directory

	4.Copying  a directory

	$cp  -R  sourcedir  targetdir

## To move a file/directory or  rename a file/directory

	$mv file1 file2  -- file1 is renamed by file2
	
	$mv dir1 dir2

## creating hidden files

	To hide file/dir  start a file name or directory name with  “.” character

	eg: $mkdir .dirname 
	
	 $cat>.filename

##Viewing list of files or ls cmds

	
	1.$ls 

	It list current directory all files and subdirectories in the ascending order based on ASCII  values


	2.$ls -a

	It list all files along  with hidden filess(.filename)

	3.$ls  -r 
	
	It list all files in reverse order i.e descending order

	4.$ls -R 

	It list all files Recursively

	5.$ls -t 
	
	It list all files based on date & time of creation

	6.$ls -l 
 
	It list all files with  detailed information as follows

	eg:
	File	File 		NO.of 	Owner	Group		Size in Date	Time	Filename
	Type	permissions	links	name	name		Bytes	
	-----------------------------------------------------------------------------------------
	d	rwxr-xr-x	2	pavan 	domain^users 	4096 	Jun 29 	10:41 	newproject
	-	rw-r--r--	1	pavan 	domain^users  	584 	Jul  6 	19:15 	mytodo.md


	Note: d ->directory, - ->file
	
	r-read,w-write ,x-execution  ( r=4 w=2 x=1)

	Default permissions for directory :rwxr-xr-x
				file	  :rw-r--r--
    
	Owner or user | group |other
	rwx r-x r-x(755)--owner has read,write,execute permissions ,group and other has  read and execute permissions

	we can change the permissions using chmod 744 <filename>
	

##Wild card  characters

	*  :matches 0 or more characters
	?  :matches any single character
	[] :matches any single character in the given list
	[-]:matches any single character in the given range

	eg:$ls d*	$ls a???	$ls [vsd]*	$ls[abcdef]*  or $ls[a-f]*   $ls *t	
	
	$cp *.js ~/javascript		$cp abc/a1/*   .
	
	$rm *		$rm -i * 	$rm [bkt]*

##head
	It displays the first "n" lines from a given file
	
	Syntax :head -n filename
	
		$head filename     --it display first 10 lines
		$head -20 filename -- it display first 20 lines	

	Note:BY default it displays first 10 lines from the file

##tail
	It displays the last "n" lines from a given file
	
		$tail -20 filename
	
##how to search the particular filename
	 
	  $locate <searchfile or word name>
	    
						or
	 
     $find|grep <searchfile or wordname>
	 
	    To ignore the case sensitive
	 
	    $locate -i pivot
	    $find|grep -i pivot
	 
	   It will give both lower case and upper case available files

#How to see memory usage of linux syster

	$top

	$free -m //it will show in mbs   $free -g OR free --it will show in kb

	$vmstat -s


