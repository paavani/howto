## Postgres
 
 Database is a collection 
 
	\h  --> help for SQL commands
	it shows list of SQL commands
	
	eg:\h alter schema
	
	\?  --> psql cmds

##Connect to the database using sudo cmd

	$sudo -u postgres psql databasename

	$sudo -u postgres psql mydb

##To know only postgresql path

    $sudo -u postgres printenv PATH
   
    Note: PATH is case sensitive,you must give it as capitals.


## To view current database connection                                                  
	\conninfo(it shows connected database name,user,port number and socket)
	or \c (it shows  connected database name and user only)

## To see list of databases
	\l or \l+ (it gives \l and including size,tablespace,description)

##To see list of schemas in database

	\dn -- It display list of schemas of connected database and owner

	\dt *.* --It display the list of schemas,tables,owner

## To see list of tables
	\d+ or \dt

	Note: By default it display the[ public schema] tables

##To see list of tables of particular schema

	syntax: \dt schemaname.*
	
	\dt practise.*

## To see table structure 
	\d tablename or 
	
## To list all tables,views and sequences in databasae 
	\dp or \z


### 1.Listing all the tables in the database

```sql
  SELECT table_name FROM information_schema.tables
  WHERE table_schema NOT IN ('information_schema','pg_catalog');
``` 
  
    The information_schema is a PostgreSQL schema available as per the SQL standards which contains a collection of views such as tables, columns etc.
	 The tables view provides the information about all tables in DB.
    
	 before enter into login db if u use this cmd then automatically ask pwd when using \c dbname
    psql -U username -d dbname -p 5433 -h  127.0.0.1 -W
   
### Roles and previledges
  
  \du
   
     It displays the list of  rolenames(users) of  DB and their Access previledges
    or
   
    select usename from pg_user;
   
    or
   
    To  know all users
    select * from pg_shadow;
   
##create a role
```sql

	CREATE ROLE role_name;

```
##How to grant permission in postgres

	syntax:	GRANT permission_type ON table_name to rolename
     
     eg: permissiontype((arwdDxt--append(insert),read(select),write(update),delete(delete),D(Truncate),x(references),t(trigger)
```sql

     GRANT all on tablename to rolename;
     GRANT insert on tablename to rolename;
```
     --after using the above cmd pls check \z
   
##Use Group roles in postgres

```sql

	CREATE role temp_users;
	
	GRANT temp_users to user1;
	
	GRANT temp_users to user3;

```

##To REMOVE THE Permission in POSTGRES

	Syntax:	REVOKE permission_type ON  table_name from user_name 
	
```sql
	REVOKE all on tablename from rolename;
	REVOKE insert on tablename from rolename;
         --check \z 
```

## To create a new trigger in PostgreSQL you need to:

	Create a trigger function using CREATE FUNCTION statement.
	
	Bind this trigger function to a table using CREATE TRIGGER statement.
	
### Creating Trigger function
```sql

	CREATE FUNCTION trigger_functionname RETURN  trigger AS
```
## To know list of functions

	Syntax:\df *functionname*
	
	\df *round* --It display list of available function of round

	\df *max*

	
## To list down all Triggers in database

```sql

	select * from pg_trigger;

```	
## To list down triggers on a particular table

```sql

	SELECT tgname FROM pg_trigger, pg_class WHERE tgrelid=pg_class.oid AND relname='table_name';

```

## Dropping trigger

```sql
	
	DROP TRIGGER trigger_name on table_name

```
	
	
	
##CREATE DATABASE

	Syntax: createdb  <DATABASE NAME>

```sql		
		createdb  testdb;
```

##Connect to the database

	$psql testdb

	Note:after connect to the database you can change to  another database using \c Target_database
##Drop database(delete the database)
	
	It delete the database

	syantax:dropdb <db name>

```sql
		dropdb testdb;
```

##ALTER DATABASE


###To change the database name
	
	Syntax:ALTER DATABASE  target_databasename  RENAME TO new_name;

```sql

	ALTER DATABASE  testdb RENAME TO testhrdb;

```

###To change the owner of db
	
	Syntax:ALTER DATABASE target_database  OWNER TO new_owner;

```sql

	ALTER DATABASE testhrdb OWNER TO hr;

```

	Note :if hr role does not exist, create it by using follwing way

```sql
	
	CREATE ROLE hr valid UNTIL 'infinity';
	
```

### To change the tablespace of db

		A tablespace is a location on disk where PostgreSQL stores data files containing database objects e.g., indexes., tables, etc. PostgreSQL uses a tablespace to map a logical name to physical location on disk.

	Two tablespaces are automatically created by initdb.
	
		1.pg_global  --It stores all global data
		2.pg_default

	The pg_global tablespace is used for shared system catalogs(it doesn't belongs to any particular database).
	
	The pg_default tablespace is the default tablespace of the template1 and template0 databases.
	
	NOTE:pg_default  will be the default tablespace for other databases as well, unless overridden by a TABLESPACE clause in CREATE DATABASE.

```sql
	
	ALTER DATABASE testhrdb
	SET TABLESPACE hr_default;
```

	if the hr_default tablespace does not exist,you can create it as follows

```sql
	
	CREATE TABLESPACE hr_default
	OWNER hr
	LOCATION '/pgdata/hr';

```
  Note: Name of the tablespace shouldn't start with pg_
  
  here  location denotes the home directory , for windows we use 'c:/pgdata/hr'

##CREATE TABLE

	CREATE TABLE item_test(itemname text,


























