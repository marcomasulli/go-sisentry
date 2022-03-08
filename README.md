# go-sisentry

My first Golang project. It launches a request to the sisense.com subdomain of choice every 60 seconds, and check if any sisense build has failed. 
Sisense offers email notifications out of the box, however with this I can push notifications to all my team almost in real time, and in addition
to the default alerts it also gives information about the errors you get.

For now this pushes notifications only to MS teams, would like to add whatsapp/telegram or other protocols. in the future.
