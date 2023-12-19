
gofritz
===========================

A command line (CLI) application and Go library, to interact with AVM Fritz!Box routers/devices
via TR064 API.

This work was inspired by https://github.com/kanimaru/fritzbox-soap-example - thanks for that.

## Configure the TR064 API and a user

This is the basic setup to be done, so an application can use this API.

1. log into http://fritz.box/ via your admin account
2. enable TR064, via menu path... 
   'home network' -> 'network settings' -> 'advanced settings' -> tick the box 'allow access for applications'
3. create a user with password, via menu path...
   'system' -> 'Fritz!box user' -> edit or create a new one and set a password
