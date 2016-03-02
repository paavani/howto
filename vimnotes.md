## Vim editor

It is classified into 3 types

	1.command mode
	2.input or insert mode
	3.Ex command mode
	
It is used for create new files or to modify existing files
	eg: vim myfile 
	If it is already exist open a file ,otherwise create a file


###Use the below commands to shift from command mode to insert mode

	 A -->places cursor at the end of the current line
	 a -->places cursor at the right side of the cursor line
	
	 I -->places cursor at the beginning of the current line
	 i -->places cursor at left side of the cursor line
	
	 O -->inserts new line of the cursor
	 o -->inserts new line below of the cursor

	"ESC" is the key to shift from insert mode to command mode

	":" is the command to shift from to EX command mode from command mode
	
## Command mode commands:
     
         k
         ^ 
         |
    h<--   -->l
         |
         j

	
	H --Beginning of the current page
	M --Middle of the current page 
	L --End of the current page

	ctrl+f --forward on the page (pgdn)
	ctrl+b --backward one (pgup)




##EX command mode commands:

  :w  --save without quit
  
  :wq --save and quit
  
  :w filename -- save with given name
  
  :q! --quit without save
  
  :n  --places cursor at nth line
  
  :$  --places cursor at last line in the file

  :set nu   --set line numbers
  
  :set nonu --to remove line numbers
  
  :/string/ --top to bottom search (n -> next occurence)
  
  :?string? --bottom to top search (N -> previous occurence)
  :undo  --undo the changes
  :redo	--redo the changes
###delete line 

	:nd  --delete nth line
	
	:$d  --delete last line
	
	:4,7d --delete 4th line to 7th line 
    

###Open multiple files in same terminal

	$vim -p filename1 filename 2 filename3
	
	moving between  the files in forward or backward
	
	 ctrl+Tab  or shift +Tab

