##Postman Rest Client

Enter a name postman rest client -> add  it to chrome  ->
left side of the google chrome -> Apps -> select POSTMAN REST CLIENT

History->  Enter URL as https://127.0.0.1:9090 ->  request as POST -> select RAW

and click on SEND button

enter the https://127.0.0.1:9090/printBarCode

write "{"ItemCode":"APPLE","ItemName":"F107","LabelQty":1,"PrintToFileOnly":true}" and click on send button

##Testing of Barcode

	First run the bar code.exe file -->open url 127.0.0.1:9090/printBarcode '(get 404 page not found)' and follow above steps for testing in postman rest client.
	In barcode.exe file the records are print. 
