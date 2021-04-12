Environment Variables!

Supported databases
* [MySQL](https://www.mysql.com/) (`DONORFIDE_DATABASE=mysql`)
	* Recommended for use on self-hosted installations
* [PostgreSQL](https://www.postgresql.org/) (`DONORFIDE_DATABASE=postgres`)
	* Recommended for use on [Heroku](https://heroku.com)
* [BigQuery](https://cloud.google.com/bigquery/) (`DONORFIDE_DATABASE=bigquery`)
	* Recommended for use on [Google App Engine](https://cloud.google.com/appengine/)
* [SQL Server](https://www.microsoft.com/en-us/sql-server?rtc=1) (`DONORFIDE_DATABASE=mssql`)
	* Recommended for use on [Azure App Service](https://azure.microsoft.com/en-us/services/app-service/)
* [SQLite](https://sqlite.org/index.html) (`DONORFIDE_DATABASE=sqlite`)
	* Not recommended for use in production environments, however, it's super easy to set up if you want to test Donorfide locally.