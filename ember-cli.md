
##Emberjs

	Ember requires Node.js 0.12 or higher and npm 2.7 or higher.

##NPM (node package manager)

	I. it is helpful for js developers to share  and reuse the code

	 for eg:
		npm install ember-cli  -g
		npm install cofee-script -g

	CLI(command line interface)

	It install the program  in /usr/local/bin/

		usr/local/bin -it allows you to run the program from the console(cli tool)

  II.Dependency Management

		when you have a node project with package.json file,you can run npm install from the project root and npm will install all the dependencies listed in the package.json 
		
		$git clone https://github.com/yourname/store.git
		
		$cd store
		
		$npm install

		After running above commands, you will see a node_modules folder containing all of the project dependencies specified in the package.json.

###FILE package.json
	
		All npm packages contain a file in the project root(package.json).This file holds various metadata relevant to the project.

		This file is used to give information to npm that allows it to identify the project as well as handle the project's dependencies

		for more info run :  npm help json

##npm config setting for behind the corporate proxy
		
		$npm config set proxy http://192.168. .:310

		$npm config set https-proxy http://192.168. .:310

###Requirements

		Install nodejs using npm(Node Package Manager)

		$npm -g install ember-cli

###after installing ember-cli  create a new project

		syntax:ember new project_name

		$ember new simpleprj

		after enter this command ,new cmd will create a directory as follows

                Project Structure
					 -----------------
					 |-- Brocfile.js
					 |-- README.md
					 |-- app
					 |-- bower.json
					 |-- bower_components
					 |-- config
					 |-- node_modules
					 |-- package.json
					 |-- public
					 |-- testem.json
					 |-- tests
					 +-- vendor
					_________________

##pods:

		pods are help with directory readability and easy of file access when developing

		If you want to use the pods in your project,go and make the following changes in config/environment.js

		var ENV = {
		  
			  podModulePrefix: 'simpleprj/pods',
			//other config stuff...
		 }	

		and in .ember-cli set  usePods as true

		{
		  "usePods": true,
	     "disableAnalytics": true
		}
 


##Ember data Adapter

###Create a adapter

```js
		$ember g adapter application -- it will generate application.js in adapters folder in app folder
```
		or If you are using pod structure in application generate adapter as follows
	
```js
		$ember g adapter application --pod
```
		Ember data provides 2 types of data Adapters

		REST Adapter: For a http server access using JSON

		FIXTURE Adapter:For loading data from local memory

	Note:By default RESTAdapter

###we write the fixtures in model  as follows

		import DS from 'ember-data';

		 var Todo = DS.Model.extend({
				name: DS.attr('string'),
				isDone: DS.attr('boolean',{defaultValue: false})
			  });

			  Todo.reopenClass({

			  FIXTURES: [
				     {
					   id: 1,
						name: "prepare list",
						isDone:true
						},

						{
						id: 2,
						name:" install Ember-cli",
						},

						{
						id: 3,
						name: "practise Ember-cli",
					    }
					]

			});

			export default Todo;

