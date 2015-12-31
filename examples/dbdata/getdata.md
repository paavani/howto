//create a table in postgres DB

BEGIN;
CREATE SCHEMA IF NOT EXISTS pp;
CREATE TABLE IF NOT EXISTS pp.Itempost (
	ItemId serial PRIMARY KEY,
	ItemName text not null,
	Uom text not null default 0,
	Amount numeric not null default 0
	);

  insert into pp.Itempost(ItemName,Uom,Amount)values('apple','1kg',100),('banana','1kg',30),('carrot','1kg',50);
COMMIT;

Note:If You are using Ubuntu 
we have two ways to login into database

#1.Using sudo command 

   $sudo -u postgres psql mydb

#2.using psql

   $psql mydb

#Note
If you create a table using sudo command,and trying to get the data from the database using golang you will get the error.

