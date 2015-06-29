## Slowly changing dimension(scd)

  	 The SCD concept is basically about how the data modifications are absorbed and maintained in a Dimension Table.
  	 
  	 Mostly we use 3 types of scd's
  	 
  	 
  	simple example:create itemmaster and rate history
  	
 	Itemmaster data:
 	
  	"ITEMID", "ITEMNAME",  "QTY",  "RATE"
  	-------   ---------  ---------  ------
  	"30",	 "mango",	"1kg",   "40"
  	"31",	 "apple",	"1kg",   "120"
  	
  	note: If you update the mango rate as 60 .we can see how the 3 types of scd display
  	
  	 
### Type 1 (Overwriting the old value)

  	 	It does not maintain history and the table is always updated with  recent data whenever there is a change,because of update, we lose the previous values
 	
 	"ITEMID", "ITEMNAME","QTY","RATE"
 	-------   ---------  ---------  -----
  	"30",	"mango",     "1kg", 	"60"
  	"31",	"apple",     "1kg",	"120"
 	
 	
###Type 2 (Creating a new additional record)

			It maintain the change history in the dimension table and create a new row when there are changes The table is vertically growing.

 			When the value of a chosen attribute changes, the current record is closed. A new record is created with the changed data values and this new record becomes the current record. Each record contains the effective time and expiration time to identify the time period between which  record is active.
 			
 	ratehistory:
 	
 	"ITEMID", "RATE", "DATEFROM", "DATEEND"
 	--------  ------  ---------- ---------
 	30,	  40, 	 23/06/2015,  25/06/2015
 	31,	  120,	 22/06/2015,	NULL
 	30,	  60,	 25/06/2015,	NULL
 	
 			

###Type 3 ( Adding a new column)

	 		It maintain the history in a separate column instead of separate rows(current value and one more previous value).
	 		
	 	"ITEMID", "previousRate",  "currentRate"
	 	--------  ---------------  ------------
  		  30,      	40,		  	60
  		

