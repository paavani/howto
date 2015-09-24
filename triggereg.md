##simple example on trigger

### create a itemmaster table
  
```sql

  create table itemmaster (
	itemid numeric primary key,
	itemname text not null,
	qty numeric default 1,
	rate numeric default 0,
	createddate date default now(),
	createdby text not null,
     );
     
```
###create a table invoice 

```sql

create table invoice (
	id numeric primary key,
	invdate timestamp  default current_timestamp,
	customer text
);

```

###create a  table rate_history

```sql
create table rate_history( 
	itemid numeric,
	rate numeric,
	datefrom timestamp,
	dateend timestamp
);

```

###create a function

```sql
	
	CREATE OR REPLACE FUNCTION trg_rate_history()
	RETURNS trigger AS $$
  BEGIN
	IF TG_OP= 'INSERT' THEN
	INSERT INTO rate_history values (NEW.itemid,NEW.rate,now(),null);
	ELSIF TG_OP = 'UPDATE' THEN
	UPDATE rate_history set dateend = now() where itemid = NEW.itemid and dateend is null;
	INSERT INTO rate_history values (NEW.itemid,NEW.rate,now(),null);
	END IF;
	
    	RETURN NEW;
  END;
$$ LANGUAGE plpgsql;

Note : PL/pgSQL is a loadable procedural language for the PostgreSQL database system.It is used to create a functions and trigger procedures
```



###Create a Trigger 

```sql

   CREATE TRIGGER rate_history
   AFTER INSERT OR UPDATE ON
   itemmaster
   FOR EACH ROW EXECUTE PROCEDURE trg_rate_history();
   
```
### If you change the logic in function or trigger  first  you drop the trigger and again create function and  trigger

```sql

	DROP trigger rate_history on itemmaster
```

###Insert the values into table

```sql

insert into itemmaster(itemid,itemname,rate,createdby) values (1,'chacos',30,'siv');
insert into itemmaster(itemid,itemname,rate,createdby) values (1,'grapes',80,'siv');
insert into itemmaster(itemid,itemname,rate,createdby) values (1,'mango',55,'anil');

insert into invoice (id,customer) values (1,'Raja');

```

### itemmaster table data

 	itemid | itemname |   qty    | rate | createddate| createdby 
 
	--------+----------+----------+------+------------+----------

	1	| chacos   |        1 |   30 | 2015-06-22 | siv               
	2 	| grapes   |        1 |   80 | 2015-06-22 | siv	                 
	3 	| mango    |        1 |   55 | 2015-06-22 | anil          
      
 
### invoice table data:
 
 
	 id |          invdate           | customer
	 
	----+----------------------------+----------
	
  	1   | 2015-06-25 16:51:25.189744 | Raja
  
 
### rate_history table data:
  
  
	 itemid|  rate|          datefrom       |   dateend      
  
   	-----+------+---------------------------+-------------
   
	 1  |   30 | 2015-06-22 08:09:48.85778  |
	 2  |   80 | 2015-06-22 08:01:43.002791 | 
	 3  |   55 | 2015-06-22 08:05:21.627771 |
     
  
### update the record with new value
 
 ```sql
 
 update itemmaster set rate = 35 where itemid = 1;
 
 ```
### after updating the record the rate_history table maintain history as follows
 
 	rate_history table data :
 	
 
	 itemid|  rate |          datefrom          |   dateend    
 
	 -----+------+----------------------------+-----------------------
	 
	   1  	|   30 | 2015-06-22 08:09:48.85778  |2015-06-23 08:00:48.85778
  	   2	|   80 | 2015-06-22 08:01:43.002791 | 
  	   3    |   55 | 2015-06-22 08:05:21.627771 |
  	   1    |   35 | 2015-06-23 08:00:48.85778  |
 

       
###before insert or update the  rows, change into upper case using trigger
```sql
	CREATE TABLE practise.new_book(id bigserial primary key,title text,author text);
```

### CREATE A trigger for to change text into upper case

```sql
	CREATE  OR REPLACE FUNCTION f_trigbooks() RETURNS trigger as
	$$
	BEGIN
		NEW.title := upper(NEW.title);
		RETURN NEW;
	END;
	$$
		LANGUAGE  plpgsql VOLATILE;

		CREATE trigger trig_books
		BEFORE INSERT  OR UPDATE  OF title,author on practise.new_book
		FOR EACH ROW
		EXECUTE PROCEDURE f_trigbooks();
```

